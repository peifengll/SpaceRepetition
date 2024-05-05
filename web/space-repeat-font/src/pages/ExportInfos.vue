<template>
    <h4 style="text-align: center;">训练数据、复习导出数据只会保存一个月，请及时保存重要数据</h4>
    <br>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column fixed prop="export_date" label="数据导出日期" width="150" />
      <el-table-column prop="isTrained" label="是否训练" width="120" />
      <el-table-column prop="columns" label="总条数" width="120" />
      <el-table-column prop="start" label="数据开始时间" width="120" />
      <el-table-column prop="end" label="数据结束时间" width="600" />
      <el-table-column fixed="right" label="Operations" width="120">
        <template #default>
          <el-button link type="primary" size="small" @click="handleClick">
            下载复习数据
          </el-button>
          <br>
          <el-button link type="primary" size="small">
             训练
          </el-button>
          <br>
          <el-button link type="primary" size="small">下载训练数据</el-button>
          
        </template>
      </el-table-column>
    </el-table>
  </template>
  
  <script lang="ts" setup>
  const handleClick = () => {
    console.log('click')
  }
  
  const tableData = [
    {
        export_date: '2024-05-03',
      isTrained: false,
      columns: 500,
      start: '2016-05-03',
      end: '2023-05-03',
      zip: 'CA 90036',
      tag: 'Home',
    }
   
  ]
  </script>
  