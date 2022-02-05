<template>
  <div class="em-docs">
    <em-header
      icon="ios-book"
      :title="page.title"
      :description="page.description"
    >
    </em-header>
    <em-keyboard-short></em-keyboard-short>
    <Back-top>
      <em-add icon="arrow-up-c" :bottom="90"></em-add>
    </Back-top>
    <div class="em-container">
      <transition name="fade">
        <div class="em-docs__content" v-show="pageAnimated">
          <transition name="fade">
            <el-affix
              :offset-top="70"
              @on-change="changeFixed"
              v-if="!isChangelog"
            >
              <el-menu mode="horizontal" class="em-docs__nav">
                <el-submenu
                  :name="'100-' + i"
                  v-for="(parent, i) in nav"
                  :key="i"
                  :index="`docs-${i}`"
                  v-if="parent&&parent.children&&parent.children.length > 0"
                >
                  <template slot="title">{{ parent.title }}</template>
                  <el-menu-item-group>
                    <el-menu-item
                      :name="o"
                      v-for="(item, o) in parent.children"
                      :key="o"
                      :index="`docs2-${o}`"
                      @click.native="go(item.id)"
                    >
                      {{ item.title }}
                    </el-menu-item>
                  </el-menu-item-group>
                </el-submenu>
                <el-menu-item name="101" @click.native="toChangelog">
                  <el-badge dot :count="readChangelog ? '0' : '1'">
                    {{ $tc("p.docs.header.title", 2) }}
                  </el-badge>
                </el-menu-item>
              </el-menu>
            </el-affix>
          </transition>
          <div
            class="docs-content"
            :class="{ 'is-fixed': isFixed, 'is-changelog': isChangelog }"
          >
            <transition-group name="fade">
              <doc
                class="markdown-body"
                ref="doc"
                key="doc"
                v-show="!isChangelog"
              >
              </doc>
              <changes
                class="markdown-body"
                ref="changelog"
                key="changelog"
                v-show="isChangelog"
              >
              </changes>
            </transition-group>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "../../styles/base/var";

.em-docs {
  .ivu-menu:after,
  .ivu-menu-item-group-title {
    display: none;
  }

  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item-active,
  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-item:hover,
  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-submenu-active,
  .ivu-menu-light.ivu-menu-horizontal .ivu-menu-submenu:hover {
    border-bottom: none;
    color: inherit;
  }

  .ivu-menu-horizontal .ivu-menu-item,
  .ivu-menu-horizontal .ivu-menu-submenu {
    transition: none;
  }

  .ivu-menu-horizontal
    .ivu-menu-submenu
    .ivu-select-dropdown
    .ivu-menu-item-selected {
    background: inherit;
  }

  .ivu-menu-horizontal,
  .ivu-menu-submenu,
  .ivu-select-dropdown,
  .ivu-menu-item-selected:hover {
    background: #f3f3f3;
  }

  .docs-content {
    padding: 20px 40px;
  }

  .ivu-affix {
    .em-docs__nav {
      margin-top: 0;
    }
  }

  .is-fixed {
    margin-top: 62px;
  }

  .markdown-body:last-child {
    padding-top: 20px;
    padding-bottom: 20px;
  }

  .markdown-body h1:target,
  .markdown-body h2:target,
  .markdown-body h3:target,
  .markdown-body h4:target {
    padding-top: 152px;
  }

  .is-changelog h1:target,
  .is-changelog h2:target,
  .is-changelog h3:target,
  .is-changelog h4:target {
    padding-top: 92px;
  }

  .em-docs__content {
    margin: 20px auto;
    border: 1px solid #eee;
    background: #fff;
    border-radius: 4px;
    box-shadow: 0 2px 3px #eee;

    ul,
    ol {
      list-style: inherit;
    }

    img {
      box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
      border-radius: 4px;
      margin: 10px 0;
    }

    th {
      background: #f7f7f7;
    }

    td,
    th {
      text-align: left !important;
    }
  }

  .em-docs__nav {
    list-style: none !important;
    background: #f8f8f8;
    border-bottom: 1px solid #eaecef;
    box-sizing: content-box;
    padding: 0 20px;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px;
  }

  .qq-qun img {
    width: 300px;
  }

  p.warning,
  p.tip {
    padding: 12px 24px 12px 20px;
    margin: 2em 0;
    border-left: 4px solid;
    background-color: #f8f8f8;
    position: relative;
    border-bottom-right-radius: 2px;
    border-top-right-radius: 2px;
    &:before {
      content: "!";
      position: absolute;
      top: 14px;
      left: -12px;
      color: #fff;
      width: 20px;
      height: 20px;
      border-radius: 100%;
      text-align: center;
      line-height: 20px;
      font-weight: bold;
      font-family: "Dosis", "Source Sans Pro", "Helvetica Neue", Arial,
        sans-serif;
    }
  }

  p.warning {
    border-left-color: #f66;
    &:before {
      background-color: #f66;
    }
  }

  p.tip {
    border-left-color: #3c763d;
    &:before {
      background-color: #3c763d;
    }
  }
}
</style>

<script>
import Changes from "./Changes.vue";
import doc from "./doc.vue";
import {nextTick} from "vue";
export default {
  name: "document",
  components: {
    Changes,
    doc,
  },
  data() {
    return {
      nav: [],
      isFixed: false,
      page: {
        title: "",
        description: "",
      },
    };
  },
  computed: {
    doc() {
      return this.$refs.doc;
    },
    isChangelog() {
      return this.$route.path === "/changelog";
    },
    readChangelog() {
      return this.$store.state.app.readChangelog;
    },
    appVersion() {
      return this.$store.state.app.version;
    },
  },
  mounted() {
    const doc = this.doc;
    console.log(doc.$el)
    console.log(`%c温迪来到`,`color:red;font-size:16px;background:transparent`)
    const docNodes = doc.$el.children;
     nextTick(() => {
      const hash = location.hash;
      if (hash) {
        location.href = hash;
      }
    });
    for (let len = docNodes.length, i = 0; i < len; i += 1) {
      const node = docNodes[i];
      const tagName = node.tagName.toLowerCase();
      if ("h2 h3".split(" ").includes(tagName)  ) {
        if (tagName === "h2") {
          this.nav.push({
            title: node.innerText,
            children: [],
          });
        } else {
          this.nav[this.nav.length - 1].children.push({
            id: node.id,
            title: node.innerText,
          });
        }
      }
    }
    this.changeRoute();
  },
  watch: {
    $route: "changeRoute",
  },
  methods: {
    go(id) {
      location.href = "#" + id;
    },
    changeFixed(isFixed) {
      this.isFixed = isFixed;
    },
    toChangelog() {
      this.$store.commit("app/SET_READ_CHANGELOG", true);
      localStorage.setItem("version", this.appVersion);
      this.$router.push("/changelog");
    },
    changeRoute() {
      this.isFixed = false;
      if (this.isChangelog) {
        this.page.title = this.$tc("p.docs.header.title", 2);
        this.page.description = this.$tc("p.docs.header.description", 2);
      } else {
        this.page.title = this.$tc("p.docs.header.title", 1);
        this.page.description = this.$tc("p.docs.header.description", 1);
      }
    },
  },
};
</script>
