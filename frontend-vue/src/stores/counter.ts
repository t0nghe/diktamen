import { defineStore } from "pinia";

export const useLoggedInStore = defineStore({
  id: "loggedIn",
  state: () => ({
    loggedIn: false,
  }),
  getters: {
    isLoggedIn: (state) => state.loggedIn,
  },
  actions: {
    piniaLogOut() {
      this.loggedIn = false;
    },
    piniaLogIn() {
      this.loggedIn = true;
    },
  },
});
