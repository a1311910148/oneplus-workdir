<template>
  <div
    class="relative"
    :class="{ 'text-green-500': currentActive == data }"
    @click="handleSelectX"
  >
    <span>{{ data }}</span>
    <span
      v-if="xOry == 'x'"
      class="absolute -bottom-0.5 h-2 w-full left-0"
      :class="{ activet: currentActive == data }"
    ></span>
    <span
      v-if="xOry == 'y'"
      class="absolute -right-0.5 h-full w-full left-0"
      :class="{ activel: currentActive == data }"
    ></span>
  </div>
</template>

<script setup>
import store from "../store/index";
const props = defineProps({
  data: {
    type: Number,
  },
  currentActive: {
    type: Number,
  },
  xOry: {
    type: String,
  },
});
const handleSelectX = () => {
  store.selectLine(props.data, props.xOry);
};
</script>

<style lang="scss" scoped>
.activet {
  color: rgb(17, 170, 17);
  border-bottom: 3px solid rgb(13, 212, 13);
}

.activel {
  color: rgb(46, 236, 46);
  border-right: 3px solid rgb(13, 212, 13);
}
</style>
