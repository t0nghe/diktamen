<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { loginMutation } from "../graphql";
import { useMutation } from "@vue/apollo-composable";

const router = useRouter();
const errorMessages = ref<string[]>([]);
const usernameField = ref("");
const passwordField = ref("");

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
  <div class="auth-container">
    <form class="auth-login-form">
      <h2>login</h2>
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
      <button @click.prevent="login">login</button>
    </form>
  </div>
  <div v-if="errorMessages.length > 0" class="auth-error-message">
    <ul>
      <li v-for="(item, idx) in errorMessages" :key="idx">{{ item }}</li>
    </ul>
  </div>
</template>

<style lang="scss">
@import "../assets/variables";
@import "../assets/base.css";
/* This will be used for the form in login view and sign up view */
.auth-container {
  input {
    background-color: $azure-secondary;
  }

  button {
    background-color: $blue-secondary;
  }
}
</style>
