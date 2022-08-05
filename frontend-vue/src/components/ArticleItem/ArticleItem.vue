<script setup lang="ts">
import { ref } from "vue";
import variables from "@/assets/variables.ts";

defineProps({
  pcPrimary: ref("#00bcd4"),
  pcSecondary: ref("#dcbc00"),
  pcProg: ref(1),
  pcGoal: ref(10),
});
</script>

<template>
  <div>
    <div class="article-item__title">
      Zelenskyj är en man som kvinnor längtar efter
    </div>
    den traditionelle mannen
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.article-item-progress {
  background-color: $blue-primary;
}

.article-item-complete {
  background-color: $blue-secondary;
}

.article-item-new {
  background-color: $yellow-beige;
}
</style>
