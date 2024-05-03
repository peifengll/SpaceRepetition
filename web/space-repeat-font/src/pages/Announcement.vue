<template>
    <div class="notice-list">
      <h2>公告列表</h2>
      <ul class="notice-items">
        <li v-for="notice in paginatedNotices" :key="notice.id" class="notice-item">
          {{ notice.title }}
        </li>
      </ul>
  
      <div class="pagination">
        <button @click="prevPage" :disabled="currentPage === 1" class="pagination-btn">上一页</button>
        <span class="current-page">{{ currentPage }}</span>
        <button @click="nextPage" :disabled="currentPage === totalPages" class="pagination-btn">下一页</button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue';
  
  interface Notice {
    id: number;
    title: string;
  }
  
  const notices: Notice[] = [
    { id: 1, title: '公告 1' },
    { id: 2, title: '公告 2' },
    { id: 3, title: '公告 3' },
    { id: 4, title: '公告 4' },
    { id: 5, title: '公告 5' },
    { id: 6, title: '公告 6' },
    { id: 7, title: '公告 7' },
    // 添加更多公告
  ];
  
  const pageSize = 5;
  const currentPage = ref(1);
  
  const totalPages = computed(() => Math.ceil(notices.length / pageSize));
  
  const paginatedNotices = computed(() => {
    const startIndex = (currentPage.value - 1) * pageSize;
    const endIndex = startIndex + pageSize;
    return notices.slice(startIndex, endIndex);
  });
  
  const nextPage = () => {
    if (currentPage.value < totalPages.value) {
      currentPage.value++;
    }
  };
  
  const prevPage = () => {
    if (currentPage.value > 1) {
      currentPage.value--;
    }
  };
  </script>
  
  <style scoped>
  .notice-list {
    margin: 20px auto;
    width: 80%;
  }
  
  .notice-items {
    list-style: none;
    padding: 0;
  }
  
  .notice-item {
    background-color: #f0f0f0;
    padding: 10px;
    margin-bottom: 5px;
  }
  
  .pagination {
    margin-top: 10px;
    text-align: center;
  }
  
  .pagination-btn {
    margin: 0 5px;
    padding: 5px 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 3px;
    cursor: pointer;
  }
  
  .current-page {
    font-weight: bold;
  }
  </style>
  