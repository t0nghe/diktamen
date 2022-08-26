<script setup lang="ts">
import { useRoute } from "vue-router";
import { ref, computed, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import TextNavButton from "../components/Interaction/TextNavButton.vue";
import ScoreCircle from "../components/Interaction/ScoreCircle.vue";
import SentenceTried from "../components/Sentence/SentenceTried.vue";
import LoadingEllipsis from "@/components/Interaction/LoadingEllipsis.vue";
import { corrSents, incorrSents } from "../graphql";
import { examineSent } from "@/types";
import TopBar from "@/components/TopBar/TopBar.vue";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
const route = useRoute();
const articleId = computed(() => {
  return route.params.id;
});

const score = ref(0);

const {
  result: resultCorrect,
  loading: loadingCorrect,
  // error: error1,
} = useQuery(
  corrSents,
  {
    articleId: articleId.value,
  },
  { fetchPolicy: "network-only" }
);

const {
  result: resultIncorrect,
  loading: loadingIncorrect,
  // error: error2,
} = useQuery(
  incorrSents,
  {
    articleId: articleId.value,
  },
  { fetchPolicy: "network-only" }
);

const allSents = ref<examineSent[]>([]);

watch(resultCorrect, () => {
  if (resultCorrect.value) {
    allSents.value = allSents.value.concat(
      resultCorrect.value.examineCorrectSents
    );
  }
});

watch(resultIncorrect, () => {
  if (resultIncorrect.value) {
    allSents.value = allSents.value.concat(
      resultIncorrect.value.examineIncorrectSents
    );
  }
});

const sentsSorted = ref<examineSent[]>({});
watch(allSents, () => {
  sentsSorted.value = [...allSents.value].sort(
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
  <top-bar state="summary" />
  <div
    class="view-container"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <score-circle v-if="score !== 0 && !isNaN(score)" :score="score" />
    <div class="view-content-wrapper">
      <div
        v-for="sent in sentsSorted"
        :key="sent.indexInArticle"
        class="summary-view-sentence"
      >
        <template v-if="sent.tryText">
          <sentence-tried
            :is-correct="true"
            :is-summary="true"
            :sent-id="sent.sentId"
            :index-in-article="sent.indexInArticle"
            :try-text="sent.tryText"
          />
        </template>
        <template v-else>
          <sentence-tried
            :is-correct="false"
            :is-summary="true"
            :sent-id="sent.sentId"
            :index-in-article="sent.indexInArticle"
            :sent-words="sent.sentWords"
          />
        </template>
      </div>
      <div class="bottom-text-button">
        <template v-if="loadingCorrect || loadingIncorrect">
          <loading-ellipsis />
        </template>
        <text-nav-button href="/articles">complete</text-nav-button>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.view-content-wrapper {
  width: 60%;
  margin-left: -5vw;
  margin-right: 15vw;

  @include for-mobile {
    margin-top: 100px;
    width: 80%;
  }
}

.bottom-text-button {
  text-align: center;
  margin: 50px;
  margin-bottom: 150px;
}
</style>
