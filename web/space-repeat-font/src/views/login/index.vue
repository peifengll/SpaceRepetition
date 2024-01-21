<!-- src\views\login\index.vue -->
<template>
  <div>
    登录
  </div>
</template>

<script setup lang="ts">
import {demo} from '../../api/common.ts'
import {onMounted} from 'vue'

onMounted(() => {
  demo().then(res => {
    console.log(res.data)
    // {"msg":"ok","status":200,"data":{"title":"Hello World","date":1649412637487}}
  })
})

</script>

<style scoped>
</style>

