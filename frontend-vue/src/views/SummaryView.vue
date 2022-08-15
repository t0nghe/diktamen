<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, computed, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import TextNavButton from "../components/Interaction/TextNavButton.vue";
import ScoreCircle from "../components/Interaction/ScoreCircle.vue";
import SentenceTried from "../components/Sentence/SentenceTried.vue";
import { corrSents, incorrSents } from "../queries";
import { examineSent } from "@/types";

const route = useRoute();
const articleId = computed(() => {
  return route.params.id;
});

const score = ref(0);

const {
  result: result1,
  loading: loading1,
  error: error1,
} = useQuery(corrSents, {
  articleId: articleId.value,
});

const {
  result: result2,
  loading: loading2,
  error: error2,
} = useQuery(incorrSents, {
  articleId: articleId.value,
});

const allSents = ref<examineSent[]>([]);

watch(result1, () => {
  if (result1.value) {
    allSents.value = allSents.value.concat(result1.value.examineCorrectSents);
  }
});

watch(result2, () => {
  if (result2.value) {
    allSents.value = allSents.value.concat(result2.value.examineIncorrectSents);
  }
});

const sentsSorted = computed<examineSent[]>(() => {
  return [...allSents.value].sort(
    (a, b) => a.indexInArticle - b.indexInArticle
  );
});

const getScore = () => {
  const count = sentsSorted.value.length;
  let sum = 0;
  for (let i = 0; i < count; i++) {
    if (sentsSorted.value[i].tryText) {
      sum += 1;
    } else {
      let thisSum = 0;
      const sentWords = sentsSorted.value[i].sentWords || undefined;
      if (typeof sentWords !== "undefined" && sentWords.length > 0) {
        for (let j = 0; j < sentWords.length; j++) {
          const thisWord = sentWords[j];
          if (
            thisWord &&
            typeof thisWord.lastInputScore !== "undefined" &&
            thisWord.isCloze
          ) {
            if (thisWord.lastInputScore === 0) {
              thisSum += 1;
            } else {
              thisSum += thisWord.lastInputScore;
            }
          }
        }
        sum += thisSum / sentWords.length;
      }
    }
  }
  return sum / count;
};

watch(sentsSorted, () => {
  score.value = getScore();
});
</script>

<template>
  <ScoreCircle v-if="score !== 0 && !isNaN(score)" :score="score" />
  <div class="view-content-wrapper">
    <div v-for="sent in sentsSorted" :key="sent.indexInArticle">
      <div v-if="sent.tryText">
        <SentenceTried
          :is-correct="true"
          :is-summary="true"
          :sent-id="sent.sentId"
          :index-in-article="sent.indexInArticle"
          :try-text="sent.tryText"
        />
      </div>
      <div v-else>
        <SentenceTried
          :is-correct="false"
          :is-summary="true"
          :sent-id="sent.sentId"
          :index-in-article="sent.indexInArticle"
          :sent-words="sent.sentWords"
        />
      </div>
    </div>
    <!-- <h2>Correct sentences:</h2>
    <div v-if="result1">{{ result1.examineCorrectSents }}</div>
    <div v-if="loading1">Loading correct sentences...</div>
    <div v-if="error1">{{ error1.message }}</div>

    <h2>Incorrect sentences:</h2>
    <div v-if="result2">{{ result2.examineIncorrectSents }}</div>
    <div v-if="loading2">Loading incorrect sentences...</div>
    <div v-if="error2">{{ error2.message }}</div> -->
  </div>
  <TextNavButton href="/">Complete</TextNavButton>
</template>

<style>
.view-content-wrapper {
  margin-right: 30px;
}
</style>
