import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createApp } from "vue";
import { createI18n } from "vue-i18n";
 import router from "./router";
import App from "./App.vue";
import "@/styles/index.less";
import store from "@/store/index";
import enLocale from "@/locale/en";
import zhLocale from "@/locale/zh-CN";
import locale from 'element-plus/lib/locale/lang/zh-cn'
import "dayjs/locale/zh-cn";

// 设置语言


//全局组件
import Add from "@/components/add.vue";
import Spots from "@/components/spots.vue";
import Header from "@/components/header.vue";
import Loading from "@/components/loading.vue";
import Placeholder from "@comp/Placeholder.vue";
import ShapeShifter from "@/components/shape-shifter/index.vue";
import KeyboardShort from "@comp/KeyboardShort.vue";
import AnimatedInteger from "@/components/animated-integer.vue";
import Econ from "@comp/Econ.vue";
import {initAPI} from "@/api";
import clickout from "@/directives/clickout";
import ShortKey from "@/directives/shortkey";


const i18n = createI18n({
  locale: localStorage.getItem("locale") || "zh-CN",
  fallbackLocale: "zh-CN",
  messages: {
    "zh-CN": {
      ...zhLocale,
    },
    en: {
      ...enLocale,
    },
  }, // something vue-i18n options here ...
});

const app = createApp(App);

//指令
app.directive('click-outside',clickout)
app.directive('shortkey',ShortKey)

//全局组件
app.component('em-add', Add);
app.component(Spots.name, Spots);
app.component(Header.name, Header);
app.component(Loading.name, Loading);
app.component(Placeholder.name, Placeholder);
app.component(ShapeShifter.name, ShapeShifter);
app.component(KeyboardShort.name, KeyboardShort);
app.component(AnimatedInteger.name, AnimatedInteger);
app.component('Econ', Econ);

app.mixin({
  data() {
    return {
      pageAnimated: false,
    };
  },
  mounted() {
    this.pageAnimated = true;
  },
});
initAPI(router)
app.use(i18n);
app.use(ElementPlus,{size:'mini',zIndex: 3000,locale});
app.use(router);
app.use(store);
app.mount("#app");
