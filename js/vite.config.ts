import path from "path";
import { defineConfig } from "vite";

export default defineConfig({
  build: {
    outDir: "../public",
    manifest:"manifest.json",
    rolldownOptions: {
      input: {
        main: path.resolve(__dirname, "src/main.ts"),
        admin: path.resolve(__dirname, "src/admin.ts"),
        overlay: path.resolve(__dirname, "src/overlay.ts"),
      },
      output: {
        entryFileNames: (chunkInfo: any) => {
          if (chunkInfo.name === "main") {
            return "js/main.js";
          }

          if (chunkInfo.name === "admin") {
            return "admin/admin.bundle.js";
          }

          if (chunkInfo.name === "overlay") {
            return "overlay/overlay.js";
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
