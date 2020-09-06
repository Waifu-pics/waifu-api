<template>
  <div v-if="!loggedin" class="centered">
    <v-card max-width="400" outlined>
      <div class="topbox">
        <h1 class="font-weight-light">Sign in</h1>
        <p>login to admin dashboard</p>
      </div>
      <v-form @submit.prevent="login">
        <v-card-text>
          <v-text-field v-model="username" outlined label="Login" name="login" type="text"/>
          <v-text-field v-model="password" outlined id="password" label="Password" name="password" type="password"/>
        </v-card-text>
        <v-btn class="centered" depressed type="submit">Login</v-btn>
      </v-form>
    </v-card>
  </div>
</template>

<style lang="scss" scoped>
.v-btn {
  margin-bottom: 5px;
}
.v-card {
  margin: auto;
}
.topbox {
  text-align: center;
  margin-top: 20px;
}
</style>

<script>
import Axios from 'axios'
import { api } from '@/functions/api.js'

export default {
  data: function () {
    return {
      loggedin: true,
    }
  },
  methods: {
    login: function () {
      const { username, password } = this
      
      Axios({
        method: "post",
        url: `/api/admin/login`,
        data: {
          username: username,
          password: password,
        },
      }).then(() => {
        this.$notification.success("You have been logged in!")
        this.$router.push('/admin')
      }).catch(() => {
        this.$notification.error("There was a problem logging in!")
      })
    },
  },
  mounted: function () {
    api.checkLoggedIn().then(() => {
      this.$router.push('/admin')
    }).catch(() => {
      this.loggedin = false
    })
  },
}
</script>
