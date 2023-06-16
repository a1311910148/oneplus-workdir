<template>
    <div class="container" @mousedown="handleMouseDown" @mouseup="handleMouseUp">
        <template v-for="(item, index) in data" :key="item.id">
            <Cell :data="item" />
        </template>
    </div>
    <ContextMenu />
</template>

<script setup>
import { ref } from 'vue'
import lodash from 'lodash';

import Cell from './cell.vue'
import ContextMenu from './contextmenu.vue';

console.log(import.meta.url);
const props = defineProps({
    data: {
        type: Array
    }
})

const emit = defineEmits(['select'])

let startEl = null;
let stopEl = null;
// 
const handleMouseDown = (e) => {
    startEl = e.target;
}
// 
const handleMouseUp = (e) => {
    stopEl = e.target
    let startxy = { x: startEl.dataset.x, y: startEl.dataset.y };
    let stopxy = { x: stopEl.dataset.x, y: stopEl.dataset.y };
    console.log(startxy, stopxy);
    emit('select', { start: startxy, stop: stopxy })
}
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    flex-wrap: wrap;
    border-collapse: collapse;
}
</style>