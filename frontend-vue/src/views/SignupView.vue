<script setup lang="ts">
import TopBar from "@/components/TopBar/TopBar.vue";
import { ref } from "vue";
import { signupMutation } from "../graphql";
import { useMutation } from "@vue/apollo-composable";
import TextNavButton from "@/components/Interaction/TextNavButton.vue";
import { useNavStore } from "../stores/navWidthStore";

const navWidthStore = useNavStore();
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

const reLetter = /[A-Za-z]/;
const reDigit = /[0-9]/;

const signupClickHandler = () => {
  const conditionsNotMet = [];
  if (usernameField.value.length < 4) {
    conditionsNotMet.push("username needs to contain at lease 4 characters");
  }
  if (!reLetter.test(usernameField.value[0])) {
    conditionsNotMet.push("username needs to start with a letter");
  }
  if (passwordField.value.length < 4) {
    conditionsNotMet.push("password needs to contain at lease 4 characters");
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
  <div
    class="onboarding-container"
    :class="
      navWidthStore.isWide
        ? 'container-position-wide-side'
        : 'container-position-narrow-side'
    "
  >
    <div class="onboarding-content onboarding-content-signup">
      <template v-if="!signUpSuccess">
        <form class="auth-signup-form">
          <table>
            <tbody>
              <tr>
                <th><label for="signup-username">username:</label></th>
                <td>
                  <input
                    type="text"
                    name="username"
                    id="signup-username"
                    v-model="usernameField"
                    @input="clearErrors"
                  />
                  <p>at least 4 characters</p>
                </td>
              </tr>
              <tr>
                <th><label for="signup-password">password:</label></th>
                <td>
                  <input
                    type="password"
                    name="password"
                    id="signup-password"
                    v-model="passwordField"
                    @input="clearErrors"
                  />
                </td>
              </tr>
              <tr>
                <th>
                  <label for="signup-password-confirm">repeat password:</label>
                </th>
                <td>
                  <input
                    type="password"
                    name="password"
                    id="signup-password-confirm"
                    v-model="passwordConfirmField"
                    @input="clearErrors"
                  />
                  <p>
                    at least 4 characters, containing<br />at least 1 letter and
                    1 digit
                  </p>
                </td>
              </tr>
              <tr>
                <th><label for="signup-password">email:</label></th>
                <td>
                  <input
                    type="email"
                    name="email"
                    id="signup-email"
                    v-model="emailField"
                    @input="clearErrors"
                  />
                </td>
              </tr>
            </tbody>
          </table>
          <button @click.prevent="signupClickHandler">sign up</button>
        </form>
        <div v-if="errorMessages.length > 0" class="auth-error-message">
          <ul>
            <li v-for="(item, idx) in errorMessages" :key="idx">{{ item }}</li>
          </ul>
        </div>
      </template>
      <template v-else>
        <h2>you have signed up</h2>
        <text-nav-button href="/login">continue to log in</text-nav-button>
      </template>
    </div>
  </div>
</template>

<style lang="scss">
@import "../assets/variables";
/* This will be used for the form in login view and sign up view */
.onboarding-content-signup {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  th {
    text-align: right;
    vertical-align: top;
  }
  td {
    text-align: left;
    vertical-align: top;
  }

  label {
    font-weight: 200;
    font-size: 1.5rem;
    line-height: 2.2rem;
  }

  p {
    font-weight: 300;
    font-size: 1rem;
    margin-left: 1rem;
    line-height: 1.5rem;
  }

  input {
    background-color: $azure-secondary;
    font-size: 1.4rem;
    height: 1.8rem;
    width: 20rem;
    color: $blue-primary;
    line-height: 2.2rem;
    margin-left: 1rem;
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
    font-weight: bold;
  }
}

.auth-signup-form {
  text-align: center;
}
</style>
