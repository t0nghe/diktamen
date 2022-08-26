<script setup lang="ts">
import { ref } from "vue";
import SentenceRev from "@/components/Sentence/SentenceRev.vue";
import { sent3words, sent3audio } from "./data.ts";

const showCtaRef = ref(false);

const showCta = () => {
  showCtaRef.value = true;
};
</script>

<template>
  <div class="onboarding-content onboarding-screen3">
    <sentence-rev
      :sent-id="3"
      :media-uri="sent3audio"
      :sent-words="sent3words"
      :finns-next="false"
      @submit-sent="showCta"
    />
  </div>
  <template v-if="showCtaRef">
    <div class="slide-desc">
      <router-link to="/test-signup">sign up</router-link> to try it out
    </div>
  </template>
</template>

<style lang="scss">
@import "@/assets/variables";

.onboarding-screen3 {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
