<style lang="less" scoped>
@import "../../styles/base/var";

.em-proj-detail {
  padding-bottom: 20px;

  td .ivu-table-cell {
    font-size: 14px;
  }

  .method-tag {
    width: 100%;
    text-align: center;
  }

  .em-proj-detail__info {
    margin: 20px 0;
    font-size: 16px;
    color: #fff;
    background: @colorPrimary;
    border-radius: 4px;
    box-shadow: 0 2px 3px #eee;

    .el-col {
      padding: 10px;
    }

    .el-col:first-child {
      padding: 30px 20px;
    }

    .el-col:last-child {
      border-top-right-radius: 4px;
      border-bottom-right-radius: 4px;
      box-shadow: 0 0 3px #000;
      text-align: center;
      background: #252d47;
      position: absolute;
      right: 0;
      top: 0;
      bottom: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      img {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        border: 2px solid #fff;
      }
    }

    .author {
      background: rgba(0, 0, 0, 0.3);
      border-radius: 20px;
      padding: 2px 10px;
      font-size: 14px;
      margin-top: 10px;
      max-width: 200px;
      .utils-ellipsis();
    }

    .tag {
      background: rgba(0, 0, 0, 0.3);
      border-radius: 20px;
      font-size: 13px;
      margin-top: 16px;
      padding-right: 10px;
      .utils-ellipsis();

      span {
        padding: 5px 10px;
        background: rgba(0, 0, 0, 0.5);
        display: inline-block;
        border-top-left-radius: 20px;
        border-bottom-left-radius: 20px;
        margin-right: 10px;
        width: 100px;
        text-align: center;
      }
    }
  }

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

  .em-proj-detail__members {
    background: #fff;
    margin-bottom: 20px;
    border-radius: 4px;
    box-shadow: 0 2px 3px #eee;
    padding: 20px 20px 0;

    h2 {
      font-size: 16px;
      margin-bottom: 10px;
    }

    img {
      width: 60px;
      height: 60px;
      border-radius: 50%;
      display: block;
      margin-bottom: 20px;
    }
  }

  .ivu-table-cell .ivu-dropdown {
    margin-left: 5px;
  }
}

