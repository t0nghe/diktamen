<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import slides from "../components/Onboarding/";
import { ref } from "vue";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
const scrIdx = ref<number>(0);
const scrArr = [slides.SlideOne, slides.SlideTwo, slides.SlideThree];
const strArr = [
  "type out what you hear",
  "review mistakes",
  "learn again in a few days",
];

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
  <div
    class="onboarding-container"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <div class="slide-desc">{{ strArr[scrIdx] }}</div>
    <component :is="scrArr[scrIdx]" @next-screen="nextHandler"></component>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.onboarding-container {
  background: url("@/assets/onboarding-bg.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  height: calc(100vh - $topbar-height);
  position: fixed;
  top: $topbar-height;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.slide-desc {
  font-size: 1.5rem;
  color: $blue-primary;
  font-weight: 200;
  margin: 10px 0px;
}

.container-position-wide-side {
  width: calc(
    100vw - $nav-sidebar-display-width-wide - $nav-sidebar-control-width
  );
  left: $nav-sidebar-display-width-wide + $nav-sidebar-control-width;
}

.container-position-narrow-side {
  width: calc(
    100vw - $nav-sidebar-display-width-narrow - $nav-sidebar-control-width
  );
  left: $nav-sidebar-display-width-narrow + $nav-sidebar-control-width;
}
</style>
