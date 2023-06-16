<template>
  <div class="relative w-20">
    <select
      class="block appearance-none outline-0 w-20 h-6 bg-white border border-gray-400"
      v-model="selectValue"
      :width="width"
    >
      <option v-for="item in data" :value="item.value">{{ item.title }}</option>
    </select>
    <div
      class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
    >
      <div
        class="pointer-events-none absolute inset-y-0 right-0 pr-1 flex items-center text-gray-700"
      >
        <img src="../../assets/select.svg" width="14" alt="" srcset="" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";
const props = defineProps({
  data: {
    type: Array,
  },
  width: {
    type: Number,
  },
});
// 选择
const selectValue = ref(props.data[0].value);
watch(selectValue, () => {});
</script>

<style lang="scss" scoped></style>
