<script setup lang="ts">
import ProgressCircle from "../../components/ArticleItem/ProgressCircle.vue";
import ArticleItem from "../../components/ArticleItem/ArticleItem.vue";
import NavSidebar from "../../components/NavSidebar/NavSidebar.vue";
import TopBar from "../../components/TopBar/TopBar.vue";
import CompleteCircle from "@/components/Interaction/CompleteCircle.vue";
import TextNavButton from "@/components/Interaction/TextNavButton.vue";
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
  <div style="text-align: center">
    <TextNavButton href="/articles/6/learn">Hello</TextNavButton>
  </div>
</template>

<style lang="scss">
h2 {
  margin-top: 4rem;
  margin-bottom: 2rem;
}
</style>
