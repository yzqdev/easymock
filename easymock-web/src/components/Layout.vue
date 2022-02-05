<template>
  <div class="em-layout">
    <em-add @click.native="$router.push('/new')"></em-add>
    <div v-shortkey="['p']" @shortkey="$router.push('/')"></div>
    <div v-shortkey="['g']" @shortkey="$router.push('/group')"></div>
    <div v-shortkey="['w']" @shortkey="$router.push('/workbench')"></div>
    <div v-shortkey="['d']" @shortkey="$router.push('/docs')"></div>
    <div v-shortkey="['n']" @shortkey="$router.push('/new')"></div>
    <div v-shortkey="['s']" @shortkey="onSearch"></div>

    <transition name="fade">
      <div class="em-layout__nav" v-show="pageAnimated">
        <!--  background-color="#495060"
          text-color="#ffffffb3"
          active-text-color="#ffffff"-->
        <el-menu router :default-active="pageKey" mode="horizontal">
          <div class="nav-logo" @click="$router.push('/')">
            <img src="/images/easy-mock.png" />
          </div>
          <div class="nav-search">
            <el-input
              v-model="searchValue"
              placeholder="Search Easy Mock"
              ref="search"
            ></el-input>
          </div>
          <el-submenu index="datatype">
            <template #title>
              <econ name="pound"></econ>
              {{ $t("c.layout.menu[0][0]") }}
            </template>
            <el-menu-item index="/">
              <econ name="person" size="10px" style="margin-right: 4px"></econ
              >{{ $t("c.layout.menu[0][1]") }}
            </el-menu-item>
            <el-menu-item index="/group">
              <econ
                name="personstalker"
                size="10px"
                style="margin-right: 4px"
              ></econ
              >{{ $t("c.layout.menu[0][2]") }}
            </el-menu-item>
          </el-submenu>
          <el-menu-item name="/workbench" index="/workbench">
            <econ name="codeworking"></econ> {{ $t("c.layout.menu[1]") }}
          </el-menu-item>
          <el-menu-item index="/dashboard">
            <econ name="ios-speedometer"></econ> {{ $t("c.layout.menu[2]") }}
          </el-menu-item>
          <el-menu-item index="/docs">
            <el-badge :is-dot="true" :count="readChangelog ? '0' : '1'">
              <econ name="iosbook"></econ> {{ $t("c.layout.menu[3]") }}
            </el-badge>
          </el-menu-item>
          <el-submenu index="egg">
            <template #title>
              <i class="iconfont el-iconfont-egg"></i>
              {{ $t("c.layout.menu[4][0]") }}
            </template>
            <li
              class="ivu-menu-item"
              @click="open('https://github.com/easy-mock/easy-mock')"
            >
              <i class="iconfont el-iconfont-link"></i> GitHub
            </li>
            <li
              class="ivu-menu-item"
              @click="open('https://github.com/easy-mock/easy-mock-cli')"
            >
              <i class="iconfont el-iconfont-link"></i>
              {{ $t("c.layout.menu[4][1]") }}
            </li>
            <li
              class="ivu-menu-item"
              @click="open('http://mockjs.com/examples.html')"
            >
              <econ name="link"></econ> {{ $t("c.layout.menu[4][2]") }}
            </li>
          </el-submenu>
          <el-submenu index="avatar" class="nav-avatar" v-show="userHeadImg">
            <template #title>
              <img :src="userHeadImg" v-show="userHeadImg" />
            </template>
            <el-menu-item name="/profile" index="/profile">
              <econ name="edit" size="14px"></econ>
              {{ $t("c.layout.menu[5][0]") }}
            </el-menu-item>
            <el-menu-item name="/log-out" @click.native="logOut">
              <econ name="log-out" size="14px"></econ>
              {{ $t("c.layout.menu[5][1]") }}
            </el-menu-item>
          </el-submenu>
          <el-menu-item class="nav-avatar" index="/login" v-show="!userHeadImg">
            <econ name="log-in-outline"></econ> {{ $t("c.layout.menu[5][2]") }}
          </el-menu-item>
        </el-menu>
      </div>
    </transition>
    <div class="em-layout__content">
      <transition name="fade" mode="out-in">
        <router-view></router-view>
      </transition>
    </div>
    <p class="em-layout__copyright" v-if="copyright && pageAnimated">
      {{ copyright }}
    </p>
  </div>
</template>

<style lang="less" scoped>
@import "../styles/base/var.less";

.em-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;

  ::v-deep(.em-layout__nav) {
    background: #495060;
    img {
      vertical-align: middle;
    }
    //  background-color="#495060"
    //text-color="#ffffffb3"
    //active-text-color="#ffffff"
    //active-text-color="#ffffff"

    .el-menu {
      max-width: @bodyWidth;
      margin: auto;
      background-color: #495060;

      .el-menu-item {
        &.is-active {

        }
      }
    }

    .el-menu.el-menu--horizontal {
      height: 60px;
      line-height: 60px;
      border-bottom: 0;
    }
    .nav-logo {
      float: left;
      margin-left: @basePadding;
      cursor: pointer;

      img {
        height: 20px;
      }
    }

    .nav-search {
      float: left;
      width: 150px;
      margin: 0 @basePadding;
    }

    .nav-avatar {
      float: right;

      img {
        border-radius: 50%;
        border: 2px solid #fff;
        width: 28px;
        height: 28px;
      }

      .ivu-select-dropdown {
        margin-right: 50px;
      }
    }

    .ivu-input-wrapper {
      vertical-align: inherit;
    }

    input {
      background: rgba(255, 255, 255, 0.125);
    }

    input:focus {
      background: rgba(255, 255, 255, 0.175);
    }

    input,
    input:hover,
    input:focus {
      color: #fff;
      border: none;
      box-shadow: none;
    }
  }

  .em-layout__content {
    flex: 1;
  }

  .em-layout__copyright {
    text-align: center;
    color: #c5c5c5;
    margin: 16px 0;
    flex: 0;
  }
}
</style>
<style lang="less">
.el-popper {
  .el-menu--popup {
    min-width: 140px;
    text-align: center;
  }
}
</style>
<script>
import config from "@/config/config";

export default {
  name: "EmLayout",
  data() {
    return {
      searchValue: "",
      pageKey: "/",
      copyright: config.copyright,
    };
  },
  computed: {
    userHeadImg() {
      return this.$store.state.user.headImg;
    },
    readChangelog() {
      return this.$store.state.app.readChangelog;
    },
  },
  mounted() {
    this.pageKey = this.$route.path;
    this.show = true;
  },
  watch: {
    $route(to) {
      this.pageKey = to.path;
    },
    searchValue: function (value) {
      // this.broadcast('group', 'query', value)
      // this.broadcast('project', 'query', value)
      // this.broadcast('projectDetail', 'query', value)
    },
  },
  methods: {
    open(url) {
      window.open(url);
    },
    logOut() {
      this.$router.push("/log-out");
    },
    onSearch() {
      this.$refs.search.focus();
    },
  },
};
</script>
