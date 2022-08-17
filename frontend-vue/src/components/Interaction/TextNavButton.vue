<script setup lang="ts">
const props = defineProps<{ href: string }>();
import { useRouter } from "vue-router";
const router = useRouter();
</script>

<template>
  <button class="text-nav-button" @click.stop="router.push(props.href)">
    <slot></slot>
  </button>
</template>

<style lang="scss">
@import "@/assets/variables";

.text-nav-button {
  text-align: center;
  background-color: $blue-primary;
  border: 2px solid $yellow-canary;
  cursor: pointer;
  padding: 0rem 3rem;
  height: 3rem;
  margin-left: 2rem;
  margin-right: 2rem;
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
  text-decoration: none;
  color: $yellow-gold;
  font-size: 1.5rem;
  font-weight: bold;
}
</style>
