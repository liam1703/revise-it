import axios from "axios";

// axios.defaults.withCredentials = true;
// axios.defaults.headers["X-Requested-With"] = "XMLHttpRequest";
// axios.defaults.baseURL = "http://127.0.0.1:8000";

const apiAxios = axios.create({
  baseURL: "http://127.0.0.1:8000",
});

export { apiAxios };
