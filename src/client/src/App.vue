<template>
  <v-app>
    <Navbar></Navbar>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<style lang="scss">
.centered {
    position: relative;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}
</style>

<script>
import Navbar from './components/Navbar.vue'
import store from '@/store/index.js'

export default {
  name: 'App',
  components: {
    Navbar,
  },
  created () {
    this.$vuetify.theme.dark = true
    // Preload
    this.$store.commit('endpoints')
  },
}
</script>
