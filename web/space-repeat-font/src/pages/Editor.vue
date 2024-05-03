<template>
  <el-container class="layout-container-demo" style="height: 670px;">


    <el-aside width="300px" style="border-right: 1px solid #bfbff4;"
    >
      <div class="demo-input-size">
        <el-input
            v-model="searchContent"
            class="w-50 m-2"
            size="large"
            placeholder="Please Input"
            :prefix-icon="Search"
            style="width: 80%"
        />
        <el-icon style="margin-left: 25px;margin-top: 25px">
          <WindPower/>
        </el-icon>
      </div>
      <el-divider/>
      <el-scrollbar style="height: auto">
        <el-menu :default-openeds="['0']" style="height: auto;" v-for="(section,item) in sections">
          <el-sub-menu :index=item.toString()>
            <template #title>
              <el-popover placement="bottom" :width="50" trigger="contextmenu">
                <el-button style="width: 130px" @click="addCard(section.id)">插入卡片</el-button>
                <br>
                <el-button style="width: 130px" @click="renameSection(section.id)">重命名</el-button>

                <br>
                <el-button style="width: 130px" @click="delFolder(section.id)">删除</el-button>


                <template #reference>

                  <span>
                     <el-icon>
                  <message/>
                </el-icon>
                    <el-text>
                {{ section.name }}
                    </el-text>

                    </span>


                </template>

              </el-popover>


            </template>
            <el-menu-item :index="card.id" v-for="card in section.cards">
              <el-popover placement="bottom" :width="50" trigger="contextmenu">
                <el-button style="width: 130px" @click="delCard(card.id)">删除卡片</el-button>
                <br>
                <el-button v-if="CheckOn(card.onlearning)" style="width: 130px" @click="AddToReview(card.id)">加入复习
                </el-button>
                <!--                <br>-->
                <!--                <el-button style="width: 130px" @click="renameSection(section.id)">重命名</el-button>-->

                <!--                <br>-->
                <!--                <el-button style="width: 130px" @click="delFolder(section.id)">删除</el-button>-->


                <template #reference>

                  <span>

                    <el-text>
                <a :href="'#card_'+section.id+'_'+card.id">
                {{ card.font }}
              </a>
                    </el-text>
                    </span>
                </template>
              </el-popover>

            </el-menu-item>

          </el-sub-menu>
        </el-menu>
        <el-menu>
          <el-icon style="margin-left: 25px">
            <Plus/>
          </el-icon>
          <el-button style="border: 0" :onclick="addSection">
            添加章节
          </el-button>
        </el-menu>
      </el-scrollbar>
    </el-aside>

    <el-container>
      <el-dialog v-model="dialogTableVisible" title="添加卡片" style="min-height: 800px">
        <template #title>
          <h1>
            添加卡片
          </h1>
          <el-radio-group v-model="cardType">
            <el-radio :label="1">问答卡片</el-radio>
            <el-radio :label="2">挖空卡片</el-radio>
          </el-radio-group>
          <el-button :onclick="StartAddCard" style="margin-left:400px">确认添加</el-button>
        </template>
        <div v-show="cardType==1">


          <el-input contenteditable="true" minlength="172px" type="textarea" v-model="inputCardType1_font">

          </el-input>
          <el-divider/>

          <el-input contenteditable="true" minlength="172px" type="textarea" v-model="inputCardType1_back">

          </el-input>
        </div>
        <div v-show="cardType==2">
          <el-input contenteditable="true" minlength="452px" type="textarea" v-model="inputCardType2">

          </el-input>

        </div>

      </el-dialog>
      <el-header style="text-align: right; font-size: 12px">
        <div class="toolbar">
          <el-dropdown>
            <el-icon style="margin-right: 8px; margin-top: 1px"
            >
              <setting
              />
            </el-icon>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :onclick="alertIntr">修改简介</el-dropdown-item>
                <el-dropdown-item  :onclick="alertDeckName">修改牌组名</el-dropdown-item>
                <el-dropdown-item  :onclick="delDeck">删除牌组</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <span>操作</span>
        </div>
      </el-header>

      <el-main>


        <el-scrollbar>
          <div v-for="(section,item) in sections">
            <el-text style="margin: 25px;" size="large" class="w-150px mb-2" truncated>
              <el-icon>
                <Memo/>
              </el-icon>
              {{ section.name }}
            </el-text>
            <el-card class="box-card" style="margin: 20px" :id="'card_'+section.id+'_'+card.id"
                     v-for="card in section.cards" contenteditable="true">
              <!--            <template #header>-->
              <!--              <div class="card-header">-->
              <!--                <span>Card name</span>-->
              <!--                <el-button class="button" text>Operation button</el-button>-->
              <!--              </div>-->
              <!--            </template>-->
              <div>
                {{ card.font }}
                <el-divider/>
                {{ card.back }}
              </div>
              <!--            <template #footer>Footer content</template>-->
            </el-card>
          </div>

        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
