<template>
  <v-container>
    <h1>Team {{ username }}</h1>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Home',

  data() {
    return {
      username: "Kahlys",
    }
  },
}
</script>
