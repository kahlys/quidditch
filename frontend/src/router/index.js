// Composables
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: () => import("@/layouts/default/Default.vue"),
    children: [
      {
        path: "",
        name: "Home",
        component: () =>
          import(/* webpackChunkName: "home" */ "@/views/Home.vue"),
        beforeEnter: (to, from, next) => {
          // TODO: add a test and use next(login) if not logged in.
          next();
        },
      },
    ],
  },
  {
    path: "/login",
    component: () => import("@/layouts/default/Default.vue"),
    children: [
      {
        path: "",
        name: "Login",
        component: () =>
          import(/* webpackChunkName: "login" */ "@/views/Login.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
