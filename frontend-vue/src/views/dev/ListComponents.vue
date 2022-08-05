<script setup lang="ts">
import ProgressCircle from "../../components/ArticleItem/ProgressCircle.vue";
import ArticleItem from "../../components/ArticleItem/ArticleItem.vue";
import { ref } from "vue";
const pcPrimary = ref("#00bcd4");
const pcSecondary = ref("#dcbc00");
const pcProg = ref(1);
const pcGoal = ref(10);
</script>

<template>
  <h2>LIST OF COMPONENTS</h2>
  <ul id="components">
    <li class="component-item">
      <h2>ProgressCircle</h2>
      <!-- TODO WRITE input elements to set circle props. -->
      Note: We still need to change the size.
      <form>
        <fieldset>
          <legend>Progress Circle</legend>
          <label for="pc_primary"
            >Primary color:
            <input
              type="color"
              id="pc_primary"
              name="primary"
              v-model="pcPrimary"
            /> </label
          ><br />
          <label for="pc_secondary"
            >Secondary color:
            <input
              type="color"
              id="pc_secondary"
              name="secondary"
              v-model="pcSecondary"
            /> </label
          ><br />
          <label for="pc_prog"
            >Progress:
            <input
              type="number"
              id="pc_prog"
              name="prog"
              v-model="pcProg"
            /> </label
          ><br />
          <label for="pc_goal"
            >Goal:
            <input
              type="number"
              id="pc_goal"
              name="goal"
              v-model="pcGoal"
            /> </label
          ><br />
        </fieldset>
      </form>
      <ProgressCircle
        :prog="pcProg"
        :goal="pcGoal"
        :primary="pcPrimary"
        :secondary="pcSecondary"
      />
    </li>
    <li class="component-item">
      <h2>ArticleItem</h2>
      <ArticleItem />
    </li>
  </ul>
</template>

<style lang="scss">
#components {
  display: flex;

  .component-item {
    padding: 1rem;
  }
}
</style>
