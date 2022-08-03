<template>
  <div>These are the cards that's due for review.</div>
  <input type="text" v-model="userInput" />
  <div v-if="$apollo.loading">loading...</div>
  <div v-else>
    <div v-for="item in displayDueSents" :key="item.sentId">
      {{ item.sentId }}
      {{ item.mediaUri }}
    </div>
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";

export default {
  name: "ReviewDueView",
  data() {
    return {
      userInput: 6,
    };
  },
  apollo: {
    displayDueSents: {
      query: gql`
        query displayDueSents($daysAhead: Int) {
          displayDueSents(daysAhead: $daysAhead) {
            sentId
            mediaUri
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
      variables(): { daysAhead: number } {
        return {
          daysAhead: parseInt(this.userInput),
        };
      },
    },
  },
};
</script>

<style></style>
