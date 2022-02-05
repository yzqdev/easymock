<template>
  <div class="em-project">
    <em-placeholder :show="projects.length === 0">
      <econ :name="keywords ? 'outlet' : 'happy-outline'"></econ>
      <p>{{ keywords ? $t("p.project.placeholder[3]") : page.placeholder }}</p>
    </em-placeholder>
    <em-keyboard-short></em-keyboard-short>
    <em-header
      :icon="page.icon"
      :title="page.title"
      :description="page.description"
    >
      <el-radio-group
        border
        v-model="filterByAuthor"
        type="button"
        @change="handleFilter"
        v-if="page.type === 0"
      >
        <el-radio-button :label="$t('p.project.filter[0]')"></el-radio-button>
        <el-radio-button :label="$t('p.project.filter[1]')"></el-radio-button>
        <el-radio-button :label="$t('p.project.filter[2]')"></el-radio-button>
      </el-radio-group>
    </em-header>
    <el-dialog v-model="removeModal.show" width="360px">
      <template #title>
        <p style="color: #f60; text-align: center">
          <econ class="informationcircled"></econ>
          <span> {{ $t("p.project.modal.delete.title") }}</span>
        </p>
      </template>
      <div>
        <p>
          {{ $tc("p.project.modal.delete.description", 1) }}
          <strong style="word-break: break-all">
            {{
              (removeModal.project.user &&
                removeModal.project.user.nick_name) ||
              (removeModal.project.group && removeModal.project.group.name)
            }}
            / {{ removeModal.project.name }}</strong
          >
        </p>
        <p>{{ $tc("p.project.modal.delete.description", 2) }}</p>
        <el-input
            size="large"
          style="margin-top: 10px"
          v-model="removeModal.inputModel"
          :placeholder="$t('p.project.modal.delete.placeholder')"
        ></el-input>
      </div>
      <template #footer>
        <el-button
          type="danger"
          size="large"
          style="width: 100%"
          :disabled="removeModal.project.name !== removeModal.inputModel"
          @click="remove"
          >{{ $t("p.project.modal.delete.button") }}</el-button
        >
      </template>
    </el-dialog>
    <transition name="fade">
      <div class="em-container em-project__list" v-show="pageAnimated">
        <el-row class="ivu-row">
          <transition-group name="list-complete">
            <el-col
              :span="6"
              class="ivu-col ivu-col-span-6 list-complete-item"
              v-for="(item, index) in projects"
              :key="index"
            >
              <!-- 检查 user.id 防止闪烁 -->
              <div
                class="em-project__item"
                @click="go(item)"
                :class="{
                  'is-join':
                    page.type === 2 ||
                    (page.type === 0 && user.id && item.user._id !== user.id),
                  'is-group': page.type === 1,
                }"
              >
                <div class="project-collect">
                  <transition name="zoom" mode="out-in">
                    <i
                      :class="
                        item.extend.is_workbench
                          ? 'el-icon-star-on'
                          : 'el-icon-star-off'
                      "
                      :key="item.extend.is_workbench"
                      @click.native.stop="handleWorkbench(item.extend)"
                    ></i>
                  </transition>
                </div>
                <h2>{{ item.name }}</h2>
                <div class="project-description">{{ item.description }}</div>
                <div class="project-url">{{ item.url }}</div>
                <div class="project-member" v-if="page.type === 0">
                  <img :src="item.user.head_img" />
                  <img
                    :src="img.head_img"
                    v-for="(img, i) in item.members"
                    v-if="i < 5"
                    :key="i"
                  />
                </div>
                <el-button-group class="project-control">
                  <el-button


                    :title="$t('p.project.control[0]')"
                    class="copy-url"
                    @click="clip(item)"
                  ><i class="iconfont el-iconfont-link"></i>
                  </el-button>
                  <el-button


                    :title="$t('p.project.control[1]')"
                    style="width: 34%"
                    @click.stop="clone(item)"
                  ><i   class="iconfont el-iconfont-ios-copy"></i></el-button>
                  <el-button

                    :title="$t('p.project.control[2]')"
                    @click.stop="removeConfirm(item)"
                  ><i class="iconfont el-iconfont-trashb"></i></el-button>
                </el-button-group>
              </div>
            </el-col>
          </transition-group>
        </el-row>
      </div>
    </transition>
    <em-loading
      @loading="loading"
      ref="loading"
      v-if="page.type !== 2"
    ></em-loading>
  </div>
</template>

<script>
import Clipboard from "clipboard";
import debounce from "lodash/debounce";
import * as api from "../api";

