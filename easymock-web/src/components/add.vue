<template>
  <transition name="fadeUp">
    <div
      class="em-add"
      :class="custClass"
      :style="custStyle"
      v-show="pageAnimated">
      <i :class="icon"></i>
    </div>
  </transition>
</template>

<script>
export default {
  name: 'EmAdd',
  props: {
    icon: {
      type: String,
      default: 'el-icon-plus'
    },
    color: {
      type: String,
      default: 'primary'
    },
    bottom: {
      type: Number,
      default: 30
    }
  },
  computed: {
    custClass () {
      return [
        this.color === 'red' ? 'is-red' : 'is-primary'
      ]
    },
    custStyle () {
      return {
        bottom: this.bottom + 'px'
      }
    }
  }
}
</script>

<style lang="less" scoped>
@import "../styles/base/var.less";

.em-add {
  width: 42px;
  height: 42px;

  border-radius: 50%;
  position: fixed;
  right: calc(@basePadding - 10px);
  text-align: center;
  color: #fff;
  line-height: 42px;
  font-size: 18px;

  cursor: pointer;
  transition:@allTransition;
  z-index: 90;
}

.em-add.is-primary {
  background: @colorPrimary;
  box-shadow: 0 10px 30px @colorPrimary+'66';
}

.em-add.is-primary:hover {
  background: @colorPrimary+'CC';
}

.em-add.is-red {
  background: @colorRed;
  box-shadow: 0 10px 30px @colorRed+'66';
}

.em-add.is-red:hover {
  background: @colorRed+'CC';
}
</style>
