<template>
  <div>
    <DocsEn v-if="visible"></DocsEn>
    <DocsZh v-if="!visible"></DocsZh>
  </div>
</template>

<script>
import docsEn from "./docs.md";
import docsZh from "./docs.zh-CN.md";
export default {
  name: "myDoc",
  components: {
    docsEn,
    docsZh,
  },
  data() {
    return {
      visible: true,
    };
  },
  mounted() {
    const locale = localStorage.getItem("locale") || "zh-CN";
    this.visible = locale === "en";
  },
};
</script>

<style scoped></style>
