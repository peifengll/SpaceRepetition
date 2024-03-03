<script setup lang="ts">

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