import { loadEnv, defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";


// https://vitejs.dev/config/
export default ({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  // expose .env as process.env instead of import.meta since jest does not not support import meta yet
  const processEnvValues = {
    "process.env": Object.entries(env).reduce(
      (prev, [key, val]) => ({
        ...prev,
        [key]: val,
      }),
      {}
    ),
  };

  return defineConfig({
    plugins: [
      vue(),
    ],
    define: processEnvValues,
  });
};
