<script setup lang="ts">
import axios from "axios";
function op() {
  console.log("op is execute")
  //发送一个`POST`请求
  axios({
    method:"POST",
    url:'http://127.0.0.1:8000/v1/login',
    data:{
      email:"admin@qq.com",
      password:"admin"
    }
  }).then(function (res){
    console.log(res)
  });
}
</script>

<template>
  <button @click="op">点这里</button>
</template>

<style scoped>

</style>