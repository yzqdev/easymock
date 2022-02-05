<template>
  <div>
    <router-view v-slot="{ Component }"
      ><transition name="fade" mode="out-in">
        <component :is="Component"></component> </transition
    ></router-view>

    <el-dialog
      width="520px"
      v-model="visible"
      title="Choose Language"
      :show-close="false"
      :close-on-press-escape="false"
      close-on-click-modal
    >
      <el-select v-model="language" style="width: 100%">
        <el-option
          v-for="item in languageList"
          :label="item.label"
          :value="item.value"
          :key="item.value"
        >
        </el-option>
      </el-select>
      <template #footer
        ><el-button type="primary" style="width: 100%" @click="settingLanguage"
          >OK</el-button
        ></template
      >
    </el-dialog>
  </div>
</template>

<script>
import languageMap from "@/locale/map";

export default {
  data() {
    return {
      visible: false,
      language: this.locale || "zh-CN",
      languageList: languageMap.list,
    };
  },
  computed: {
    locale() {
      return localStorage.getItem("locale");
    },
    appVersion() {
      return this.$store.state.app.version;
    },
  },
  mounted() {
    if (this.appVersion === localStorage.getItem("version")) {
      this.$store.commit("app/SET_READ_CHANGELOG", true);
    }
    if (localStorage.getItem("easy-mock_token")) {
      this.$store.commit(
        "user/SET_VALUE",
        JSON.parse(localStorage.getItem("user"))
      );
    }
    if (this.locale) return;
    this.visible = true;
  },
  methods: {
    settingLanguage() {
      localStorage.setItem("locale", this.language);
      this.$i18n.locale = this.language;
      this.visible = false;
    },
  },
};
</script>
