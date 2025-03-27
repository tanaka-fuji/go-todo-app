<script setup>
import { inject, onBeforeMount, nextTick } from 'vue';
import { ref } from 'vue';

const httpClient = inject('httpClient');

const createItemForm = ref({});
const items = ref([]);
const text = ref('');

const textRules = [
  (value) => {
    if (1 <= value.trim().length && value.trim().length <= 100) {
      return true;
    }
    return "メッセージ本文は、1文字以上100文字以内で入力してください。";
  }
];

const createItem = async (e) => {
  e.preventDefault();
  const validResult = await createItemForm.value.validate();
  if (!validResult.valid) {
    return;
  }
  try {
    const res = await httpClient.post('/items', {
      text: text.value,
    });
    if (res.status === 201) {
      items.value.push(res.data);
      text.value = '';
      await nextTick();
      createItemForm.value.resetValidation();
    }
  } catch (error) {
    console.error(error);
  }
};

const deleteItem = async (id) => {
  try {
    const res = await httpClient.delete(`/items/${id}`);
    if (res.status === 204) {
      setTimeout(() => {
        items.value = items.value.filter((item) => item.ID !== id);
      }, 200);
    }
  } catch (error) {
    console.error(error);
  }
}

onBeforeMount(async () => {
  try {
    const res = await httpClient.get('/items');
    if (res.status === 200 && Array.isArray(res.data)) {
      items.value = res.data;
    }
  } catch (error) {
    console.error(error);
  }
});
</script>

<template>
  <v-card class="mx-auto" max-width="1000" flat>
    <v-toolbar color="blue">
      <v-btn icon="mdi-menu" color="white"></v-btn>
      <v-toolbar-title>ToDoアプリ</v-toolbar-title>
      <v-btn icon="mdi-magnify" color="white"></v-btn>
    </v-toolbar>

    <div class="pt-5 pb-2 bg-light-blue-lighten-5">
      <v-form ref="createItemForm" @submit.prevent="createItem" class="d-flex">
        <v-text-field v-model="text" :rules="textRules" class="ps-4"maxlength="100" rows="1"
          variant="outlined" @keydown.enter="createItem"></v-text-field>
        <v-btn type="submit" class="ma-3" color="info">
          <v-icon icon="mdi-plus" />
        </v-btn>
      </v-form>
    </div>

    <v-list>
      <v-list-item
        v-for="item in items"
        :key="item.ID"
        :value="item.ID"
        active-class="text-blue-darken-1"
        class="py-3"
        style="border-bottom: 1px solid #e0e0e0;"
        @click="deleteItem(item.ID)"
      >
        <template v-slot:prepend="{ isSelected }">
          <v-list-item-action>
            <v-icon v-if="isSelected" icon="mdi-checkbox-marked" color="blue" />
            <v-icon v-else icon="mdi-checkbox-blank-outline" color="black" />
          </v-list-item-action>
        </template>
        <v-list-item-title class="ms-4">{{ item.text }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-card>
</template>