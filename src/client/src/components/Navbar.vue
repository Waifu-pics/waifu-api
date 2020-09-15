<template>
  <div>
    <v-navigation-drawer v-model="drawer" app clipped >
      <v-list v-for="route in routes" :key="route" dense nav>
        <v-list-item link :to="route.link">
          <v-list-item-action>
            <v-icon>{{route.icon}}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{route.title}}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <v-list v-for="route in bottomlink" :key="route" dense nav>
          <v-list-item :href="route.link">
            <v-list-item-action>
              <v-icon>{{route.icon}}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>{{route.title}}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </template>
    </v-navigation-drawer>

    <v-app-bar app clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>WAIFU.PICS</v-toolbar-title>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data: () => ({
    drawer: false,
    routes: [
      {
        title: "Home", 
        link: "/",
        icon: "mdi-home", 
      },
      {
        title: "NSFW", 
        link: "/nsfw/waifu",
        icon: "mdi-domino-mask",
      },
      {
        title: "More",
        link: "/more",
        icon: "mdi-dots-horizontal",
      },
      {
        title: "Uploads",
        link: "/upload",
        icon: "mdi-upload",
      },
      {
        title: "Docs",
        link: "/docs",
        icon: "mdi-api",
      },
      {
        title: "Dashboard",
        link: "/admin",
        icon: "mdi-view-dashboard",
      },
    ],
    bottomlink: [
      {
        title: "Github", 
        link: "https://github.com/Riku32/waifu.pics",
        icon: "mdi-github",
      },
    ],
  }),
}
</script>
