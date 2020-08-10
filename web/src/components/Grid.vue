<template>
  <div id="photos">
    <div v-for="image in images" v-bind:key="image">
      <img :src="'https://i.waifu.pics/' + image">
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
    }
  },
  mounted: function() {
    this.getImages()
  },
  methods: {
    getImages: function () {
      Axios({
        method: "post",
        url: "https://waifu.pics/api/many/sfw",
        data: {
          exclude: this.exclude,
        },
      }).then((response) => {
        response.data.data.map((file) => {
          this.exclude.push(file)
        })
        this.images = response.data.data
      })
    },
  },
}
</script>

<style lang="scss">
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
