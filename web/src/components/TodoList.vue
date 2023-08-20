<template>
  <el-row :gutter="24">
    <el-col :span="4"/>
    <el-col :span="16">
      <el-container>
        <el-header></el-header>
        <el-main>
          <el-row :gutter="24">
            <el-col :span="20">
              <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 6 }" v-model="newTodo" placeholder="Add a new task"></el-input>
            </el-col>
            <el-col :span="4">
              <el-button @click="addTodo" type="primary">添加</el-button>
            </el-col>
          </el-row>
          <el-divider></el-divider>
          <el-table
              :data="todos"
              style="width: 100%"
          >
            <el-table-column fixed="left" prop="id" label="序号" width="80"/>
            <el-table-column prop="content" label="内容" width="300">
              <template #default="scope">
                <el-input
                    type="textarea"
                    :autosize="{ minRows: 1, maxRows: 4 }"
                    v-model="scope.row.content"
                    @blur="updateTodo(scope.row)"
                ></el-input>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="90">
              <template #default="scope">
                <el-tag round size="large" :type="getStatusTagType(scope.row.status)">
                  {{ getStatusLabel(scope.row.status) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="create_at" label="创建时间" width="170"/>
            <el-table-column prop="finish_at" label="完成时间" width="170"/>
<!--            <el-table-column prop="update_at" label="更新时间" width="180"/>-->
            <el-table-column align="center" fixed="right" label="操作" width="150">
              <template #default="scope">
                <el-button v-show="scope.row.status==='N'" type="primary" size="small" @click="finishTodo(scope.row.id)">完成</el-button>
                <el-button type="danger" size="small" @click="deleteTodo(scope.row.id)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-main>
        <el-footer></el-footer>
      </el-container>
    </el-col>
    <el-col :span="4"/>
  </el-row>






</template>

<script setup lang="ts">

import {list, add, del, finish, update} from "../api";
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
    content: newTodo.value
  }).then(response => {
    if (response.data.success) {
      listTodos()
    }
  })
}

function updateTodo(data:{}) {
  update(data).then(response => {
    if (response.data.success) {
      listTodos()
    }
  })
}

function finishTodo(id: number) {
  finish(id).then(response => {
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

function getStatusLabel(status:string) {
  if (status === 'Y') {
    return '已完成';
  } else if (status === 'N') {
    return '未完成';
  } else {
    return '未知状态';
  }
}
function getStatusTagType(status:string) {
  return status === 'Y' ? 'success' : 'danger'; // 根据状态选择不同的标签类型
}


</script>

<style>

</style>


