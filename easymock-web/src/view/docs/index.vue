<template>
  <changes ref="doc"></changes>

</template>

<script>
//todo 这里要改
import Changes from "./Changes.vue";
import myDoc from "./doc.vue";
import { nextTick } from "vue";
export default {
  name: "document",
  components: {
    Changes,
    myDoc
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
    console.log(doc);
    console.log(
      `%c温迪来到`,
      `color:red;font-size:16px;background:transparent`
    );
    const docNodes = doc.$el

  },

  methods: {}
};
</script>

<style scoped></style>
