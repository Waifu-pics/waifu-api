<template>
  <v-container fill-height fluid style="text-align:center;">
    <v-row align="center" justify="center">
      <v-col>
        <v-card :loading="card.loading" class="mx-auto" max-width="500" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">Meme Generator</v-list-item-title>
              <i><p class="desc">Waifu meme generator. Why does this exist you ask? Good question.</p></i>
              <v-select v-model="endpoint" dense outlined label="Endpoint" :items="endpoints"/>
              <v-checkbox label="NSFW" v-model="nsfw" @change="update" style="margin-left: 10px;"/>
              <div>
                <v-text-field v-model="text.top" outlined dense label="Top text"/>
                <v-text-field v-model="text.bottom" outlined dense label="Bottom text"/>
              </div>
              <v-btn v-on:click="generate()" :disabled="card.loading" depressed>Generate</v-btn>
            </v-list-item-content>
          </v-list-item>
          <v-expansion-panels v-model="open" multiple>
            <v-expansion-panel>
              <v-expansion-panel-header>Image</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-list-item class="bottom">
                  <div style="text-align: center;" v-if="buffer != ''">
                    <img style="max-width:600px; max-height:600px;" :src="buffer"/>
                  </div>
                  <div v-else>
                    <p>No image has been generated yet</p>
                  </div>
                </v-list-item>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { api } from '@/functions/api.js'

export default {
  data () {
    return {
      // This is an empty image
      open: [0],
      buffer: "",
      card: {
        loading: false,
      },
      endpoints: [],
      endpoint: "waifu",
      nsfw: false,
      text: {
        top: "",
        bottom: "",
      },
    }
  },
  mounted: function () {
    this.update()
  },
  methods: {
    update: function () {
      api.getEndpoints((res, err) => {
        this.endpoints = (this.nsfw ? res.nsfw : res.sfw)
        this.endpoint = this.endpoints[0]
      })
    },
    generate: function () {
      this.card.loading = true
      api.generateImage(this.endpoint, this.nsfw, this.text, (res, err) => {
        if (err) {
          this.card.loading = false

          return this.$notification.error(err)
        }

        let bytes = new Uint8Array(res)
        let binary = bytes.reduce((data, b) => data += String.fromCharCode(b), '')
        this.buffer = "data:image/jpeg;base64," + btoa(binary)

        this.card.loading = false
      })
    },
  },
}
</script>

<style lang="scss" scoped>
img {
  display: block;
  width: auto;
  height: 400px;
  max-width: 100%;
  max-height: 50%;
  margin: 20px auto;
}
.centered {
  max-width: 500px;
  max-height: 500px;
  text-align: center;
}
.desc {
  margin-top: 5px;
  margin-bottom: -5px;
}
.v-select {
  margin-top: 20px;
}
</style>