</style>
<template>
  <div class="em-proj-detail">
    <em-header
      i="cube"
      :title="project.name"
      :description="page.description"
      :nav="nav"
      v-model="pageName"
    >
    </em-header>
    <div v-shortkey="['tab']" @shortkey="handleKeyTab()"></div>
    <em-keyboard-short v-model="keyboards"></em-keyboard-short>
    <el-backtop>
      <em-add i="arrow-up-c" :bottom="90"></em-add>
    </el-backtop>
    <transition name="fade" mode="out-in">
      <project
        v-if="pageName === $t('p.detail.nav[1]')"
        key="a"
        :project-data="project"
      ></project>
    </transition>
    <transition name="fade" mode="out-in">
      <div
        class="em-container"
        v-if="pageAnimated && pageName === $t('p.detail.nav[0]')"
        key="b"
      >
        <div class="em-proj-detail__info">
          <el-row>
            <el-col :span="19">
              <em-spots :size="6"></em-spots>
              {{ project.description }}
              <p class="tag">
                <span>Base URL</span>
                {{ baseUrl }}
              </p>
              <p class="tag">
                <span>Project ID</span>
                {{ project._id }}
              </p>
            </el-col>
            <el-col :span="5">
              <div>
                <img
                  :src="
                    group ? '/images/group-default.png' : project.user.head_img
                  "
                />
                <p class="author">
                  {{ group ? group.name : project.user.nick_name }}
                </p>
              </div>
            </el-col>
          </el-row>
        </div>
        <div class="em-proj-detail__switcher">
          <ul>
            <li
              @click="openEditor()"
              v-shortkey="['ctrl', 'n']"
              @shortkey="openEditor()"
            >
              <i class="el-icon-plus"></i> {{ $t("p.detail.create.action") }}
            </li>
            <li
              @click="handleWorkbench"
              v-shortkey="['ctrl', 'w']"
              @shortkey="handleWorkbench"
            >
              <transition name="zoom" mode="out-in">
                <i
                  :type="
                    project.extend.is_workbench
                      ? 'android-star'
                      : 'android-star-outline'
                  "
                  :key="project.extend.is_workbench"
                ></i>
              </transition>
              {{ $t("p.detail.workbench") }}
            </li>
            <li
              @click="updateBySwagger"
              v-shortkey="['ctrl', 's']"
              @shortkey="updateBySwagger"
            >
              <i type="loop"></i> {{ $t("p.detail.syncSwagger.action") }}
            </li>
            <li @click="download">
              <i type="code-download"></i> {{ $tc("p.detail.download", 1) }}
            </li>
          </ul>
        </div>
        <div class="em-proj-detail__members" v-if="project.members.length">
          <em-spots :size="6"></em-spots>
          <h2><i type="person-stalker"></i> {{ $t("p.detail.member") }}</h2>
          <el-row :gutter="20">
            <el-col
              :span="2"
              v-for="(item, index) in project.members"
              :key="index"
            >
              <img :src="item.head_img" :title="item.nick_name" />
            </el-col>
          </el-row>
        </div>
        <el-table
          border
          :data="list"
          :default-sort="{ prop: 'url', order: 'descending' }"
          @selection-change="selectionChange"
          :highlight-current-row="true"
        >
          <el-table-column type="expand">
            <template #default="{ row }">
              <MockExpand :mock="row"></MockExpand>
            </template>
          </el-table-column>
          <el-table-column type="selection" width="55"> </el-table-column>
          <el-table-column
            label="Method"
            prop="title"
            :filters="[
              { text: 'get', value: 'get' },
              { text: 'post', value: 'post' },
              { text: 'put', value: 'put' },
              { text: 'delete', value: 'delete' },
              { text: 'patch', value: 'patch' },
            ]"
            :filter-method="filterHandler"
          >
            <template #default="{ row }">
              <el-tag class="method-tag" :color="getColor(row.method)">
                {{ row.method.toUpperCase() }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column sortable="" label="Url" prop="url">
          </el-table-column>
          <el-table-column
            :label="$t('p.detail.columns[0]')"
            prop="description"
          >
          </el-table-column>
          <el-table-column label="操作" prop="action">
            <template #default="{ row }">
              <div>
                <el-button-group>
                  <el-tooltip :content="$t('p.detail.action[0]')">
                    <el-button size="small" @click="preview(row)"
                      ><i class="el-icon-eye"></i
                    ></el-button>
                  </el-tooltip>
                  <el-tooltip :content="$t('p.detail.action[1]')">
                    <el-button size="small" @click="openEditor(row)"
                      ><i class="el-icon-edit"></i
                    ></el-button>
                  </el-tooltip>
                  <el-tooltip :content="$t('p.detail.action[2]')">
                    <el-button
                      size="small"
                      class="copy-url"
                      @click="clip(row.url)"
                      ><i type="link"></i
                    ></el-button>
                  </el-tooltip>
                </el-button-group>
                <el-dropdown>
                  <el-button size="small"><i type="more"></i></el-button>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item @click="clone(row)"
                        ><i class="el-icon-document-copy"></i>
                        {{ $t("p.detail.action[3]") }}</el-dropdown-item
                      >
                      <el-dropdown-item @click="download(row._id)"
                        ><i class="el-icon-download"></i>
                        {{ $tc("p.detail.download", 2) }}</el-dropdown-item
                      >
                      <el-dropdown-item @click="remove(row._id)"
                        ><i class="el-icon-delete"></i>
                        {{ $t("p.detail.action[4]") }}</el-dropdown-item
                      >
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </transition>
  </div>
</template>

<script>
import Clipboard from "clipboard";
import debounce from "lodash/debounce";

import * as api from "../../api";
import Project from "@/view/Project.vue";
import MockExpand from "@/view/project-detail/mock-expand.vue";

export default {
  name: "projectDetail",
  data() {
    return {
      pageName: this.$t("p.detail.nav[0]"),
      selection: [],
      keywords: "",
      nav: [
        { title: this.$t("p.detail.nav[0]"), i: "android-list" },
        { title: this.$t("p.detail.nav[1]"), i: "gear-a" },
      ],
      keyboards: [
        {
          category: this.$t("p.detail.keyboards[0].category"),
          list: [
            {
              description: `${this.$t("p.detail.nav[0]")}/${this.$t(
                "p.detail.nav[1]"
              )}`,
              shorts: ["tab"],
            },
          ],
        },
        {
          category: this.$t("p.detail.keyboards[1].category"),
          list: [
            {
              description: this.$t("p.detail.keyboards[1].list[0]"),
              shorts: ["ctrl", "n"],
            },
            {
              description: this.$t("p.detail.keyboards[1].list[1]"),
              shorts: ["ctrl", "w"],
            },
            {
              description: this.$t("p.detail.keyboards[1].list[2]"),
              shorts: ["ctrl", "s"],
            },
          ],
        },
      ],
      // el-columns: [
      //   {
      //     type: 'expand',
      //     width: 50,
      //     render: (h, params) => {
      //       return h(MockExpand, {
      //         props: {
      //           mock: params.row
      //         }
      //       })
      //     }
      //   },
      //   { type: 'selection', width: 60, align: 'center' },
      //   {
      //     title: 'Method',
      //     width: 110,
      //     key: 'method',
      //     filters: [
      //       { label: 'get', value: 'get' },
      //       { label: 'post', value: 'post' },
      //       { label: 'put', value: 'put' },
      //       { label: 'delete', value: 'delete' },
      //       { label: 'patch', value: 'patch' }
      //     ],
      //     filterMethod (value, row) {
      //       return row.method.indexOf(value) > -1
      //     },
      //     render: (h, params) => {
      //       const color = {
      //         get: 'blue',
      //         post: 'green',
      //         delete: 'red',
      //         put: 'yellow',
      //         patch: 'yellow'
      //       }
      //       return <tag class="method-tag" color={color[params.row.method]}>
      //         {params.row.method.toUpperCase()}
      //       </tag>
      //     }
      //   },
      //   { title: 'URL', width: 420, ellipsis: true, sortable: true, key: 'url' },
      //   { title: this.$t('p.detail.el-columns[0]'), ellipsis: true, key: 'description' },
      //   {
      //     title: this.$t('p.detail.el-columns[1]'),
      //     key: 'action',
      //     width: 160,
      //     align: 'center',
      //     render: (h, params) => {
      //       return (
      //         <div>
      //           <el-button-group>
      //             <el-button size="small" title={this.$t('p.detail.action[0]')} onClick={this.preview.bind(this, params.row)}><i type="eye"></i></el-button>
      //             <el-button size="small" title={this.$t('p.detail.action[1]')} onClick={this.openEditor.bind(this, params.row)}><i type="edit"></i></el-button>
      //             <el-button size="small" title={this.$t('p.detail.action[2]')} class="copy-url" onClick={this.clip.bind(this, params.row.url)}><i type="link"></i></el-button>
      //           </el-button-group>
      //           <dropdown>
      //             <el-button size="small"><i type="more"></i></el-button>
      //             <dropdown-menu slot="list">
      //               <dropdown-item nativeOnClick={this.clone.bind(this, params.row)}><i type="ios-copy"></i> {this.$t('p.detail.action[3]')}</dropdown-item>
      //               <dropdown-item nativeOnClick={this.download.bind(this, params.row._id)}><i type="ios-download"></i> {this.$tc('p.detail.download', 2)}</dropdown-item>
      //               <dropdown-item nativeOnClick={this.remove.bind(this, params.row._id)}><i type="trash-b"></i> {this.$t('p.detail.action[4]')}</dropdown-item>
      //             </dropdown-menu>
      //           </dropdown>
      //         </div>
      //       )
      //     }
      //   }
      // ]
    };
  },
  created() {
    this.$store.commit("mock/INIT_REQUEST");
    return this.$store.dispatch("mock/FETCH", this.$route);
  },
  mounted() {
    // this.$on('query', debounce((keywords) => {
    //   this.keywords = keywords
    // }, 500))
  },
  computed: {
    project() {
      return this.$store.state.mock.project;
    },
    list() {
      const list = this.$store.state.mock.list;
      const reg = this.keywords && new RegExp(this.keywords, "i");
      return reg
        ? list.filter(
            (item) =>
              reg.test(item.name) || reg.test(item.url) || reg.test(item.method)
          )
        : list;
    },
    page() {
      return {
        description: this.project?.user
          ? this.$t("p.detail.header.description[0]")
          : this.$t("p.detail.header.description[1]"),
      };
    },
    baseUrl() {
      const baseUrl = location.origin + "/mock/" + this.project._id;
      return this.project.url === "/" ? baseUrl : baseUrl + this.project.url;
    },
    group() {
      return this.project.group;
    },
  },
  methods: {
    filterHandler(value, row, column) {
      return row.method.includes(value);
    },
    handleKeyTab() {
      this.pageName =
        this.pageName === this.$t("p.detail.nav[1]")
          ? this.$t("p.detail.nav[0]")
          : this.$t("p.detail.nav[1]");
    },
    clip(mockUrl) {
      const clipboard = new Clipboard(".copy-url", {
        text: () => {
          return this.baseUrl + mockUrl;
        },
      });
      clipboard.on("success", (e) => {
        e.clearSelection();
        clipboard.destroy();
        this.$message.success(this.$t("p.detail.copySuccess"));
      });
    },
    preview(mock) {
      window.open(this.baseUrl + mock.url + "#!method=" + mock.method);
    },
    selectionChange(selection) {
      this.selection = selection;
    },
    download(mockId) {
      if (typeof mockId === "string") {
        const ids = this.selection.length
          ? this.selection.map((item) => item._id)
          : [mockId];
        api.mock.export(ids);
      } else {
        api.mock.export(this.project._id);
      }
    },
    updateBySwagger() {
      if (!this.project.swagger_url) {
        this.$message.warning(this.$t("p.detail.syncSwagger.warning"));
        return;
      }
      this.$confirm({
        title: this.$t("confirm.title"),
        content: this.$t("p.detail.syncSwagger.confirm"),
        onOk: () => {
          api.project
            .updateSwagger({
              data: { id: this.project._id },
            })
            .then((res) => {
              if (res.data.success) {
                const syncErrorURLs = res.data.data.syncErrorURLs;
                if (syncErrorURLs.length) {
                  this.$Notice.success({
                    title: this.$t("p.detail.syncSwagger.syncResult"),
                    desc: this.$t("p.detail.syncSwagger.success"),
                  });
                  this.$Notice.warning({
                    title: this.$t("p.detail.syncSwagger.syncFailed.title"),
                    duration: 0,
                    desc: `${syncErrorURLs.join(", ")} ${this.$t(
                      "p.detail.syncSwagger.syncFailed.desc"
                    )}`,
                  });
                } else {
                  this.$message.success(
                    this.$t("p.detail.syncSwagger.success")
                  );
                }
                this.$store.commit("mock/SET_REQUEST_PARAMS", {
                  pageIndex: 1,
                });
                this.$store.dispatch("mock/FETCH", this.$route);
              }
              return res;
            });
        },
      });
    },
    remove(mockId) {
      const ids = this.selection.length
        ? this.selection.map((item) => item._id)
        : [mockId];
      this.$confirm({
        title: this.$t("confirm.title"),
        content:
          ids.length > 1
            ? this.$t("p.detail.remove.confirm[0]")
            : this.$t("p.detail.remove.confirm[1]"),
        onOk: () => {
          api.mock
            .delete({
              data: { project_id: this.project._id, ids },
            })
            .then((res) => {
              if (res.data.success) {
                this.$message.success(this.$t("p.detail.remove.success"));
                this.$store.commit("mock/SET_REQUEST_PARAMS", {
                  pageIndex: 1,
                });
                this.$store.dispatch("mock/FETCH", this.$route);
              }
            });
        },
      });
    },
    handleWorkbench() {
      this.$store.dispatch("project/WORKBENCH", this.project.extend);
    },
    clone(mock) {
      this.$store
        .dispatch("mock/CREATE", {
          route: this.$route,
          ...mock,
          url: `${mock.url}_copy_${new Date().getTime()}`,
        })
        .then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t("p.detail.create.success"));
          }
        });
    },
    getColor(item) {
      const color = {
        get: "blue",
        post: "green",
        delete: "red",
        put: "yellow",
        patch: "yellow",
      };
      return color[item];
    },
    openEditor(mock) {
      if (mock) {
        this.$store.commit("mock/SET_EDITOR_DATA", {
          mock,
          baseUrl: this.baseUrl,
        });
        this.$router.push(`/editor/${this.project._id}/${mock._id}`);
      } else {
        this.$router.push(`/editor/${this.project._id}`);
      }
    },
  },
  components: {
    MockExpand,
    Project,
  },
};
</script>
