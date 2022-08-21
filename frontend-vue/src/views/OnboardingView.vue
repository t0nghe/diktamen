<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import screens from "../components/Onboarding/";
import { ref } from "vue";

const scrIdx = ref<number>(0);
const scrArr = [screens.ScreenOne, screens.ScreenTwo, screens.ScreenThree];

const prevHandler = () => {
  if (scrIdx.value === 0) {
    scrIdx.value = 2;
  } else {
    scrIdx.value--;
  }
};

const nextHandler = () => {
  if (scrIdx.value === 2) {
    scrIdx.value = 0;
  } else {
    scrIdx.value++;
  }
};
</script>

<template>
  <top-bar state="waitlist" ahref="https://forms.gle/nvc46wgCnDTZty3d6" />
  <component :is="scrArr[scrIdx]"></component>
  <button @click="prevHandler">Previous</button>
  <button @click="nextHandler">Next</button>
  <p>Go to <router-link to="/test-signup">Sign up</router-link></p>
</template>
