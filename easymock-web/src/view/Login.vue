<template>
  <div class="em-index">
    <transition name="zoom">
      <div class="em-index__login" v-if="page === 0">
        <img src="/images/easy-mock.png" />
        <p>{{ $tc("p.login.description", 1) }}</p>
        <p>{{ $tc("p.login.description", 2) }}</p>
        <transition name="fadeUp" mode="out-in">
          <el-button
            type="primary"
            style="width: 100%"
            @click.stop="start"
            v-if="!isLogin"
            key="start"
            >{{ $tc("p.login.form.button", 1) }}</el-button
          >
          <el-button
            type="success"
            style="width: 100%"
            @click.stop="login"
            v-else
            key="login"
            >{{ $tc("p.login.form.button", 2) }}</el-button
          >
        </transition>
        <transition name="fadeLeft">
          <div v-show="isLogin" v-click-outside="onClickOutside">
            <el-input
              size="large"
              v-if="ldap"
              :placeholder="$tc('p.login.form.placeholder', 2)"
              ref="user"
              v-model="userName"
              @on-enter="login"
            ></el-input>
            <el-input
              size="large"
              v-if="!ldap"
              :placeholder="$tc('p.login.form.placeholder', 1)"
              ref="user"
              v-model="userName"
              @on-enter="login"
            ></el-input>
            <el-input
              size="large"
              :placeholder="$t('p.login.form.password')"
              type="password"
              v-model="password"
              @on-enter="login"
            ></el-input>
          </div>
        </transition>
      </div>
    </transition>

    <div
      class="em-index__section em-index__section--login"
      style="z-index: 6"
      :class="{ 'is-old': page > 0 }"
    >
      <transition name="fade">
        <div
          class="fullscreen"
          :class="{ 'is-login': isLogin }"
          ref="wallpaper"
          v-show="wallpaperVisible"
        ></div>
      </transition>
      <div class="links">
        <router-link to="/docs" class="link">Document</router-link>
        <a
          href="https://github.com/easy-mock/easy-mock-cli"
          target="_blank"
          class="link"
          >CLI</a
        >
        <a
          href="https://github.com/easy-mock/easy-mock"
          target="_blank"
          class="link"
          >GitHub</a
        >
        <p v-if="copyright">{{ copyright }}</p>
      </div>
      <transition name="fade">
        <div class="fullscreen-by" v-if="wallpaperCopyright">
          <div v-if="wallpaperCopyright.name === 'Bing'">
            Photo by
            <a :href="wallpaperCopyright.link" target="_blank">
              <strong>{{ wallpaperCopyright.name }}</strong>
            </a>
          </div>
          <div v-else>
            Photo by
            <a :href="wallpaperCopyright.link" target="_blank">
              <strong>{{ wallpaperCopyright.name }}</strong>
            </a>
            <strong> / </strong>
            <a href="https://unsplash.com" target="_blank">
              <strong>Unsplash</strong>
            </a>
            <a :href="wallpaperCopyright.link" target="_blank" class="avatar">
              <img :src="wallpaperCopyright.profileImage" />
            </a>
          </div>
        </div>
      </transition>
      <div class="about-btn" @click="page = 1">
        {{ $tc("p.login.about", 1) }}
      </div>
    </div>

    <div
      class="em-index__section section-about"
      style="z-index: 5"
      :class="{ 'is-old': page > 1 }"
    >
      <!--      <em-shape-shifter v-if="page === 1"></em-shape-shifter>-->
      <div class="feature-list">
        <transition-group name="fadeDown">
          <div class="section-title" key="a" v-show="featureVisible">
            Easy Mock
          </div>
          <div class="section-description" key="b" v-show="featureVisible">
            {{ $tc("p.login.about", 2) }}
          </div>
        </transition-group>
        <el-row :gutter="100">
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <img
                    src="/images/icon-swagger.png"
                    style="margin-left: 1px"
                  />
                </div>
                <h2>{{ $tc("p.login.feature[0]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[0]", 2) }}</p>
              </div>
            </transition>
          </el-col>
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <img src="/images/icon-mock.png" style="margin-left: 1px" />
                </div>
                <h2>{{ $tc("p.login.feature[1]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[1]", 2) }}</p>
              </div>
            </transition>
          </el-col>
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <img src="/images/icon-command.png" />
                </div>
                <h2>{{ $tc("p.login.feature[2]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[2]", 2) }}</p>
              </div>
            </transition>
          </el-col>
        </el-row>
        <el-row :gutter="100">
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <i type="ios-book"></i>
                </div>
                <h2>{{ $tc("p.login.feature[3]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[3]", 2) }}</p>
              </div>
            </transition>
          </el-col>
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <i type="ribbon-b"></i>
                </div>
                <h2>{{ $tc("p.login.feature[4]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[4]", 2) }}</p>
              </div>
            </transition>
          </el-col>
          <el-col span="8">
            <transition name="zoom">
              <div v-show="featureVisible">
                <div class="feature-icon">
                  <i type="lightbulb"></i>
                </div>
                <h2>{{ $tc("p.login.feature[5]", 1) }}</h2>
                <p>{{ $tc("p.login.feature[5]", 2) }}</p>
              </div>
            </transition>
          </el-col>
        </el-row>
      </div>
    </div>

    <div class="em-index__pagination">
      <div class="dot" :class="{ active: page === 0 }" @click="page = 0"></div>
      <div class="dot" :class="{ active: page === 1 }" @click="page = 1"></div>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "../styles/base/var";

