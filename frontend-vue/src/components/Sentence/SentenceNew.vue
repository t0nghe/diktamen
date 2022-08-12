<script setup lang="ts">
import { sentWord } from "../../types";

const props = defineProps<{
  sentId: number;
  indexInArticle: number;
  tryText?: string;
  sentWords?: sentWord[];
}>();
</script>

<template>
  <div class="learn-new-sent">
    <div class="learn-sent-index-new">{{ props.indexInArticle }}</div>
    <span
      v-for="word in props.sentWords"
      :key="`${word.indexInSent}` + word.wordform"
      class="learn-sent-word"
    >
      <span v-if="!word.isCloze" style="font-weight: bold">
        {{ word.wordform }}
      </span>
      <template v-else>
        {{ "_".repeat(word.length) }}
      </template>
    </span>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.learn-new-sent {
  display: inline-block;
  font-size: 1.2rem;
  /* color and shadow */
  background-color: $yellow-pale;
  border: 1px solid $yellow-beige;
  box-shadow: 2px 2px $yellow-gold;
  margin: 8px;

  .learn-sent-index-new {
    font-size: 1rem;
    padding: 0.1rem 0.3rem;
    margin-right: 1rem;
    float: left;
    /* colors */
    color: $blue-primary;
  }

  .learn-sent-word {
    margin: 0.1rem 0.2rem;
  }
}
</style>
