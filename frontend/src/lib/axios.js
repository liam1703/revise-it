import axios from "axios";

axios.defaults.withCredentials = true;
axios.defaults.headers["X-Requested-With"] = "XMLHttpRequest";
axios.defaults.baseURL = "http://localhost:8000";

const apiAxios = axios.create({
  baseURL: "http://localhost:8000",
});

export { apiAxios };