.em-index {
  color: #fff;

  .em-index__pagination {
    left: 40px;
    position: absolute;
    top: 50%;
    z-index: 2147483643;
    transform: translateY(-50%);

    .dot {
      border-radius: 50%;
      box-shadow: inset 0 0 0 3px #fff;
      cursor: pointer;
      height: 16px;
      margin-bottom: 10px;
      opacity: 0.4;
      position: relative;
      width: 16px;
      box-sizing: border-box;
    }

    .dot.active {
      background: #fff;
      box-shadow: none;
    }

    .dot:hover,
    .dot.active {
      opacity: 1;
    }
  }

  .em-index__login {
    display: block;
    left: 50%;
    margin-top: -165px;
    margin-left: -150px;
    position: absolute;
    top: 50%;
    width: 300px;
    box-sizing: border-box;
    z-index: 7;
    text-align: center;
    font-size: 17px;
    font-weight: 700;

    img {
      width: 230px;
      margin-bottom: 15px;
    }

    button {
      height: 45px;
      margin: 30px 0 20px 0;
    }

    input {
      margin-bottom: 10px;
      font-weight: 400;
    }

    input,
    input:hover,
    input:focus {
      border-color: #fff;
      box-shadow: none;
    }
  }

  .em-index__section {
    background: #56bc8a;
    color: #fff;
    height: 100%;
    left: 0;
    overflow: hidden;
    position: absolute;
    top: 0;
    text-align: center;
    width: 100%;
    transform: translateY(0);
    transition: transform 0.7s cubic-bezier(0.825, 0, 0.5, 1);

    .is-old {
      transform: translateY(-100%);
    }

    &--login {
      a {
        color: #fff;
      }

      .fullscreen {
        position: fixed;
        top: 0;
        right: 0;
        left: 0;
        bottom: 45px;

        &:before {
          content: "";
          position: fixed;
          top: 0;
          right: 0;
          bottom: 0;
          left: 0;
          opacity: 0.6;
          background-color: rgba(0, 0, 0, 0.3);
          transition: all 0.5s;
        }

        &.is-login:before {
          background-color: rgba(0, 0, 0, 1);
        }
      }

      .about-btn {
        background: #56bc8a;
        box-shadow: inset 0 3px 0 rgba(0, 0, 0, 0.1);
        bottom: -45px;
        color: #fff;
        cursor: pointer;
        height: 45px;
        font-weight: 700;
        left: 0;
        line-height: 45px;
        position: absolute;
        text-align: center;
        width: 100%;
        z-index: 2;
        transition: bottom 0.3s ease-out 1.7s;
        bottom: 0;
      }

      .links,
      .fullscreen-by {
        position: absolute;
        bottom: 45px;
        padding: 0 12px;
        margin: 0 0 13px 0;
        line-height: 25px;
        z-index: 90210;
        transition: opacity 0.2s linear;
        text-align: left;
      }

      .fullscreen-by {
        right: 0;

        .avatar {
          display: inline-block;
          border-radius: 2px;
          vertical-align: -8px;
          margin: 0;
          padding: 0;
          height: 24px;
          width: 24px;
          background-size: cover;
          background-position: 50% 50%;
          background-size: 24px;
          box-shadow: 0 1px 1px rgba(0, 0, 0, 0.3);
          margin-left: 10px;
        }

        img {
          display: block;
          width: 24px;
          height: 24px;
        }
      }

      .link {
        display: inline-block;
        color: #fff;
        text-decoration: none;
        font-size: 14px;
        padding-right: 5px;
        padding-left: 5px;
        margin-left: -5px;
        margin-right: 12px;
        font-weight: normal;
      }
    }
  }

  .section-about {
    text-align: left;

    .section-title {
      font-size: 38px;
      font-weight: 700;
    }

    .section-description {
      font-size: 18px;
      margin-bottom: 100px;
    }

    .feature-list {
      width: 70%;
      height: 570px;
      position: absolute;
      top: 50%;
      left: 50%;
      margin-top: -285px;
      margin-left: -35%;
      font-size: 13px;
    }

    .ivu-col {
      min-height: 200px;
      margin-bottom: 20px;
    }

    .feature-icon {
      width: 60px;
      height: 60px;
      border-radius: 50%;
      background: #fff;
      margin: auto;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 10px;
      font-size: 22px;
      color: @colorPrimary;
      box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);

      img {
        width: 25px;
      }

      i {
        font-size: 32px;
      }
    }

    h2 {
      font-size: 22px;
      font-weight: 700;
      margin-bottom: 10px;
      text-align: center;
    }
  }
}
</style>

