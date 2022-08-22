<script setup lang="ts">
import { RouterView } from "vue-router";
import NavSidebar from "./components/NavSidebar/NavSidebar.vue";
</script>

<template>
  <nav-sidebar :wide="true" />
  <div class="main">
    <router-view />
  </div>
</template>

<script lang="ts">
export default {
  name: "App",
  components: {
    RouterView,
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
