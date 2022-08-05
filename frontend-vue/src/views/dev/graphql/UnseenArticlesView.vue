<template>
  <div>
    <h3>Unseen Articles View</h3>
    <div v-if="$apollo.loading">loading...</div>
    <div v-else>
      <div v-for="item in listUserUnseenArticles" :key="item.id">
        {{ item.articleTitle }}<br />{{ item.articleDescription }}
        {{ item.userFinishedIndex }} / {{ item.articleSentCount }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";

export default {
  name: "UnseenArticlesView",
  data() {
    return {
      listUserUnseenArticles: [
        { articleId: 0, articleTitle: "", articleSentCount: 0, articleDescription: "", userId: 0 },
      ],
    };
  },
  apollo: {
    listUserUnseenArticles: {
      query: gql`
        query unseenArticles {
          listUserUnseenArticles {
            articleId
            articleTitle
            articleSentCount
            articleDescription
            userFinishedIndex
          }
        }
      `,
    },
  },
};
</script>

<style></style>
