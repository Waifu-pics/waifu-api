<template>
  <div v-if="loggedin">
    <v-container>
      <v-card class="mx-auto" style="margin-top: 20px; margin-bottom: 10px;" max-width="400" outlined>
        <v-list-item three-line>
          <v-list-item-content>
            <v-text-field outlined dense v-model="query" label="Search"/>
            <v-checkbox style="margin-bottom:20px; margin-left:20px;" label="NSFW" v-model="nsfw" @change="update"/>
            <v-select v-model="endpoint" dense outlined label="Endpoint" :items="endpoints"/>
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

    <v-bottom-navigation v-if="deletelist.length > 0" :value="activeBtn" inset app>
      <v-btn v-on:click="verifymany(true)">
        <v-icon>mdi-check</v-icon>
      </v-btn>

      <v-btn v-on:click="verifymany(false)">
        <v-icon>mdi-delete</v-icon>
      </v-btn>

      <v-btn v-on:click="deletelist = []">
        <v-icon>mdi-select-off</v-icon>
      </v-btn>
    </v-bottom-navigation>

    <v-btn fab large dark fixed bottom right v-on:click="logout()">
      <v-icon>mdi-logout-variant</v-icon>
    </v-btn>
  </div>
</template>

<script>
import Axios from 'axios'
import { api } from '@/functions/api.js'
import Filebox from '@/components/File.vue'
import store from '@/store/index.js'

export default {
  data () {
    return {
      res: [],
      endpoints: [],
      deletelist: [],
      loggedin: false,
      endpoint: "waifu",
      nsfw: false,
      verified: false,
      verbtn: false,
      query: "",
    }
  },
  components: {
    Filebox,
  },
  methods: {
    verifymany: function (verify) {
      let endpoint = verify ? "verify" : "delete"

      Axios({
        method: "post",
        url: `/api/admin/${endpoint}`,
        data: {
          files: this.deletelist,
        },
      }).then((res) => {
        this.$notification.success(res.data.message)

        this.deletelist = []

        this.search()
      }).catch((error) => {
        this.deletelist = []

        this.$notification.error(error.response.data.message)
      })
    },
    logout: function () {
      // Remove cookie
      this.loggedin = false
      document.cookie = "auth-token= ; expires = Thu, 01 Jan 1970 00:00:00 GMT"

      // Send notif and go to login
      this.$notification.success("You have been logged out!")
      this.$router.push('/admin/login')
    },
    search: function () {
      Axios({
        method: "post",
        url: `/api/admin/list`,
        data: {
          endpoint: this.endpoint,
          nsfw: this.nsfw,
          query: this.query,
          verified: this.verbtn,
        },
      }).then((response) => {
        let query = []
        this.verified = this.verbtn

        if (response.data.files) {
          response.data.files.map(file => {
            query.push({
              file: {
                name: file.name,
                url: file.url,
              },
              verified: this.verified,
            })
          })
        }
        this.res = query
      })
    },
    update: function () {
      this.endpoints = (this.nsfw ? this.$store.getters.endpoints.nsfw : this.$store.getters.endpoints.sfw)
      this.endpoint = this.endpoints[0]
    },
  },
  mounted: function () {
    api.checkLoggedIn().then(() => {
      this.loggedin = true
      this.update()
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
