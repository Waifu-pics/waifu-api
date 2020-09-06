<template>
  <div>
    <v-card>
      <v-img :src="file.url" class="white--text align-end" gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)" height="250px">
        <v-card-title v-text="file.name"></v-card-title>
      </v-img>

      <v-card-actions>
        <v-checkbox style="margin-bottom: -20px; margin-top: -5px; margin-left: 5px;" v-model="$parent._data.deletelist" :value="file.name"></v-checkbox>

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
    file: Object,
    verified: Boolean,
  },
  data: () => ({
  }),
  methods: {
    verify: function (verify) {
      let endpoint = verify ? "verify" : "delete"
      let plural = verify ? "verified" : "deleted"

      Axios({
        method: "post",
        url: `/api/admin/${endpoint}`,
        data: {
          files: [
            this.file,
          ],
        },
      }).then((res) => {
        this.$notification.success(`${this.file.name} was successfuly ${plural}!`)

        this.$parent.search()
      }).catch(() => {
        this.$notification.error(`${this.file.name} could not be ${plural}!`)
      })
    },
    open: function () {
      window.open(this.file)
    },
  },
}
</script>
