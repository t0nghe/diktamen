<template>
  <div>
    <ScoreCircle :score="78" :finished="true" />
  </div>
</template>

<script lang="ts">
import { gql } from "graphql-tag";
import ScoreCircle from "@/components/Interaction/ScoreCircle.vue";
// TODO: This needs to come from routing param
const articleId = 6;

export default {
  name: "SummaryView",
  components: { ScoreCircle },
};
</script>
