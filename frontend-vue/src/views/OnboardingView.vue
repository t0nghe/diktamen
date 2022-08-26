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
  <top-bar state="waitlist" />
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

  @include for-mobile {
    position: fixed;
    top: $mobile-topbar-height;
    height: calc(100% - $mobile-topbar-height - $mobile-navbar-height);
  }
}

.slide-desc {
  font-size: 1.5rem;
  color: $blue-primary;
  margin: 10px 0px;
  text-shadow: 2px 2px 4px $azure-secondary, -2px -2px 4px $azure-secondary;
}

.container-position-wide-side {
  width: calc(
    100vw - $nav-sidebar-display-width-wide - $nav-sidebar-control-width
  );
  left: $nav-sidebar-display-width-wide + $nav-sidebar-control-width;

  @include for-mobile {
    width: 100vw;
    left: 0px;
  }
}

.container-position-narrow-side {
  width: calc(
    100vw - $nav-sidebar-display-width-narrow - $nav-sidebar-control-width
  );
  left: $nav-sidebar-display-width-narrow + $nav-sidebar-control-width;

  @include for-mobile {
    width: 100vw;
    left: 0px;
  }
}
</style>
