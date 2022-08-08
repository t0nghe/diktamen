<script setup lang="ts">
import { ref, computed, ComputedRef } from "vue";
import variables from "@/assets/variables.ts";
import ProgressCircle from "./ProgressCircle.vue";

const props = defineProps<{
  id: number;
  title: string;
  description: string;
  progress: number;
  goal: number;
}>();

const progCircleForProgress = {
  primary: variables.yellowGold,
  secondary: variables.blueSecondary, // change to yellow (or not??)
  descBgColor: variables.bluePrimary,
};

const progCircleForNew = {
  primary: variables.yellowGold,
  secondary: variables.azureSecondary,
  descBgColor: variables.yellowPale,
};

const progCircleForComplete = {
  primary: variables.yellowCanary,
  secondary: variables.azurePrimary,
  descBgColor: variables.blueSecondary,
};

const status = ref("");
const className = ref("");

if (props.progress === 0) {
  status.value = "new";
  className.value = "article-item-new";
} else if (props.progress === props.goal) {
  status.value = "complete";
  className.value = "article-item-complete";
} else {
  status.value = "progress";
  className.value = "article-item-progress";
}

const progStyling: ComputedRef<{
  primary: string;
  secondary: string;
  descBgColor: string;
}> = computed(() => {
  if (status.value === "new") {
    return progCircleForNew;
  } else if (status.value === "complete") {
    return progCircleForComplete;
  } else {
    return progCircleForProgress;
  }
});

const emit = defineEmits(["go-to-article"]);
const clickHandler = () => {
  if (status.value !== "complete") {
    emit("go-to-article", props.id);
  }
};
</script>

<template>
  <div class="article-item" :class="className">
    <div class="article-item__summary">
      <div class="article-item__title" @click="clickHandler">
        {{ title }}
      </div>
      <div class="article-item__description">
        {{ description }}
      </div>
    </div>
    <div class="article-item__progress">
      <ProgressCircle
        :goal="goal"
        :prog="progress"
        :primary="progStyling.primary"
        :secondary="progStyling.secondary"
        :desc-bg-color="progStyling.descBgColor"
      />
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.article-item {
  /* display and size */
  width: 1130px;
  height: auto;
  display: flex;
  flex-direction: row;
  align-content: space-between;

  /* color: changing depending on state of article */
  &.article-item-progress {
    border: 1px solid $blue-primary;
    box-shadow: 10px 10px 0px 0px $blue-secondary,
      10px 10px 0px 2px $blue-primary;
    background-color: $blue-primary;

    .article-item__title {
      color: $yellow-canary;
      border-bottom: 2px solid $yellow-cream;
      cursor: pointer;
    }

    .article-item__description {
      color: $yellow-canary;
    }
  }

  &.article-item-complete {
    border: 1px solid $blue-secondary;
    box-shadow: 10px 10px 0px 0px $blue-primary,
      10px 10px 0px 2px $blue-secondary;
    background-color: $blue-secondary;

    .article-item__title {
      color: $yellow-canary;
      border-bottom: 2px solid $yellow-cream;
      cursor: auto;
    }

    .article-item__description {
      color: $yellow-canary;
    }
  }

  &.article-item-new {
    border: 1px solid $yellow-gold;
    box-shadow: 10px 10px 0px 0px $yellow-cream, 10px 10px 0px 2px $yellow-gold;
    background-color: $yellow-pale;

    .article-item__title {
      color: $blue-primary;
      border-bottom: 2px solid $blue-secondary;
      cursor: pointer;
    }

    .article-item__description {
      color: $blue-secondary;
    }
  }

  &__summary {
    width: 990px;
    height: auto;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: flex-begin;
    align-items: flex-begin;
  }

  &__title {
    font-size: 2rem;
    font-weight: bold;
  }

  &__description {
    font-size: 1.2rem;
    line-height: 1.5rem;
    padding-top: 10px;
  }

  &__progress {
    width: 140px;
    height: 140px;
    padding: 10px;
  }
}
</style>
