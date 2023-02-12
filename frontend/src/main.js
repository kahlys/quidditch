/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

import App from "./App.vue"; // Components
import { createApp } from "vue"; // Composables
import { registerPlugins } from "@/plugins"; // Plugins

import axios from "axios";
import router from "./router";

axios.defaults.withCredentials = true;

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      console.log("HALLO");
      router.push("/login");
    }
    return Promise.reject(error);
  }
);

const app = createApp(App);

registerPlugins(app);

app.mount("#app");
