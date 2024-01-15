import { createApp } from 'vue'
import './index.css'
import router from "./router";
import App from './App.vue'

// Styles
import "./assets/styles.css";

createApp(App).use(router).mount('#app')
