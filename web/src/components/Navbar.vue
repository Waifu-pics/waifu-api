<template>
  <div>
    <nav role="navigation" aria-label="main navigation">
      <router-link class="home" to="/"><span>HOME</span></router-link>
      <span ref="hidelink" class="hidden">
        <router-link to="/nsfw">NSFW</router-link>
      </span>
      <a class="icon" @click="collapse()">
        <font-awesome-icon :icon="hamburger"/>
      </a>
    </nav>
  </div>
</template>

<script>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faBars } from '@fortawesome/free-solid-svg-icons'

export default {
  computed: {
    hamburger () {
      return faBars
    },
  },

  methods: {
    collapse: function () {
      const x = this.$refs.hidelink

      if (!x.classList.contains("hidden")) {
        x.classList.add("hidden")

        return
      }

      x.classList.remove("hidden")

      return
    },
  },

  components: {
    FontAwesomeIcon,
  },
}
</script>

<style lang="scss" scoped>
nav {
    text-align: center;
    background-color: #303031;
    position: fixed;
    overflow: hidden;
    z-index: 1;
    width: 100%;
    opacity: 0.9;
    backdrop-filter: blur(5px);

    a {
      display: inline-block;
      color: #f2f2f2;
      text-align: center;
      padding: 7px 15px;
      text-decoration: none;
      font-size: 17px;
      z-index: 2;
    }

    a.icon {
      display: none;
    }
}

.router-active:not(.title), .router-active:not(.title):focus {
  background-color: #2896d6;
}

@media only screen and (max-width: 500px) {
  .router-active:not(.title), .router-active:not(.title):focus {
    background-color: #3a3a3a;
    color: white;
  }

  nav {
    .home {
      span {
        display: none;
      }
    }

    .home::after {
      content: "WAIFU.PICS";
    }


    a {
        padding: 13px 16px;
        display: block;
    }

    a.icon {
        background: #2896d6;
        display: block;
        position: absolute;
        right: 0;
        top: 0;
    }

    .hidden {
        display: none;
    }
  }
}
</style>
