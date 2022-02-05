<template>
  <div class="em-dashboard">
    <em-header
      :spots="6"
      icon="ios-speedometer"
      :title="$t('p.dashboard.header.title')"
      :description="$t('p.dashboard.header.description')"
    >
    </em-header>
    <em-keyboard-short></em-keyboard-short>
    <div class="em-container em-dashboard__content">
      <el-row :gutter="20">
        <el-col :span="12">
          <transition name="fadeLeft">
            <div
              class="em-dashboard__item em-dashboard__item--key"
              v-show="pageAnimated"
            >
              <h2>
                <Icon type="stats-bars"></Icon>
                {{ $tc("p.dashboard.total.mockUse", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="total.mockUseCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.total.mockUse", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
        <el-col :span="6">
          <transition name="fadeRight">
            <div class="em-dashboard__item" v-show="pageAnimated">
              <h2>
                <Icon type="cube"></Icon>
                {{ $tc("p.dashboard.total.project", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="total.projectCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.total.project", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
        <el-col :span="6">
          <transition name="fadeRight">
            <div class="em-dashboard__item" v-show="pageAnimated">
              <h2>
                <Icon type="link"></Icon> {{ $tc("p.dashboard.total.mock", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="total.mockCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.total.mock", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <transition name="fadeLeft">
            <div
              class="em-dashboard__item em-dashboard__item--key"
              v-show="pageAnimated"
            >
              <em-spots :size="6"></em-spots>
              <h2>
                <Icon type="person"></Icon>
                {{ $tc("p.dashboard.total.user", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="total.userCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.total.user", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
        <el-col :span="12">
          <transition name="fadeRight">
            <div class="em-dashboard__item" v-show="pageAnimated">
              <em-spots :size="6"></em-spots>
              <h2>
                <Icon type="person-add"></Icon>
                {{ $tc("p.dashboard.today.user", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="today.userCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.today.user", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <transition name="fadeLeft">
            <div
              class="em-dashboard__item em-dashboard__item--key"
              v-show="pageAnimated"
            >
              <h2>
                <Icon type="stats-bars"></Icon>
                {{ $tc("p.dashboard.today.mockUse", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="today.mockUseCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.today.mockUse", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
        <el-col :span="6">
          <transition name="fadeRight">
            <div class="em-dashboard__item" v-show="pageAnimated">
              <h2>
                <Icon type="cube"></Icon>
                {{ $tc("p.dashboard.today.project", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="today.projectCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.today.project", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
        <el-col :span="6">
          <transition name="fadeRight">
            <div class="em-dashboard__item" v-show="pageAnimated">
              <h2>
                <Icon type="link"></Icon> {{ $tc("p.dashboard.today.mock", 1) }}
              </h2>
              <p class="number">
                <em-animated-integer
                  :value="today.mockCount"
                ></em-animated-integer>
                <span>{{ $tc("p.dashboard.today.mock", 2) }}</span>
              </p>
            </div>
          </transition>
        </el-col>
      </el-row>
      <transition name="fadeUp">
        <div class="em-dashboard__users" v-show="pageAnimated">
          <el-row>
            <el-col :span="6">
              <em-spots :size="6"></em-spots>
              <div class="em-dashboard__users__title">
                <Icon type="quote"></Icon>
              </div>
            </el-col>
            <el-col :span="18">
              <el-row :gutter="10" style="padding: 0 10px">
                <el-col span="2" v-for="(item, i) in users" :key="i">
                  <img :src="item.head_img" :title="item.nick_name" />
                </el-col>
              </el-row>
            </el-col>
          </el-row>
        </div>
      </transition>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "../styles/base/var";

.em-dashboard {
  .em-dashboard__content {
    margin: 20px auto;
  }

  .em-dashboard__item,
  .em-dashboard__users {
    width: 100%;
    background: #fff;
    box-shadow: 0 2px 3px #eee;
    border-radius: 4px;
    margin-bottom: 20px;
  }

  .em-dashboard__item--key {
    background-image: linear-gradient(
      45deg,
      @colorPrimary,
      #2d8cf099
    );
    color: #fff;
    box-shadow: 0 2px 3px #2d8cf066;
  }

  .em-dashboard__item {
    padding: 20px;
    position: relative;
    h2 {
      font-weight: 500;
      font-size: 16px;

      i {
        font-size: 26px;
        vertical-align: bottom;
        margin-right: 6px;
      }
    }

    .number {
      font-size: 32px;

      span:last-child {
        font-size: 22px;
        margin-left: 10px;
      }
    }
  }

  .em-dashboard__users {
    position: relative;

    .em-dashboard__title {
      height: 190px;
      border-top-left-radius: 4px;
      border-bottom-left-radius: 4px;
      color: #495060;
      font-size: 36px;
      text-align: center;
      line-height: 190px;
    }

    img {
      display: block;
      width: 50px;
      height: 50px;
      border-radius: 4px;
      margin-top: 10px;
    }
  }
}
</style>

<script>
export default {
  name: "dashboard",
  computed: {
    total() {
      return this.$store.state.dashboard.total;
    },
    today() {
      return this.$store.state.dashboard.today;
    },
    users() {
      return this.$store.state.dashboard.users;
    },
  },
  created() {
    return this.$store.dispatch("dashboard/FETCH");
  },
};
</script>
