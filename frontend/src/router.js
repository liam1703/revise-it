import { createRouter, createWebHistory } from "vue-router";
import HelloWorld from "./views/HelloWorld.vue";


export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: HelloWorld,
    },
    {
      path: "/testing",
      component: test,
    },
  ],
});
