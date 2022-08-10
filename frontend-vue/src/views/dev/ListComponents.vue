<script setup lang="ts">
import ProgressCircle from "../../components/ArticleItem/ProgressCircle.vue";
import ArticleItem from "../../components/ArticleItem/ArticleItem.vue";
import NavSidebar from "../../components/NavSidebar/NavSidebar.vue";
import TopBar from "../../components/TopBar/TopBar.vue";
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
const topBarDone = ref(1);
const topBarToGo = ref(10);

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
  <h2>NavSidebar</h2>
  <NavSidebar :wide="true" />
  <h2>TopBar</h2>
  <h3>TopBar - waitlist</h3>
  <TopBar state="waitlist" ahref="https://dope.com/" />
  <h3>TopBar - articles</h3>
  <TopBar state="articles" />
  <h3>TopBar - learn</h3>
  <TopBar state="learn" title="Nu kör vi!" />
  <h3>TopBar - summary</h3>
  <TopBar state="summary" />
  <h2>Dear TopBar review, aren't you a little pain dans le cul?</h2>
  <form>
    <fieldset>
      <legend>customize your progress, my friend</legend>
      <label for="topBarDone"
        >You've reviewed this many cards:
        <input type="number" v-model="topBarDone"
      /></label>
      <br />
      <label for="topBarToGo"
        >You need to review this many cards:
        <input type="number" v-model="topBarToGo"
      /></label>
    </fieldset>
  </form>
  <br /><br /><br />
  <TopBar state="review" :revcount="topBarDone" :duecount="topBarToGo" />
</template>

<style lang="scss">
h2 {
  margin-top: 4rem;
  margin-bottom: 2rem;
}
</style>
