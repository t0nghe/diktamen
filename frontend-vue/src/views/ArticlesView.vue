<script setup lang="ts">
import ArticleItem from "../components/ArticleItem/ArticleItem.vue";
import LoadingEllipsis from "@/components/Interaction/LoadingEllipsis.vue";
import { useQuery } from "@vue/apollo-composable";
import { userArticleType } from "../types";
import { seenArticles, unseenArticles } from "../graphql";
import { useRouter } from "vue-router";
import TopBar from "@/components/TopBar/TopBar.vue";

const router = useRouter();

const {
  result: resultSeen,
  loading: loadingSeen, // TODO: make a svg to indicate loading
  // error: errorSeen, // TODO: handle these possible errors
} = useQuery(seenArticles);

const {
  result: resultUnseen,
  loading: loadingUnseen,
  // error: errorUnseen,
} = useQuery(unseenArticles);

const navLearnHandler = (data: number) => {
  router.push(`/articles/${data}/learn`);
};

const navSummaryHandler = (data: number) => {
  router.push(`/articles/${data}/summary`);
};
</script>

<template>
  <top-bar state="articles" />
  <div v-if="resultSeen && resultSeen.listUserArticles">
    <h2>continue...</h2>
    <div v-for="art in resultSeen.listUserArticles" :key="art.articleId">
      <template v-if="art.userFinishedIndex < art.articleSentCount">
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
    v-if="
      resultUnseen &&
      resultUnseen.listUserUnseenArticles &&
      resultUnseen.listUserUnseenArticles.length > 0
    "
  >
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
        @go-to-article="navLearnHandler"
      />
    </div>
  </div>
  <div v-if="resultSeen && resultSeen.listUserArticles">
    <h2>finished...</h2>
    <div v-for="art in resultSeen.listUserArticles" :key="art.articleId">
      <template v-if="art.userFinishedIndex === art.articleSentCount">
        <ArticleItem
          :id="art.articleId"
          :title="art.articleTitle"
          :description="art.articleDescription"
          :progress="art.userFinishedIndex"
          :goal="art.articleSentCount"
          @go-to-article="navSummaryHandler"
        />
      </template>
    </div>
  </div>
  <div v-if="loadingSeen || loadingUnseen" style="text-align: center">
    <loading-ellipsis />
  </div>
</template>

<style></style>
