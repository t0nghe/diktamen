<script setup lang="ts">
const props = defineProps<{ href: string }>();
</script>

<template>
  <div class="text-nav-button">
    <router-link :to="props.href">
      <slot></slot>
    </router-link>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.text-nav-button {
  background-color: $blue-primary;
  border: 2px solid $yellow-canary;
  display: inline-block;

  a {
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
}
</style>
