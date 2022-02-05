<template>
  <div class="em-group">
    <em-add
      icon="person-add"
      color="red"
      :bottom="90"
      @click.native="openModal"
    ></em-add>
    <div v-shortkey="['ctrl', 'c']" @shortkey="openModal"></div>
    <em-keyboard-short v-model="keyboards"></em-keyboard-short>
    <el-dialog
      class-name="em-group-modal"
      v-model="modalShow"
      title="创建团队"
      @on-ok="submit"
      :closable="false"
    >
      <el-tabs v-model="tabName">
        <el-tab-pane
          :label="$tc('p.group.modal.tab.create', 0)"
          name="create"
          :disabled="tabName === 'rename'"
        >
          <el-form :label-width="64" @submit.native.prevent>
            <el-form-item :label="$tc('p.group.modal.tab.create', 1)">
              <el-input
                v-model="groupName"
                :placeholder="$tc('p.group.modal.tab.create', 2)"
                @on-enter="submit"
                ref="inputCreate"
              ></el-input>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane
          :label="$tc('p.group.modal.tab.join', 0)"
          name="join"
          :disabled="tabName === 'rename'"
        >
          <el-form :label-width="64" @submit.native.prevent>
            <el-form-item :label="$tc('p.group.modal.tab.join', 1)">
              <el-input
                v-model="groupName"
                @on-enter="submit"
                :placeholder="$tc('p.group.modal.tab.join', 2)"
              ></el-input>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane
          :label="$tc('p.group.modal.tab.edit', 0)"
          name="rename"
          :disabled="tabName !== 'rename'"
        >
          <el-form :label-width="64" @submit.native.prevent>
            <el-form-item :label="$tc('p.group.modal.tab.edit', 1)">
              <el-input
                v-model="groupName"
                @on-enter="submit"
                :placeholder="$tc('p.group.modal.tab.edit', 2)"
              ></el-input>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
      <template #footer>
    <span class="dialog-footer">
      <el-button size="small" @click="modalShow = false">取 消</el-button>
      <el-button size="small" type="primary" @click="modalShow = false">确 定</el-button>
    </span>
      </template>
    </el-dialog>
    <em-placeholder :show="groups.length === 0">
      <Icon :type="keywords ? 'outlet' : 'happy-outline'"></Icon>
      <p>
        {{
          keywords
            ? $tc("p.group.placeholder", 1)
            : $tc("p.group.placeholder", 2)
        }}
      </p>
    </em-placeholder>
    <em-header
      icon="person-stalker"
      :title="$t('p.group.header.title')"
      :description="$t('p.group.header.description')"
    >
    </em-header>
    <transition name="fade">
      <div class="em-container em-group__list" v-show="pageAnimated">
        <div class="ivu-row">
          <transition-group name="fadeUp">
            <div
              class="ivu-col ivu-col-span-6"
              v-for="(item, index) in groups"
              :key="index"
            >
              <div
                class="em-group__item"
                @click="$router.push(`/group/${item._id}?name=${item.name}`)"
              >
                <h2>{{ item.name }}</h2>
                <el-button-group class="group-control">
                  <el-button

                    icon="edit"
                    @click.stop="rename(item)"
                  ></el-button>
                  <el-button

                    icon="trash-b"
                    @click.stop="remove(item)"
                  ></el-button>
                </el-button-group>
              </div>
            </div>
          </transition-group>
        </div>
      </div>
    </transition>
  </div>
</template>

<style lang="less" scoped>
@import "../styles/base/var";

.em-group {
  .em-group__list {
    margin-top: @basePadding;

    .el-row {
      margin-left: -10px;
      margin-right: -10px;
    }

    .el-col {
      padding-left: 10px;
      padding-right: 10px;
    }
  }

  .em-group__item {
    padding: 20px 14px;
    box-shadow: 0 1px 5px #e5e5e5;
    border-radius: 4px;
    transition: var(--all-transition);
    position: relative;
    margin-bottom: var(--base-padding);
    cursor: pointer;
    background: @colorRed;

    h2,
    .group-collect,
    i {
      color: #fff;
    }

    h2 {
      font-weight: 700;
      font-size: 14px;
      .utils-ellipsis();
      .utils-user-select();
    }

    &:hover {
      background: @colorRed+ "E6";
    }

    .ivu-btn {
      border-color: #fff;
    }

    .ivu-btn:hover {
      background: #fff;
      i {
        color: @colorRed;
      }
    }

    .group-control {
      margin-top: 10px;
      width: 100%;

      .ivu-btn {
        width: 50%;
      }
    }
  }
}

.em-group-modal {
  .ivu-form-item {
    margin-bottom: 0;
  }

  .ivu-modal-footer {
    border-top: none;
  }
}
</style>

<script>
import debounce from "lodash/debounce";

export default {
  name: "group",
  data() {
    return {
      groupName: "",
      renameGroup: null,
      modalShow: false,
      tabName: "create",
      keywords: "",
      keyboards: [
        {
          category: this.$t("p.group.keyboards[0].category"),
          list: [
            {
              description: this.$tc("p.group.keyboards[0].list", 0),
              shorts: ["ctrl", "c"],
            },
          ],
        },
      ],
    };
  },
  asyncData({ store }) {
    return store.dispatch("group/FETCH");
  },
  mounted() {
    // this.$on(
    //   "query",
    //   debounce((keywords) => {
    //     this.keywords = keywords;
    //   }, 500)
    // );
  },
  computed: {
    groups() {
      const list = this.$store.state.group.list;
      const keywords = this.keywords;
      return keywords
        ? list.filter((item) => new RegExp(keywords, "i").test(item.name))
        : list;
    },
  },
  methods: {
    openModal() {
      this.tabName = "create";
      this.groupName = "";
      this.modalShow = true;
      this.$nextTick(() => {
        this.$refs.inputCreate.focus();
      });
    },
    submit() {
      this.modalShow = false;
      if (this.tabName === "create") {
        this.$store.dispatch("group/ADD", this.groupName).then((body) => {
          if (body.success)
            this.$message.success(this.$t("p.group.create.success"));
        });
      } else if (this.tabName === "join") {
        this.$store.dispatch("group/JOIN", this.groupName).then((body) => {
          if (body.success) {
            this.$message.success(
              this.$t("p.group.join.success", { groupName: this.groupName })
            );
          } else {
            this.$message.warning(
              this.$t("p.group.join.warning", { groupName: this.groupName })
            );
          }
        });
      } else {
        this.$store
          .dispatch("group/RENAME", {
            id: this.renameGroup._id,
            name: this.groupName,
          })
          .then((body) => {
            if (body.success)
              this.$message.success(this.$t("p.group.update.success"));
          });
      }
    },
    remove(group) {
      this.$confirm({
        title: this.$t("confirm.title"),
        content: this.$t("p.group.confirm.delete.content", {
          name: group.name,
        }),
        onOk: () => {
          this.$store.dispatch("group/REMOVE", group._id).then((body) => {
            if (body.success) {
              this.$message.success(this.$t("p.group.remove.success"));
            }
          });
        },
      });
    },
    rename(group) {
      this.tabName = "rename";
      this.modalShow = true;
      this.groupName = group.name;
      this.renameGroup = group;
    },
  },
};
</script>
