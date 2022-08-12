<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <header>
    <div>LOGO</div>
    <div>
      Sanity check: <span>{{ randomInteger }}</span>
    </div>
    <nav class="nav-links">
      <RouterLink to="/">Sentence Components</RouterLink>
      <RouterLink to="/list-components">List of components</RouterLink>
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
    </nav>
  </header>

  <RouterView />
</template>

<script lang="ts">
import gql from "graphql-tag";

export default {
  name: "App",
  components: {
    RouterLink,
    RouterView,
  },
  data() {
    return {
      randomInteger: -1,
    };
  },
  apollo: {
    randomInteger: gql`
      query {
        randomInteger
      }
    `,
  },
};
</script>

<style lang="scss">
.nav-links {
  display: flex;
  flex-wrap: wrap;

  a {
    margin: 0.2rem;
    padding: 0.2rem;
    border: 1px solid #ccc;
    border-radius: 0.5rem;
  }
}
</style>
