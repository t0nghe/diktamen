<template>
  <div>On this page, new users can sign up.</div>
</template>

<script lang="ts">
// import gql from 'graphql-tag'

export default {
  name: "SignupView",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    // login() {
    //   this.$apollo.mutate({
    //     mutation: gql`
    //       mutation login($username: String!, $password: String!) {
    //         userLogIn(input: {username: $username, password: $password}) {
    //           token
    //         }
    //       }
    //     `,
    //     variables: {
    //       username: this.username,
    //       password: this.password
    //     }
    //   }).then(({ data }) => {
    //     localStorage.setItem('token', data.userLogIn.token)
    //     this.$router.push('/')
    //   })
    // }
  },
};
</script>

<style lang="scss"></style>
