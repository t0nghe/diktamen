<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { loginMutation } from "../graphql";
import { useMutation } from "@vue/apollo-composable";
import { useLoginStore } from "../stores/loginStore";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
const router = useRouter();
const errorMessages = ref<string[]>([]);
const usernameField = ref("");
const passwordField = ref("");
const loginStore = useLoginStore();

const clearErrors = () => {
  errorMessages.value = [];
};

const { mutate: userLogin } = useMutation(loginMutation);

const login = () => {
  if (usernameField.value.length < 4) {
    errorMessages.value.push("username needs to contain at lease 4 characters");
    return null;
  }
  if (passwordField.value.length < 6) {
    errorMessages.value.push("password needs to contain at lease 6 characters");
    return null;
  }

  userLogin({
    username: usernameField.value,
    password: passwordField.value,
  })
    .then((res) => {
      if (res && res.data && res.data.userLogIn && res.data.userLogIn.token) {
        const obtainedToken = res.data.userLogIn.token;
        window.localStorage.setItem("token", obtainedToken);
        loginStore.piniaLogIn();
        router.push("/articles");
      } else if (res && !res.data && res.errors && res.errors.length > 0) {
        const message =
          res.errors[0].message === "sql: no rows in result set"
            ? "user not found"
            : res.errors[0].message;
        errorMessages.value.push(message ?? "something went wrong");
      } else {
        errorMessages.value.push("something went wrong");
      }
    })
    .catch((error) => (errorMessages.value = [`${error}`]));
};
</script>

<template>
  <top-bar state="heading" heading="login" />
  <div
    class="onboarding-container onboarding-container-login"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <div class="onboarding-content">
      <form class="auth-login-form">
        <div>
          <label for="login-username"
            >username:
            <input
              type="text"
              name="username"
              id="login-username"
              v-model="usernameField"
              @input="clearErrors"
            />
          </label>
        </div>
        <div>
          <label for="login-password"
            >password:
            <input
              type="password"
              name="password"
              id="login-password"
              v-model="passwordField"
              @input="clearErrors"
            />
          </label>
        </div>
        <div>
          <button @click.prevent="login">login</button>
        </div>
        <div class="auth-error-message">
          <ul v-if="errorMessages.length > 0">
            <li v-for="(item, idx) in errorMessages" :key="idx">{{ item }}</li>
          </ul>
        </div>
      </form>
    </div>
  </div>
</template>

<style lang="scss">
@import "../assets/variables";
/* This will be used for the form in login view and sign up view */
.auth-login-form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 2rem;

  div {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }

  label {
    font-weight: 200;
    font-size: 1.5rem;
    line-height: 2.2rem;
  }

  input {
    background-color: $azure-secondary;
    font-size: 1.4rem;
    height: 1.8rem;
    color: $blue-primary;
    line-height: 2.2rem;
    margin-left: 0.3rem;
    border: 1px solid $azure-primary;
    box-shadow: 2px 2px $blue-secondary;
  }

  button {
    text-align: center;
    background-color: $blue-primary;
    border: 2px solid $yellow-canary;
    cursor: pointer;
    padding: 0rem 3rem;
    height: 3rem;
    margin-left: 2rem;
    margin-right: 2rem;
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
    text-decoration: none;
    color: $yellow-gold;
    font-size: 1.2rem;
    /* font-weight: bold; */
  }
}
</style>
