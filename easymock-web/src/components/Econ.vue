<template>
  <i
    class="iconfont"
    :class="`el-iconfont-${name}`"
    :style="{ color: color, fontSize: size }"
  ></i>
</template>

<script>
export default {
  name: "EIcon",

  props: {
    name: {
      type: String,
      require: true,
    },
    color: {
      type: String,
    },
    size: {
      type: String,
    },
  },
};
</script>
