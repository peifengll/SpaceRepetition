<template>
  <h4 style="text-align: center;">训练数据、复习导出数据只会保存十条，导出新数据时请及时保存重要数据</h4>
  <h4 style="text-align: center;">复习数据条数少于1000的文件训练不会给结果的，请注意！</h4>
  <br>
  <el-table :data="tableData" style="width: 100%" >
    <el-table-column fixed prop="exportDate" label="数据导出日期" width="150" />
    <el-table-column prop="isTrained" label="是否训练" width="120" />
    <el-table-column prop="columns" label="总条数" width="120" />
    <el-table-column prop="dataStartTime" label="数据开始时间" width="120" />
    <el-table-column prop="state" label="状态" width="120" />
    <el-table-column prop="dataEndTime" label="数据结束时间" width="600" />
 
    
    <el-table-column fixed="right" label="Operations" width="120" >
      <template #default="scope">
       
     <el-button v-if="scope.row.revlogPath!='' "  link type="primary" size="small" @click="download(scope.row.revlogPath)">
          下载复习数据
        </el-button>
        <br>
        <el-button v-if="scope.row.revlogPath!='' "  link type="primary" @click="trainData(scope.row.revlogPath,scope.row.id)" size="small">
          训练
        </el-button>
        <br>
        <el-button v-if="scope.row.trainDataPath!='' " link type="primary" size="small" @click="download(scope.row.trainDataPath)">下载训练数据</el-button>

      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import { onBeforeMount } from "vue";
import request, { SelfUrl } from '@/utils/request.ts'

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
};
const tableData = ref([
  {
    "id": "3",
    "userId": "BLNOXd36vf",
    "revlogPath": "D:\\work\\code\\go\\SpaceRepetition\\storage\\revlog_datas/BLNOXd36vf/1714923922.csv",
    "trainDataPath": "",
    "dataStartTime": "2023-05-03T00:00:00+08:00",
    "dataEndTime": "2024-05-08T00:00:00+08:00",
    "exportDate": "2024-05-05T00:00:00+08:00",
    "columns": 33,
    "state": 0
  }
])
onBeforeMount(() => {
  request.get("v1/decks/exportinfos", {}).then((res: any) => {
    let data = res.data.data
    for(let i=0;i<data.length;i++){
       data[i].isTrained= (data[i].trainDataPath!="")
       data[i].exportDate =formatDate(data[i].exportDate)
       data[i].dataStartTime =formatDate(data[i].dataStartTime)
       data[i].dataEndTime =formatDate(data[i].dataEndTime)
       if( data[i].state==0 ){
        data[i].state="无任务"
       }else if(data[i].state==1){
        data[i].state="等待导出数据中"
       }else if(data[i].state==1){
        data[i].state="等待训练数据中"
       }

    }
    tableData.value = data
    console.log(data)
  })

});


const  download=(path:String)=>{
  window.location.href='http://localhost:8000/v1/file?filepath='+path
  // request.get("/v1/file", {
  // })

}

const  trainData=(path:String,epid:Number)=>{
  
  request.get("v1/decks/review/train?filepath="+path+"&epid=" +epid, {}).then(res=>{
    console.log("训练完毕")
  })

}


const handleClickDownExportInfo = (row:any) => {

  request.post("/v1/decks/exportfile", {
    file_path: row.revlogPath,

  }).then(res => {
    console.log(res)
   
  }).catch(error => {
   
  })


  console.log('click',row)
}

</script>