import { createApp } from "vue";
import { createPinia } from "pinia";

import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
  ApolloLink,
  concat,
} from "@apollo/client/core";
import { createApolloProvider } from "@vue/apollo-option";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

const token = localStorage.getItem("token") || "";
console.log("token: ", token);

const authMiddleware = new ApolloLink((operation, forward) => {
  operation.setContext({
    headers: { authorization: `bearer ${token}` },
  });
  return forward(operation);
});

const httpLink = createHttpLink({
  uri: import.meta.env.VITE_DEV_GRAPHQL_SERVER,
});

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: concat(authMiddleware, httpLink),
  cache,
});

const apolloProvider = createApolloProvider({
  defaultClient: apolloClient,
});

app.use(createPinia());
app.use(router);
app.use(apolloProvider);

app.mount("#app");