<script>
import config from "@/config/config";
import * as api from "../api";
import { nextTick } from "vue";
let resizeTimer;

export default {
  name: "index",
  components: {},
  data() {
    return {
      ldap: config.ldap,
      isLogin: false,
      page: 0,
      userName: localStorage.getItem("last-user"),
      password: "",
      copyright: config.copyright,
      featureVisible: false,
      wallpaperVisible: false,
    };
  },

  mounted() {
    this.$store.dispatch("wallpaper/FETCH");
    console.log(this.$store.state);
    console.log(`%c哦草`, `color:red;font-size:16px;background:transparent`);
    const img = new Image();
    img.src = this.$store.state.wallpaper.url;
    img.onload = () => {
      this.wallpaperVisible = true;
      nextTick(() => {
        this.$refs.wallpaper.style.background = `url(${img.src})`;
        this.$refs.wallpaper.style.backgroundSize = "cover";
        this.$refs.wallpaper.style.backgroundPosition = "50% 50%";
      });
    };
  },
  computed: {
    wallpaperCopyright() {
      return this.$store.state.wallpaper.copyright;
    },
  },
  watch: {
    page: function (current) {
      clearTimeout(resizeTimer);
      if (current === 1) {
        resizeTimer = setTimeout(() => {
          this.featureVisible = true;
        }, 3000);
      } else {
        this.featureVisible = false;
      }
    },
  },
  methods: {
    onClickOutside() {
      if (!this.userName && !this.password) {
        this.isLogin = false;
      }
    },
    start() {
      this.isLogin = true;
      this.$nextTick(() => {
        this.$refs.user.focus();
      });
    },
    login() {
      api.u
        .login({
          messageUnless: ["用户不存在"],
          data: {
            name: this.userName,
            password: this.password,
          },
        })
        .then((res) => {
          const body = res.data;
          if (body.success) {
            this.$store.commit("user/SET_VALUE", body.data);
            localStorage.setItem("user", JSON.stringify(body.data));
            localStorage.setItem("last-user", this.userName);
            localStorage.setItem(
              config.storageNamespace + "token",
              body.data.token
            );
            this.$router.push("/");
          }
        })
        .catch((res) => {
          if (res.data.message === "用户不存在") {
            this.$confirm(
              this.$t("p.login.confirm.register.content"),
              this.$t("confirm.title"),
                {type:'warning'}
            ).then(() => {
              this.register();
            }).catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              });
            });
          }
        });
    },
    register() {
      api.u
        .register({
          data: {
            name: this.userName,
            password: this.password,
          },
        })
        .then((res) => {
          if (res.data.success) {
            this.$message.success(this.$t("p.login.confirm.register.success"));
            this.login();
          }
        });
    },
  },
};
</script>
