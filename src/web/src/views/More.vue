<template>
  <div class="about centered">
    <v-card class="mx-auto" max-width="400" outlined>
      <v-list-item three-line>
        <v-list-item-content>
          <v-list-item-title class="headline mb-1">Pages</v-list-item-title>
          <v-list-item-subtitle>Here are all the different pages you can explore.</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-expansion-panels v-model="panel" multiple>
        <v-expansion-panel>
          <v-expansion-panel-header>SFW</v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-list-item class="bottom">
              <v-btn outlined v-for="endpoint in endpoints.sfw" :key="endpoint" :to="`/sfw/${endpoint}`" text>{{endpoint}}</v-btn>
            </v-list-item>
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header>NSFW</v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-list-item class="bottom">
              <v-btn outlined v-for="endpoint in endpoints.nsfw" :key="endpoint" :to="`/nsfw/${endpoint}`" text>{{endpoint}}</v-btn>
            </v-list-item>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card>
  </div>
</template>

<style lang="scss" scoped>
.v-card {
  text-align: center;

  .v-list-item {
    display: block;
  }

  .v-btn {
    width: 90px;
    margin: 6px;
  }

  .bottom {
    margin-bottom: 20px;
  }
}
</style>

<script>
import { api } from '@/functions/api.js'

export default {
  name: 'More',
  data () {
    return {
      endpoints: {
        sfw: {},
        nsfw: {},
      },
    }
  },
  created () {
    api.getEndpoints((res) => {
      this.endpoints.sfw = res.sfw
      this.endpoints.nsfw = res.nsfw
    })
  },
}
</script>
