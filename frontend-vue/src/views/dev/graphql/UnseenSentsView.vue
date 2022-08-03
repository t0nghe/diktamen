<template>
  <div>Sentences the user has NOT seen.</div>
  <input type="number" v-model="userInput" />
  <div v-if="$apollo.loading">loading...</div>
  <div v-else>
    <div v-for="sent in displayUnseenSents" :key="sent.mediaUri">
      {{ sent.sentId }}
      {{ sent.indexInArticle }}
      {{ sent.mediaUri }}
      <div v-if="sent.sentWords.length > 0">
        <span
          v-for="word in sent.sentWords"
          :key="word.length + word.indexInSent"
          style="padding: 0.2em"
        >
          {{ word.wordform }}
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";

export default {
  name: "UnseenSentsView",
  data() {
    return {
      userInput: 6,
    };
  },
  apollo: {
    displayUnseenSents: {
      query: gql`
        query displayUnseenSents($articleId: Int!) {
          displayUnseenSents(articleId: $articleId) {
            sentId
            indexInArticle
            mediaUri
            sentWords {
              length
              isCloze
              wordform
              indexInSent
            }
          }
        }
      `,
      variables(): { articleId: number } {
        return {
          articleId: parseInt(this.userInput),
        };
      },
    },
  },
};
</script>

<style></style>
