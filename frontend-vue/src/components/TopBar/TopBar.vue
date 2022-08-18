<script setup lang="ts">
import { ref } from "vue";
import logoWide from "../../assets/logo-wide.svg";
import ReviewProgress from "./ReviewProgress.vue";

const props = defineProps<{
  state: "waitlist" | "review" | "articles" | "learn" | "summary" | "heading";
  title?: string | null;
  revcount?: number | null;
  duecount?: number | null;
  ahref?: string | null;
  heading?: string | null;
}>();

const className = ref("");
const heading = ref("");

switch (props.state) {
  case "waitlist":
    className.value = "top-bar-waitlist";
    heading.value = "join waitlist";
    break;
  case "review":
    className.value = "top-bar-review";
    heading.value = "review";
    break;
  case "articles":
    className.value = "top-bar-articles";
    heading.value = "learning progress";
    break;
  case "learn":
    className.value = "top-bar-learn";
    heading.value = props.title ? `learning: ${props.title}` : "learning...";
    break;
  case "summary":
    className.value = "top-bar-summary";
    heading.value = "summary";
    break;
  case "heading":
    className.value = "top-bar-summary";
    heading.value = props.heading ?? "";
    break;
  default:
    className.value = "top-bar-summary";
    heading.value = "";
}
</script>

<template>
  <div class="top-bar-wrapper" :class="className">
    <div class="top-bar-logo"><img :src="logoWide" alt="logo" /></div>
    <template v-if="props.state !== 'waitlist'">
      <template v-if="props.state === 'review'">
        <ReviewProgress :revcount="props.revcount" :duecount="props.duecount" />
      </template>
      <div v-else class="top-bar-heading">{{ heading }}</div>
    </template>
    <div v-else class="top-bar-heading">
      <a :href="props.ahref ? props.ahref : ''">{{ heading }}</a>
    </div>
    <div>
      <!-- empty div to arrange -->
    </div>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.top-bar-waitlist {
  background-color: $azure-primary;
  color: $blue-primary;

  a {
    text-decoration: none;
    color: $blue-primary;
    cursor: pointer;
    font-weight: bold;
  }
}

.top-bar-articles {
  background-color: $yellow-cream;
  color: $blue-primary;

  .top-bar-logo {
    /* background-color: $blue-primary !important; */
    background: linear-gradient(
      to right,
      $blue-primary 0 50%,
      $blue-secondary 50% 100%
    ) !important;
    text-align: left;
    padding-right: 4px;
  }
}

.top-bar-learn {
  background-color: $azure-primary;
  color: $blue-primary;
}

.top-bar-summary {
  background-color: $blue-primary;
  color: $yellow-cream;
}

.top-bar-review {
  background: linear-gradient(
    to right,
    $blue-secondary 0 $logo-width,
    transparent $logo-width 100%
  );
  color: $yellow-cream;
}

.top-bar-wrapper {
  height: $topbar-height;
  width: 100vw;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 5;

  .top-bar-heading {
    font-size: 1.5rem;
    font-weight: bold;
  }

  .top-bar-logo {
    background: linear-gradient(
      to right,
      $blue-primary 0 50%,
      transparent 50% 100%
    );
    height: 100%;

    img {
      height: $topbar-height;
      width: $logo-width;
    }
  }
}
</style>
