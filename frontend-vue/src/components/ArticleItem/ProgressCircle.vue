<script setup lang="ts">
import { computed, ComputedRef } from "vue";
const props = defineProps<{
  primary: string;
  secondary: string;
  prog: number;
  goal: number;
}>();

// Normalized progress and goal values.
const normProg: ComputedRef<number> = computed(
  (): number => Math.min(Math.max(props.prog, 0), props.goal)
);
const normGoal: ComputedRef<number> = computed(
  (): number => Math.max(props.goal, 0)
);
</script>

<template>
  <div class="progress">
    <div class="pie"></div>
    <div class="desc">{{ normProg }}/{{ normGoal }}</div>
  </div>
</template>

<style lang="scss" scoped>
.progress {
  position: relative;
  /* These are manually changed in CSS. */
  --progress-width: 100px;
  --progress-stroke: 20px;
  /* These are props passed in from parent component. */
  --primary-color: v-bind(primary);
  --secondary-color: v-bind(secondary);
  /* These doesn't work. Even after adding `ref`, percentage and degsStrings still doesn't change with user input. */
  /* --p: v-bind(percentage);
  --v: v-bind(degsString); */
  /* Trying to calculate inside CSS using calc(). */
  --goal: v-bind(normGoal);
  --prog: v-bind(normProg);
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