export default {
  name: "project",
  data() {
    return {
      filterByAuthor: this.$t("p.project.filter[0]"),
      cliped: false,
      removeModal: {
        show: false,
        project: {},
        inputModel: "",
      },
    };
  },
  created() {
    this.$store.commit("project/INIT_REQUEST");
    this.$store.dispatch("project/INIT_PAGE", this.$route);
    this.$store.dispatch("project/FETCH");
  },
  mounted() {
    // this.$on(
    //   "query",
    //   debounce((keywords) => {
    //     this.$store.dispatch("project/QUERY", keywords);
    //   }, 500)
    // );
  },
  computed: {
    page() {
      const route = this.$route;
      switch (route.fullPath) {
        case "/workbench":
          return {
            title: this.$t("p.project.header.title[2]"),
            description: this.$t("p.project.header.description[2]"),
            placeholder: this.$t("p.project.placeholder[2]"),
            icon: "codeworking",
            type: 2, // 0.个人项目 1.团队项目 2.工作台
          };
        case "/":
          return {
            title: this.$t("p.project.header.title[0]"),
            description: this.$t("p.project.header.description[0]"),
            placeholder: this.$t("p.project.placeholder[0]"),
            icon: "person",
            type: 0,
          };
        default:
          const groupName = (route.query && route.query.name) || "";
          return {
            title: this.$t("p.project.header.title[1]", { groupName }),
            description: this.$t("p.project.header.description[1]", {
              groupName,
            }),
            placeholder: this.$t("p.project.placeholder[1]"),
            icon: "personstalker",
            type: 1,
          };
      }
    },
    projects() {
      return this.$store.state.project.list;
    },
    user() {
      return this.$store.state.user;
    },
    keywords() {
      return this.$store.state.project.keywords;
    },
  },
  watch: {
    $route: function () {
      this.filterByAuthor = this.$t("p.project.filter[0]");
      this.$store.commit("project/INIT_REQUEST");
      this.$store.dispatch("project/INIT_PAGE", this.$route);
      this.$store.dispatch("project/FETCH");
    },
  },
  methods: {
    go(project) {
      if (!this.cliped) {
        this.$router.push(`/project/${project._id}`);
      }
    },
    clip(project) {
      const clipboard = new Clipboard(".copy-url", {
        text() {
          return location.origin + "/mock/" + project._id + project.url;
        },
      });
      this.cliped = true;
      clipboard.on("success", (e) => {
        e.clearSelection();
        clipboard.destroy();
        this.cliped = false;
        this.$message.success(this.$t("p.project.copySuccess"));
      });
    },
    handleFilter(value) {
      let filterByAuthor = 0;
      if (value === this.$t("p.project.filter[1]")) {
        filterByAuthor = 1;
      } else if (value === this.$t("p.project.filter[2]")) {
        filterByAuthor = 2;
      }
      this.$store.commit("project/INIT_REQUEST");
      this.$store.commit("project/SET_REQUEST_PARAMS", { filterByAuthor });
      this.$store.dispatch("project/FETCH");
    },
    handleWorkbench(projectExtend) {
      this.$store.dispatch("project/WORKBENCH", projectExtend);
    },
    removeConfirm(project) {
      this.removeModal.show = true;
      this.removeModal.project = project;
      this.removeModal.inputModel = "";
    },
    remove() {
      const projectId = this.removeModal.project._id;
      this.$store.dispatch("project/REMOVE", projectId).then(() => {
        this.removeModal.show = false;
        this.$message.success(
          this.$t("p.project.deleteSuccess", {
            name: this.removeModal.project.name,
          })
        );
        this.$store.commit("project/SET_REQUEST_PARAMS", { pageIndex: 1 });
        this.$store.dispatch("project/FETCH");
      });
    },
    clone(project) {
      return api.project
        .copy({
          data: { id: project._id },
        })
        .then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t("p.project.cloneSuccess"));
            this.$store.commit("project/SET_REQUEST_PARAMS", { pageIndex: 1 });
            this.$store.dispatch("project/FETCH");
          }
        });
    },
    loading() {
      this.$store.dispatch("project/FETCH").then((data) => {
        this.$refs.loading.stop();
        if (data && data.length === 0) {
          this.$refs.loading.destroy();
        }
      });
    },
  },
};
</script>
<style lang="less" scoped>
@import "../styles/base/var";

.em-project {
  .em-project__list {
    margin-top: var(--base-padding);

    .ivu-row {
      margin-left: -10px;
      margin-right: -10px;
    }

    .ivu-col {
      padding-left: 10px;
      padding-right: 10px;
    }
  }

  .em-project__item {
    background: #fff;
    padding: 0 14px 20px 14px;
    box-shadow: 0 1px 5px #e5e5e5;
    border-radius: 4px;
    transition: @allTransition;
    position: relative;
    margin-bottom: @basePadding;
    cursor: pointer;

    &:hover {
      background: #fbfbfb;
    }

    &.is-join,
    &.is-group {
      h2,
      .project-collect,
      i {
        color: #fff;
      }

      .project-url,
      .project-description,
      .project-member,
      .ivu-btn:hover {
        background: #fff;
        border-color: #fff;
      }

      .ivu-btn {
        border-color: #fff;
      }
    }

    &.is-join {
      background: @colorPrimary;

      &:hover {
        background: #2d8cf0E6;
      }
      .ivu-btn:hover i {
        color: @colorPrimary;
      }
    }

    .is-group {
      background: @colorRed;

      &:hover {
        background: @colorRed+ "E6";
      }
      .ivu-btn:hover i {
        color: @colorRed;
      }
    }

    h2 {
      font-weight: 700;
      font-size: 14px;
    }

    .project-collect {
      text-align: center;
      font-size: 26px;
      color: @colorRed;

      i {
        cursor: pointer;
      }
    }

    h2,
    .project-url,
    .project-description,
    .project-member {
      .utils-ellipsis();
      .utils-user-select();
    }

    .project-url,
    .project-description,
    .project-member {
      padding: 10px;
      border: 1px solid #eee;
      border-radius: 4px;
      margin-top: 10px;
      background: #f5f5f5;

      &:before {
        content: "";
        display: block;
        width: 20px;
        height: 2px;
        background: @colorRed;
        border-radius: 4px;
        margin-bottom: 3px;
      }
    }

    .project-description:before {
      background: @colorPrimary;
    }

    .project-member {
      display: flex;

      &:before {
        display: none;
      }

      img {
        width: 20px;
        height: 20px;
        display: block;
        border-radius: 4px;
        margin-right: 12px;
      }
    }

    .project-control {
      margin-top: 10px;
      margin-left: 1px;
      width: 100%;

      .el-button {
        width: 33%;
        i{
          font-size: 12px;
        }
      }
    }
  }
}

.list-complete-item {
  transition: all 1s;
}
.list-complete-enter,
.list-complete-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
.list-complete-leave-active {
  position: absolute;
}
</style>
