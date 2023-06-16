<template>
  <div>
    <div class="p-4 pb-8">
      <baseSelect
        :data="[
          { title: '11', value: 11 },
          { title: '10', value: 10 },
          { title: '9', value: 9 },
        ]"
      />
      <iconButton src="http://101.43.18.212:8883/1683725465308.svg" />
      <iconButton src="http://101.43.18.212:8883/1683725659391.svg" />
      <iconButton src="http://101.43.18.212:8883/1683725728052.svg" />
      <activeBox src="http://101.43.18.212:8883/1683726308691.svg" />
    </div>
    <LayOut>
      <CellBox :data="data" @select="handleSelect" />
    </LayOut>
  </div>
</template>

<script setup>
import { ref } from "vue";

import LayOut from "./components/layout.vue";
import CellBox from "./components/cellbox.vue";
import baseSelect from "./components/toolbar/select.vue";
import iconButton from "./components/toolbar/iconButton.vue";
import activeBox from "./components/toolbar/activeBox.vue";
import store from "./store/index";

const data = store.data;
console.log(import.meta.url);
const handleSelect = ({ start, stop }) => {
  console.log(1);
  console.log(start, stop);
  // 生成副本
  let dataclone = data.value;
  dataclone.forEach((el) => {
    el.select = false;
    if (
      el.x >= start.x &&
      el.x <= stop.x &&
      el.y >= start.y &&
      el.y <= stop.y
    ) {
      el.select = true;
    }
  });
  data.value = dataclone;
};
</script>

<style scoped></style>
