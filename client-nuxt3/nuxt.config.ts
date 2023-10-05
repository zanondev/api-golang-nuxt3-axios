import vuetify from "vite-plugin-vuetify";

export default defineNuxtConfig({
  devtools: { enabled: false },
  ssr: false,
  css: ["vuetify/lib/styles/main.sass"],
  build: {
    transpile: ["vuetify"],
  },
  vite: {
    ssr: {
      noExternal: ["vuetify"],
    },
  },
  hooks: {
    "vite:extendConfig": (config) => {
      config.plugins?.push(
        vuetify({
          styles: { configFile: "./settings.scss" },
        })
      );
    },
  },
});
