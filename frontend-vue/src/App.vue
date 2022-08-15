<script setup lang="ts">
import { RouterView } from "vue-router";
import TopBar from "./components/TopBar/TopBar.vue";
import NavSidebar from "./components/NavSidebar/NavSidebar.vue";
</script>

<template>
  <TopBar state="waitlist" />
  <NavSidebar />
  <div class="main">
    <RouterView />
  </div>
</template>

<script lang="ts">
export default {
  name: "App",
  components: {
    RouterView,
    TopBar,
    NavSidebar,
  },
};
</script>

<style lang="scss">
@import "@/assets/variables";

.main {
  top: $topbar-height;
  left: $logo-width;
  width: calc(100vw - $logo-width);
  margin: 20px;
}
</style>
