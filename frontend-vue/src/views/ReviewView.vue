<script setup lang="ts">
import { computed, ref } from "vue";
import { dueSents, mutationTrySent } from "../graphql";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { useRoute } from "vue-router";
import LoadingEllipsis from "@/components/Interaction/LoadingEllipsis.vue";
import { reviewSentType, sentWord } from "@/types";
import SentenceRev from "@/components/Sentence/SentenceRev.vue";
import SentenceTried from "@/components/Sentence/SentenceTried.vue";
import TextNavButton from "@/components/Interaction/TextNavButton.vue";
import TopBar from "@/components/TopBar/TopBar.vue";

const route = useRoute();
const daysAhead = computed<number>(() => {
  // Routes that direct here are:
  // - /review or /review/ahead
  if (route.params.days) {
    return 7;
  } else {
    return 0;
  }
});

const {
  result: resultDue,
  loading: loadingDue,
  error: errorDue,
} = useQuery(dueSents, { days: daysAhead.value });

const { mutate: trySent } = useMutation(mutationTrySent);

const dueSentsArray = computed<reviewSentType[] | null>(() => {
  if (!loadingDue.value && !errorDue.value && resultDue.value) {
    return resultDue.value.displayDueSents;
  } else {
    return null;
  }
});

const wrongSentsArray = ref<reviewSentType[]>([]);
const showSummary = ref<boolean>(false);

const revIndex = ref<number>(0);
const finnsNext = computed<boolean>(() => {
  if (revIndex.value < dueSentsArray.value.length - 1) {
    return true;
  } else {
    return false;
  }
});

const submitSentHandler = (payload: {
  sentId: number;
  userInputWords: { indexInSent: number; inputText: string }[];
}) => {
  const { sentId, userInputWords } = payload;
  console.log("This is the payload we get: ", payload);
  const tryTextArray = [];
  const activeSentence = dueSentsArray.value[revIndex.value];
  let userInputCorrect = true;
  const sentWords: sentWord[] = activeSentence.sentWords;
  const wordCount = sentWords.length;

  for (let k = 1; k <= wordCount; k++) {
    const userInputWordAtK = userInputWords.find(
      (item) => item.indexInSent === k
    );
    const wordAtK = sentWords.find((item) => item.indexInSent === k) || null;

    if (userInputWordAtK) {
      tryTextArray.push(userInputWordAtK.inputText);

      if (
        wordAtK &&
        wordAtK.wordform &&
        wordAtK.wordform !== userInputWordAtK.inputText
      ) {
        console.log("wordAtK.wordform", wordAtK.wordform);
        console.log("userInputWordAtK", userInputWordAtK);
        userInputCorrect = false;
      }
    } else {
      tryTextArray.push(
        sentWords.find((word) => word.indexInSent === k).wordform || ""
      );
    }
  }

  if (!userInputCorrect) {
    const wrongSent = constructNew(tryTextArray, activeSentence);
    wrongSentsArray.value.push(wrongSent);
  }

  trySent({
    sentId: sentId,
    userInputJson: JSON.stringify(tryTextArray),
  });

  if (finnsNext.value) {
    revIndex.value++;
  } else {
    showSummary.value = true;
  }
};

const constructNew = (tryTextArray, activeSentence) => {
  const wrongSentWords: sentWord[] = [];
  activeSentence.sentWords.forEach((word: sentWord, index: number) => {
    const newWord = { ...word };
    newWord.lastInputText = tryTextArray[index];
    wrongSentWords.push(newWord);
  });

  const wrongSent = { ...activeSentence };
  wrongSent.sentWords = wrongSentWords;
  return wrongSent;
};
</script>

<template>
  <template v-if="dueSentsArray && dueSentsArray.length > 0">
    <top-bar
      state="review"
      :duecount="dueSentsArray.length"
      :revcount="revIndex + 1"
    />
  </template>
  <template v-else>
    <top-bar state="heading" heading="review" />
  </template>

  <div class="review-view-container">
    <div v-if="loadingDue" style="text-align: center">
      <loading-ellipsis />
    </div>
    <template v-if="dueSentsArray && !showSummary">
      <template v-for="(item, index) in dueSentsArray" :key="item.sentId">
        <template v-if="index === revIndex">
          <sentence-rev
            :sent-id="item.sentId"
            :media-uri="item.mediaUri"
            :sent-words="item.sentWords"
            :finns-next="finnsNext"
            @submit-sent="submitSentHandler"
          />
        </template>
      </template>
    </template>
    <template v-if="showSummary">
      <img
        src="@/assets/buttons/circledStar.svg"
        alt="Star"
        class="congrats-star"
      />
      <h2>Congrats!</h2>
      <p>You've finished reviewing due sentences today.</p>
      <template
        v-if="wrongSentsArray && wrongSentsArray.length > 0 && showSummary"
      >
        <div class="rev-list-of-errors">
          <p>Here are some mistakes:</p>
          <div
            v-for="(item, index) in wrongSentsArray"
            :key="item.sentId"
            class="rev-list-of-errors"
          >
            <SentenceTried
              :is-correct="false"
              :is-summary="false"
              :sent-id="item.sentId"
              :index-in-article="index + 1"
              :sent-words="item.sentWords"
            />
          </div>
        </div>
      </template>
      <div class="complete-button">
        <text-nav-button href="/articles">complete</text-nav-button>
      </div>
    </template>
    <template v-if="!dueSentsArray || dueSentsArray.length === 0">
      <img
        src="@/assets/buttons/circledStar.svg"
        alt="Star"
        class="congrats-star"
      />
      <h2>Congrats!</h2>
      <p>You don't have due sentences to review now.</p>
      <div class="complete-button">
        <text-nav-button href="/articles">complete</text-nav-button>
      </div>
    </template>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.review-view-container {
  margin-top: 100px;
  padding-right: 20%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.congrats-star {
  width: 150px;
  height: 150px;
}

h2,
p {
  font-weight: bold;
  color: $blue-primary;
}

.complete-button {
  margin-top: 30px;
}

.rev-list-of-errors {
  max-width: 500px;
}
</style>
