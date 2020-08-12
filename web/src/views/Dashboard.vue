<template>
  <div v-if="loggedin">
    <v-container>
      <v-card class="mx-auto" style="margin-top: 20px; margin-bottom: 10px;" max-width="400" outlined>
        <v-list-item three-line>
          <v-list-item-content>
            <v-text-field outlined dense v-model="query" label="Search"/>
            <v-select v-model="endpoint" dense outlined label="Endpoint" :items="endpoints" @change="changePoint"/>
            <v-spacer/>
            <v-row justify="space-around">
              <v-checkbox label="Verified" class="vercheck" v-model="verbtn"/>
            </v-row>
            <v-btn v-on:click="search()" depressed>Search</v-btn>
          </v-list-item-content>
        </v-list-item>
      </v-card>
    </v-container>
    <div>
      <v-container fluid>
        <v-row dense justify="center">
          <v-col v-for="file in res" :key="file" style="max-width: 400px;">
            <Filebox :file="file.file" :verified="file.verified"></Filebox>
          </v-col>
        </v-row>
      </v-container>
    </div>
    <v-btn fab large dark fixed bottom right v-on:click="getImages(false)">
      <v-icon>mdi-logout-variant</v-icon>
    </v-btn>
  </div>
</template>

<script>
import Axios from 'axios'
import { api } from '@/functions/api.js'
import Filebox from '@/components/File.vue'

export default {
  data () {
    return {
      res: [],
      endpoints: [],
      loggedin: false,
      endpoint: "sfw",
      verified: false,
      verbtn: false,
      query: "",
    }
  },
  components: {
    Filebox,
  },
  methods: {
    search: function () {
      Axios({
        method: "post",
        url: `${process.env.VUE_APP_APIROOT}/api/admin/list`,
        data: {
          endpoint: this.endpoint,
          query: this.query,
          verified: this.verbtn,
        },
      }).then((response) => {
        let query = []
        this.verified = this.verbtn

        if (response.data.files) {
          response.data.files.map(file => {
            query.push({
              file: file,
              verified: this.verified,
            })
          })
        }

        this.res = query
      })
    },
    changePoint: function (newpoint) {
      this.endpoint = newpoint
    },
  },
  mounted: function () {
    api.checkLoggedIn().then(() => {
      this.loggedin = true
      this.endpoints = api.getEndpoints(false)
      this.search()
    }).catch(() => {
      this.$router.push('/admin/login')
    })
  },
}
</script>

<style lang="scss" scoped>
.vercheck {
  margin-top: -5px;
}
</style>
