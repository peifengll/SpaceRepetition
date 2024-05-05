<template>
    <el-table
      :data="tableData"
      :default-sort="{ prop: 'date', order: 'descending' }"
      style="width: 100%"
    >
      <el-table-column prop="id" label="索引" sortable width="180" />
      <el-table-column prop="card_num" label="复习卡片数" sortable width="180" />
      <el-table-column prop="easy_num" label="easy评分次数" sortable width="180" />
      <el-table-column prop="good_num" label="good评分次数" sortable width="180" />
      <el-table-column prop="again_num" label="again评分次数" sortable width="180" />
      <el-table-column prop="hard_num" label="hard评分次数" sortable width="180" />
      <el-table-column prop="record_date" label="日期" sortable width="180" />
      <!-- <el-table-column prop="name" label="Name" width="180" />
      <el-table-column prop="address" label="Address" :formatter="formatter" /> -->
    </el-table>
    <div class="example-pagination-block">
   
    
    <el-pagination layout="prev, pager, next" :total="50" />
  </div>
  </template>
  
  <script lang="ts" setup>
  import type { TableColumnCtx } from 'element-plus'
  
  interface User {
    id :number
    card_num:number
    easy_num:number
    good_num:number
    again_num:number
    hard_num:number
    record_date:string
  }
  
  const formatter = (row: User, column: TableColumnCtx<User>) => {
    return row.address
  }
  const pagenum=ref(5)
  const tableData: User[] = [
  {
    "id": 4,
    "card_num": 10,
    "easy_num": 5,
    "good_num": 3,
    "again_num": 1,
    "hard_num": 1,
    "record_date": "2024-05-01",
  },
  {
    "id": 5,
    "card_num": 15,
    "easy_num": 8,
    "good_num": 4,
    "again_num": 2,
    "hard_num": 1,
    "record_date": "2024-04-029",
  },
  {
    "id": 6,
    "card_num": 12,
    "easy_num": 6,
    "good_num": 3,
    "again_num": 2,
    "hard_num": 1,
    "record_date": "2024-04-30",
  },
  {
    "id": 7,
    "card_num": 20,
    "easy_num": 10,
    "good_num": 6,
    "again_num": 3,
    "hard_num": 1,
    "record_date": "2024-05-01",
  },
  {
    "id": 8,
    "card_num": 18,
    "easy_num": 9,
    "good_num": 5,
    "again_num": 3,
    "hard_num": 1,
    "record_date": "2024-05-02",
  },
  {
    "id": 9,
    "card_num": 22,
    "easy_num": 11,
    "good_num": 7,
    "again_num": 3,
    "hard_num": 1,
    "record_date": "2024-05-03",
  },
  {
    "id": 10,
    "card_num": 25,
    "easy_num": 12,
    "good_num": 8,
    "again_num": 4,
    "hard_num": 1,
    "record_date": "2024-05-04",
  },
  
]
  </script>
  <style scoped>
  .example-pagination-block + .example-pagination-block {
    margin-top: 10px;
  }
  .example-pagination-block .example-demonstration {
    margin-bottom: 16px;
  }
  </style>