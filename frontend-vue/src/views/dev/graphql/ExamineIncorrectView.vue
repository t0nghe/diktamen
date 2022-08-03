<template>
  <h3>User dictations that are NOT correct.</h3>
  <div>Sentence ID: <input type="number" v-model="userInput" /></div>
  <div v-if="$apollo.loading">loading...</div>
  <div v-else>
    <div v-for="sent in examineIncorrectSents" :key="sent.id">
      {{ sent.id }}
      {{ sent.indexInArticle }}
      <div v-if="sent.sentWords.length > 0">
        <span
          v-for="word in sent.sentWords"
          :key="word.wordform + word.indexInSent"
          style="padding: 0.2em"
        >
          {{ word.wordform }}
          <span
            v-if="word.isCloze && word.lastInputScore < 0.998"
            style="color: red"
          >
            {{ word.lastInputText }}
          </span>
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";
type articleIdType = {
  articleId: number;
};

export default {
  name: "ExamineIncorrectView",
  data() {
    return {
      userInput: 6,
    };
  },
  apollo: {
    examineIncorrectSents: {
      query: gql`
        query examineIncorrectSents($articleId: Int!) {
          examineIncorrectSents(articleId: $articleId) {
            sentId
            indexInArticle
            sentWords {
              length
              isCloze
              wordform
              indexInSent
              lastInputText
              lastInputScore
            }
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
