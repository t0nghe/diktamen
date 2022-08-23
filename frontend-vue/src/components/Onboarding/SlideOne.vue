<script setup lang="ts">
import { ref, onMounted } from "vue";
import PlayPauseRev from "@/components/Interaction/PlayPauseRev.vue";
import SentenceTried from "@/components/Sentence/SentenceTried.vue";
import SentenceTrying from "@/components/Sentence/SentenceTrying.vue";
import SentenceNew from "@/components/Sentence/SentenceNew.vue";
import ArrowOnboarding from "./ArrowBlueOnboarding.vue";
import { sent1, sent2, sent3words, sent3audio, sent4words } from "./data.ts";

const playPauseCounter = ref(0);
const downArrowClicked = ref(0);

const emit = defineEmits<{ (e: "next-screen") }>();

onMounted(() => {
  downArrowClicked.value++;
});
</script>

<template>
  <div class="onboarding-content onboarding-screen1">
    <div class="onboarding-screen1-left">
      <sentence-tried
        :is-correct="true"
        :is-summary="false"
        :sent-id="1"
        :index-in-article="1"
        :try-text="sent1"
      />
      <sentence-tried
        :is-correct="true"
        :is-summary="false"
        :sent-id="2"
        :index-in-article="2"
        :try-text="sent2"
      />
      <sentence-trying
        :sent-id="3"
        :index-in-article="3"
        :sent-words="sent3words"
        :parent-arrow-click="downArrowClicked"
        @play-sound="playPauseCounter++"
        @submit-sent="emit('next-screen')"
      />
      <sentence-new
        :sent-id="4"
        :index-in-article="4"
        :sent-words="sent4words"
      />
    </div>
    <div class="onboarding-screen1-right">
      <play-pause-rev
        :media-uri="sent3audio"
        :parent-play-pause="playPauseCounter"
        :dimension="200"
      />
    </div>
    <div class="onboarding-screen1-bottom arrow-onboarding">
      <arrow-onboarding @click="emit('next-screen')" />
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.onboarding-content {
  padding: 30px 30px;
  background-color: rgba($yellow-beige, 0.85);
  height: 60vh;
  min-height: 480px;
  width: 60vw;
  border: 2px solid $yellow-canary;
  box-shadow: 5px 5px $yellow-gold;

  .arrow-onboarding {
    button {
      cursor: pointer;
      border: none;
      cursor: pointer;
      background: transparent;
    }
  }
}

.onboarding-screen1 {
  display: grid;
  grid-auto-columns: auto 250px;
  grid-auto-rows: auto 95px;
  gap: 20px;

  .onboarding-screen1-left {
    grid-column: 1;
    grid-row: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
  }

  .onboarding-screen1-right {
    grid-column: 2;
    grid-row: 1;

    display: flex;
    justify-content: center;
    align-items: center;
  }

  .onboarding-screen1-bottom {
    grid-column: 1/3;
    grid-row: 2;
    /* centering the downward arrow */
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
