<script setup lang="ts">
import request from "../utils/request.ts";
import {ref} from 'vue'
import {ElMessage} from "element-plus";

// 定义 Card 结构
const cardSchema = {
  id: Number,
  record_id: String,
  font: String,
  originfont: String,
  back: String,
  deckid: String,
  typeof: Number
};

// 定义 deck 结构
const DeckSchema = {
  id: Number,
  name: String,
  cards: [cardSchema]
};
const NowDeckId = ref(-1)
const NowReviewId = ref(-1)
const NowCard = ref({
  "id": 3,
  "record_id": "1",
  "font": "无所胃，我会被删掉",
  "originfont": "我其实是一号卡片",
  "back": "你其实是二号卡片",
  "deckid": "1",
  "typeof": 1,
})
const NowDeckNum = ref(0)
const Decks = ref([
  {
    "id": 1,
    "name": "shabi",
    "cards": [{
      "id": 3,
      "record_id": "1",
      "font": "无所胃，我会被删掉",
      "originfont": "我其实是一号卡片",
      "back": "你其实是二号卡片",
      "deckid": "1",
      "typeof": 1,
    }]
  }
]);
const getAllReview = () => {
  request.get("/v1/review/", {}).then(res => {
    // console.log("res.data.data", res.data.data)
    Decks.value = res.data.data
    // console.log("decks", Decks)
    // console.log("id", Decks.value[0].id)
    // console.log("len(decks)", Decks.value.length)
    for (let i = 0; i < Decks.value.length; i++) {
      options.value.push(
          {
            label: Decks.value[i].name,
            value: Decks.value[i].name,
            id: i,
          },
      )
      // console.log("iter::", Decks.value[i])
    }


  }).catch(res => {
    console.log("报错", res)
  })
}
const activeAnswer = ref(false)
const deckNowName = ref("")
const options = ref([])
// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getAllReview()
})


const handleUpdate = (value, option) => {
  NowDeckId.value = option.id
  NowDeckNum.value = Decks.value[option.id].cards.length
  NowReviewId.value = 0
}

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


const Op = (opt: number) => {
  console.log("id", Decks.value[NowDeckId.value].cards[NowReviewId.value].id)
  console.log("record_id", Decks.value[NowDeckId.value].cards[NowReviewId.value].record_id)
  console.log("opt", opt)
  request.put("/v1/review/option", {
    id: Decks.value[NowDeckId.value].cards[NowReviewId.value].id,
    record_id: Decks.value[NowDeckId.value].cards[NowReviewId.value].record_id,
    opt: opt,
  }).then((res) => {
    if (res.data.data.finished == false) {
      Decks.value[NowDeckId.value].cards = Decks.value[NowDeckId.value].cards.splice(NowReviewId.value, 1)
      NowDeckNum.value = NowDeckNum.value - 1
      if (NowDeckNum.value == 0) {
        //   如果这个牌组已经复习完了,删掉现在这个牌组
        Decks.value = Decks.value.splice(NowDeckId.value, 1)
        if (Decks.value.length==0){
          console.log("该结束了！！！！！！！！！！~~~~~~~~~~~~~~")
          return;
        }

        NowDeckId.value=(NowDeckId.value+1)%Decks.value.length
        NowDeckNum.value = Decks.value[NowDeckId.value].cards.length
      }
      NowReviewId.value = (NowReviewId.value + 1) % NowDeckNum.value

      console.log(Decks.value[NowDeckId.value].cards, NowReviewId.value)

      return
    }
    console.log(res.data.data.finished)
  }).catch((res) => {
    console.log(res.data)
  })
  NowReviewId.value = (NowReviewId.value + 1) % NowDeckNum.value
  activeAnswer.value = false
}

</script>

<template>

  <div class="common-layout">
    <el-container>
      <el-header>
        <div style="margin-left: calc(80%/2)">
          <n-popselect v-model:value="deckNowName" :options="options" trigger="click" @update:value="handleUpdate">
            <n-button>{{ deckNowName || '选择牌组' }}</n-button>
          </n-popselect>
          {{ NowDeckNum }}
        </div>

      </el-header>
      <el-main style="height: 560px">

        <!--        {{Decks}}-->
        <div v-if="NowReviewId!=-1">
          <div v-if="Decks[NowDeckId].cards[NowReviewId].typeof==1">
            <div>
              {{ Decks[NowDeckId].cards[NowReviewId].font }}
            </div>
            <div v-if="activeAnswer">
              <el-divider></el-divider>
              {{ Decks[NowDeckId].cards[NowReviewId].font }}
            </div>

          </div>
          <div v-if="Decks[NowDeckId].cards[NowReviewId].typeof==2">
            <div v-show="!activeAnswer">
              {{ Decks[NowDeckId].cards[NowReviewId].font }}
            </div>
            <div v-show="activeAnswer">
              {{ Decks[NowDeckId].cards[NowReviewId].originfont }}
            </div>

          </div>
        </div>
      </el-main>
      <n-switch class="my-switch" v-model:value="activeAnswer" size="large">
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