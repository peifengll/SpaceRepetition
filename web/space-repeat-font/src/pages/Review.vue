<script setup lang="ts">
import request from "../utils/request.ts";
import {Deck, Card, ReviewItem,insertCard} from "../utils/datastruct.ts";
import {ref} from 'vue'


let decks = [
  {
    "id": 10,
    "name": "测试牌组号",
    "num": 2,
    "cards": [
      {
        "id": 14,
        "record_id": 46,
        "font": "二号二号的测试卡片",
        "originfont": "二号二号的测试卡片###这是二号",
        "back": "这是二号",
        "deckid": 10,
        "typeof": 1
      },
      {
        "id": 15,
        "record_id": 47,
        "font": "二号测试",
        "originfont": "二号测试###这个的目的是为了验证跳转问题",
        "back": "这个的目的是为了验证跳转问题",
        "deckid": 10,
        "typeof": 1
      }
    ]
  },
] as [Deck];

const getAllReview = () => {
  request.get("/v1/review/", {}).then(res => {
    let reinfo = res.data.data
    for (let i = 0; i < reinfo.length; i++) {
      reinfo[i].num = reinfo[i].cards.length
    }

    decks = reinfo as [Deck]
    console.log("注意这个decks : ", decks)
    //  删掉占位的那一个
    options.value.splice(0, 1)
   
    for (let i = 0; i < decks.length; i++) {
      options.value.push({
            label: decks[i].name,
            value: decks[i].name,
            id: i,
          },
      )
      // console.log("iter::", Decks.value[i])
    }
    if (options.value.length>0){
      deckNowName.value=options.value[0].label
      NowReviewInfo.value={
        "carditem":0,
        "deckitem":0,
      } as ReviewItem;
      NoeDeckNum.value=decks[NowReviewInfo.value.deckitem].num
      ReviewCard.value=decks[NowReviewInfo.value.deckitem].cards[0] as Card
    }
  }).catch(res => {
    console.log("报错", res)
  })
}


const activeAnswer = ref(false)
const deckNowName = ref("")
const options = ref([{
  label: "测试",
  value: "测试",
  id: 0,
}])
// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getAllReview()
  isLoadFinish.value=true
  console.log(NowReviewInfo.value)
})



const isEasy = () => {
  console.log("easy")
  Op(4)
}

const isGood = () => {
  console.log("good")
  Op(3)
}

const isHard = () => {
  console.log("hard")
  Op(2)
}

const isAgain = () => {
  console.log("again")
  Op(1)
}
const ReviewCard=ref<Card>({
  "id": 0,
  "record_id": 0,
  "font": "占位！！",
  "originfont": "占位！！",
  "back": "占位！！",
  "deckid": 0,
  "typeof": 0
} as Card)
const NowReviewInfo=ref<ReviewItem>({
  "carditem":-1,
  "deckitem":-1,
} as ReviewItem)
const NoeDeckNum=ref(0)
const isLoadFinish =ref(false)
const handleUpdate = (value: string, option: any) => {
  console.log("handler uodate ", value)
  console.log("handler opt id ", option.id)
  console.log(option)
  console.log("ReviewCard: ",ReviewCard.value==null)
  NowReviewInfo.value.carditem=0
  NowReviewInfo.value.deckitem=option.id
  ReviewCard.value=decks[NowReviewInfo.value.deckitem].cards[NowReviewInfo.value.carditem] as Card
  deckNowName.value=option.label
  NoeDeckNum.value=decks[NowReviewInfo.value.deckitem].num
  // NowDeckId.value = option.id
  // NowDeckNum.value = decks[option.id].cards.length
  // NowReviewId.value = 0
}


