<script setup lang="ts">
import RobotSwitch from "@/components/tabs/RobotSwitch.vue";
import {ref, onBeforeMount, onMounted} from 'vue'
import type {TabsPaneContext} from 'element-plus'
import request from "../utils/request.ts";
import {ElMessage, ElMessageBox} from 'element-plus'
import {Select} from '@element-plus/icons-vue'

import {Search} from '@element-plus/icons-vue'
import {bind} from "lodash-es";
import router from "@/router";
import { time } from "console";

const Decks = ref([])
const activeName = ref('')
const isFolderCanDelete = ref(false)


const input3 = ref('')

const getLoginData = () => {
  request.get("/v1/decks/", {}).then(res => {
    Decks.value = res.data.data
    activeName.value = Decks.value[0].id
  })

}

function formatDate(dateString: string): string {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = date.getMonth() + 1; // 月份从0开始计数，所以加1
  const day = date.getDate();
  return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
}

const annElessage=(index:number,maxl:number,data:any)=>{
  setTimeout(function () {
          ElNotification({
            title: data[index].title,
            dangerouslyUseHTMLString: true,
            message: '<strong>' + data[index].content+ '</strong> <br> <i>'+ formatDate(data[index].publishTime)+'</i>',
            offset: 100,
          })
          if (index<maxl-1){
            annElessage(index+1,maxl,data)
          }
        }, 4000);
}
const getAnns = () => {
  request.get("v1/pushanns", {}).then((res) => {
    let data = res.data.data
    if(data==null){
      return 
    }
    annElessage(0,data.length,data)
  })
}
// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getLoginData()
  
})
onMounted(()=>{
  getAnns()
})

const addFloder = (flodername: string) => {
  request.post("/v1/decks/floder", {name: flodername}).then((res) => {
    console.log(res.data)
  })
  console.log("增加文件夹", flodername)
}
const addDecks = (deckname: string, floderid: number) => {
  request.post("/v1/decks/deck", {
    name: deckname,
    floderid: floderid,
    introduction: "请编辑",
  }).then((res) => {
    console.log("增加牌组成功", deckname)
  })
  console.log(deckname)
}
const open = () => {
  ElMessageBox.prompt('请输入文件夹名称', '添加文件夹', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    // inputPattern:
    //     /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
    inputErrorMessage: 'Invalid Email',
  })
      .then(({value}) => {
        addFloder(value)
        ElMessage({
          type: 'success',
          message: `文件夹${value}添加成功`,
        })
        location.reload();
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: 'Input canceled',
        })
      })
}
const open2 = () => {
  ElMessageBox.prompt('请输入牌组名称', '添加牌组', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    // inputPattern:
    //     /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
    inputErrorMessage: 'Invalid Email',
  })
      .then(({value}) => {
        if (activeName.value == "") {
          ElMessage.error("请先选中一个文件夹噢")
          return
        }
        addDecks(value, activeName.value)
        ElMessage({
          type: 'success',
          message: `牌组${value}添加成功`,
        })
        console.log("当前的文件夹是", activeName.value)
        location.reload();
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: 'Input canceled',
        })
      })
}

const changeEditMode = () => {
  isFolderCanDelete.value = !isFolderCanDelete.value
  if (isFolderCanDelete.value == true) {
    ElMessage({
      type: 'success',
      message: `进入编辑模式，可删除文件夹`,
    })
  } else if (isFolderCanDelete.value == false) {
    ElMessage({
      type: 'success',
      message: `退出编辑模式`,
    })
  }
}

const open3 = (id: number) => {
  //  删除文件夹
  ElMessageBox.confirm(
      '确定删除文件夹吗',
      '删除文件夹',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true,
      }
  )
      .then(() => {
        console.log(id)
        request.delete("/v1/decks/floder/" + id).then((res) => {
          if (res.data.code == 0) {
            console.log(res.data)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
            location.reload()
            return
          }

        }).catch((res) => {
          ElMessage({
            type: 'error',
            message: '删除失败:'+res.toString(),
          })
        })


      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: 'Delete canceled',
        })
      })
}
const delFolder = (id: number) => {
  console.log(id)
  open3(id)

}

const open4 = (id: number) => {
  ElMessageBox.prompt('输入新名字', '重命名', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    center: true,

  })
      .then(({value}) => {
        if (value == null ||value=="") {
          ElMessage({
            type: 'info',
            message: '未检测到输入，更改未进行',
          })
          return
        }

        console.log("val: ", value)
        request.put("/v1/decks/floder", {
          id: id,
          name: value,
        }).then((res) => {
          ElMessage({
            type: 'success',
            message: `该文件夹新名字更改为:${value}`,
          })
          console.log(res.data)
          location.reload()
        }).catch((res) => {
          console.log(res.data)
          ElMessage({
            type: 'info',
            message: 'Delete canceled',
          })
        })

      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: 'Input canceled',
        })
      })
}

