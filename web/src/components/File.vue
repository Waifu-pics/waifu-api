<template>
  <div>
    <v-card>
      <v-img :src="cdnroot + file" class="white--text align-end" gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)" height="250px">
        <v-card-title v-text="file"></v-card-title>
      </v-img>

      <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn icon v-on:click="open()">
          <v-icon>mdi-launch</v-icon>
        </v-btn>

        <v-btn icon v-on:click="verify(false)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>

        <v-btn v-if="!verified" v-on:click="verify(true)" icon>
          <v-icon>mdi-check</v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
import Axios from 'axios'

export default {
  name: 'Filebox',
  props: {
    file: String,
    verified: Boolean,
  },
  data () {
    return {
      apiroot: process.env.VUE_APP_APIROOT,
      cdnroot: process.env.VUE_APP_CDNROOT,
    }
  },
  methods: {
    verify: function (verify) {
      let endpoint = verify ? "verify" : "delete"
      let plural = verify ? "verified" : "deleted"

      Axios({
        method: "post",
        url: `${this.apiroot}/api/admin/${endpoint}`,
        data: {
          file: this.file,
        },
      }).then((res) => {
        this.$notification.success(`${this.file} was successfuly ${plural}!`)

        this.$parent.search()
      }).catch(() => {
        this.$notification.error(`${this.file} could not be ${plural}!`)
      })
    },
    open: function () {
      window.open(this.cdnroot + this.file)
    },
  },
}
</script>
