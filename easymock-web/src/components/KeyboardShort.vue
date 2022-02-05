<template>
  <div class="em-keyboard-short"
      v-shortkey="['shift', '?']" @shortkey="show">
    <el-dialog
      v-model="visible"
      :title="$t('c.keyboardShort.modalTitle')"
      :transfer="false">
      <div class="em-keyboard-short__content">
        <el-row :gutter="20">
          <el-col :span="12">
            <div>
              <h2>{{$t('c.keyboardShort.keyboards[0].category')}}</h2>
              <ul>
                <li>
                  <code class="keyboard-short">p</code>
                  {{$t('c.keyboardShort.keyboards[0].list[0]')}}
                </li>
                <li>
                  <code class="keyboard-short">g</code>
                  {{$t('c.keyboardShort.keyboards[0].list[1]')}}
                </li>
                <li>
                  <code class="keyboard-short">w</code>
                  {{$t('c.keyboardShort.keyboards[0].list[2]')}}
                </li>
                <li>
                  <code class="keyboard-short">d</code>
                  {{$t('c.keyboardShort.keyboards[0].list[3]')}}
                </li>
              </ul>
            </div>
            <div>
              <h2>{{$t('c.keyboardShort.keyboards[1].category')}}</h2>
              <ul>
                <li>
                  <code class="keyboard-short">n</code>
                  {{$t('c.keyboardShort.keyboards[1].list[0]')}}
                </li>
                <li>
                  <code class="keyboard-short">s</code>
                  {{$t('c.keyboardShort.keyboards[1].list[1]')}}
                </li>
              </ul>
            </div>
          </el-col>
          <el-col :span="12">
            <div v-for="(item, index) in value" :key="KeyboardShort">
              <h2>{{item.category}}</h2>
              <ul>
                <li v-for="(item, index) in item.list" :key="KeyboardShort">
                  <code
                    class="keyboard-short"
                    v-for="(short, i) in item.shorts"
                    :key="i">{{short}}</code>
                  {{item.description}}
                </li>
              </ul>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-dialog>
  </div>
</template>

<style lang="less" scoped>
@import '../styles/base/var.less';

@c keyboard-short {
  @d content {
    h2 {
      font-weight: 700;
    }

    li {
      margin: 5px 0;
    }

    .keyboard-short {
      display: inline-block;
      padding: 3px 5px;
      line-height: 10px;
      color: #444d56;
      vertical-align: middle;
      background-color: #fafbfc;
      border: solid 1px #d1d5da;
      border-bottom-color: #c6cbd1;
      border-radius: 3px;
      box-shadow: inset 0 -1px 0 #c6cbd1;
      margin-right: 5px;
    }
  }
}

</style>

<script>
export default {
  name: 'EmKeyboardShort',
  data () {
    return {
      visible: false
    }
  },
  props: {
    value: {
      type: Array
    }
  },
  methods: {
    show () {
      this.visible = !this.visible
    }
  }
}
</script>
