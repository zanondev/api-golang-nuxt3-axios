// plugins/axios.js

import { defineNuxtPlugin } from 'nuxt3';

export default defineNuxtPlugin(({ app }) => {
  // Configure o Axios aqui, se necessário
  app.config.globalProperties.$axios.defaults.baseURL = 'http://localhost:8080'; // Substitua pela URL da sua API
});
