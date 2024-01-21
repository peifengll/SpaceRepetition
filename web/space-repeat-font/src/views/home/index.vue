<!-- src\views\home\index.vue -->
<template>
  <div>
    首页
  </div>
  <beauty-button @click="store.setUser">
    {{ store.count }} => {{ store.doubleCount }}
  </beauty-button>
  <div>count 永远是 0：{{ store.count }}</div>
  <div>doubleCount 永远是 0：{{ doubleCount }}</div>

  <el-divider/>
  <p>Double count is {{ counter.count }}</p>
  <p>Double count is {{ counter.token }}</p>
  <button @click="countadd">aaa</button>
  <el-divider/>

</template>

<script setup lang="ts">

import {useCounterStore} from "@/stores/counter.ts";

const counter = useCounterStore()



function countadd() {
  counter.count++
  counter.token="sasas"+counter.count
}

import useStore from "../../stores";
import BeautyButton from "@/components/tabs/BeautyButton.vue";
const store = useStore()
const { count, doubleCount } = store
// 在 script-setup 中使用解构赋值要通过 defineExpose 暴漏出去
defineExpose({
  count
})
</script>

<style scoped>

</style>

