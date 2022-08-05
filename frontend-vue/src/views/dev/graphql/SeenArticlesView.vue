<template>
  <div>
    <h3>Seen Articles View</h3>
    <div v-if="$apollo.loading">loading...</div>
    <div v-else>
      <div v-for="item in listUserArticles" :key="item.id">
        {{ item.articleTitle }}<br />{{ item.articleDescription }}
        {{ item.userFinishedIndex }} / {{ item.articleSentCount }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";
// import { seenArticleType } from '../types'

export default {
  name: "SeenArticlesView",
  data() {
    return {
      listUserArticles: [
        {
          articleId: 0,
          articleTitle: "",
          articleSentCount: 0,
          articleDescription: "",
          userId: 0,
        },
      ],
    };
  },
  apollo: {
    listUserArticles: {
      query: gql`
        query listUserArticles {
          listUserArticles {
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
