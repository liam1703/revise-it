import { createRouter, createWebHistory } from "vue-router";
import CreateFlashcard from "./views/CreateFlashcard.vue";
import Home from "./views/Home.vue"
import Login from "./views/Login.vue"



export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Home,
    },
    {
      path: "/login",
      component: Login
    },
    {
      path:"/create-new-flashcard",
      component: CreateFlashcard,
    }
  ],
});
