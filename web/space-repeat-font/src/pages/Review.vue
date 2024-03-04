<script setup lang="ts">
import request from "../utils/request.ts";
import {ref} from "vue";

const getAllReview=()=>{
  console.log("9999")
  request.get("/v1/review/", {}).then(res => {
    console.log("zzzz")
    console.log(res)
  }).catch(res=>{
    console.log("报错")
    console.log(res)
  })
  console.log("asas")
}

// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getAllReview()
  console.log("66666666666")
})

</script>

<template>
    <div class="common-layout">
      <el-container>
        <el-header>Header</el-header>
        <el-main style="height: 560px">Main</el-main>
        <el-footer >
          <el-row class="mb-4">
            <el-button class="bebig" size="large" type="success">Easy</el-button>
            <el-button class="bebig" size="large" type="primary">Good</el-button>
            <el-button class="bebig" size="large" type="warning">Hard</el-button>
            <el-button class="bebig" size="large" type="danger">Again</el-button>
          </el-row>
        </el-footer>
      </el-container>
    </div>


</template>

<style scoped>

.bebig {
  min-height: 40px;
  min-width: calc(90%/4);
  margin: auto;
  border: #e6a23c;
}
</style>