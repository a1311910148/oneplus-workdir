// vite.config.js
import { defineConfig } from "vite";
// import VitePluginStyleInject from "vite-plugin-style-inject";

export default defineConfig({
  // plugins: [VitePluginStyleInject()],
  build: {
    lib: {
      entry: "./src/main.ts",
      name: "youhouzi",
      fileName: "you",
    },
    cssCodeSplit: true,
  },
});
