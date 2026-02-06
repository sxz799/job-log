<template>
  <div class="space-y-4 pb-20">
    <!-- Input Card -->
    <div class="bg-white dark:bg-gray-800 rounded-2xl p-4 shadow-sm border border-gray-100 dark:border-gray-700">
      <div class="flex gap-2">
        <el-input
          v-model="newTodo"
          :autosize="{ minRows: 1, maxRows: 3 }"
          type="textarea"
          class="flex-1 !text-base"
          placeholder="有什么需要记录的吗?"
          resize="none"
          @keyup.enter.ctrl="addTodo" 
        />
        <el-button 
          type="primary" 
          @click="addTodo" 
          class="!h-auto !px-4 !rounded-xl"
          :disabled="!newTodo.trim()"
        >
          添加
        </el-button>
      </div>
    </div>

    <!-- Todo List -->
    <div class="space-y-3">
      <h3 class="text-sm font-semibold text-gray-500 dark:text-gray-400 px-1">待办列表</h3>
      
      <transition-group name="list" tag="div" class="space-y-3">
        <div 
          v-for="item in todos" 
          :key="item.id"
          class="bg-white dark:bg-gray-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-gray-700 hover:shadow-md transition-all duration-300"
          :class="{'opacity-75 bg-gray-50 dark:bg-gray-800/50': item.status === 'Y'}"
        >
          <div class="flex items-start gap-3">
            <!-- Status Checkbox -->
            <button 
              @click="item.status === 'Y' ? unfinishTodo(item) : finishTodo(item)"
              class="mt-1 w-6 h-6 rounded-full border-2 flex items-center justify-center transition-colors duration-200"
              :class="item.status === 'Y' ? 'bg-green-500 border-green-500' : 'border-gray-300 dark:border-gray-600 hover:border-blue-500'"
            >
              <el-icon v-if="item.status === 'Y'" class="text-white" :size="14"><Check /></el-icon>
            </button>

            <!-- Content -->
            <div class="flex-1 min-w-0">
              <el-input
                v-model="item.content"
                type="textarea"
                autosize
                resize="none"
                class="w-full !bg-transparent !border-0 !shadow-none !p-0 text-sm leading-relaxed"
                :class="{'line-through text-gray-400': item.status === 'Y', 'text-gray-700 dark:text-gray-200': item.status === 'N'}"
                @blur="updateTodo(item)"
              />
              <div class="flex items-center gap-2 mt-2">
                 <span class="text-xs text-gray-400 font-mono">#{{ item.id }}</span>
                 <span v-if="item.finish_at && item.status === 'Y'" class="text-xs text-green-500/70">Done {{ item.finish_at.split(' ')[0] }}</span>
              </div>
            </div>

            <!-- Delete Action -->
            <el-popconfirm
              confirm-button-text="Yes"
              cancel-button-text="No"
              title="要删除记录吗?"
              @confirm="deleteTodo(item.id)"
              width="200"
            >
              <template #reference>
                <button class="text-gray-400 hover:text-red-500 transition-colors p-1">
                  <el-icon :size="18"><Delete /></el-icon>
                </button>
              </template>
            </el-popconfirm>
          </div>
        </div>
      </transition-group>

      <!-- Empty State -->
      <div v-if="todos.length === 0" class="text-center py-10 text-gray-400">
        <el-icon :size="48" class="mb-2 opacity-50"><Check /></el-icon>
        <p class="text-sm">暂无记录</p>
      </div>
    </div>

    <!-- Pagination -->
    <div class="flex justify-center pt-4">
      <el-pagination
        size="small"
        background
        layout="prev, pager, next"
        :current-page="page"
        :page-size="pageSize"
        :total="total"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import {list, add, del, update} from "../api/todo";
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
      newTodo.value = "" // Clear input on success
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

<style scoped>
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

/* Customize Element Input to look like plain text */
:deep(.el-textarea__inner) {
  box-shadow: none !important;
  background: transparent !important;
  padding: 0 !important;
  border: none !important;
}
</style>
