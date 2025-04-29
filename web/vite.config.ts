import Tailwindcss from "@tailwindcss/vite";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import React from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";
import TSConfigPaths from "vite-tsconfig-paths";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    TanStackRouterVite({ target: "react", autoCodeSplitting: true }),
    React(),
    Tailwindcss(),
    TSConfigPaths({ projectDiscovery: "lazy" }),
  ],
});
