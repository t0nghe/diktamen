import { defineStore } from "pinia";

export const useNavStore = defineStore({
  id: "navWidthStore",
  state: () => ({
    wide: true,
  }),
  getters: {
    isWide: (state) => state.wide,
  },
  actions: {
    toggle() {
      this.wide = !this.wide;
    },
  },
});
