<template>
  <div class="rounded hover:bg-gray-300 opacity-50 inline-block">
    <img :src="src" width="20" />
  </div>
</template>

<script setup>
const props = defineProps({
  src: {
    type: String,
  },
});
</script>

<style lang="scss" scoped></style>
