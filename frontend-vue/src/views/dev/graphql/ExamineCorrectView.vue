<template>
  <h3>User dictations that are correct.</h3>
  <div>Sentence ID: <input type="number" v-model="userInput" /></div>
  <div v-if="$apollo.loading">loading...</div>
  <div v-else>
    <div v-for="sent in examineCorrectSents" :key="sent.id">
      {{ sent.id }}
      {{ sent.indexInArticle }}
      {{ sent.tryText }}
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";
type articleIdType = {
  articleId: number;
};

export default {
  name: "ExamineCorrectView",
  data() {
    return {
      userInput: 6,
    };
  },
  apollo: {
    examineCorrectSents: {
      query: gql`
        query examineCorrectSents($articleId: Int!) {
          examineCorrectSents(articleId: $articleId) {
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
