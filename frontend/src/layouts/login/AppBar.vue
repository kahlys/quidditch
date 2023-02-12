<template>
  <v-app-bar>
    <v-app-bar-title>
      Quidditch Manager
    </v-app-bar-title>
  </v-app-bar>
</template>

<script>
export default {
  data: () => ({
    drawer: false,
    group: null,
    items: [
      {
        title: 'Home',
        path: '/',
      },
      {
        title: 'Team',
        path: '/team',
      },
    ],
  }),

  watch: {
    group() {
      this.drawer = false
    },
  },
}
</script>
