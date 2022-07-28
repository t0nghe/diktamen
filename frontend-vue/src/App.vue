<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <header>
      <nav>
        <RouterLink to="/">Home</RouterLink>
      </nav>
  </header>

  <hr />

  <RouterView />

  <div>
    {{ randomInteger }}
  </div>
</template>

<script lang="ts">
import gql from 'graphql-tag'

export default {
  name: 'App',
  components: {
    RouterLink,
    RouterView
  },
  data() {
    return {
      randomInteger: 0
    }
  },
  apollo: {
    randomInteger: gql`
      query {
        randomInteger
      }
    `
  }
}
</script>

<style>
</style>
