<template>
  <div class="about centered">
    <v-card class="mx-auto" max-width="400" outlined>
      <v-list-item three-line>
        <v-list-item-content>
          <v-list-item-title class="headline mb-1">Pages</v-list-item-title>
          <v-list-item-subtitle>Here are all the different pages you can explore.</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-list-item class="bottom">
        <v-btn outlined v-for="endpoint in endpoints" :key="endpoint" :to="endpoint" text>{{endpoint}}</v-btn>
      </v-list-item>
    </v-card>
  </div>
</template>

<style lang="scss">
.vertical-center {
  margin: 0;
  position: absolute;
  top: 50%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
}

.v-card {
  text-align: center;

  .v-list-item {
    display: block;
  }

  .v-btn {
    width: 100px;
    margin: 6px;
  }

  .bottom {
    margin-bottom: 20px;
  }
}
</style>

<script>
import testdata from '@/assets/endpoints.js'

export default {
  name: 'More',
  data () {
    return {
      endpoints: [],
    }
  },
  created () {
    this.endpoints = testdata
  },
}
</script>
