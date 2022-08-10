<script setup lang="ts">
import { ref } from "vue";

const props = defineProps<{
  state: "waitlist" | "review" | "articles" | "learn" | "summary";
  title?: string | null;
  revcount?: number | null;
  duecount?: number | null;
  ahref?: string | null;
}>();

const className = ref("classy");
</script>

<template>
  <div id="top-bar-wrapper" :class="className">
    &nbsp;
  </div>
</template>

<style lang="scss"></style>
