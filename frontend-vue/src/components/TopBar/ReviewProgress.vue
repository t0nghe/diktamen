<script setup lang="ts">
import { computed } from "vue";
const props = defineProps<{
  revcount: number | null | undefined;
  duecount: number | null | undefined;
}>();

const revCountNum = computed(() => (props.revcount ? props.revcount : 0));
const dueCountNum = computed(() => (props.duecount ? props.duecount : 1));
const pct = computed(() => {
  const pctNum = (revCountNum.value * 100) / dueCountNum.value;
  return `${pctNum}%`;
});
</script>

<template>
  <div class="review-progress">
    <div></div>
    <div class="review-heading">review</div>
    <div class="review-progress-numbers">
      {{ revCountNum }}/{{ dueCountNum }}
    </div>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.review-progress {
  position: relative;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  width: calc(100vw - $logo-width);
  height: $topbar-height;

  background: linear-gradient(
    to right,
    $blue-secondary 0 v-bind(pct),
    $blue-primary v-bind(pct) 100%
  );

  .review-heading {
    font-size: 1.5rem;
    font-weight: bold;
    text-align: center;
    width: auto;
  }

  .review-progress-numbers {
    font-size: 1.5rem;
    font-weight: 200;
    text-align: right;
    padding-right: 10px;
  }
}
</style>
