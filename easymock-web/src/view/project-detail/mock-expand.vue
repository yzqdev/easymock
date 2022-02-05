<style lang="less" scoped>

.em-mock-expand {
  background: #fff;
  border-radius: 4px;
  box-shadow: 0 1px 5px #ddd;
  padding: 20px;

  h2 {
    font-size: 14px;
    font-weight: 600;
  }

  p {
    margin-top: 6px;
    margin-bottom: 10px;
    background: #41444e;
    color: #fff;
    padding: 6px;
    border-radius: 4px;
    font-size: 13px;
  }

  pre {
    margin: 0;
  }
}

.em-mock__data-type-expand {
  background: #41444e;
  padding: 10px;
  color: #fff;
  border-radius: 4px;
  overflow: scroll;

  p,
  pre {
    margin: 0;
    padding: 0;
  }
}
</style>
<template>
  <div class="em-mock-expand">
    <h2>Method</h2>
    <p>{{mock.method}}</p>
    <h2>URL</h2>
    <p>{{mock.url}}</p>
    <h2>{{$t('p.detail.expand.description')}}</h2>
    <p>{{mock.description}}</p>
    <el-tabs value="request" v-if="mock.parameters || mock.response_model">
      <el-tab-pane :label="$t('p.detail.expand.tab[0]')" name="request" v-if="mock.parameters">
        <Table :columns="columnsRequest" :data="request"></Table>
      </el-tab-pane>
      <el-tab-pane :label="$t('p.detail.expand.tab[1]')" name="response" v-if="mock.response_model">
        <Table :columns="columnsResponse" :data="response"></Table>
      </el-tab-pane>
      <el-tab-pane label="Class Model" name="class" v-if="mock.response_model && entities.js.length">
        <el-collapse>
          <el-card>
            JavaScript
            <template #content>
              <article v-for="(item, i) in entities.js" :key="i">
              <pre>{{item}}</pre>
              </article>
            </template>
          </el-card>
          <el-card>
            Objective-C
            <template #content>
              <article v-for="(item, i) in entities.oc" :key="i">
              <pre>{{item}}</pre>
              </article>
            </template>
          </el-card>
        </el-collapse>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import {
  getJavaScriptEntities,
  getObjectiveCEntities
} from 'swagger-parser-mock/lib/entity'
import jsBeautify from 'js-beautify/js/lib/beautify'
import DataTypeExpand from './data-type-expand.vue'

export default {
  props: {
    mock: {}
  },
  data () {
    return {
      columnsRequest: [
        {
          type: 'expand',
          width: 50,
          render: (h, params) => h(DataTypeExpand, { props: { data: params.row.parameter } })
        },
        { title: this.$t('p.detail.expand.columnsRequest[0]'), key: 'name' },
        { title: this.$t('p.detail.expand.columnsRequest[1]'), key: 'description' },
        { title: this.$t('p.detail.expand.columnsRequest[2]'), key: 'paramType' },
        { title: this.$t('p.detail.expand.columnsRequest[3]'), key: 'dataType' }
      ],
      columnsResponse: [
        {
          type: 'expand',
          width: 50,
          render: (h, params) => h(DataTypeExpand, { props: { data: params.row.response } })
        },
        { title: this.$t('p.detail.expand.columnsResponse[0]'), key: 'code' },
        { title: this.$t('p.detail.expand.columnsResponse[1]'), key: 'message' }
      ]
    }
  },
  computed: {
    request () {
      const parameters = this.mock.parameters ? JSON.parse(this.mock.parameters) : []
      return parameters.map(parameter => ({
        name: parameter.name,
        description: parameter.description || this.$t('p.detail.expand.defaultDescription'),
        paramType: parameter.in,
        dataType: this.getParamDataType(parameter),
        parameter: parameter
      }))
    },
    response () {
      const responseModel = this.mock.response_model ? JSON.parse(this.mock.response_model) : {}
      return Object.keys(responseModel).map(code => {
        const response = responseModel[code]
        return {
          code: code,
          message: response.message || response.description || this.$t('p.detail.expand.defaultDescription'),
          response: response
        }
      })
    },
    entities () {
      const res = this.response.filter(o => {
        const code = o.code.toString()
        return code === '200' || code === 'default'
      })[0]
      const response = res ? res.response : {}

      return {
        js: getJavaScriptEntities(response).map(o => jsBeautify.js_beautify(o, { indent_size: 2 })),
        oc: getObjectiveCEntities(response)
      }
    }
  },
  methods: {
    getParamDataType (parameter) {
      const { type, schema } = parameter
      if (type) return type
      if (schema && schema.type) {
        return schema.type === 'array' ? schema.items.type : schema.type
      }
    }
  }
}
</script>