const Op = (opt: number) => {
  console.log("opt !!!!!", opt ,ReviewCard.value)
  request.put("/v1/review/option", {
    id: ReviewCard.value.id,
    record_id: ReviewCard.value.record_id,
    opt: opt,
  }).then((res:any) => {
    let resp=res.data.data
    console.log("复习请求发起之后的返回值：",resp)
    decks[NowReviewInfo.value.deckitem].cards.shift()
    console.log("ss",decks[NowReviewInfo.value.deckitem])
    if (resp.finished == false) {
      console.log("false 执行")
      ReviewCard.value.record_id=resp.rid
      insertCard(decks[NowReviewInfo.value.deckitem],ReviewCard.value)
      ReviewCard.value=decks[NowReviewInfo.value.deckitem].cards[0] as Card
      console.log("复习操作请求之后：", res.data)
      console.log("开始")
    }else{
      decks[NowReviewInfo.value.deckitem].num--
      NoeDeckNum.value=decks[NowReviewInfo.value.deckitem].num
      // 如果消耗完了，就要换牌组了
      if (NoeDeckNum.value==0){
        for(let i=0;i<options.value.length;i++){
          if(options.value[i].id== NowReviewInfo.value.deckitem){
            options.value.splice(i,1)//把现在这个从里边删掉
            break;
          }
        }
        // 删完options之后，更换牌组来复习
        if(options.value.length>0){
          NowReviewInfo.value.deckitem=options.value[0].id
          ReviewCard.value= decks[NowReviewInfo.value.deckitem].cards[0] as Card
          NoeDeckNum.value=decks[NowReviewInfo.value.deckitem].num
          deckNowName.value=options.value[0].label
        }else{
          // 啥都没得就毕业把
          NowReviewInfo.value.deckitem=-1
          NowReviewInfo.value.carditem=-1
        }
      }else{
        // 没消耗完的时候
        ReviewCard.value= decks[NowReviewInfo.value.deckitem].cards[0] as Card
      }
      
    }
    console.log(res.data.data.finished)
  }).catch((res:any) => {

    console.log("报错了！！",res)
  })

  activeAnswer.value = false
}

</script>

<template>

  <div class="common-layout">
    <el-container v-if="isLoadFinish==true">
      <el-header>
        <div style="margin-left: calc(80%/2)" v-if="NowReviewInfo.deckitem!=-1"  >
          <n-popselect v-model:value="deckNowName" :options="options" trigger="click" @update:value="handleUpdate">
            <n-button>{{ deckNowName }}</n-button>
          </n-popselect>
          {{ NoeDeckNum }}
<!--          {{ NowReviewInfo.deckitem }}-->
        </div>

      </el-header>
      <el-main style="height: 560px" v-if="NowReviewInfo.deckitem==-1">
        <h1>复习完了小子 </h1>
      </el-main>
      <el-main style="height: 560px" v-if="NowReviewInfo.deckitem!=-1">

        <div v-if="NowReviewInfo.carditem!=-1">
          <div v-if="ReviewCard.typeof==1">
            <div>
              {{ ReviewCard.font}}
            </div>
            <div v-if="activeAnswer">
              <el-divider></el-divider>
              {{ ReviewCard.back }}
            </div>

          </div>
          <div v-if="ReviewCard.typeof==2">
            <div v-show="!activeAnswer">
              {{ ReviewCard.font }}
            </div>
            <div v-show="activeAnswer">
              {{ ReviewCard.originfont }}
            </div>

          </div>
        </div>
      </el-main>
    
      <div v-if="NowReviewInfo.deckitem!=-1">
        <n-switch class="my-switch" v-model:value="activeAnswer" size="large" >
        <template #icon>
          🤔
        </template>
      </n-switch>
      <el-row v-if="activeAnswer" class="mb-4 floating-row">
        <el-button class="bebig" size="large" type="success" @click="isEasy">Easy</el-button>
        <el-button class="bebig" size="large" type="primary" @click="isGood">Good</el-button>
        <el-button class="bebig" size="large" type="warning" @click="isHard">Hard</el-button>
        <el-button class="bebig" size="large" type="danger" @click="isAgain">Again</el-button>
      </el-row>
      </div>
     

    </el-container>
  </div>


</template>

<style scoped>

.bebig {
  min-height: 40px;
  min-width: calc(90% / 4);
  margin: auto;
  border: #e6a23c;
}

.floating-row {
  position: fixed;
  bottom: 0;
  width: 97.5%;
  padding: 10px;
}

.my-switch {
  position: fixed;
  bottom: 30%;
  right: 10%;
  transform: translateY(50%);
}

</style>