<template>
  <div class="em-editor">
    <div class="em-editor__editor">
      <div ref="codeEditor"></div>
    </div>
    <div class="panel-info">
      <em-spots :size="10"></em-spots>
      <div class="wrapper">
        <h2>{{isEdit ? $t('p.detail.editor.title[0]') : $t('p.detail.editor.title[1]')}}</h2>
        <div class="em-editor__form">
          <el-form label-position="top">
            <el-form-item label="Method">
              <el-select v-model="temp.method">
                <el-option   v-for="item in methods" :value="item.value" :key="item.value">{{ item.label }}</el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="URL">
              <el-input v-model="temp.url">
                <span slot="prepend">/</span>
              </el-input>
            </el-form-item>
            <el-form-item :label="$t('p.detail.columns[0]')">
              <el-input v-model="temp.description"></el-input>
            </el-form-item>
            <el-form-item :label="$t('p.detail.editor.autoClose')" v-if="isEdit">
              <el-switch v-model="autoClose"></el-switch>
            </el-form-item>
            <el-form-item>
              <el-button type="primary"  style="width: 100%" @click="submit">{{isEdit ? $t('p.detail.editor.action[0]') : $t('p.detail.editor.action[1]')}}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="em-editor__control">
          <div class="em-proj-detail__switcher">
            <ul>
              <li @click="format">{{$t('p.detail.editor.control[0]')}}</li>
              <li @click="preview" v-if="isEdit">{{$t('p.detail.editor.control[1]')}}</li>
              <li @click="close">{{$t('p.detail.editor.control[2]')}}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '../../styles/base/var';

.em-editor {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 999;
  display: flex;
  overflow: hidden;

  & > div {
    height: 100%;
    flex: 1;
  }

  .em-editor__editor {
    overflow: hidden;
    display: flex;
    flex-direction: column;

    & > div {
      height: 100%;
    }
  }

  & > div:last-child {
    flex: 0 0 350px;
  }

  .panel-info {
    background: @colorPrimary;
    padding: 20px 20px;
    overflow: auto;
    display: flex;
    align-items: center;
    box-shadow: 0 0 3px #000;

    h2 {
      color: #fff;
      font-size: 13px;
      text-align: center;
      padding: 6px;
      background: rgba(0,0,0,.5);
      border-radius: 30px;
      margin-bottom: 20px;
    }

    .wrapper {
      width: 100%;
    }
  }

.em-editor__form {
  padding: 20px 20px 1px;
  background: #fff;
  box-shadow: 0 2px 3px #777;
  border-radius: 4px;
  margin-bottom: 20px;
}

.em-editor__control {
  .em-proj-detail__switcher {
    border-radius: 4px;
    margin-bottom: 20px;
    color: #fff;
    box-shadow: 0 2px 3px #bbb;
    font-size: 13px;
    .utils-user-select();

    ul {
      display: table;
      table-layout: fixed;
      width: 100%;
    }

    li:first-child {
      border-top-left-radius: 4px;
      border-bottom-left-radius: 4px;
    }

    li:last-child {
      border-top-right-radius: 4px;
      border-bottom-right-radius: 4px;
    }

    li {
      background: #252d47;
      padding: 15px 0;
      text-align: center;
      display: table-cell;
      cursor: pointer;
      transition: @allTransition;

      &:hover {
        background: #252d47;
      }
    }
  }
}
}

</style>

<script>
import * as api from '../../api'
import jsBeautify from 'js-beautify/js/lib/beautify'
import * as ace from 'brace'
import 'brace/mode/javascript';
import 'brace/theme/monokai';
import 'brace/ext/language_tools'
import 'brace/ext/searchbox'
// import './snippets'
// if (typeof window !== 'undefined') {
//   ace = require('brace')
//   require('brace/mode/javascript')
//   require('brace/theme/monokai')
//   require('brace/ext/language_tools')
//   require('brace/ext/searchbox')
//   require('./snippets')
// }

