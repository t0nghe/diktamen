<script setup lang="ts">
import ProgressCircle from "../../components/ArticleItem/ProgressCircle.vue";
import ArticleItem from "../../components/ArticleItem/ArticleItem.vue";
import { ref } from "vue";
const pcPrimary = ref("#00bcd4");
const pcSecondary = ref("#dcbc00");
const pcProg = ref(1);
const pcGoal = ref(10);
const articleItemData = {
  articleId: 6,
  articleTitle: "Zelenskyj är en man som kvinnor längtar efter",
  articleSentCount: 17,
  articleDescription:
    "När han nu väl tar plats, den traditionelle mannen, väcker han både kvinnors åtrå och mäns önskan att vara som han.",
  userFinishedIndex: 3,
};

const eventHandler = (data: number) => {
  alert("clicked");
  console.log(data);
};
</script>

<template>
  <h1>LIST OF COMPONENTS</h1>
  <hr />
  <h2>ProgressCircle</h2>
  Note: We still need to change the size.
  <form>
    <fieldset>
      <legend>Progress Circle</legend>
      <label for="pc_primary"
        >Primary color:
        <input
          type="color"
          id="pc_primary"
          name="primary"
          v-model="pcPrimary"
        /> </label
      ><br />
      <label for="pc_secondary"
        >Secondary color:
        <input
          type="color"
          id="pc_secondary"
          name="secondary"
          v-model="pcSecondary"
        /> </label
      ><br />
      <label for="pc_prog"
        >Progress:
        <input
          type="number"
          id="pc_prog"
          name="prog"
          v-model="pcProg"
        /> </label
      ><br />
      <label for="pc_goal"
        >Goal:
        <input
          type="number"
          id="pc_goal"
          name="goal"
          v-model="pcGoal"
        /> </label
      ><br />
    </fieldset>
  </form>
  <ProgressCircle
    :prog="pcProg"
    :goal="pcGoal"
    :primary="pcPrimary"
    :secondary="pcSecondary"
    desc-bg-color="white"
  />
  <hr />
  <h2>ArticleItem - in progress (trying click handler)</h2>
  <ArticleItem
    :id="articleItemData.articleId"
    :title="articleItemData.articleTitle"
    :description="articleItemData.articleDescription"
    :goal="articleItemData.articleSentCount"
    :progress="articleItemData.userFinishedIndex"
    @go-to-article="eventHandler"
  />
  <h2>ArticleItem - complete</h2>
  <ArticleItem
    :id="articleItemData.articleId"
    :title="articleItemData.articleTitle"
    :description="articleItemData.articleDescription"
    :goal="10"
    :progress="10"
  />
  <h2>ArticleItem - new</h2>
  <ArticleItem
    :id="articleItemData.articleId"
    :title="articleItemData.articleTitle"
    :description="articleItemData.articleDescription"
    :goal="10"
    :progress="0"
  />
  <h2>...</h2>
</template>

<style lang="scss">
h2 {
  margin-top: 4rem;
  margin-bottom: 2rem;
}
</style>
