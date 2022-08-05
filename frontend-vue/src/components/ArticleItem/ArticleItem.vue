<script setup lang="ts">
import variables from "@/assets/variables.ts";
console.log(variables)
console.log(typeof variables)
</script>

<template>
<div class="complete">
hello
</div>
</template>

<style lang="scss">
@import "@/assets/variables";

.complete {
    background-color: $blue-secondary;
}
</style>