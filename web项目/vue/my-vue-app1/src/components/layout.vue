<template>
  <div class="container">
    <!-- top -->
    <div class="top">
      <template v-for="item in arr" :key="item">
        <cellHeader
          :data="item"
          :current-active="store.activeline.x"
          xOry="x"
        />
      </template>
    </div>
    <!--  down  -->
    <div class="down">
      <!--  left -->
      <div class="left">
        <template v-for="item in arr" :key="item">
          <cellHeader
            :data="item"
            :current-active="store.activeline.y"
            xOry="y"
          />
        </template>
      </div>
      <!-- content  -->
      <div class="content">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup>
import store from "../store/index";
import cellHeader from "./cellHeader.vue";

let arr = Array(50)
  .fill(1)
  .map((el, index) => index + 1);

const handleSelectX = (item) => {
  console.log(item);
  store.selectLine(item);
};
const handleSelectY = (item) => {
  console.log(item);
};
</script>

<style lang="scss" scoped>
.container {
  width: 3070px;
}

.top {
  padding-left: 40px;
  background: #eaebec;
  white-space: nowrap;

  div {
    width: 60px;
    text-align: center;
    box-sizing: border-box;
    border-left: 1px solid gray;
    display: inline-block;
    background: #e8e9eb;
    cursor: url("../assets/vue.svg"), auto;
    user-select: none;
    border-bottom: 1px solid gray;
  }
}

.down {
  display: flex;

  .left {
    display: flex;
    flex-direction: column;
    text-align: center;
    background: #e9eaec;
    box-sizing: border-box;
    border-top: 1px solid gray;
    transform: translateY(-1px);

    div {
      width: 40px;
      height: 19px;
      user-select: none;
      border-bottom: 1px solid gray;
      box-sizing: border-box;
      display: inline-flex;
      align-items: center;
      justify-content: center;
    }

    // &:nth-child(1) {
    //     border-top: 1px solid gray;
    //     height: 17px !important;
    // }
    // .left::before {
    //     height: 25px;
    //     box-sizing: border-box;
    //     border-bottom: 1px solid gray;
    //     content: '';
    //     display: block;
    // }
  }

  .content {
    border-collapse: collapse;
  }
}
</style>
