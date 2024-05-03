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
    <el-form :model="form" label-width="auto" style="max-width: 600px" label-position="left">
    <el-form-item label="用户名">
      <el-input v-model="form.name" />
    </el-form-item>
    <el-form-item label="性别">
      <el-select v-model="form.gender" placeholder="please select your gender">
        <el-option label="男" value=1 />
        <el-option label="女" value=2 />
      </el-select>
    </el-form-item>
   
    <el-form-item label="年龄" prop="age">
      <el-input  v-model.number="form.age" oninput="value=value.replace(/[^\d.]/g,'')" />
    </el-form-item>
    <el-form-item label="最大复习间隔" prop="interval">
      <el-input  v-model.number="form.interval" oninput="value=value.replace(/[^\d.]/g,'')" />
    </el-form-item>
    <el-form-item label="回忆成功的概率" prop="request_retention">
      <el-input  v-model.number="form.request_retention" oninput="value=value.replace(/[^\d.]/g,'')" />
    </el-form-item>
    
    <el-form-item label="复习参数">
      <el-input v-model="form.reviewparms" />
    </el-form-item>
    <el-form-item label="个性签名">
      <el-input v-model="form.introduction" type="textarea" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Create</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
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
import request,{SelfUrl}from '@/utils/request.ts'
import {ElMessage, ElMessageBox} from "element-plus";
import useStore from "../../stores";
const toHomeIndex = () => {
  // location.reload()
  window.location.href=SelfUrl
}

onBeforeMount(() => {
  console.log("组件挂载前");
});

import { reactive } from 'vue'

// do not use same name with ref
const form = reactive({
  gender:'',
  name: '',
  introduction:'',
  age:'',
  interval:'',
  reviewparms:'',
  request_retention:'',
})

const onSubmit = () => {
  console.log("form.name",form.name)
  console.log("form.gender",form.gender)
  console.log("form.introduction",form.introduction)
  console.log("form.age",form.age)
  console.log("form.interval",form.interval)
  console.log("form.reviewparms",form.reviewparms)
  console.log("form.request_retention",form.request_retention)

  request.put("/v1/user", {
    gender:Number(form.gender) ,
    age: Number(form.age) ,
    username:form.name,
    introduction:form.introduction,
    maxInterval:Number(form.interval)  ,
    weights:form.reviewparms,
    requestRetention:  parseFloat(form.request_retention),
  }).then(res => {
    console.log(res)
    if (res.status == 200) {
      console.log(res.data.data.accessToken)
      ElMessage({
        showClose: true,
        message: '修改成功',
        type: 'success',
      })
    }
  }).catch(error => {
    console.log(error)
    ElMessage.error('修改失败'+error)

  })

}


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
