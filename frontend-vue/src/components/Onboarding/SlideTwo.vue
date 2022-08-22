<script setup lang="ts">
import ArrowOnboarding from "./ArrowBlueOnboarding.vue";
import SentenceTried from "@/components/Sentence/SentenceTried.vue";
import { sent1, sent2words, sent3words, sent4text } from "./data.ts";
const emit = defineEmits<{ (e: "next-screen") }>();
</script>

<template>
  <div class="onboarding-content onboarding-screen2">
    <div class="onboarding-screen2-left">
      <sentence-tried
        :sent-id="1"
        :index-in-article="1"
        :is-summary="true"
        :is-correct="true"
        :try-text="sent1"
      />
      <sentence-tried
        :sent-id="2"
        :index-in-article="2"
        :is-summary="true"
        :is-correct="false"
        :sent-words="sent2words"
      />
      <sentence-tried
        :sent-id="3"
        :index-in-article="3"
        :is-summary="true"
        :is-correct="false"
        :sent-words="sent3words"
      />
      <sentence-tried
        :sent-id="4"
        :index-in-article="4"
        :is-summary="true"
        :is-correct="true"
        :try-text="sent4text"
      />
    </div>
    <div class="onboarding-screen2-right">
      <div class="onboarding-score-circle">91</div>
    </div>
    <div class="onboarding-screen2-bottom arrow-onboarding">
      <arrow-onboarding @click="emit('next-screen')" />
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";
/* onboarding-content and arrow-onboarding are defined in `slideone` */
.onboarding-screen2 {
  display: grid;
  grid-auto-columns: auto 250px;
  grid-auto-rows: auto 95px;
  gap: 20px;

  .onboarding-screen2-left {
    grid-column: 1;
    grid-row: 1;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    padding: 50px;
  }

  .onboarding-screen2-right {
    grid-column: 2;
    grid-row: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;

    .onboarding-score-circle {
      width: 220px;
      height: 220px;
      border: 14px solid $blue-secondary;
      border-radius: 50%;
      display: flex;
      justify-content: center;
      align-items: center;

      font-size: 6rem;
      color: $blue-secondary;
      font-weight: bold;
    }
  }

  .onboarding-screen2-bottom {
    grid-column: 1/3;
    grid-row: 2;
    /* centering the downward arrow */
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
