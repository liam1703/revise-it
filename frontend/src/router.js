import { createRouter, createWebHistory } from "vue-router";
import HelloWorld from "./views/HelloWorld.vue";
import Testing from "./views/Testing.vue";


export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: HelloWorld,
    },
    {
      path: "/testing",
      component: Testing,
    },
  ],
});
