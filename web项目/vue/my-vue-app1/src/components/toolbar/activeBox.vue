<template>
  <div
    ref="tigger"
    class="relative flex items-center justify-between w-6 hover:bg-gray-400 hover:bg-opacity-50 p-1 rounded"
    @click="isshow = true"
  >
    <!--主 icon -->
    <img :src="src" class="w-4" />
    <!-- 下拉icon -->
    <div class="flex items-center">
      <img class="w-2" src="../../assets/select.svg" />
    </div>
    <!-- 下拉框 -->
    <div
      v-if="isshow"
      style="transform: translateX(-50%)"
      class="absolute top-full left-1/2"
    >
      <slot>2222</slot>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
const isshow = ref(false);
const tigger = ref();
const props = defineProps({
  src: {
    type: String,
  },
});
window.onclick = (e) => {
  if (!tigger.value.contains(e.target)) {
    isshow.value = false;
  }
};
</script>

<style lang="scss" scoped></style>
