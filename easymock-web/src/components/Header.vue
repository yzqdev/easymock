<template>
  <transition name="fade">
    <div
      class="em-header"
      :class="{ 'em-header--fixed': isFixed }"
      v-show="pageAnimated"
    >
      <el-affix @change="changeFixed">
        <div class="em-header__content">
          <transition name="fade">
            <em-spots :size="spots" v-if="routeChanged"></em-spots>
          </transition>
          <div class="em-container">
            <el-row :gutter="20">
              <el-col  :span="1">
                <transition name="fade">
                  <div class="logo"><econ :name="icon" v-show="routeChanged"></econ></div>
                </transition>
              </el-col>
              <el-col :span="16" class="em-header__info">
                <transition-group name="fade">
                  <h2 key="a" v-show="routeChanged">{{ title }}</h2>
                  <p key="b" v-show="routeChanged">{{ description }}</p>
                </transition-group>
              </el-col>
              <el-col :span="6" class="em-header__control">
                <slot v-if="!nav"></slot>
              </el-col>
            </el-row>
          </div>
          <div class="em-header__nav" v-if="nav">
            <div class="em-container">
              <div style="float: right">
                <div
                  class="em-header__nav__item"
                  v-for="(item, i) in nav"
                  :class="{ 'is-active': value === item.title }"

                  :key="i"
                  @click="$emit('input', item.title)"
                >
                  <econ :name="item.icon" v-if="item.icon"></econ> {{ item.title }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </el-affix>
    </div>
  </transition>
</template>

<style lang="less" scoped>
@import "../styles/base/var.less";

 .em-header {
  height: 110px;
  position: relative;

  ::v-deep(.el-affix) {
    .em-header__content {
      padding: 30px 0;
    }
  }

  .em-header--fixed {
    z-index: 20;
  }

  .em-header__nav {
    position: absolute;
    bottom: -1px;
    width: 100%;
    font-size: 13px;

    .em-container{
      .em-header__nav__item {
        float: left;
        height: 40px;
        line-height: 40px;
        padding: 0 15px;
        color: #586069;
        cursor: pointer;
        border: solid transparent;
        border-width: 3px 1px 1px;
        border-radius: 3px 3px 0 0;

        i {
          margin-right: 2px;
        }

        &.is-active {
          color: #24292e;
          background: #f5f7fa;
          border-color: @colorPrimary #e1e4e8 transparent;
        }
      }
    }
  }

  .logo {
    width: 50px;
    height: 50px;
    background: @colorPrimary;
    border-radius: 4px;
    box-shadow: 0 10px 50px #2d8cf066;
    color: #fff;
    font-size: 28px;
    text-align: center;
    line-height: 50px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .em-header__content {
    padding: 30px 0;
    background: #fff;
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
    overflow-y: hidden;
    width: 100%;
    transition: var(--all-transition);
    position: relative;
  }

  .em-header__control {
    position: relative;
    top: 10px;

    &,
    .ivu-radio-group {
      float: right;
    }
  }

  .em-header__info {
    margin-top: 2px;
    margin-left: 20px;

    h2 {
      font-size: 18px;
      font-weight: 700;
    }

    p {
      font-size: 12px;
      color: #9a9ca0;
      font-weight: 400;
    }
  }
}
</style>

<script>
import {nextTick} from "vue";

export default {
  name: "EmHeader",
  props: {
    title: String,
    description: String,
    icon: String,
    spots: Number,
    nav: Array,
    value: {},
  },
  data() {
    return {
      routeChanged: true,
      isFixed: false,
    };
  },
  watch: {
    title () {
      this.routeChanged = false;
      nextTick(() => {
        this.routeChanged = true;
      });
    },
  },
  methods: {
    changeFixed(isFixed) {
      this.isFixed = isFixed;
    },
  },
};
</script>
