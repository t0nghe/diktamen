<script setup lang="ts">
const props = defineProps<{ score: number }>();
</script>

<template>
  <div class="summary-score-circle">
    {{ Math.floor(props.score * 100) }}
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.summary-score-circle {
  /* width: calc(var(--max-dimension) * 0.35);
  height: calc(var(--max-dimension) * 0.38); */
  width: 600px;
  height: 600px;
  left: 60vw;
  top: 30vh;
  position: fixed;
  background: transparent;
  z-index: -2;

  border: 32px solid $blue-secondary;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;

  font-size: 20rem;
  color: $blue-secondary;
  font-weight: bold;
}
</style>
