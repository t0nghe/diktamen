<template>
  <div class="loading-dots">
    <svg
      enable-background="new 0 0 342.382 342.382"
      height="342.382"
      viewBox="0 0 342.382 342.382"
      width="342.382"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        class="loading-dot-left"
        d="m45.225 125.972c-24.941 0-45.225 20.284-45.225 45.219 0 24.94 20.284 45.219 45.225 45.219 24.926 0 45.219-20.278 45.219-45.219 0-24.935-20.293-45.219-45.219-45.219z"
      />
      <path
        class="loading-dot-middle"
        d="m173.409 125.972c-24.938 0-45.225 20.284-45.225 45.219 0 24.94 20.287 45.219 45.225 45.219 24.936 0 45.226-20.278 45.226-45.219 0-24.935-20.29-45.219-45.226-45.219z"
      />
      <path
        class="loading-dot-right"
        d="m297.165 125.972c-24.932 0-45.222 20.284-45.222 45.219 0 24.94 20.29 45.219 45.222 45.219 24.926 0 45.217-20.278 45.217-45.219 0-24.935-20.291-45.219-45.217-45.219z"
      />
    </svg>
  </div>
</template>

<style lang="scss" scoped>
@import "@/assets/variables";

svg {
  width: 120px;
  height: auto;
}
.loading-dot-left {
  animation-duration: 3s;
  animation-timing-function: ease-in-out;
  animation-name: loading-left;
  animation-iteration-count: infinite;
  fill: $blue-secondary;

  @keyframes loading-left {
    0% {
      fill: $yellow-gold;
    }
    32% {
      fill: $yellow-gold;
    }
    33% {
      fill: $blue-secondary;
    }
    100% {
      fill: $blue-secondary;
    }
  }
}

.loading-dot-middle {
  animation-duration: 3s;
  animation-timing-function: ease-in-out;
  animation-name: loading-middle;
  animation-iteration-count: infinite;
  fill: $blue-secondary;

  @keyframes loading-middle {
    0% {
      fill: $blue-secondary;
    }
    32% {
      fill: $blue-secondary;
    }
    33% {
      fill: $yellow-gold;
    }
    66% {
      fill: $yellow-gold;
    }
    67% {
      fill: $blue-secondary;
    }
    100% {
      fill: $blue-secondary;
    }
  }
}

.loading-dot-right {
  animation-duration: 3s;
  animation-timing-function: ease-in-out;
  animation-name: loading-right;
  animation-iteration-count: infinite;
  fill: $blue-secondary;

  @keyframes loading-right {
    0% {
      fill: $blue-secondary;
    }
    66% {
      fill: $blue-secondary;
    }
    67% {
      fill: $yellow-gold;
    }
    100% {
      fill: $yellow-gold;
    }
  }
}
</style>
