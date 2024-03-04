<script setup lang="ts">
import request from "../utils/request.ts";
import {ref} from 'vue'

// 定义 Card 结构
const cardSchema = {
  id: Number,
  record_id: String,
  font: String,
  originfont: String,
  back: String,
  deckid: String
};

// 定义 deck 结构
const DeckSchema = {
  id: Number,
  name: String,
  cards: [cardSchema]
};
const NowDeck=ref([])
const NowDeckNum=ref(0)
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
      "deckid": "1"
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
            id:i,
          }
      )
      // console.log("iter::", Decks.value[i])
    }
  }).catch(res => {
    console.log("报错", res)
  })
}
const deckNowName = ref("")
const options = ref([])
// 在页面加载时调用 API 获取数据
onBeforeMount(() => {
  getAllReview()
})



const handleUpdate = (value, option) => {
  NowDeck.value=Decks.value[option.id]
  NowDeckNum.value=NowDeck.value.cards.length
}

</script>

<template>

  <div class="common-layout">
    <el-container>
      <el-header>
        <n-popselect v-model:value="deckNowName" :options="options" trigger="click" @update:value="handleUpdate">
          <n-button >{{ deckNowName || '选择牌组' }}</n-button>
        </n-popselect>
        {{NowDeckNum}}
      </el-header>
      <el-main style="height: 560px">
        {{ NowDeck}}
      </el-main>
      <el-footer>
        <el-row class="mb-4">
          <el-button class="bebig" size="large" type="success">Easy</el-button>
          <el-button class="bebig" size="large" type="primary">Good</el-button>
          <el-button class="bebig" size="large" type="warning">Hard</el-button>
          <el-button class="bebig" size="large" type="danger">Again</el-button>
        </el-row>
      </el-footer>
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
</style>