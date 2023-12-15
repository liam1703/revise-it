import { createRouter, createWebHistory } from "vue-router";
import CreateFlashcard from "./views/CreateFlashcard.vue";



export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: CreateFlashcard,
    },
  ],
});
