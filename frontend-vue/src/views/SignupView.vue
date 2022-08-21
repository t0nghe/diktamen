<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { signupMutation } from "../graphql";
import { useMutation } from "@vue/apollo-composable";
import TextNavButton from "@/components/Interaction/TextNavButton.vue";

const signUpSuccess = ref(false);
const errorMessages = ref<string[]>([]);
const usernameField = ref("");
const passwordField = ref("");
const passwordConfirmField = ref("");
const emailField = ref("");

const { mutate: userSignup } = useMutation(signupMutation);

const clearErrors = () => {
  errorMessages.value = [];
};

const reSymbol = /[!@$&*_\-.,"#%'()+/:;<=>?[\\\]^{|}~\]]/;
const reLetter = /[A-Za-z]/;
const reDigit = /[0-9]/;

const signupClickHandler = () => {
  const conditionsNotMet = [];
  if (usernameField.value.length < 4) {
    conditionsNotMet.push("username needs to contain at lease 4 characters");
  }
  if (passwordField.value.length < 6) {
    conditionsNotMet.push("password needs to contain at lease 6 characters");
  }

  if (!reSymbol.test(passwordField.value)) {
    conditionsNotMet.push("password needs to include at least 1 symbol");
  }
  if (!reLetter.test(passwordField.value)) {
    conditionsNotMet.push("password needs to include at least 1 letter");
  }
  if (!reDigit.test(passwordField.value)) {
    conditionsNotMet.push("password needs to include at least 1 digit");
  }
  if (passwordField.value !== passwordConfirmField.value) {
    conditionsNotMet.push("passwords do not match");
  }
  if (conditionsNotMet.length > 0) {
    errorMessages.value = conditionsNotMet;
    return null;
  } else {
    userSignup({
      username: usernameField.value,
      password: passwordField.value,
      email: emailField.value ?? "",
    })
      .then((res) => {
        if (
          res &&
          res.data &&
          res.data.userSignUp &&
          res.data.userSignUp.success
        ) {
          signUpSuccess.value = true;
        } else if (res && !res.data && res.errors && res.errors.length > 0) {
          res.errors.forEach((item) =>
            errorMessages.value.push(item.message || "something went wrong")
          );
        } else {
          errorMessages.value.push("something went wrong");
        }
      })
      .catch((err) => (errorMessages.value = [`${err}`]));
  }
};
</script>

<template>
  <top-bar state="heading" heading="sign up" />
  <div class="auth-container">
    <template v-if="!signUpSuccess">
      <form class="auth-signup-form">
        <h2>sign up</h2>
        <p>
          Username needs to be at least 4 characters. Password needs to be at
          least 6 characters and should contain at least 1 symbol, 1 letter and
          1 digit.
        </p>
        <div>
          <label for="signup-username"
            >username:
            <input
              type="text"
              name="username"
              id="signup-username"
              v-model="usernameField"
              @input="clearErrors"
            />
          </label>
        </div>
        <div>
          <label for="signup-password"
            >password:
            <input
              type="password"
              name="password"
              id="signup-password"
              v-model="passwordField"
              @input="clearErrors"
            />
          </label>
        </div>
        <div>
          <label for="signup-password-confirm"
            >confirm password:
            <input
              type="password"
              name="password"
              id="signup-password-confirm"
              v-model="passwordConfirmField"
              @input="clearErrors"
            />
          </label>
        </div>
        <div>
          <label for="signup-password"
            >email (optional):
            <input
              type="email"
              name="email"
              id="signup-email"
              v-model="emailField"
              @input="clearErrors"
            />
          </label>
        </div>
        <button @click.prevent="signupClickHandler">sign up</button>
      </form>
      <div v-if="errorMessages.length > 0" class="auth-error-message">
        <ul>
          <li v-for="(item, idx) in errorMessages" :key="idx">{{ item }}</li>
        </ul>
      </div>
    </template>
    <template v-else>
      <h2>You've signed up!</h2>
      <text-nav-button href="/login">continue to log in</text-nav-button>
    </template>
  </div>
</template>

<style lang="scss">
@import "../assets/variables";
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
