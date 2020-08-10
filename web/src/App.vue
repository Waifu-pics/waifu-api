<template>
  <div id="app">
    <Navbar></Navbar>
    <router-view/>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'

export default {
  name: 'App',
  components: {
    Navbar,
  },
}
</script>

<style lang="scss">
html {
  --background: #1e1c22;
  --foreground: #27242d;
  --accent: #2896d6;

  background-color: var(--background);
  scrollbar-width: thin;
  scrollbar-color: var(--accent) var(--background);
  font-family: Manrope;
  color: white;

  -webkit-tap-highlight-color: transparent;
}

::-webkit-scrollbar {
  width: 7px;
  height: 7px;
  background-color: transparent;
}
  
::-webkit-scrollbar-thumb {
  background: var(--accent);
}

a {
  text-decoration: none;
}

* { 
  margin: 0 !important; 
}
</style>
