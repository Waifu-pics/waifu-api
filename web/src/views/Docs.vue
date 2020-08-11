<template>
  <div class="centered">
    <v-card class="mx-auto" max-width="450" outlined>
      <v-list-item three-line>
        <v-list-item-content>
          <v-list-item-title class="headline mb-1">API Docs</v-list-item-title>
          <v-list-item-subtitle>Here are the docs that should explain how to use the API</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-list-item class="bottom">
        <v-card class="mx-auto" max-width="400" tile>
          <v-list>
            <v-subheader>ENDPOINTS</v-subheader>
            <v-list-item v-for="endpoint in endpoints" :key="endpoint" :inactive="inactive">
              <v-list-item-content>
                <v-list-item-title v-html="endpoint"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-list-item>
    </v-card>
  </div>
</template>

<script>
import Axios from 'axios'
import { api } from '@/functions/api.js'

export default {
  name: 'More',
  data () {
    return {
      endpoints: [],
    }
  },
  mounted: function () {
    this.endpoints = api.getEndpoints(true)
  },
}
</script>

<style lang="scss" scoped>
.centered {
  text-align: center;
}
.v-list{
  height: 300px;
  width: 350px;
  overflow-y: auto;
  text-align: start;
}
.bottom {
  margin-bottom: 30px;
}
</style>
