<template>
  <el-row :gutter="24">
    <el-col :span="4" :xs="0"/>
    <el-col :span="16" :xs="24">
      <el-container>
        <el-header></el-header>
        <el-main>
          <el-row >
            <el-col :span="1" :xs="0"/>
            <el-col :span="20" :xs="18">
              <el-input
                  :resize="'none'"
                  type="textarea"
                  :autosize="{ minRows: 2, maxRows: 6 }"
                  v-model="newTodo"
                  placeholder="添加待办事项... "
              ></el-input>
            </el-col>
            <el-col :span="2" :xs="6" style="display: flex; flex-direction: column; justify-content: center; align-items: center;">
              <el-button
                  size="large"
                  :auto-insert-space="true"
                  @click="addTodo"
                  type="primary">添加</el-button>
            </el-col>
            <el-col :span="1" :xs="0"/>
          </el-row>
          <el-divider></el-divider>
          <el-table
              :data="todos"
              style="width: 100%"
          >
            <el-table-column prop="id" label="序号" width="60">
              <template #default="scope">
                <el-tag round size="small" :type="getStatusTagType(scope.row.status)">
                  {{ scope.row.id }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="content" label="内容">
              <template #default="scope">
                <el-input
                    :resize="'none'"
                    type="textarea"
                    :autosize="{ minRows: 1, maxRows: 6 }"
                    v-model="scope.row.content"
                    @blur="updateTodo(scope.row)"
                ></el-input>
              </template>
            </el-table-column>
            <el-table-column v-if="false" prop="status" label="状态" width="90">
              <template #default="scope">
                <el-tag round size="large" :type="getStatusTagType(scope.row.status)">
                  {{ getStatusLabel(scope.row.status) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column v-if="false" prop="create_at" label="创建时间" width="170"/>
            <el-table-column v-if="false" prop="finish_at" label="完成时间" width="170"/>
            <el-table-column v-if="false" prop="update_at" label="更新时间" width="180"/>
            <el-table-column label="操作" width="120">
              <template #default="scope">
                <el-button v-if="scope.row.status==='N'" type="primary"
                           @click="finishTodo(scope.row)" :icon="Check" circle>
                </el-button>
                <el-button v-if="scope.row.status==='Y'" type="warning"
                           @click="unfinishTodo(scope.row)" :icon="RefreshLeft" circle>
                </el-button>
                <el-popconfirm
                    confirm-button-text="确定"
                    cancel-button-text="取消"
                    title="确定要删除吗?"
                    @confirm="deleteTodo(scope.row.id)"
                >
                  <template #reference>
                    <el-button type="danger" circle :icon="Delete"></el-button>
                  </template>
                </el-popconfirm>

              </template>
            </el-table-column>
          </el-table>
          <el-pagination
              small
              :style="{'justify-content':'center'}"
              :background="true"
              :hide-on-single-page="false"
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[5, 10, 30, 50]"
              :total="total"
              layout=" sizes, prev, pager, next"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
          />
        </el-main>
        <el-footer></el-footer>
      </el-container>
    </el-col>
    <el-col :span="4" :xs="0"/>
  </el-row>


</template>

<script setup lang="ts">

import {list, add, del, update} from "../api";
import {ref} from "vue";
import {Check, Delete, RefreshLeft} from '@element-plus/icons-vue'
import {ElMessage} from "element-plus";


class Todo {
  id: number;
  content: string;
  status: string;
  finish_at: string;
  update_at: string
}

const todos = ref([]);
const newTodo = ref("");
const page = ref(1)
const total = ref(0)
const pageSize = ref(5)

function listTodos() {
  list({
    page: page.value,
    pageSize: pageSize.value,
  }).then(response => {
    console.log(response)
    todos.value = response.data.list
    total.value = response.data.total
  })
}

function addTodo() {
  if (!newTodo.value) {
    ElMessage("内容不可为空！")
    return
  }
  add({
    content: newTodo.value
  }).then(response => {
    if (response.success) {
      listTodos()
      ElMessage.success(response.message)
    } else {
      ElMessage.error(response.message)
    }
  })
}

function updateTodo(data: Todo) {
  data.update_at = formatDateTime()
  update(data.id, data).then(response => {
    if (response.success) {
      listTodos()
      ElMessage.success(response.message)
    } else {
      ElMessage.error(response.message)
    }
  })
}

function finishTodo(data: Todo) {
  data.status = 'Y'
  data.finish_at = formatDateTime()
  update(data.id, data).then(response => {
    if (response.success) {
      listTodos()
      ElMessage.success(response.message)
    } else {
      ElMessage.error(response.message)
    }
  })
}

function unfinishTodo(data: Todo) {
  data.status = 'N'
  data.finish_at = ' '
  update(data.id, data).then(response => {
    if (response.success) {
      listTodos()
      ElMessage.success(response.message)
    } else {
      ElMessage.error(response.message)
    }
  })
}

function deleteTodo(index: number) {
  del(index).then(response => {
    if (response.success) {
      listTodos()
      ElMessage.success(response.message)
    } else {
      ElMessage.error(response.message)
    }

  })
}


listTodos()

const handleSizeChange = (val: number) => {
  pageSize.value = val
  listTodos()
}

const handleCurrentChange = (val: number) => {
  page.value = val
  listTodos()
}


function getStatusLabel(status: string) {
  if (status === 'Y') {
    return '已完成';
  } else if (status === 'N') {
    return '未完成';
  } else {
    return '未知状态';
  }
}

function getStatusTagType(status: string) {
  return status === 'Y' ? 'success' : 'danger'; // 根据状态选择不同的标签类型
}


function formatDateTime() {
  const currentDate = new Date();
  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, '0'); // Month is zero-based, so add 1
  const day = String(currentDate.getDate()).padStart(2, '0');
  const hours = String(currentDate.getHours()).padStart(2, '0');
  const minutes = String(currentDate.getMinutes()).padStart(2, '0');
  const seconds = String(currentDate.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

</script>

<style>

</style>


