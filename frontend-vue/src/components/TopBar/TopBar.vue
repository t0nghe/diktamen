<script setup lang="ts">
import { ref } from "vue";
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
    heading.value = "";
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
    <div class="top-bar-logo">
      <router-link to="/">
        <img alt="logo" src="@/assets/logo.svg" class="top-bar-logo-image" />
      </router-link>
    </div>
    <template v-if="props.state === 'review'">
      <ReviewProgress :revcount="props.revcount" :duecount="props.duecount" />
    </template>
    <div v-else class="top-bar-heading">{{ heading }}</div>
    <div>
      <!-- empty div to arrange -->
    </div>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.top-bar-waitlist {
  background-color: $blue-secondary;
  color: $blue-primary;

  a {
    text-decoration: none;
    color: $blue-primary;
    cursor: pointer;
    font-weight: bold;
  }
}

.top-bar-articles {
  background-color: $azure-primary;
  color: $blue-primary;

  .top-bar-logo {
    background: linear-gradient(
      to right,
      $blue-primary 0 50%,
      $azure-primary 50% 100%
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
  background-color: $blue-secondary;
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

  @include for-mobile {
    height: $mobile-topbar-height;
  }

  .top-bar-heading {
    font-size: 1.5rem;
    font-weight: bold;
    margin-left: -10vw;

    @include for-mobile {
      font-size: 1.2rem;
      font-weight: 500;
    }
  }

  .top-bar-logo {
    background: linear-gradient(
      to right,
      $blue-primary 0 30%,
      transparent 30% 100%
    );
    height: 100%;

    .top-bar-logo-image {
      height: $topbar-height;
      width: $logo-width;

      @include for-mobile {
        height: $mobile-topbar-height;
        width: $mobile-logo-width;
      }
    }
  }
}
</style>
