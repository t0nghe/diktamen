<script setup lang="ts">
import { ref, watch, computed } from "vue";
const props = defineProps<{
  mediaUri: string;
  parentPlayPause: number;
  dimension: number;
}>();
const revAudioPlayerRef = ref<HTMLMediaElement | null>(null);

const playing = ref(false);

const parentPlayPauseCounter = computed(() => props.parentPlayPause);
watch(parentPlayPauseCounter, () => {
  if (!playing.value) {
    revAudioPlayerRef.value?.play();
    playing.value = true;
  }
});

const dimCss = computed(() => `${props.dimension}px`);

const playHandler = () => {
  playing.value = true;
};
const stopHandler = () => {
  playing.value = false;
};

const toggle = () => {
  if (playing.value) {
    revAudioPlayerRef.value?.pause();
    playing.value = false;
  } else {
    revAudioPlayerRef.value?.play();
    playing.value = true;
  }
};
</script>

<template>
  <div
    class="rev-audio"
    :class="playing ? 'rev-audio-bg-pause' : 'rev-audio-bg-play'"
    @click="toggle"
  >
    <audio
      :src="props.mediaUri"
      ref="revAudioPlayerRef"
      autoplay="false"
      preload="auto"
      style="display: none"
      @play="playHandler"
      @pause="stopHandler"
      @end="stopHandler"
    ></audio>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.rev-audio {
  width: v-bind(dimCss);
  height: v-bind(dimCss);
  cursor: pointer;

  @include for-mobile {
    transform: scale(0.8, 0.8);
  }
}

.rev-audio-bg-pause {
  background: url(@/assets/buttons/buttonPause.svg);
  background-size: cover;
}

.rev-audio-bg-play {
  background: url(@/assets/buttons/buttonPlay.svg);
  background-size: cover;
}
</style>
