<template>
    <input @click="handleClick" :style="{ background: data.color }" :class="{ active: data.select }" type="text"
        v-model="data.content" :data-id="data.id" :data-x="data.x" :data-y="data.y">
</template>

<script setup>
import store from '../store/index'
const props = defineProps({
    data: {
        type: Object
    }
})
const handleClick = () => {
    let x = props.data.x
    let y = props.data.y
    store.activeline.x = x
    store.activeline.y = y
}
</script>

<style  scoped>
input {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-size: 16px;
    font-family: Arial, sans-serif;
    line-height: 1.5;
    color: #333;
    text-align: left;
    text-decoration: none;
    list-style: none;
    border: none;
    outline: none;
    background-color: transparent;
    height: 100%;
    width: 60px;
    border: 1px solid rgb(167, 163, 163);
    display: inline-block !important;
    border-top: 0;
    border-right: 0;
}

input:hover {
    cursor: cell;
}

input::backdrop {
    display: none;
}

.active {
    background: #cfcfcfa6;
}
</style>