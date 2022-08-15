import { createApp, h } from "vue";
import { createPinia } from "pinia";

import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client/core";
import { provideApolloClient } from "@vue/apollo-composable";

import App from "./App.vue";
import router from "./router";
import "@/assets/base.css";

const getHeaders = () => {
  const headers: { Authorization?: string; "Content-Type"?: string } = {};
  const token = localStorage.getItem("token") || "";
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }
  headers["Content-Type"] = "application/json";

  console.log(headers);
  return headers;
};

const httpLink = createHttpLink({
  uri: import.meta.env.VITE_DEV_GRAPHQL_SERVER,
  fetch: (uri: RequestInfo, options: RequestInit) => {
    options.headers = getHeaders();
    return fetch(uri, options);
  },
});

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
  defaultOptions: {
    query: {
      errorPolicy: "all",
    },
    mutate: {
      errorPolicy: "all",
    },
  },
});

const app = createApp({
  setup() {
    provideApolloClient(apolloClient);
  },
  render: () => h(App),
});

app.use(createPinia());
app.use(router);

app.mount("#app");
