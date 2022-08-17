<template>
  <div class="complete-circle">
    <img src="@/assets/buttons/circledTick.svg" alt="finished" />
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.complete-circle {
  /* width: calc(var(--max-dimension) * 0.35);
  height: calc(var(--max-dimension) * 0.38); */
  width: 600px;
  height: 600px;
  left: 60vw;
  top: 30vh;
  position: fixed;
  background: transparent;
  z-index: -2;

  img {
    width: 600px;
    height: 600px;
  }
}
</style>
