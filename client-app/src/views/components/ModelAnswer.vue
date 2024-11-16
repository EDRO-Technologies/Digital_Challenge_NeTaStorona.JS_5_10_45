<template>
  <p v-html="props.answer"></p>
</template>

<script setup lang="ts">
const props = defineProps<{
  answer: any;
}>();
</script>

<style scoped></style>
