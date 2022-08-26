<script setup lang="ts">
import { useRoute } from "vue-router";
import { computed, ref } from "vue";
import {
  unseenSents,
  seenSents,
  singleArticle,
  mutationTrySent,
} from "../graphql";
import { useQuery, useMutation } from "@vue/apollo-composable";
import SentenceTried from "../components/Sentence/SentenceTried.vue";
import SentenceNew from "../components/Sentence/SentenceNew.vue";
import SentenceTrying from "../components/Sentence/SentenceTrying.vue";
import ArrowYellow from "../components/Interaction/ArrowYellow.vue";
import PlayPause from "@/components/Interaction/PlayPause.vue";
import CompleteCircle from "@/components/Interaction/CompleteCircle.vue";
import LoadingEllipsis from "@/components/Interaction/LoadingEllipsis.vue";
import TextNavButton from "@/components/Interaction/TextNavButton.vue";
import TopBar from "@/components/TopBar/TopBar.vue";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
const route = useRoute();
const articleId = computed<string>(() => {
  if (typeof route.params.id === "string") {
    return route.params.id;
  } else {
    return "0";
  }
});

const {
  result: unseenResult,
  loading: unseenLoading,
  // error: unseenError,
} = useQuery(unseenSents, {
  articleId: articleId.value,
});

const {
  result: seenResult,
  loading: seenLoading,
  // error: seenError,
} = useQuery(seenSents, {
  articleId: articleId.value,
});

const {
  result: articleResult,
  loading: articleLoading,
  // error: articleError,
} = useQuery(singleArticle, {
  articleId: articleId.value,
});

const { mutate: trySent } = useMutation(mutationTrySent, {
  refetchQueries: ["displaySeenSents", "displayUnseenSents", "getUserArticle"],
});

const arrowClickCounter = ref(0);
const arrowClickHandler = () => {
  arrowClickCounter.value++;
};

const spacebarPressCount = ref(0);
const spacebarPressHandler = () => {
  spacebarPressCount.value++;
};

const thisArticle = computed(() => {
  if (articleResult.value && articleResult.value.getUserArticle) {
    // console.log(articleResult.value.getUserArticle);
    return articleResult.value.getUserArticle;
  } else {
    return null; // THIS MIGHT BE PROBLEMATIC.
  }
});

const learningIndex = computed(() => {
  // console.log(thisArticle.value);
  // console.log(thisArticle.value.userFinishedIndex);
  if (thisArticle.value && thisArticle.value.userFinishedIndex) {
    return thisArticle.value.userFinishedIndex + 1;
  } else {
    return 1;
  }
});

const activeSentence = computed(() => {
  if (
    unseenResult &&
    unseenResult.value &&
    unseenResult.value.displayUnseenSents
  ) {
    return unseenResult.value.displayUnseenSents.find(
      (item) => item.indexInArticle === learningIndex.value
    );
  } else {
    return null;
  }
});

const submitSentHandler = (payload: {
  sentId: number;
  userInputWords: { indexInSent: number; inputText: string }[];
}) => {
  const { sentId, userInputWords } = payload;
  const tryTextArray = [];
  const sentWords = activeSentence.value.sentWords;
  const wordCount = sentWords.length;

  for (let k = 1; k <= wordCount; k++) {
    const userInputWordAtK = userInputWords.find(
      (item) => item.indexInSent === k
    );
    console.log(userInputWordAtK);
    if (userInputWordAtK) {
      tryTextArray.push(userInputWordAtK.inputText);
      console.log(tryTextArray);
    } else {
      tryTextArray.push(
        sentWords.find((word) => word.indexInSent === k).wordform || ""
      );
    }
  }

  trySent({
    sentId: sentId,
    userInputJson: JSON.stringify(tryTextArray),
  });
};
</script>

<template>
  <top-bar state="learn" :title="thisArticle.articleTitle ?? ''" />

  <div
    class="view-container view-container-learn"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <template v-if="seenResult && seenResult.displaySeenSents">
      <div
        v-for="sent in seenResult.displaySeenSents"
        :key="sent.sentId"
        class="learn-view-section"
      >
        <template v-if="sent.indexInArticle <= thisArticle.userFinishedIndex">
          <sentence-tried
            :is-correct="true"
            :is-summary="false"
            :sent-id="sent.sentId"
            :index-in-article="sent.indexInArticle"
            :try-text="sent.tryText"
          />
        </template>
      </div>
    </template>

    <div
      v-if="articleLoading || unseenLoading || seenLoading"
      style="text-align: center"
    >
      <loading-ellipsis />
    </div>
    <div
      v-if="
        activeSentence &&
        thisArticle.userFinishedIndex < thisArticle.articleSentCount
      "
      class="learn-view-section"
    >
      <sentence-trying
        @play-sound="spacebarPressHandler"
        @submit-sent="submitSentHandler"
        :sent-id="activeSentence.sentId"
        :index-in-article="activeSentence.indexInArticle"
        :sent-words="activeSentence.sentWords"
        :parent-arrow-click="arrowClickCounter"
      />
    </div>

    <template v-if="unseenResult && unseenResult.displayUnseenSents">
      <div
        v-for="sent in unseenResult.displayUnseenSents"
        :key="sent.sentId"
        class="learn-view-section"
      >
        <template v-if="sent.indexInArticle > learningIndex">
          <sentence-new
            :sent-id="sent.sentId"
            :index-in-article="sent.indexInArticle"
            :sent-words="sent.sentWords"
          />
        </template>
      </div>
    </template>

    <template
      v-if="
        activeSentence &&
        activeSentence.mediaUri &&
        thisArticle.userFinishedIndex < thisArticle.articleSentCount
      "
      ><play-pause
        :media-url="activeSentence.mediaUri"
        :parent-click-play="spacebarPressCount"
      />
    </template>

    <arrow-yellow
      v-if="thisArticle.userFinishedIndex < thisArticle.articleSentCount"
      @click-down-arrow="arrowClickHandler"
    />

    <template
      v-if="thisArticle.userFinishedIndex >= thisArticle.articleSentCount"
    >
      <complete-circle />
      <text-nav-button :href="`/articles/${articleId}/summary`"
        >summary</text-nav-button
      >
    </template>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.view-container-learn {
  align-items: left !important;
  justify-content: left !important;
  padding-bottom: 150px;
  height: calc(100vh - $topbar-height - 100px) !important;

  @include for-mobile {
    padding-top: 20%;
  }
}

.learn-view-section {
  width: 60%;
  justify-self: flex-start !important;
  margin-left: -25vw;

  @include for-mobile {
    width: 80%;
    margin: 0 15% 0 5%;
  }
}
</style>
