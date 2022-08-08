<script setup lang="ts">
import { ref } from "vue";
const myString = ref("");
myString.value = window.localStorage.getItem("textToRender") || "";
</script>

<template>
  <div>{{ myString }}</div>
</template>

<style>
</style>
