<script setup lang="ts">
import ArticleItem from "../components/ArticleItem/ArticleItem.vue";
import { computed } from "vue";
import { useQuery } from "@vue/apollo-composable";
import { userArticleType } from "../types";
import { seenArticles, unseenArticles } from "../queries";
import { useRouter } from "vue-router";

const router = useRouter();

const {
  result: resultSeen,
  // loading: loadingSeen, // TODO: make a svg to indicate loading
  // error: errorSeen, // TODO: handle these possible errors
} = useQuery(seenArticles);

const {
  result: resultUnseen,
  // loading: loadingUnseen,
  // error: errorUnseen,
} = useQuery(unseenArticles);

const unfinishedArticles = computed<userArticleType[]>(() => {
  if (resultSeen.value && resultSeen.value.listUserArticles) {
    return resultSeen.value.listUserArticles.filter(
      (item: userArticleType) => item.userFinishedIndex <= item.articleSentCount
    );
  } else {
    return [];
  }
});

const finishedArticles = computed<userArticleType[]>(() => {
  if (resultSeen.value && resultSeen.value.listUserArticles) {
    return resultSeen.value.listUserArticles.filter(
      (item: userArticleType) =>
        item.userFinishedIndex === item.articleSentCount
    );
  } else {
    return [];
  }
});

const navHandler = (data: number) => {
  router.push(`/articles/${data}/learn`);
};
</script>

<template>
  <div v-if="unfinishedArticles">
    <h2>continue...</h2>
    <div v-for="art in unfinishedArticles" :key="art.articleId">
      <ArticleItem
        :id="art.articleId"
        :title="art.articleTitle"
        :description="art.articleDescription"
        :progress="art.userFinishedIndex"
        :goal="art.articleSentCount"
        @go-to-article="navHandler"
      />
    </div>
  </div>
  <div v-if="resultUnseen">
    <h2>new...</h2>
    <div
      v-for="art in resultUnseen.listUserUnseenArticles"
      :key="art.articleId"
    >
      <ArticleItem
        :id="art.articleId"
        :title="art.articleTitle"
        :description="art.articleDescription"
        :progress="art.userFinishedIndex"
        :goal="art.articleSentCount"
        @go-to-article="navHandler"
      />
    </div>
  </div>
  <div
    v-if="
      typeof finishedArticles !== 'undefined' && finishedArticles.length > 0
    "
  >
    <h2>you've finished...</h2>
    <div v-for="art in finishedArticles" :key="art.articleId">
      <ArticleItem
        :id="art.articleId"
        :title="art.articleTitle"
        :description="art.articleDescription"
        :progress="art.userFinishedIndex"
        :goal="art.articleSentCount"
      />
    </div>
  </div>
</template>

<style></style>
