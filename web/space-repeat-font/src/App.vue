<!-- src\App.vue -->
<template>
  <el-config-provider :locale="locale">
<!--    <index></index>-->
<!--    <test-axios></test-axios>-->
    <router-view />
  </el-config-provider>
</template>

<script setup lang="ts">
import locale from 'element-plus/es/locale/lang/zh-cn'
import TestAxios from "@/pages/TestAxios.vue";
// import index from './views/login/index.vue'
</script>

