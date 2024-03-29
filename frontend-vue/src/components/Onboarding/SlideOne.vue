<script setup lang="ts">
import { ref, onMounted, inject } from "vue";
import PlayPauseRev from "@/components/Interaction/PlayPauseRev.vue";
import SentenceTried from "@/components/Sentence/SentenceTried.vue";
import SentenceTrying from "@/components/Sentence/SentenceTrying.vue";
import SentenceNew from "@/components/Sentence/SentenceNew.vue";
import ArrowOnboarding from "./ArrowBlueOnboarding.vue";
import { sent2, sent3words, sent3audio, sent4words } from "./data.ts";

const playPauseCounter = ref(0);
const downArrowClicked = ref(0);

const emit = defineEmits<{ (e: "next-screen") }>();
</script>

<template>
  <div class="onboarding-content onboarding-screen1">
    <div class="onboarding-screen1-left">
      <sentence-tried
        :is-correct="true"
        :is-summary="false"
        :sent-id="2"
        :index-in-article="1"
        :try-text="sent2"
      />
      <!-- {{ $screen.width }} -->
      <sentence-trying
        :sent-id="3"
        :index-in-article="2"
        :sent-words="sent3words"
        :parent-arrow-click="downArrowClicked"
        @play-sound="playPauseCounter++"
        @submit-sent="emit('next-screen')"
      />
      <sentence-new
        :sent-id="4"
        :index-in-article="3"
        :sent-words="sent4words"
      />
    </div>
    <div
      class="onboarding-screen1-right-overlay"
      @click="playPauseCounter++"
    ></div>
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
  padding: 20px 20px;
  background-color: rgba($yellow-beige, 0.85);
  height: 60vh;
  min-height: 430px;
  width: 60vw;
  border: 2px solid $yellow-canary;
  box-shadow: 5px 5px $yellow-gold;

  @media (min-width: 401px) and (max-width: 1023px) {
    width: 85vw;
    height: 35vh;
    padding: 10px;
    background-color: rgba($yellow-beige, 0.75);
  }

  @media (max-width: 400px) {
    width: 95vw;
    height: 50%;
    padding: 5px;
  }

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

  @include for-mobile {
    display: flex;
    position: relative;
  }

  .onboarding-screen1-left {
    grid-column: 1;
    grid-row: 1/3;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;

    @media (min-width: 401px) and (max-width: 1023px) {
      padding-top: 35px;
      padding-left: 5%;
      padding-right: 10%;
      z-index: 3;
    }

    @media (max-width: 400px) {
      padding-top: 25px;
      padding-left: 3%;
      padding-right: 5%;
      z-index: 3;
    }
  }

  .onboarding-screen1-right-overlay {
    @include for-mobile {
      cursor: pointer;
      position: absolute;
      top: 0%;
      left: calc(50% - 20px);
      z-index: 4;
      width: 160px;
      height: 160px;
    }
  }

  .onboarding-screen1-right {
    grid-column: 2;
    grid-row: 1;
    display: flex;
    justify-content: center;
    align-items: center;

    @include for-mobile {
      position: absolute;
      top: 0%;
      left: calc(50% - 20px);
      z-index: 2;
    }
  }

  .onboarding-screen1-bottom {
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
