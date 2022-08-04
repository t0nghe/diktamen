<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import ListComponents from "./views/dev/ListComponents.vue";
</script>

<template>
  <header>
    <div>LOGO</div>
    <!-- <div>
      Sanity check: <span>{{ randomInteger }}</span>
    </div> -->
    <!-- <nav class="nav-links">
      <RouterLink to="/dev/components">List of Components</RouterLink>
      <RouterLink to="/dev/graphql">Log In</RouterLink>
      <RouterLink to="/dev/graphql/signup">Sign Up</RouterLink>
      <RouterLink to="/dev/graphql/articles/seen">Seen Articles</RouterLink>
      <RouterLink to="/dev/graphql/articles/unseen">Unseen Articles</RouterLink>
      <RouterLink to="/dev/graphql/sents/seen">Seen Sentences</RouterLink>
      <RouterLink to="/dev/graphql/sents/unseen">Unseen Sentences</RouterLink>
      <RouterLink to="/dev/graphql/examine/correct"
        >Examine Correct Dictations</RouterLink
      >
      <RouterLink to="/dev/graphql/examine/incorrect"
        >Examine Incorrect Dictations</RouterLink
      >
      <RouterLink to="/dev/graphql/review/due">Review Due Cards</RouterLink>
    </nav> -->
  </header>
  <ListComponents />

  <!-- <RouterView /> -->
</template>

<script lang="ts">
import gql from "graphql-tag";

export default {
  name: "App",
  components: {
    RouterLink,
    RouterView,
  },
  // data() {
  //   return {
  //     randomInteger: -1,
  //   };
  // },
  // apollo: {
  //   randomInteger: gql`
  //     query {
  //       randomInteger
  //     }
  //   `,
  // },
};
</script>

<style lang="scss">
.nav-links {
  display: flex;
  justify-content: space-around;

  a {
    margin: 0.5rem;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.5rem;
  }
}
</style>
