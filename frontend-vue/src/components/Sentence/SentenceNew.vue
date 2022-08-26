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
    <div class="learn-sent-fake-text">
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
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.learn-new-sent {
  /* display: inline-block; */
  font-size: 1.2rem;
  /* color and shadow */
  background-color: rgba($yellow-pale, 0.75);
  border: 1px solid rgba($yellow-beige, 0.75);
  box-shadow: 2px 2px rgba($yellow-gold, 0.75);
  margin: 8px;
  display: flex;
  flex-direction: row;

  @include for-mobile {
    font-size: 1rem;
  }

  .learn-sent-index-new {
    font-size: 1rem;
    padding: 0.1rem 0.3rem;
    margin-right: 1rem;
    color: $blue-primary;
  }

  .learn-sent-fake-text {
    word-wrap: break-word;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
  }

  .learn-sent-word {
    margin: 0.1rem 0.2rem;
    /* word-wrap: break-word; */
  }
}
</style>
