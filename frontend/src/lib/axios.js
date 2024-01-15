import axios from "axios";

axios.defaults.withCredentials = true;
axios.defaults.headers["X-Requested-With"] = "XMLHttpRequest";
axios.defaults.baseURL = process.env.VITE_BASE_URL;

const apiAxios = axios.create({
  baseURL: process.env.VITE_BASE_URL,
});

export { apiAxios };
