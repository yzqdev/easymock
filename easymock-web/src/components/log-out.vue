<template>
  <transition name="fade">
    <p v-show="pageAnimated" style="padding: 20px;">{{$t('c.logOut.text')}}</p>
  </transition>
</template>

<script>
 import conf from '@/config/config'

export default {
  beforeMount () {
    this.logOut()
  },
  methods: {
    logOut () {
     localStorage.removeItem('user')
     localStorage.removeItem(conf.storageNamespace + 'token')
      this.$store.commit('user/SET_VALUE', {})
      this.$router.push('/login')
    }
  }
}
</script>
