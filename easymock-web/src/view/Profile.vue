<template>
  <div class="em-profile">
    <em-header
      icon="edit"
      :title="$t('p.profile.header.title')"
      :description="$t('p.profile.header.description')"
    >
    </em-header>
    <em-keyboard-short></em-keyboard-short>
    <el-dialog
      :title="$t('p.profile.modal.title')"
      v-model="visible"
      width="400"
    >
      <img :src="form.headImg" style="width: 100%" v-show="form.headImg" />
    </el-dialog>
    <transition name="fade">
      <div class="em-container" v-show="pageAnimated">
        <div class="em-profile__content">
          <el-row :gutter="20">
            <el-col span="18">
              <el-form
                label-position="top"
                :model="form"
                :rules="rules"
                ref="form"
              >
                <el-form-item size="large" :label="$t('p.profile.form.language')">
                  <el-select v-model="language">
                    <el-option
                      v-for="item in languageList"
                      :value="item.value"
                      :key="item.value"
                      :label="item.label"
                    >
                    </el-option>
                  </el-select>
                </el-form-item>
                <el-form-item :label="$t('p.profile.form.nickName')">
                  <el-input v-model="form.nickName"></el-input>
                </el-form-item>
                <el-form-item
                  :label="$t('p.profile.form.password')"
                  v-show="!ldap"
                >
                  <el-input type="password" v-model="form.password"></el-input>
                </el-form-item>
                <el-form-item
                  :label="$t('p.profile.form.passwordCheck')"
                  prop="passwordCheck"
                  v-show="!ldap"
                >
                  <el-input
                    type="password"
                    v-model="form.passwordCheck"
                  ></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="update">{{
                    $t("p.profile.form.update")
                  }}</el-button>
                </el-form-item>
              </el-form>
            </el-col>
            <el-col span="6">
              <p>{{ $t("p.profile.form.avatar") }}</p>
              <img
                class="avatar"
                :src="form.headImg"
                v-show="form.headImg"
                :alt="form.nickName"
                :title="form.nickName"
                @click="visible = true"
              />
              <el-upload
                :show-file-list="false"
                :accept="['jpg', 'jpeg', 'png']"
                :on-success="handleSuccess"
                :headers="uploadHeaders"
                :on-format-error="handleFormError"
                :action="uploadAPI"
              >
                <el-button   plain icon="ios-cloud-upload-outline" long>{{
                  $t("p.profile.form.upload")
                }}</el-button>
              </el-upload>
            </el-col>
          </el-row>
        </div>
      </div>
    </transition>
  </div>
</template>

<style lang="less" scoped>
@import "../styles/base/var";

.em-profile {
  .em-profile__content {
    padding: 80px 165px;
    margin-top: 20px;
    width: 100%;
    background: #fff;
    box-shadow: 0 2px 3px #eee;
    border-radius: 4px;
    margin-bottom: 20px;
  }

  .avatar {
    width: 100%;
    border-radius: 4px;
    margin: 5px 0;
    cursor: pointer;
  }

  .ivu-upload {
    width: 100%;
  }
}
</style>

<script>
import config from "@/config/config";
import * as api from "../api";
import languageMap from "../locale/map.json";

export default {
  name: "profile",
  components: {},
  data() {
    const validatePassCheck = (rule, value, callback) => {
      if (value !== this.form.password) {
        callback(new Error(this.$t("p.profile.validateError")));
      } else {
        callback();
      }
    };

    return {
      ldap: config.ldap,
      visible: false,
      language: localStorage.getItem("locale") || "zh-CN",
      languageList: languageMap.list,
      uploadAPI: "/api/upload",
      form: {
        headImg: this.$store.state.user.headImg,
        nickName: this.$store.state.user.nickName,
        password: "",
        passwordCheck: "",
      },
      rules: {
        passwordCheck: [{ trigger: "blur", validator: validatePassCheck }],
      },
    };
  },
  computed: {
    uploadHeaders() {
      return {
        Authorization: "Bearer " + this.$store.state.user.token,
      };
    },
  },
  methods: {
    handleFormError(file) {
      this.$Notice.warning({
        title: this.$tc("p.profile.formatError", 1),
        desc: this.$tc("p.profile.formatError", 2, { name: file.name }),
      });
    },
    handleSuccess(response, file, fileList) {
      this.form.headImg = response.data.path;
    },
    update() {
      const data = {
        nick_name: this.form.nickName,
        head_img: this.form.headImg,
      };

      if (this.form.password) {
        data.password = this.form.password;
      }

      this.$refs.form.validate((valid) => {
        if (valid) {
          api.u.update({ data }).then((res) => {
            if (res.data.success) {
              localStorage.setItem("locale", this.language);
              this.$i18n.locale = this.language;
              this.$message.success({
                title: this.$tc("p.profile.updateSuccess", 1),
                content: this.$tc("p.profile.updateSuccess", 2),
                onOk: () => {
                  this.$router.push("/log-out");
                },
              });
            }
          });
        }
      });
    },
  },
};
</script>
