<template>
  <transition name="fade">
    <div class="em-container" v-show="pageAnimated">
      <div class="em-new__content">
        <el-form label-position="top" :model="form" ref="formValidate">
          <el-form-item :label="$tc('p.new.form.name', 1)">
            <template #label>
              {{ $tc("p.new.form.name", 1) }}
              <el-tooltip :content="$tc('p.new.form.name', 2)">
                <Icon type="help-circled"></Icon>
              </el-tooltip>
            </template>
            <el-row>
              <el-col :span="7">
                <el-form-item>
                  <el-select v-model="form.groupId" :disabled="isEdit">
                    <el-option
                      v-for="item in groups"
                      :value="item.value"
                      :key="item.value"
                      :label="item.label"
                      > </el-option
                    >
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="1" style="text-align: center">/</el-col>
              <el-col :span="16">
                <el-form-item>
                  <el-input
                    v-model="form.projectName"
                    placeholder="example"
                    ref="projectName"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item :label="$tc('p.new.form.url', 1)">
            <template #label>
              {{ $tc("p.new.form.url", 1) }}
              <el-tooltip :content="$tc('p.new.form.url', 2)">
                <Icon type="help-circled"></Icon>
              </el-tooltip>
            </template>
            <el-input v-model="projectUrl" placeholder="example">
              <template #prepend>
                <span  >/</span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item
            :label="$tc('p.new.form.description', 1)"
            class="em-new__form-hr"
          >
            <el-input
              v-model="form.projectDesc"
              :placeholder="$tc('p.new.form.description', 2)"
            ></el-input>
          </el-form-item>
          <el-form-item :label="$tc('p.new.form.swagger', 0)">
            <template #label>
              {{ $tc("p.new.form.swagger", 0) }}
              <span>({{ $tc("p.new.form.swagger", 1) }})</span>
            </template>
            <el-select v-model="swaggerType" class="em-new__swagger-type">
              <el-option value="URL">URL</el-option>
              <el-option value="Upload">Upload</el-option>
            </el-select>
            <el-input
              v-if="swaggerType === 'URL'"
              v-model="form.projectSwagger"
              placeholder="http://example.com/swagger.json"
            ></el-input>
            <el-upload
              type="drag"
              :headers="uploadHeaders"
              :show-file-list="false"
              :accept="['json', 'yml', 'yaml']"
              :action="uploadAPI"
              :on-success="handleSwaggerUploadSuccess"
              :on-format-error="handleSwaggerUploadError"
              v-if="swaggerType === 'Upload'"
            >
              <div style="padding: 20px 0">
                <Icon
                  type="ios-cloud-upload"
                  size="52"
                  style="color: #3399ff"
                ></Icon>
                <p>JSON / YAML</p>
              </div>
            </el-upload>
            <article class="em-new__form-description">
              {{ $tc("p.new.form.swagger", 2) }}
              <router-link to="/docs#swagger"
                ><Icon type="help-circled"></Icon
              ></router-link>
            </article>
          </el-form-item>
          <el-form-item
            :label="$t('p.new.form.member[0]')"
            class="em-new__form-hr"
          >
            <template #label>
              {{ $t("p.new.form.member[0]") }}
              <span
                >({{
                  isGroup
                    ? $tc("p.new.form.member[1]", 2)
                    : $tc("p.new.form.member[1]", 1)
                }})</span
              >
            </template>
            <el-select
              v-model="form.projectMembers"
              multiple
              filterable
              remote
              :disabled="isGroup"
              :placeholder="$t('p.new.form.member[2]')"
              :remote-method="remote"
              :loading="remoteLoading"
            >
              <el-option
                v-for="option in users"
                :value="option.value"
                :key="option.value"
                >{{ option.label }}</el-option
              >
            </el-select>
          </el-form-item>
          <el-form-item :class="{ 'em-new__form-hr': isEdit }">
            <el-button type="primary" style="width: 100%;" @click="submit">{{
              isEdit
                ? $t("p.new.form.button.update")
                : $t("p.new.form.button.create")
            }}</el-button>
          </el-form-item>
          <el-form-item :label="$tc('p.new.form.confirm', 0)" v-if="isEdit">
            <el-input
              v-model="confirmName"
              :placeholder="$tc('p.new.form.confirm', 1)"
            ></el-input>
            <article class="em-new__form-description">
              {{ $tc("p.new.form.confirm", 2) }}
              <strong style="word-break: break-all">
                {{
                  (projectData.user && projectData.user.nick_name) ||
                  (projectData.group && projectData.group.name)
                }}
                / {{ projectData.name }}
              </strong>
            </article>
          </el-form-item>
          <el-form-item v-if="isEdit">
            <el-button
              type="error"
             style="width: 100%;"
              @click="remove"
              :disabled="confirmName !== projectData.name"
              >{{ $t("p.new.form.button.delete") }}</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </transition>
</template>

<style lang="less" scoped>
@import "index";
</style>

<script>
import * as api from "../../api";
import {nextTick} from "vue";

