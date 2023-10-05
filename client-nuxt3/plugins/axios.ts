import axios from "axios";
export default defineNuxtPlugin((nuxtApp) => {
  const defaultUrl = "http://localhost:8080"; // Remove the angle brackets around the URL

  let api = axios.create({
    baseURL: defaultUrl, 
    headers: {
      common: {},
    },
  });
  return {
    provide: {
      api: api,
    },
  };
});
