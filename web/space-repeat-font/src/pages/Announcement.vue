<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column prop="publishTime" label="Date" width="180" />
    <el-table-column prop="title" label="标题" width="180" />
    <el-table-column prop="content" label="内容" />
    <el-table-column fixed="right" label="Operations" width="180" >
      <template #default="scope">
    
        <el-button v-if="scope.row.revlogPath!='' "  link type="primary" @click="onRead(scope.row)" size="small">
          确认
        </el-button>
       
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import {ref, onBeforeMount, onMounted} from 'vue'
import request from "../utils/request.ts";
onBeforeMount(() => {
  getAnns()
  
})
function formatDate(dateString: string): string {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = date.getMonth() + 1; // 月份从0开始计数，所以加1
  const day = date.getDate();
  return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
}
const getAnns = () => {
  request.get("v1/pushanns", {}).then((res) => {
    let data = res.data.data
    if(data==null){
      tableData.value=data
      return 
    }
    for (let i=0;i<data.length;i++){
      data[i].publishTime  =formatDate( data[i].publishTime )
    }
    tableData.value=data
    // annElessage(0,data.length,data)
  })
}
const onRead =(row:any)=>{

  console.log(row)
  request.post("/v1/readann", {
    id:row.id,
  }).then(res => {
    console.log(res)
    getAnns()
  }).catch(error => {

  })

}
const tableData = ref( [
  {
    "id": 1,
    "admin_id": 2,
    "title": "测试公告",
    "content": "这只是一条测试",
    "publish_time": "2024-04-16 21:24:01"
  },
  {
    "id": 2,
    "admin_id": 2,
    "title": "测试公告2",
    "content": "这也只是一条测试",
    "publish_time": "2024-04-16 21:30:08"
  },
  {
    "id": 4,
    "admin_id": 2,
    "title": "测试公告2",
    "content": "这也只是一条测试",
    "publish_time": "2024-05-03 16:17:06"
  },
  {
    "id": 5,
    "admin_id": 2,
    "title": "测试公告2",
    "content": "这也只是一条测试",
    "publish_time": "2024-05-03 16:17:57"
  },
  {
    "id": 6,
    "admin_id": 2,
    "title": "测试公告44",
    "content": "这也只是一条测试俄日一",
    "publish_time": "2024-05-03 18:14:36"
  }
])

</script>
