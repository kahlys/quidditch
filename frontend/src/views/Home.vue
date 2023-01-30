<template>
  <v-card class="mx-auto" max-width="900">

    <v-toolbar>
      <v-toolbar-title>Register</v-toolbar-title>
    </v-toolbar>

    <v-container>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field v-model="name" :rules="nameRules" label="Name" required></v-text-field>
        <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>
        <v-text-field v-model="password" :rules="passwordRules" label="Password" required></v-text-field>
        <v-btn :disabled="!valid" color="success" class="mr-4" @click="validate">
          Register
        </v-btn>
        <span>Already have an account ? <a href="">Login</a></span>
      </v-form>
    </v-container>
  </v-card>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Home',

  data() {
    return {
      valid: true,
      name: '',
      email: '',
      password: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length >= 3) || 'Name must be at least 3 characters',
      ],
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
        v => (v && v.length >= 8) || 'Password must be at least than 8 characters',
      ],
    }
  },

  methods: {
    validate() {
      axios
        .post("/api/register",
          {
            "name": this.name,
            "email": this.email,
            "password": this.password,
          },
        )
      // .then(response => {
      //   this.result = response.data.exist
      // })

      axios.get("/api/home")
    },
  },
}
</script>
