<script setup lang="ts">
// import { ref } from "vue";
const props = defineProps<{
  primary: string;
  secondary: string;
  prog: number;
  goal: number;
}>();

// This did NOT make CSS reactive to props.
// That's why we are using calc().
// let percentage = ref(Math.round((props.prog * 100) / props.goal));
// let degs = ref(Math.round(percentage.value * (18 / 5)) - 180);
// let degsString = ref(`${degs.value}deg`);
</script>

<template>
  <div class="progress">
    <div class="pie"></div>
    <div class="desc">{{ prog }}/{{ goal }}</div>
  </div>
</template>

<style lang="scss" scoped>
.progress {
  position: relative;
  /* These are manually changed in CSS. */
  --progress-width: 100px;
  --progress-stroke: 20px;
  /* These are props passed in from parent component. */
  --primary-color: v-bind(props.primary);
  --secondary-color: v-bind(props.secondary);
  /* These doesn't work. Even after adding `ref`, percentage and degsStrings still doesn't change with user input. */
  /* --p: v-bind(percentage);
  --v: v-bind(degsString); */
  /* Trying to calculate inside CSS using calc(). */
  --prog: v-bind(props.prog);
  --goal: v-bind(props.goal);
  --p: calc(var(--prog) / var(--goal) * 100);
  --v: calc((var(--p) * (18 / 5) - 180) * 1deg);
}

.desc {
  position: absolute;
  width: calc(var(--progress-width) - var(--progress-stroke) * 2);
  height: calc(var(--progress-width) - var(--progress-stroke) * 2);
  top: var(--progress-stroke);
  left: var(--progress-stroke);
  font-size: 1rem;
  font-weight: bold;
  color: var(--secondary-color);
  background-color: white;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.pie {
  position: absolute;
  top: 0;
  left: 0;
  width: var(--progress-width);
  height: var(--progress-width);
  display: flex;
  border-radius: 50%;
  overflow: hidden;
  background: linear-gradient(
    to right,
    var(--secondary-color) 50%,
    var(--primary-color) 0
  );
  justify-content: center;
  align-items: center;

  &::before,
  &::after {
    content: "";
    width: 50%;
    padding-top: 100%;
    transform: rotate(var(--v));
  }

  // secondary color
  &::before {
    background: linear-gradient(var(--secondary-color) 0 0) 0 /
      calc((50 - var(--p)) * 1%);
    transform-origin: right;
  }

  // primary color
  &::after {
    background: linear-gradient(var(--primary-color) 0 0) 0 /
      calc((var(--p) - 50) * 1%);
    transform-origin: left;
  }
}
</style>
