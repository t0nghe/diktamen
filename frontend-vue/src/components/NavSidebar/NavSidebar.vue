<script setup lang="ts">
import { useRouter } from "vue-router";
import { useLoginStore } from "../../stores/loginStore";
import { useNavStore } from "../../stores/navWidthStore";

const router = useRouter();
const loginStore = useLoginStore();
const navWidthStore = useNavStore();

const goToLogin = () => {
  router.push("/login");
};

const goToSignup = () => {
  router.push("/test-signup");
};

const goToLearn = () => {
  router.push("/articles");
};
const goToReview = () => {
  router.push("/review");
};

const logout = () => {
  window.localStorage.removeItem("token");
  loginStore.piniaLogOut();
  router.push("/");
};
</script>

<template>
  <!-- LOGGED IN -->
  <div v-if="loginStore.isLoggedIn" class="nav-sidebar-container">
    <div
      v-if="navWidthStore.isWide"
      id="sidebar-control-area"
      @click="navWidthStore.toggle()"
    ></div>
    <div v-else id="sidebar-control-area" @click="navWidthStore.toggle()"></div>
    <div
      id="sidebar-display-area"
      :class="navWidthStore.isWide ? 'sidebar-wide' : 'sidebar-narrow'"
    >
      <div class="learn-review-container">
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
          <span v-if="navWidthStore.isWide">learn</span>
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
          <span v-if="navWidthStore.isWide">review</span>
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
        <span v-if="navWidthStore.isWide">log out</span>
      </div>
    </div>
  </div>
  <!-- NOT LOGGED IN -->
  <div v-else class="nav-sidebar-container">
    <div
      v-if="navWidthStore.isWide"
      id="sidebar-control-area"
      @click="navWidthStore.toggle()"
    ></div>
    <div v-else id="sidebar-control-area" @click="navWidthStore.toggle()"></div>
    <div
      id="sidebar-display-area"
      :class="navWidthStore.isWide ? 'sidebar-wide' : 'sidebar-narrow'"
    >
      <div class="log-in-out-positioning-gap">
        <!-- This empty div is used to push log in button to the bottom. -->
      </div>
      <div class="log-in-out-positioning-container">
        <div id="sidebar-button-signup" @click="goToSignup">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 512 512"
            class="svg-button-login"
          >
            <path
              d="M490.3 40.4C512.2 62.27 512.2 97.73 490.3 119.6L460.3 149.7L362.3 51.72L392.4 21.66C414.3-.2135 449.7-.2135 471.6 21.66L490.3 40.4zM172.4 241.7L339.7 74.34L437.7 172.3L270.3 339.6C264.2 345.8 256.7 350.4 248.4 353.2L159.6 382.8C150.1 385.6 141.5 383.4 135 376.1C128.6 370.5 126.4 361 129.2 352.4L158.8 263.6C161.6 255.3 166.2 247.8 172.4 241.7V241.7zM192 63.1C209.7 63.1 224 78.33 224 95.1C224 113.7 209.7 127.1 192 127.1H96C78.33 127.1 64 142.3 64 159.1V416C64 433.7 78.33 448 96 448H352C369.7 448 384 433.7 384 416V319.1C384 302.3 398.3 287.1 416 287.1C433.7 287.1 448 302.3 448 319.1V416C448 469 405 512 352 512H96C42.98 512 0 469 0 416V159.1C0 106.1 42.98 63.1 96 63.1H192z"
            />
          </svg>
          <span v-if="navWidthStore.isWide">sign up for testing</span>
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
          <span v-if="navWidthStore.isWide">log in</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import "@/assets/variables";

.nav-sidebar-container {
  display: flex;
  flex-direction: row;
  justify-content: left;
  align-items: center;
  width: auto;
  height: calc(100vh - $topbar-height);
  position: fixed;
  top: $topbar-height;
  left: 0;

  @include for-mobile {
    width: 100%;
    height: $mobile-navbar-height;
    top: calc(100% - $mobile-navbar-height);
    /* display: flex;
    flex-direction: column; */
    z-index: 10;
  }

  #sidebar-control-area {
    width: $nav-sidebar-control-width;
    height: calc(100vh - $topbar-height);
    background-color: $blue-primary;
    cursor: pointer;

    @include for-mobile {
      display: none;
    }
  }

  #sidebar-control-area:hover {
    background-color: $blue-secondary;
  }

  #sidebar-display-area.sidebar-wide {
    width: $nav-sidebar-display-width-wide;

    @include for-mobile {
      width: 100vw;
    }
  }

  #sidebar-display-area.sidebar-narrow {
    width: $nav-sidebar-display-width-narrow;

    @include for-mobile {
      width: 100vw;
    }
  }

  #sidebar-display-area {
    height: calc(100vh - $topbar-height);
    background-color: $blue-primary;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    @include for-mobile {
      height: $mobile-navbar-height;
      flex-direction: row;
      justify-content: space-between;
      align-items: space-between;
    }
  }

  #sidebar-nav-learn,
  #sidebar-nav-review,
  #sidebar-button-logout,
  #sidebar-button-login,
  #sidebar-button-signup {
    margin-top: 30px;
    cursor: pointer;
    display: flex;
    flex-direction: row;
    align-items: center;

    @include for-mobile {
      margin-top: 0px;
      height: $mobile-navbar-height;
      margin-left: 10px;
    }

    span {
      margin-left: 30px;
      font-size: 1.5rem;
      color: $azure-secondary;
      height: 32px;
      line-height: 32px;
      font-weight: bold;

      @include for-mobile {
        margin: 5px;
        font-size: 1.2rem;
      }
    }
  }

  #sidebar-button-logout {
    margin-bottom: 20px;

    span {
      font-size: 1.4rem;

      @include for-mobile {
        font-size: 1.2rem;
      }
    }

    @include for-mobile {
      margin-bottom: 0px;
      padding: 0px;
      margin-right: 10px;
    }
  }

  #sidebar-button-signup {
    padding-bottom: 30px;

    span {
      font-size: 1.3rem;
    }

    @include for-mobile {
      padding: 0px;
      margin-left: 10px;
    }
  }

  #sidebar-button-login {
    margin-bottom: 20px;

    span {
      font-size: 1.3rem;
    }

    @include for-mobile {
      margin-bottom: 0px;
      padding: 0px;
      margin-right: 10px;
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

  .log-in-out-positioning-gap {
    @include for-mobile {
      display: none;
    }
  }

  .log-in-out-positioning-container {
    @include for-mobile {
      width: 100%;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
    }
  }

  .learn-review-container {
    @include for-mobile {
      display: flex;
      flex-direction: row;
    }
  }
}
</style>
