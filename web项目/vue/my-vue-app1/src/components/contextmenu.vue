<template>
    <div class="menu" :style="position">
        <div>复制</div>
        <div>粘贴</div>
        <div @click="setColor">颜色</div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import store from '../store/index'

const position = ref({ top: 0, left: 0 })

let currentCellId = ''

window.oncontextmenu = (e) => {
    e.preventDefault()
    currentCellId = e.target.dataset.id
    let x = e.pageX
    let y = e.pageY
    position.value.left = x + 'px'
    position.value.top = y + 'px'
}
window.onclick = (e) => {
    if (e.target.dataset.id) {
        position.value.left = -1000 + 'px'
        position.value.top = -1000 + 'px'
    }
}
const setColor = () => {
    console.log(currentCellId);
    store.setColor(currentCellId, 'red')
}
</script>

<style lang="scss" scoped>
.menu {
    position: fixed;
    top: 100px;
    left: 100px;
    background: #f8f6f6;
    box-shadow: 0 0 8px gray;
    border-radius: 5px;
    padding: 5px 0;

    div {
        padding: 2px 25px;
        cursor: pointer;
    }

    div:hover {
        background: #ebebeb;
    }
}
</style>