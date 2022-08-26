<script setup lang="ts">
import { RouterView } from "vue-router";
import NavSidebar from "./components/NavSidebar/NavSidebar.vue";
</script>

<template>
  <nav-sidebar :wide="true" />
  <router-view />
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
