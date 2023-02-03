<template>
    <v-card v-if="register" class="mx-auto" max-width="900">

        <v-toolbar>
            <v-toolbar-title>Register</v-toolbar-title>
        </v-toolbar>

        <v-container>
            <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field v-model="name" :rules="nameRules" label="Name" required></v-text-field>
                <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>
                <v-text-field v-model="password" :rules="passwordRules" label="Password" required></v-text-field>
                <v-btn :disabled="!valid" color="success" class="mr-4" @click="apiRegister">
                    Register
                </v-btn>
                <span>Already have an account ? <a href="" @click.prevent="toLogin">Login</a></span>
            </v-form>
        </v-container>
    </v-card>
    <v-card v-else class="mx-auto" max-width="900">

        <v-toolbar>
            <v-toolbar-title>Login</v-toolbar-title>
        </v-toolbar>

        <v-container>
            <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>
                <v-text-field v-model="password" :rules="passwordRules" label="Password" required></v-text-field>
                <v-btn :disabled="!valid" color="success" class="mr-4" @click="apiLogin">
                    Login
                </v-btn>
                <span>No account ? <a href="" @click.prevent="toRegister">Register</a></span>
            </v-form>
        </v-container>
    </v-card>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Login',

    data() {
        return {
            register: true,
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
        toLogin() {
            this.register = false
        },
        toRegister() {
            this.register = true
        },
        apiRegister() {
            axios
                .post("/api/register", {
                    "name": this.name,
                    "email": this.email,
                    "password": this.password,
                }).then(response => {
                    if (response.status === 200) {
                        this.$router.push('/')
                    }
                }).catch(error => {
                    console.log(error);
                });
        },
        apiLogin() {
            axios
                .post("/api/login", {
                    "email": this.email,
                    "password": this.password,
                }).then(response => {
                    if (response.status === 200) {
                        this.$router.push('/')
                    }
                }).catch(error => {
                    console.log(error);
                });
        },
    },
}
</script>