import {ref, onBeforeMount} from 'vue'

import {Menu as IconMenu, Message, Setting} from '@element-plus/icons-vue'
import {Search} from '@element-plus/icons-vue'
import request from "@/utils/request.ts";

const searchContent = ref('')
const item = {
  date: '2016-05-02',
  name: 'Tom',
  address: 'No. 189, Grove St, Los Angeles',
}
const tableData = ref(Array.from({length: 20}).fill(item))
import {defineProps} from 'vue';
import {Console} from "inspector";
import {ElMessage, ElMessageBox} from "element-plus";
import { SelfUrl } from '../utils/request';


const deckId = defineProps(['id']);
const alertIntr = ()=>{
  console.log(deckId.id)
  open3(deckId.id)
  console.log("666")
}
const alertDeckName = ()=>{
  console.log(deckId.id)
  open4(deckId.id)
  console.log("666")
}

const delDeck = ()=>{
  console.log(deckId.id)
  open2(deckId.id)
  console.log("666")
}



const sections = ref([])


const open2 = (id: number) => {
  //  删除牌组
  ElMessageBox.confirm(
      '确定删除此牌组吗（删除之后词牌组所有记录均会被删除，该过程不可逆）',
      '删除牌组',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true,
      }
  )
      .then(() => {
        console.log(id)
        request.delete("/v1/decks/deck/" + id).then((res:any) => {
          if (res.data.code == 0) {
            console.log(res.data)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
            window.location.href=SelfUrl
            return
          }

        }).catch((res:any) => {
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

//  修改牌组简介对应的
const open3 = (id: number) => {
  ElMessageBox.prompt('输入新简介', '修改简介', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    center: true,

  }).then(({value}) => {
        if (value == null ||value=="") {
          ElMessage({
            type: 'info',
            message: '未检测到输入，更改未进行',
          })
          return
        }
        console.log("val: ", value)
        request.put("/v1/decks/deck", {
          id:  Number(id),
          introduction:value,
        }).then((res:any) => {
          ElMessage({
            type: 'success',
            message: `该牌组简介更改为:${value}`,
          })
          console.log(res.data)
          getDeckAndAllDetail()
        }).catch((res:any) => {
          console.log(res.data)
          ElMessage({
            type: 'info',
            message: 'option canceled',
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

//  修改牌组名字对应的
const open4 = (id: number) => {
  ElMessageBox.prompt('输入新名字', '重命名', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    center: true,

  }).then(({value}) => {
        if (value == null ||value=="") {
          ElMessage({
            type: 'info',
            message: '未检测到输入，更改未进行',
          })
          return
        }
        console.log("val: ", value)
        request.put("/v1/decks/deck", {
          id:  Number(id),
          name: value,
          introduction:null,
          floderid:null,
        }).then((res:any) => {
          ElMessage({
            type: 'success',
            message: `该牌组新名字更改为:${value}`,
          })
          console.log(res.data)
          getDeckAndAllDetail()
        }).catch((res:any) => {
          console.log(res.data)
          ElMessage({
            type: 'info',
            message: 'option canceled',
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


const getDeckAndAllDetail = () => {
  if (deckId.id == null || deckId.id == "") {
    ElMessage.error("bad request")
    return
  }
  request.get("/v1/decks/deck/" + deckId.id).then((res) => {
    // console.log(res.data.data.sections)
    sections.value = res.data.data.sections
    // console.log("aaa", sections.value)
  }).catch((res) => {
    console.log(res)
  })
}
// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getDeckAndAllDetail()

})


const openAddSection = () => {
  ElMessageBox.prompt('输入章节名字', '新增章节', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    center: true,

  })
      .then(({value}) => {
        if (value == null || value == "") {
          ElMessage({
            type: 'info',
            message: '未检测到输入，操作未进行',
          })
          return
        }

        console.log("val: ", value)
        request.post("/v1/decks/deck/section", {
          deckid: Number(deckId.id),
          name: value,
        }).then((res) => {
          ElMessage({
            type: 'success',
            message: `章节:${value}添加成功`,
          })
          console.log(res.data)
          getDeckAndAllDetail()
          // location.reload()
        }).catch((res) => {
          console.log(res)
          ElMessage({
            type: 'info',
            message: 'Add canceled',
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

const addSection = () => {
  openAddSection()
}

const renameSection = (id: number) => {
  openRenameSection(id)
}


const dialogTableVisible = ref(false)
const cardType = ref(1)
const inputCardType2 = ref('sadasdasdas')
const inputCardType1_font = ref('正面(问题)')
const inputCardType1_back = ref('背面(答案)')
let toAddSectionId = -1


const delCard = (id: number) => {
  openDelCard(id)
}

const openDelCard = (id: number) => {
  //  删除文件夹
  ElMessageBox.confirm(
      '确定删除该卡片吗',
      '删除卡片',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true,
      }
  )
      .then(() => {
        console.log(id)
        request.delete("/v1/decks/card/" + id).then((res) => {
          if (res.data.code == 0) {
            console.log(res.data)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
            getDeckAndAllDetail()
            // location.reload()
            return
          }

        }).catch((res) => {
          ElMessage({
            type: 'error',
            message: '删除失败:' + res.toString(),
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


const addCard = (id: number) => {
  console.log(id)
  toAddSectionId = id
  dialogTableVisible.value = true
}
const StartAddCard = () => {
  if (cardType.value == 1) {
    if (inputCardType1_font.value != "" && inputCardType1_back.value != "") {
      request.post("/v1/decks/card", {
        font: inputCardType1_font.value,
        originfont: inputCardType1_font.value + "###" + inputCardType1_back.value,
        back: inputCardType1_back.value,
        typeof: 1,
        deckid: Number(deckId.id),
        sectionid: toAddSectionId
      }).then((res) => {
        ElMessage.success("添加成功")
        getDeckAndAllDetail()
      }).catch((res) => {
        ElMessage.error("添加card失败", res)
      })
    } else {
      ElMessage.error("输入不能为空")
    }
  } else if (cardType.value == 2) {
    if (inputCardType2.value != "") {
      request.post("/v1/decks/card", {
        font: inputCardType2.value,
        originfont: inputCardType2.value,
        typeof: 2,
        deckid: Number(deckId.id),
        sectionid: toAddSectionId
      }).then((res) => {
        ElMessage.success("添加成功")
        getDeckAndAllDetail()
      }).catch((res) => {
        ElMessage.error("添加card失败", res)
      })
    } else {
      ElMessage.error("输入不能为空")
    }

  }
  inputCardType2.value = 'sadasdasdas'
  inputCardType1_font.value = '正面(问题)'
  inputCardType1_back.value = '背面(答案)'
  dialogTableVisible.value = false
}


const openRenameSection = (id: number) => {
  ElMessageBox.prompt('输入新名字', '重命名', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    center: true,

  })
      .then(({value}) => {
        if (value == null || value == "") {
          ElMessage({
            type: 'info',
            message: '未检测到输入，更改未进行',
          })
          return
        }
        console.log("val: ", value)
        request.put("/v1/decks/deck/section", {
          id: id,
          name: value,
        }).then((res) => {
          ElMessage({
            type: 'success',
            message: `该章节新名字更改为:${value}`,
          })
          console.log(res.data)
          getDeckAndAllDetail()
          // location.reload()
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


const openDel = (id: number) => {
  //  删除文件夹
  ElMessageBox.confirm(
      '确定删除该章节吗',
      '删除章节',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
        center: true,
      }
  )
      .then(() => {
        console.log(id)
        request.delete("/v1/decks/deck/section/" + id).then((res) => {
          if (res.data.code == 0) {
            console.log(res.data)
            ElMessage({
              type: 'success',
              message: '删除成功',
            })
            getDeckAndAllDetail()
            // location.reload()
            return
          }

        }).catch((res) => {
          ElMessage({
            type: 'error',
            message: '删除失败:' + res.toString(),
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
  openDel(id)
}

const CheckOn = (id: number) => {
  return id == 0;
}

const AddToReview = (id: number) => {
  request.put("/v1/decks/card/toreview", {
    id: parseInt(id),
  }).then((res) => {
    ElMessage({
      type: 'success',
      message: `加入复习成功`,
    })
    getDeckAndAllDetail()
  }).catch((res) => {
    ElMessage({
      type: 'error',
      message: '加入复习失败:' + res.toString(),
    })
  })
}

</script>

<style scoped>
.layout-container-demo .el-header {
  position: relative;
  /*background-color: var(--el-color-primary-light-7);*/
  color: var(--el-text-color-primary);
}

.layout-container-demo .el-aside {
  color: var(--el-text-color-primary);
  /*background: var(--el-color-primary-light-8);*/
}

.layout-container-demo .el-menu {
  border-right: none;
}

.layout-container-demo .el-main {
  padding: 0;
}

.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}


/*卡片部分*/
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: auto;
}

a {
  text-decoration: none; /* 去除默认的下划线 */
  outline: none; /* 去除旧版浏览器的点击后的外虚线框 */
  color: #000; /* 去除默认的颜色和点击后变化的颜色 */
}
</style>
