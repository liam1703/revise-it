import { createRouter, createWebHistory } from "vue-router";
import CreateFlashcard from "./views/CreateFlashcard.vue";
import Home from "./views/Home.vue"



export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Home,
    },
    {
      path:"/create-new-flashcard",
      component: CreateFlashcard,
    }
  ],
});
