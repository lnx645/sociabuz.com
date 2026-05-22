import path from "path";
import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import { svelte, vitePreprocess } from "@sveltejs/vite-plugin-svelte";
export default defineConfig({
  server:{
    cors:true,
  },
  plugins: [
    tailwindcss(),
    svelte({
      preprocess: vitePreprocess(),
    }),
  ],
  build: {
    outDir: "../public",
    manifest: "manifest.json",
    emptyOutDir: true,
    rolldownOptions: {
      input: {
        main: path.resolve(__dirname, "src/main.ts"),
        admin: path.resolve(__dirname, "src/admin.ts"),
        overlay: path.resolve(__dirname, "src/overlay.ts"),
      },
      output: {
        entryFileNames: (chunkInfo: any) => {
          if (chunkInfo.name === "main") {
            return "js/main[hash].js";
          }

          if (chunkInfo.name === "admin") {
            return "admin/admin.[hash].js";
          }

          if (chunkInfo.name === "overlay") {
            return "overlay/overlay[hash].js";
          }

          return "js/[name].js";
        },
        assetFileNames: (assetInfo: any) => {
          const ext = assetInfo.name.split(".").pop();

          if (ext === "css") {
            return "css/app.css";
          }

          return "assets/[name].[ext]";
        },
      },
    },
  },
});
