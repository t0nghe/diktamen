<template>
  <main>
    <form>
      <input type="text" v-model="username" placeholder="Username" />
      <input type="password" v-model="password" placeholder="Password" />
      <button type="submit" @click.prevent="login">Login</button>
    </form>
  </main>
</template>

<script lang="ts">
import gql from "graphql-tag";

export default {
  name: "LoginView",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    login() {
      this.$apollo
        .mutate({
          mutation: gql`
            mutation login($username: String!, $password: String!) {
              userLogIn(input: { username: $username, password: $password }) {
                token
              }
            }
          `,
          variables: {
            username: this.username,
            password: this.password,
          },
        })
        .then(({ data }) => {
          localStorage.setItem("token", data.userLogIn.token);
          this.$router.push("/");
        });
    },
  },
};
</script>

<style lang="scss"></style>
