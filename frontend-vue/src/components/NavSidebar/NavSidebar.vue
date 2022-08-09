<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
const props = defineProps<{ wide: boolean }>();
const router = useRouter();

const wideMode = ref(true);
const loggedIn = ref(false);
if (window.localStorage.getItem("token") !== null) {
  loggedIn.value = true;
}
// eslint-disable-next-line vue/no-setup-props-destructure
wideMode.value = props.wide;

const toggleWide = () => {
  wideMode.value = !wideMode.value;
};

const goToLogin = () => {
  router.push("/login");
};

const goToLearn = () => {
  router.push("/articles");
};
const goToReview = () => {
  router.push("/review");
};

const logout = () => {
  window.localStorage.removeItem("token");
  loggedIn.value = false;
  router.push("/");
};
</script>

<template>
  <!-- LOGGEIN IN -->
  <div v-if="loggedIn" class="nav-sidebar-container">
    <div v-if="wideMode" id="sidebar-control-area" @click="toggleWide"></div>
    <div v-else id="sidebar-control-area" @click="toggleWide"></div>
    <div
      id="sidebar-display-area"
      :class="wideMode ? 'sidebar-wide' : 'sidebar-narrow'"
    >
      <div>
        <div id="sidebar-nav-learn" @click="goToLearn">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="svg-nav-learn"
            viewBox="0 0 550 550"
          >
            <path
              d="M464 64C490.5 64 512 85.49 512 112V176C512 202.5 490.5 224 464 224H48C21.49 224 0 202.5 0 176V112C0 85.49 21.49 64 48 64H464zM448 128H320V160H448V128zM464 288C490.5 288 512 309.5 512 336V400C512 426.5 490.5 448 464 448H48C21.49 448 0 426.5 0 400V336C0 309.5 21.49 288 48 288H464zM192 352V384H448V352H192z"
            />
          </svg>
          <span v-if="wideMode">learn</span>
        </div>
        <div id="sidebar-nav-review" @click="goToReview">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="svg-nav-review"
            viewBox="0 0 512 512"
          >
            <path
              d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM371.8 211.8C382.7 200.9 382.7 183.1 371.8 172.2C360.9 161.3 343.1 161.3 332.2 172.2L224 280.4L179.8 236.2C168.9 225.3 151.1 225.3 140.2 236.2C129.3 247.1 129.3 264.9 140.2 275.8L204.2 339.8C215.1 350.7 232.9 350.7 243.8 339.8L371.8 211.8z"
            />
          </svg>
          <span v-if="wideMode">review</span>
        </div>
      </div>
      <div id="sidebar-button-logout" @click="logout">
        <svg
          class="svg-button-logout"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 512 512"
        >
          <path
            d="M160 416H96c-17.67 0-32-14.33-32-32V128c0-17.67 14.33-32 32-32h64c17.67 0 32-14.33 32-32S177.7 32 160 32H96C42.98 32 0 74.98 0 128v256c0 53.02 42.98 96 96 96h64c17.67 0 32-14.33 32-32S177.7 416 160 416zM502.6 233.4l-128-128c-12.51-12.51-32.76-12.49-45.25 0c-12.5 12.5-12.5 32.75 0 45.25L402.8 224H192C174.3 224 160 238.3 160 256s14.31 32 32 32h210.8l-73.38 73.38c-12.5 12.5-12.5 32.75 0 45.25s32.75 12.5 45.25 0l128-128C515.1 266.1 515.1 245.9 502.6 233.4z"
          />
        </svg>
        <span v-if="wideMode">log out</span>
      </div>
    </div>
  </div>
  <!-- NOT LOGGED IN -->
  <div v-else class="nav-sidebar-container">
    <div v-if="wideMode" id="sidebar-control-area" @click="toggleWide"></div>
    <div v-else id="sidebar-control-area" @click="toggleWide"></div>
    <div
      id="sidebar-display-area"
      :class="wideMode ? 'sidebar-wide' : 'sidebar-narrow'"
    >
      <div>
        <!-- This empty div is used to push log in button to the bottom. -->
      </div>
      <div id="sidebar-button-login" @click="goToLogin">
        <svg
          class="svg-button-login"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 512 512"
        >
          <path
            d="M416 32h-64c-17.67 0-32 14.33-32 32s14.33 32 32 32h64c17.67 0 32 14.33 32 32v256c0 17.67-14.33 32-32 32h-64c-17.67 0-32 14.33-32 32s14.33 32 32 32h64c53.02 0 96-42.98 96-96V128C512 74.98 469 32 416 32zM342.6 233.4l-128-128c-12.51-12.51-32.76-12.49-45.25 0c-12.5 12.5-12.5 32.75 0 45.25L242.8 224H32C14.31 224 0 238.3 0 256s14.31 32 32 32h210.8l-73.38 73.38c-12.5 12.5-12.5 32.75 0 45.25s32.75 12.5 45.25 0l128-128C355.1 266.1 355.1 245.9 342.6 233.4z"
          />
        </svg>
        <span v-if="wideMode">log in</span>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import "../../assets/variables";

.nav-sidebar-container {
  display: flex;
  flex-direction: row;
  justify-content: left;
  align-items: center;
  width: auto;
  height: calc(100vh - $topbar-height);

  #sidebar-control-area {
    width: $nav-sidebar-control-width;
    height: calc(100vh - $topbar-height);
    background-color: $blue-primary;
    cursor: pointer;
  }

  #sidebar-display-area.sidebar-wide {
    width: $nav-sidebar-display-width-wide;
  }

  #sidebar-display-area.sidebar-narrow {
    width: $nav-sidebar-display-width-narrow;
  }

  #sidebar-display-area {
    height: calc(100vh - $topbar-height);
    background-color: $blue-primary;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  #sidebar-nav-learn,
  #sidebar-nav-review,
  #sidebar-button-logout,
  #sidebar-button-login {
    margin-top: 30px;
    cursor: pointer;
    display: flex;
    flex-direction: row;
    align-items: center;

    span {
      margin-left: 14px;
      font-size: 1.5rem;
      color: $azure-secondary;
      height: 32px;
      line-height: 32px;
      font-weight: bold;
    }
  }

  #sidebar-button-logout {
    margin-bottom: 20px;

    span {
      font-size: 1.4rem;
    }
  }

  #sidebar-button-login {
    margin-bottom: 20px;

    span {
      font-size: 1.4rem;
    }
  }

  .svg-nav-learn,
  .svg-nav-review,
  .svg-button-logout,
  .svg-button-login {
    width: 32px;
    height: 32px;
    margin: 8px;

    path {
      fill: $azure-secondary;
    }
  }
}
</style>
