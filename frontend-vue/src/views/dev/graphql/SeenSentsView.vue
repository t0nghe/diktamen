<template>
  <div>Sentences the user has seen.</div>
  <input type="number" v-model="userInput" />
  <div v-if="$apollo.loading">loading...</div>
  <div v-else>
    <div v-for="item in displaySeenSents" :key="item.sentId">
      {{ item.sentId }}
      {{ item.indexInArticle }}
      {{ item.tryText }}
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";
type articleIdType = {
  articleId: number;
};

export default {
  name: "SeenSentsView",
  data() {
    return {
      userInput: 6,
      displaySeenSents: [{ sentId: 0, indexInArticle: 0, tryText: "" }],
    };
  },
  apollo: {
    displaySeenSents: {
      query: gql`
        query displaySeenSents($articleId: Int!) {
          displaySeenSents(articleId: $articleId) {
            sentId
            indexInArticle
            tryText
          }
        }
      `,
      variables(): articleIdType {
        return {
          articleId: parseInt(this.userInput),
        };
      },
    },
  },
};
</script>

<style></style>
