<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import ArticleItem from "../components/ArticleItem/ArticleItem.vue";
import LoadingEllipsis from "@/components/Interaction/LoadingEllipsis.vue";
import { useQuery } from "@vue/apollo-composable";
import { userArticleType } from "../types";
import { seenArticles, unseenArticles } from "../graphql";
import { useRouter } from "vue-router";
import TopBar from "@/components/TopBar/TopBar.vue";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
const router = useRouter();

onMounted(() => {
  navWidthStore.$patch({ wide: false });
});

const {
  result: resultSeen,
  loading: loadingSeen, // TODO: make a svg to indicate loading
  // error: errorSeen, // TODO: handle these possible errors
} = useQuery(seenArticles, null, {
  // You need to explicitly state variables are null.
  fetchPolicy: "network-only",
});

const {
  result: resultUnseen,
  loading: loadingUnseen,
  // error: errorUnseen,
} = useQuery(unseenArticles, null, {
  fetchPolicy: "network-only",
});

const navLearnHandler = (data: number) => {
  router.push(`/articles/${data}/learn`);
};

const navSummaryHandler = (data: number) => {
  router.push(`/articles/${data}/summary`);
};

// eslint-disable-next-line vue/return-in-computed-property
const continueSentsArray = computed(() => {
  if (resultSeen.value && resultSeen.value.listUserArticles) {
    return resultSeen.value.listUserArticles.filter(
      (art) => art.userFinishedIndex < art.articleSentCount
    );
  } else {
    return [];
  }
});

// eslint-disable-next-line vue/return-in-computed-property
const finishedSentsArray = computed(() => {
  if (resultSeen.value && resultSeen.value.listUserArticles) {
    return resultSeen.value.listUserArticles.filter(
      (art) => art.userFinishedIndex === art.articleSentCount
    );
  } else {
    return [];
  }
});
</script>

<template>
  <top-bar state="articles" />
  <div
    class="view-container"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <div
      v-if="continueSentsArray && continueSentsArray.length > 0"
      class="articles-progress-section"
    >
      <h2>continue...</h2>
      <div v-for="art in continueSentsArray" :key="art.articleId">
        <ArticleItem
          :id="art.articleId"
          :title="art.articleTitle"
          :description="art.articleDescription"
          :progress="art.userFinishedIndex"
          :goal="art.articleSentCount"
          @go-to-article="navLearnHandler"
        />
      </div>
    </div>
    <div
      v-if="
        resultUnseen &&
        resultUnseen.listUserUnseenArticles &&
        resultUnseen.listUserUnseenArticles.length > 0
      "
      class="articles-progress-section"
    >
      <h2>new...</h2>
      <div
        v-for="art in resultUnseen.listUserUnseenArticles"
        :key="art.articleId"
      >
        <template v-if="art.userFinishedIndex === 0">
          <ArticleItem
            :id="art.articleId"
            :title="art.articleTitle"
            :description="art.articleDescription"
            :progress="art.userFinishedIndex"
            :goal="art.articleSentCount"
            @go-to-article="navLearnHandler"
          />
        </template>
      </div>
    </div>
    <div
      v-if="finishedSentsArray && finishedSentsArray.length > 0"
      class="articles-progress-section"
    >
      <h2>finished...</h2>
      <div v-for="art in finishedSentsArray" :key="art.articleId">
        <ArticleItem
          :id="art.articleId"
          :title="art.articleTitle"
          :description="art.articleDescription"
          :progress="art.userFinishedIndex"
          :goal="art.articleSentCount"
          @go-to-article="navSummaryHandler"
        />
      </div>
    </div>
    <div v-if="loadingSeen || loadingUnseen" style="text-align: center">
      <loading-ellipsis />
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.view-container {
  height: calc(100vh - $topbar-height);
  position: fixed;
  top: $topbar-height;

  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  overflow-y: scroll;
}

.articles-progress-section {
  margin-top: 20px;
}
</style>
