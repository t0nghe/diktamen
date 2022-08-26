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
        :sent-id="2"
        :index-in-article="1"
        :is-summary="true"
        :is-correct="false"
        :sent-words="sent2words"
      />
      <sentence-tried
        :sent-id="3"
        :index-in-article="2"
        :is-summary="true"
        :is-correct="false"
        :sent-words="sent3words"
      />
      <sentence-tried
        :sent-id="4"
        :index-in-article="3"
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
  grid-auto-columns: auto 220px;
  grid-auto-rows: auto 85px;
  gap: 10px;

  @include for-mobile {
    display: flex;
    position: relative;
  }

  .onboarding-screen2-left {
    grid-column: 1;
    grid-row: 1;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    padding: 20px;

    @include for-mobile {
      padding-top: calc((35% + 45px) / 2);
      padding-left: 5%;
      padding-right: 10%;
      z-index: 3;
    }
  }

  .onboarding-screen2-right {
    grid-column: 2;
    grid-row: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;

    @include for-mobile {
      position: absolute;
      top: 5%;
      left: calc(50% - 40px);
      z-index: 2;
    }

    .onboarding-score-circle {
      width: 200px;
      height: 200px;
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

    @media (min-width: 401px) and (max-width: 1023px) {
      position: absolute;
      top: calc(100% - 65px);
      left: calc(50% - 20px);
      z-index: 3;
    }

    @media (max-width: 400px) {
      position: absolute;
      top: calc(100% - 75px);
      left: calc(50% - 25px);
      z-index: 3;
    }
  }
}
</style>