const renameFolder = (id: number) => {
  open4(id)
  console.log("id: ", id)
  // request.delete("/v1/decks/floder/" + id).then((res) => {
  //   console.log(res.data)
  // })
}

const toEditor=(row:any)=>{
  console.log("row and column: ",row.id)

  router.push({ name: 'editor', query: { id: row.id } });
}
</script>

<template>
  <div class="common-layout" >
   
<!--    <el-divider/>-->
    <el-container style="height: 50px">
      <el-container>
        <el-aside style="overflow: hidden" width="500px">
          <!-- <el-input
              v-model="input3"
              placeholder="搜索牌组"
              class="input-with-select"
          >

            <template #append>
              <el-button :icon="Search"/>
            </template>
          </el-input> -->

          <!--          <robot-switch style="width: 40px; height: auto" ></robot-switch>-->
        </el-aside>
        <el-main style="overflow: hidden;  display: flex; justify-content: flex-end; align-items: center;">
          <button type="button" class="button" @click="open">
            <span class="button__text">添加文件夹</span>
            <span class="button__icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" stroke-width="2"
                 stroke-linejoin="round" stroke-linecap="round" stroke="currentColor" height="24" fill="none"
                 class="svg">
                <line y2="19" y1="5" x2="12" x1="12"></line>
                <line y2="12" y1="12" x2="19" x1="5"></line>
            </svg>
        </span>
          </button>
          <p>&nbsp;&nbsp;&nbsp;</p>
          <button type="button" class="button" @click="open2">
            <span class="button__text">添加牌组</span>
            <span class="button__icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" stroke-width="2"
                 stroke-linejoin="round" stroke-linecap="round" stroke="currentColor" height="24" fill="none"
                 class="svg">
                <line y2="19" y1="5" x2="12" x1="12"></line>
                <line y2="12" y1="12" x2="19" x1="5"></line>
            </svg>
        </span>
          </button>

        </el-main>


      </el-container>
    </el-container>
    <el-tabs type="border-card" v-model="activeName" ref="tabsRef" style="height: 530px"
             @tab-remove="delFolder"
             addable
    >
      <template #addIcon>
        <el-icon v-show="isFolderCanDelete" @click="changeEditMode"><Select/></el-icon>
        <el-icon v-show="!isFolderCanDelete" @click="changeEditMode">
          <Edit/>
        </el-icon>
      </template>
      <el-tab-pane
          :closable=isFolderCanDelete
          v-for="item in Decks"
          :key="item.id"
          :label="item.name"
          :name="item.id"
      >
        <template #label>
        <span class="custom-tabs-label">

          <div v-show="activeName==item.id">
    <el-popover placement="bottom" :width="50" trigger="click">
      <el-button style="width: 130px" @click="renameFolder(item.id)">重命名</el-button>
      <br>
      <el-button style="width: 130px" @click="delFolder(item.id)">删除</el-button>
      <template #reference>
<!--        <el-button style="margin-right: 16px">-->
          <span @click="console.log('cao')">{{ item.name }}</span>
        <!--    <span v-show="activeName!=item.id" @click="console.log('cao')">{{ item.name }}</span>-->

        <!--        </el-button>-->
      </template>

    </el-popover>
          </div>

    <span v-show="activeName!=item.id" @click="console.log('cao')">{{ item.name }}</span>

        </span>
        </template>

        <el-table :data="item.decks" style="width: 100%" :key="item.name"  @row-click="toEditor" >
          <el-table-column prop="name" label="牌组名" width="180"/>
          <el-table-column prop="cardnum" label="卡片数" width="180"/>
          <el-table-column prop="learnnumber" label="已加入复习个数" width="180"/>
          <el-table-column prop="introduction" label="简介"/>

        </el-table>
      </el-tab-pane>

    </el-tabs>  
    <!--    <el-divider/>-->

  </div>


</template>


<style scoped>


.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}


.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}


.button {
  position: relative;
  width: 150px;
  height: 40px;
  cursor: pointer;
  display: flex;
  align-items: center;
  border: 1px solid #08080b;
  background-color: #bfbff4;
}

.button, .button__icon, .button__text {
  transition: all 0.3s;
}

.button .button__text {
  transform: translateX(30px);
  color: #fff;
  font-weight: 600;
}

.button .button__icon {
  position: absolute;
  transform: translateX(109px);
  height: 100%;
  width: 39px;
  background-color: #496ae4;
  display: flex;
  align-items: center;
  justify-content: center;
}

.button .svg {
  width: 30px;
  stroke: #fff;
}

.button:hover {
  background: #34974d;
}

.button:hover .button__text {
  color: transparent;
}

.button:hover .button__icon {
  width: 148px;
  transform: translateX(0);
}

.button:active .button__icon {
  background-color: #2e8644;
}

.button:active {
  border: 1px solid #2e8644;
}


</style>