export default {
  name: "newProject",
  data() {
    return {
      uploadAPI: "/api/upload",
      swaggerType: "URL",
      remoteLoading: false,
      users: [],
      groups: [],
      projectUrl: "",
      confirmName: "",
      form: {
        groupId: "",
        projectId: "",
        projectName: "",
        projectUrl: "",
        projectDesc: "",
        projectSwagger: "",
        projectMembers: [],
      },
    };
  },
  props: {
    projectData: null,
  },
  mounted() {
    const proj = this.projectData;
    this.$nextTick(() => {
      this.$refs.projectName.focus();
    });
    if (proj) {
      this.remoteLoading = true; // 回填文案显示异常，应该是 iview 的 bug
      this.users = proj.members.map((member) => ({
        value: member._id,
        label: member.nick_name,
      }));
      this.form.projectId = proj._id;
      this.form.projectName = proj.name;
      this.form.projectDesc = proj.description;
      this.form.projectSwagger = proj.swagger_url;
      this.projectUrl = proj.url.slice(1); // remove /
       nextTick(() => {
        this.remoteLoading = false;
        this.form.projectMembers = this.users.map((u) => u.value);
      });
      if (proj.group) {
        this.groups = [{ value: proj.group._id, label: proj.group.name }];
        this.form.groupId = proj.group._id;
      } else {
        this.groups = [{ value: proj.user._id, label: proj.user.nick_name }];
        this.form.groupId = proj.user._id;
      }
    } else {
      this.fetchGroup().then((groups) => {
        if (groups.length < 2) {
          this.form.groupId = this.user.id;
        }
      });
    }
  },
  computed: {
    user() {
      return this.$store.state.user;
    },
    isEdit() {
      return !!this.projectData;
    },
    isGroup() {
      if (this.projectData) {
        return !!this.projectData.group;
      } else {
        return this.form.groupId !== this.user.id;
      }
    },
    uploadHeaders() {
      return {
        Authorization: "Bearer " + this.user.token,
      };
    },
  },
  methods: {
    handleSwaggerUploadSuccess(response) {
      const data = response.data;
      this.form.projectSwagger = data.path;
      this.swaggerType = "URL";
      if (data.expire && data.expire !== -1) {
        this.$message.success({
          content: this.$tc("p.new.uploadSuccess", 2, { date: data.expire }),
          duration: 5,
        });
      } else {
        this.$message.success(this.$tc("p.new.uploadSuccess", 1));
      }
    },
    handleSwaggerUploadError() {
      this.$message.error(this.$t("p.new.formatError"));
    },
    convertUrl(url) {
      const newUrl = "/" + url;
      return newUrl === "/"
        ? "/"
        : newUrl.replace(/\/\//g, "/").replace(/\/$/, "");
    },
    fetchGroup() {
      return api.group.getList().then((res) => {
        if (res.data.success) {
          this.groups = [
            { value: this.user.id, label: this.user.nickName },
          ].concat(
            res.data.data.map((o) => ({
              value: o._id,
              label: o.name,
            }))
          );
        }
        return this.groups;
      });
    },
    submit() {
      const data = {
        id: this.form.projectId,
        name: this.form.projectName,
        group: this.form.groupId,
        swagger_url: this.form.projectSwagger,
        description: this.form.projectDesc,
        url: this.convertUrl(this.projectUrl),
        members: this.isGroup ? [] : this.form.projectMembers,
      };

      if (this.isEdit) {
        api.project.update({ data }).then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t("p.new.form.success.update"));
            this.$store.commit("mock/SET_REQUEST_PARAMS", { pageIndex: 1 });
            this.$store.dispatch("mock/FETCH", this.$route);
          }
        });
      } else {
        if (this.form.groupId === "") {
          this.$message.error({
            content: this.$t("p.new.form.error.groupIsNull"),
            duration: 5,
          });
          return;
        }

        if (data.group === this.user.id) {
          data.group = "";
        }

        api.project
          .create({
            data: data,
          })
          .then((res) => {
            if (res.data.success) {
              this.$message.success(this.$t("p.new.form.success.create"));
              if (data.group) {
                const group = this.groups.filter(
                  (item) => item.value === data.group
                )[0];
                this.$router.push(`/group/${group.value}?name=${group.label}`);
              } else {
                this.$router.push("/");
              }
            }
          });
      }
    },
    remote(query) {
      if (query) {
        this.remoteLoading = true;
        api.u
          .getList({
            params: {
              keywords: query,
            },
          })
          .then((res) => {
            this.remoteLoading = false;
            if (res.data.success) {
              const list = res.data.data.map((item) => {
                return {
                  value: item._id,
                  label: item.nick_name,
                  userName: item.name,
                };
              });
              this.users = list.filter((item) => {
                const nickName = item.label.toLowerCase();
                const userName = item.userName.toLowerCase();
                const q = query.toLowerCase();
                return nickName.indexOf(q) > -1 || userName.indexOf(q) > -1;
              });
            }
          });
      } else {
        this.users = [];
      }
    },
    remove() {
      const projectId = this.projectData._id;
      this.$store.dispatch("project/REMOVE", projectId).then(() => {
        this.$message.success(
          this.$t("p.new.form.success.delete", { name: this.projectData.name })
        );
        this.$router.push("/");
      });
    },
  },
};
</script>
