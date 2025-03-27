<script setup>
import axios from 'axios';
import { provide } from 'vue';
import Items from '../components/Items.vue';

const httpClient = axios.create({
  baseURL: `${location.protocol}//${location.host}/api`,
  timeout: 1000,
  headers: {
    'Content-Type': 'application/json',
  },
});

httpClient.interceptors.response.use(
  response => response,
  error => {
    switch (error.status) {
      case 400:
        alert('リクエストが不正です。');
        break;
      case 404:
        alert('アクセスしようとしたページが見つかりません。');
        break;
      case 500:
        alert('サーバー内部エラーが発生しました。時間を空けて再度アクセスしてください。');
        break;
      default:
        alert('予期せぬエラーが発生しました。時間を空けて再度アクセスしてください。')
        break;
    }
    return error;
  }
)

provide('httpClient', httpClient);
</script>

<template>
  <Items />
</template>