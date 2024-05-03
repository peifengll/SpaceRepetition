<template>
  <el-menu
      style="width: 100%"
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      :route="true"
      @select="handleSelect"
  >
    <el-menu-item index="0" @click="toHomeIndex">
      <img
          style="width: 100px"
          src="https://element-plus.org/images/element-plus-logo.svg"
          alt="Element logo"
      />
    </el-menu-item>
    <div class="flex-grow"/>
    <el-menu-item index="3"  @click="toReviewPage" >开始学习</el-menu-item>
    <el-menu-item index="3"  @click="openConfirmLoginOut" >登出</el-menu-item>
    <el-sub-menu index="2">
      <template #title>Workspace</template>
      <el-menu-item index="2-1">item one</el-menu-item>
      <el-menu-item index="2-2">item two</el-menu-item>
      <el-menu-item index="2-3">item three</el-menu-item>
      <el-sub-menu index="2-4">
        <template #title>item four</template>
        <el-menu-item index="2-4-1">item one</el-menu-item>
        <el-menu-item index="2-4-2">item two</el-menu-item>
        <el-menu-item index="2-4-3">item three</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>
  </el-menu>
  <el-button type="primary" style="margin-left: 16px" @click="drawer = true">
    open
  </el-button>

  <el-drawer v-model="drawer" title="I am the title" :with-header="false">
    <span>Hi there!</span>
  </el-drawer>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import { onBeforeMount} from "vue";
const drawer = ref(false)

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
import {SelfUrl}from '@/utils/request.ts'
import {ElMessage, ElMessageBox} from "element-plus";
import useStore from "../../stores";
const toHomeIndex = () => {
  // location.reload()
  window.location.href=SelfUrl
}

onBeforeMount(() => {
  console.log("组件挂载前");
});


const loginOut = () => {
  const store = useStore();
  store.token=""
  // location.reload()
  window.location.href=SelfUrl
}
const open1 = () => {
  ElNotification({
    title: 'Title',
    message: h('i', { style: 'color: teal' }, 'This is a reminder'),
  })
}

const openConfirmLoginOut = () => {
  ElMessageBox.confirm(
    '确认退出登陆吗?',
    'Warning',
    {
      confirmButtonText: 'Yes',
      cancelButtonText: 'Cancel',
      type: 'warning',
    }
  )
    .then(() => {
      loginOut()

      ElMessage({
        type: 'success',
        message: 'loginout completed',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'loginout canceled',
      })
    })
}



const toReviewPage = () => {
  // location.reload()
  window.location.href=SelfUrl+"review"
}


</script>

<style>
.flex-grow {
  flex-grow: 1;

}
</style>
