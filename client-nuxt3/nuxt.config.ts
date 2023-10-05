import vuetify from "vite-plugin-vuetify";

export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ["vuetify/lib/styles/main.sass"],
  build: {
    transpile: ["vuetify"],
  },
  vite: {
    ssr: {
      noExternal: ["vuetify"],
    },
  },
  plugins: [
    '~/plugins/axios.js',
  ],
  hooks: {
    "vite:extendConfig": (config) => {
      config.plugins?.push(
        vuetify({
          styles: { configFile: "./settings.scss" },
        })
      );
    },
  },
})
