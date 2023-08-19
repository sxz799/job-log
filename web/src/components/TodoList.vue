<template>
  <div>
    <el-input type="textarea" v-model="newTodo" placeholder="Add a new task"></el-input>
    <el-button @click="addTodo" type="primary">Add</el-button>
    <el-divider></el-divider>
    <el-table :data="todos" stripe style="width: 100%">
      <el-table-column prop="id" label="序号" width="120"/>
      <el-table-column prop="message" label="内容" width="300"/>
      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
          <el-button type="danger" size="small" @click="deleteTodo(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">

import {list, add, del} from "../api";
import {ref} from "vue";

const todos = ref([]);
const newTodo = ref("");

function listTodos() {
  list().then(response => {
    todos.value = response.data.data
  })
}

function addTodo() {
  add({
    message: newTodo.value
  }).then(response => {
    if (response.data.success) {
      listTodos()
    }
  })
}

function deleteTodo(index: number) {
  del(index).then(response => {
    if (response.data.success) {
      listTodos()
    }
  })
}

listTodos()
</script>


