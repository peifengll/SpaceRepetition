<template>
  <el-table :data="tableData" :default-sort="{ prop: 'date', order: 'descending' }" style="width: 100%">
    <el-table-column prop="id" label="索引" sortable width="180" />
    <el-table-column prop="cardNum" label="复习卡片数" sortable width="180" />
    <el-table-column prop="easyNum" label="easy评分次数" sortable width="180" />
    <el-table-column prop="goodNum" label="good评分次数" sortable width="180" />
    <el-table-column prop="againNum" label="again评分次数" sortable width="180" />
    <el-table-column prop="hardNum" label="hard评分次数" sortable width="180" />
    <el-table-column prop="recordDate" label="日期" sortable width="180" />
    <!-- <el-table-column prop="name" label="Name" width="180" />
      <el-table-column prop="address" label="Address" :formatter="formatter" /> -->
  </el-table>
  <div class="example-pagination-block">


    <el-pagination :current-page="searchData.current" :page-size="searchData.limit" :total="total" :pager-count="6"
      style="text-align: center;margin-top: 20px;" layout="jumper, prev, pager, next, total"
      @current-change="getData" />

  </div>
</template>

<script lang="ts" setup>
import type { TableColumnCtx } from 'element-plus'
import { onBeforeMount } from "vue";
import request, { SelfUrl } from '@/utils/request.ts'

const searchData = reactive({
  current: 1,
  limit: 5
})

const total = ref(10)
// searchData-向后端分页查询的对象，即当前页和每页总数


//传入分页参数
function pageQuery(current:number,limit:number){
	// 模仿分页查询，将表格的数据裁切一下
	
	//     1     2     3
	//下标 0-9 10-19 20-29
	let begin=current*limit-limit
	//这里不减一是因为，slice方法裁切是左闭右开数组
	let end=current*limit
	tableData.value=storgedata.value.slice(begin,end)
}

function getData(val = 1){
	searchData.current=val
	// 先把数据搞上
  
  // tableData.value=storgedata.value.slice(0,5)
	pageQuery(searchData.current,searchData.limit)
}




const tableData = ref([
  {
    "id": 4,
    "card_num": 10,
    "easy_num": 5,
    "good_num": 3,
    "again_num": 1,
    "hard_num": 1,
    "record_date": "2024-05-01",
  },

])

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
};
const storgedata=ref([])

onBeforeMount(() => {
  request.get("v1/review/data", {}).then((res: any) => {
    let data = res.data.data
    
    for (let i = 0; i < data.length; i++) {
      data[i].recordDate = formatDate(data[i].recordDate)

    }
    total.value=data.length
    storgedata.value=data
    tableData.value=storgedata.value.slice(0,5)
    console.log(data)
  })

});

</script>
<style scoped>
.example-pagination-block+.example-pagination-block {
  margin-top: 10px;
}

.example-pagination-block .example-demonstration {
  margin-bottom: 16px;
}
</style>