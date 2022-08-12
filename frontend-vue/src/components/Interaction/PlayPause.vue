<script setup lang="ts">
import { ref, watch, computed } from "vue";
const props = defineProps<{ mediaUrl: string; parentClickPlay: number }>();
const audioPlayerRef = ref<HTMLMediaElement | null>(null);

const playing = ref(false);

const parentClickPlayCount = computed(() => props.parentClickPlay);
watch(parentClickPlayCount, () => {
  if (!playing.value) {
    audioPlayerRef.value.play();
    playing.value = true;
  }
});

const playHandler = () => {
  playing.value = true;
};

const stopHandler = () => {
  playing.value = false;
};

const toggle = () => {
  if (playing.value) {
    audioPlayerRef.value.pause();
    playing.value = false;
  } else {
    audioPlayerRef.value.play();
    playing.value = true;
  }
};
</script>

<template>
  <div class="audio-play-pause" @click="toggle">
    <audio
      :src="props.mediaUrl"
      ref="audioPlayerRef"
      autoplay="false"
      preload="auto"
      style="display: none"
      @play="playHandler"
      @pause="stopHandler"
      @end="stopHandler"
    ></audio>
  </div>
  <div class="audio-play-pause__bg">
    <img
      v-if="playing"
      src="@/assets/buttons/buttonPause.svg"
      alt="Pause button"
    />
    <img v-else src="@/assets/buttons/buttonPlay.svg" alt="Play button" />
  </div>
</template>

<style lang="scss">
.audio-play-pause {
  --max-dimension: max(100vw, 100vh);
  width: calc(var(--max-dimension) * 0.35);
  height: calc(var(--max-dimension) * 0.38);
  left: calc(var(--max-dimension) * 0.7);
  top: calc(var(--max-dimension) * 0.24);
  position: fixed;
  background: transparent;
  z-index: 2;
}

.audio-play-pause__bg {
  --max-dimension: max(100vw, 100vh);

  width: 1px;
  height: 1px;
  left: calc(var(--max-dimension) * 0.7);
  top: calc(var(--max-dimension) * 0.24);
  position: fixed;
  z-index: -2;

  img {
    width: calc(var(--max-dimension) * 0.65);
    height: calc(var(--max-dimension) * 0.65);
    transform: translate(
      calc(var(--max-dimension) * -0.18),
      calc(var(--max-dimension) * -0.14)
    );
  }
}
</style>
