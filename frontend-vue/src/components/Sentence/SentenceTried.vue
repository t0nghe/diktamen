<script setup lang="ts">
import { computed } from "vue";
import { sentWord } from "../../types";

const props = defineProps<{
  isCorrect: boolean;
  isSummary: boolean;
  sentId: number;
  indexInArticle: number;
  tryText?: string;
  sentWords?: sentWord[];
}>();

const indexClassName = computed(() => {
  if (!props.isCorrect) {
    return "learn-sent-index-incorrect";
  } else if (props.isCorrect && props.isSummary) {
    return "learn-sent-index-correct-summary";
  } else {
    return "learn-sent-index-correct";
  }
});
</script>

<template>
  <div class="learn-sent-tried">
    <div :class="indexClassName">
      {{ props.indexInArticle }}
    </div>
    <div v-if="props.isCorrect" class="learn-sent-text">
      {{ props.tryText ? props.tryText : "" }}
    </div>
    <div v-else class="learn-sent-composite-text">
      <span
        class="learn-sent-word"
        v-for="sentword in props.sentWords"
        :key="sentword['indexInSent']"
      >
        <!-- We can't use lastInputScore for this condition because that score comes from the backend. And we might not have submitted the user input there yet. -->
        <template
          v-if="
            (sentword.isCloze &&
              sentword.lastInputText === sentword.wordform) ||
            !sentword.isCloze
          "
          >{{ sentword.wordform }}</template
        >
        <template v-else>
          <div>
            <div class="word-canonical">{{ sentword.wordform }}</div>
            <div class="word-error">{{ sentword.lastInputText }}</div>
          </div>
        </template>
      </span>
    </div>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.learn-sent-tried {
  font-family: sans-serif;
  font-weight: 200;
  font-size: 1.2rem;
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  color: $blue-primary;
  margin: 4px;
  background-color: rgba($yellow-beige, 0.75);
  box-shadow: 2px 2px rgba(#f0f0f0, 0.75);
  padding: 4px;
}

.learn-sent-index-correct {
  /* sizes and spacing */
  font-size: 1rem;
  padding: 0.1rem 0.3rem;
  margin-right: 1rem;
  /* colors */
  background-color: $blue-secondary;
  color: $yellow-pale;
  box-shadow: 2px 2px $blue-primary;
}

.learn-sent-index-correct-summary {
  /* sizes and spacing */
  font-size: 1rem;
  padding: 0.1rem 0.3rem;
  margin-right: 1rem;
  /* colors */
  background-color: $yellow-pale;
  color: $blue-secondary;
  box-shadow: 2px 2px $yellow-gold;
}

.learn-sent-index-incorrect {
  /* sizes and spacing */
  font-size: 1rem;
  padding: 0.1rem 0.3rem;
  margin-right: 1rem;
  /* colors */
  background-color: $red-primary;
  color: $yellow-pale;
  box-shadow: 2px 2px $yellow-gold;
}

.learn-sent-composite-text {
  display: flex;
  flex-direction: row;
  align-items: space-between;
  flex-wrap: wrap;

  .learn-sent-word {
    margin: 0.1rem 0.2rem;
  }

  .word-error {
    color: $red-primary;
    text-decoration: line-through;
    background-color: $yellow-beige;
    font-style: italic;
    box-shadow: 1px 1px $red-primary;
  }

  .word-canonical {
    background-color: $azure-secondary;
  }
}
</style>
