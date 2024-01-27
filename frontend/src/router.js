import { createRouter, createWebHistory } from "vue-router";
import CreateFlashcard from "./views/CreateFlashcard.vue";
import Home from "./views/Home.vue"
import Login from "./views/Login.vue"
import SignUp from "./views/SignUp.vue"
import MyDecks from "./views/MyDecks.vue"




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
      path: "/sign-up",
      component: SignUp
    },
    {
      path: "/my-decks",
      component: MyDecks
    },
    {
      path:"/create-new-flashcard",
      component: CreateFlashcard,
    }
  ],
});
