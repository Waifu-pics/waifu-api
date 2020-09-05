<template>
  <div v-if="is404" class="centered">
    <img src="@/assets/404.png">
    <h1>Are you lost?</h1>
    <v-btn outlined :to="'/'" text>Go Home!</v-btn>
  </div>
  <div v-else>
    <v-btn fab large dark fixed bottom right v-on:click="getImages(false)">
      <v-icon>mdi-refresh</v-icon>
    </v-btn>
    <div id="photos">
      <div v-for="image in images" v-bind:key="image">
        <img :src="image">
      </div>
    </div>
  </div>
</template>

<script>
import Axios from 'axios'

export default {
  data: function () {
    return {
      exclude: [],
      images: [],
      is404: false,
    }
  },
  watch: {
    '$route.params.endpoint' () {
      this.is404 = false
      this.exclude = []
      this.getImages(true)
    },
    '$route.params.type' () {
      this.is404 = false
      this.exclude = []
      this.getImages(true)
    },
  },
  mounted: function () {
    this.getImages(true)
  },
  methods: {
    getImages: function (first) {
      const { type, endpoint } = this.$route.params

      Axios({
        method: "post",
        url: `/api/many/${type === undefined ? 'sfw' : type}/${endpoint === undefined ? 'waifu' : endpoint}`,
        data: {
          exclude: this.exclude,
        },
      }).then((response) => {
        response.data.files.map((file) => {
          this.exclude.push(file)
        })
        this.images = response.data.files
      }).catch((response) => {
        if (first) {
          this.is404 = true
        } else {
          this.exclude = []
          this.getImages(false)
        }
      })
    },
  },
}
</script>

<style lang="scss" scoped>
.centered {
  text-align: center;
  
  h1 {
    margin-bottom: 5px;
  }
}
#photos {
    line-height: 0;

    -webkit-column-count: 5;
    -webkit-column-gap:   0px;
    -moz-column-count:    5;
    -moz-column-gap:      0px;
    column-count:         5;
    column-gap:           0px;

    img {
      width: 100% !important;
      height: auto !important;
    }

    @media (max-width: 1200px) {
      -moz-column-count:    4;
      -webkit-column-count: 4;
      column-count:         4;
    }
    @media (max-width: 1000px) {
      -moz-column-count:    3;
      -webkit-column-count: 3;
      column-count:         3;
    }
    @media (max-width: 800px) {
      -moz-column-count:    2;
      -webkit-column-count: 2;
      column-count:         2;
    }
    @media (max-width: 400px) {
      -moz-column-count:    1;
      -webkit-column-count: 1;
      column-count:         1;
    }
}
</style>