export default {
  name: 'editor',
  data () {
    return {
      codeEditor: null,
      autoClose: true,
      methods: [
        { label: 'get', value: 'get' },
        { label: 'post', value: 'post' },
        { label: 'put', value: 'put' },
        { label: 'delete', value: 'delete' },
        { label: 'patch', value: 'patch' }
      ],
      temp: {
        url: '',
        mode: '{"data": {}}',
        method: 'get',
        description: ''
      }
    }
  },
  computed: {
    mockData () {
      return this.$store.state.mock.editorData.mock
    },
    baseUrl () {
      return this.$store.state.mock.editorData.baseUrl
    },
    projectId () {
      return this.$route.params.projectId
    },
    isEdit () {
      return !!this.$route.params.id && this.mockData
    }
  },
  beforeRouteEnter (to, from, next) {
    if (from.matched.length === 0) { // 防止编辑页刷新导致的显示异常（直接进入项目主页）
      return next({
        path: `/project/${to.params.projectId}`,
        replace: true
      })
    }
    next()
  },
  mounted () {
    this.codeEditor = ace.edit(this.$refs.codeEditor)
    this.codeEditor.getSession().setMode('ace/mode/javascript')
    this.codeEditor.setTheme('ace/theme/monokai')
    this.codeEditor.setOption('tabSize', 2)
    this.codeEditor.setOption('fontSize', 15)
    this.codeEditor.setOption('enableLiveAutocompletion', true)
    this.codeEditor.setOption('enableSnippets', true)
    this.codeEditor.clearSelection()
    this.codeEditor.getSession().setUseWorker(false)
    this.codeEditor.on('change', this.onChange)
    this.codeEditor.commands.addCommand({
      name: 'save',
      bindKey: {win: 'Ctrl-S', mac: 'Command-S'},
      exec: () => {
        this.submit()
      }
    })

    if (this.isEdit) {
      this.autoClose = true
      this.temp.url = this.mockData.url.slice(1) // remove /
      this.temp.mode = this.mockData.mode
      this.temp.method = this.mockData.method
      this.temp.description = this.mockData.description
    }

    this.$nextTick(() => {
      this.codeEditor.setValue(this.temp.mode)
      this.format()
    })
  },
  methods: {
    convertUrl (url) {
      const newUrl = '/' + url
      return newUrl === '/'
        ? '/'
        : newUrl.replace(/\/\//g, '/').replace(/\/$/, '')
    },
    format () {
      const context = this.codeEditor.getValue()
      let code = /^http(s)?/.test(context)
        ? context
        : jsBeautify.js_beautify(context, { indent_size: 2 })
      this.codeEditor.setValue(code)
    },
    onChange () {
      this.temp.mode = this.codeEditor.getValue()
    },
    close () {
      this.$store.commit('mock/SET_EDITOR_DATA', {mock: null, baseUrl: ''})
      this.$router.replace(`/project/${this.projectId}`)
    },
    submit () {
      const mockUrl = this.convertUrl(this.temp.url)

      try {
        const value = (new Function(`return ${this.temp.mode}`))() // eslint-disable-line
        if (!value) {
          this.$message.error(this.$t('p.detail.editor.submit.error[0]'))
          return
        } else if (typeof value !== 'object') {
          throw new Error()
        }
      } catch (error) {
        if (!/^http(s)?:\/\//.test(this.temp.mode)) {
          this.$message.error(error.message || this.$t('p.detail.editor.submit.error[1]'))
          return
        }
      }

      if (this.isEdit) {
        api.mock.update({
          data: {
            ...this.temp,
            id: this.mockData._id,
            url: mockUrl
          }
        }).then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t('p.detail.editor.submit.updateSuccess'))
            if (this.autoClose) this.close()
          }
        })
      } else {
        api.mock.create({
          data: {
            ...this.temp,
            url: mockUrl,
            project_id: this.projectId
          }
        }).then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t('p.detail.create.success'))
            this.close()
          }
        })
      }
    },
    preview () {
      window.open(this.baseUrl + this.mockData.url + '#!method=' + this.mockData.method)
    }
  }
}
</script>
