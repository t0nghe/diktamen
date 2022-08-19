import { defineStore } from "pinia";

export const useLoginStore = defineStore({
  id: "loggedInStore",
  state: () => ({
    loginStatus: false,
  }),
  getters: {
    isLoggedIn: (state) => state.loginStatus,
  },
  actions: {
    piniaLogOut() {
      this.loginStatus = false;
    },
    piniaLogIn() {
      this.loginStatus = true;
    },
  },
});
