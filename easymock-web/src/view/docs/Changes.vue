<template>
<div>
  <changelog ref="changeLog"></changelog>
</div>
</template>

<script>
import changelog from './CHANGELOG.md'
export default {
  name: 'Changes',
  components:{
    changelog
  },
  mounted () {
    const lis = this.$refs.changeLog.$el.querySelectorAll('li')
    for (let len = lis.length, i = 0; i < len; i += 1) {
      lis[i].innerHTML = lis[i].innerHTML.replace(/#(\d+)/g, '<a href="https://github.com/easy-mock/easy-mock/issues/$1" target="_blank">#$1</a>')
    }
  }
}
</script>

<style scoped>

</style>
