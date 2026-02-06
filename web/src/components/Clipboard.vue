<template>
  <div class="space-y-4 pb-20">
    <!-- Input Card -->
    <div class="bg-white dark:bg-gray-800 rounded-2xl p-4 shadow-sm border border-gray-100 dark:border-gray-700">
      <div class="flex justify-between items-center mb-3">
        <h2 class="text-sm font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wider">新记录</h2>
        <div class="flex space-x-2">
          <el-button size="small" type="primary" plain @click="handleCopy({})" :icon="CopyDocument" circle />
          <el-button size="small" type="danger" plain @click="handleReset" :icon="CloseBold" circle />
        </div>
      </div>
      
      <el-input
        v-model="content"
        :autosize="{ minRows: 4, maxRows: 12 }"
        type="textarea"
        class="w-full !text-base"
        @change="updateClipboardData"
        placeholder="请输入..."
      />
      
      <div class="mt-4 flex justify-end">
        <el-button type="success" @click="addClipboardLog" :icon="Plus" class="!rounded-xl !px-6 w-full sm:w-auto">
          Record
        </el-button>
      </div>
    </div>

    <!-- History List -->
    <div class="space-y-3">
      <h3 class="text-sm font-semibold text-gray-500 dark:text-gray-400 px-1">历史记录</h3>
      
      <transition-group name="list" tag="div" class="space-y-3">
        <div 
          v-for="item in logs" 
          :key="item.id"
          class="bg-white dark:bg-gray-800 rounded-xl p-4 shadow-sm border border-gray-100 dark:border-gray-700 hover:shadow-md transition-shadow"
        >
          <div class="flex justify-between items-start mb-2">
            <span class="text-xs text-gray-400 font-mono">{{ item.create_at }}</span>
            <div class="flex space-x-1">
              <el-button 
                type="primary" 
                link
                size="small"
                :icon="CopyDocument" 
                @click="handleCopy(item)"
              />
              <el-popconfirm
                confirm-button-text="Yes"
                cancel-button-text="No"
                :icon="InfoFilled"
                title="Delete this record?"
                @confirm="handleDelete(item)"
                width="200"
              >
                <template #reference>
                  <el-button type="danger" link size="small" :icon="Delete" />
                </template>
              </el-popconfirm>
            </div>
          </div>
          
          <div 
            class="text-sm text-gray-700 dark:text-gray-300 whitespace-pre-wrap break-words font-mono bg-gray-50 dark:bg-gray-900/50 p-3 rounded-lg border border-gray-100 dark:border-gray-800"
            v-html="item.content"
          ></div>
        </div>
      </transition-group>

      <!-- Empty State -->
      <div v-if="logs.length === 0" class="text-center py-10 text-gray-400">
        <el-icon :size="48" class="mb-2 opacity-50"><DocumentCopy /></el-icon>
        <p class="text-sm">No history records</p>
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
import {onMounted, ref} from 'vue';
import {ElMessage} from 'element-plus';
import useClipboard from 'vue-clipboard3';
import {CloseBold, CopyDocument, Delete, InfoFilled, Plus} from '@element-plus/icons-vue';
import {add, del, get, list} from '../api/clipboard';

const {toClipboard} = useClipboard();

const content = ref('');
const logs = ref([]);
const ws = ref(null)

const page = ref(1)
const total = ref(0)
const pageSize = ref(5)


onMounted(() => {
  handleWebsocket();
  listClipboardLogs();
});


function handleWebsocket() {
  const url = import.meta.env.VITE_BASE_URL + '/api/clipboard/ws'
  // 把url中的https替换为wss:// http替换为ws://
  const wsUrl = url.replace('https://', 'wss://').replace('http://', 'ws://')
  ws.value = new WebSocket(wsUrl);
  ws.value.onopen = () => {
    console.log('ws链接成功!');
    getClipboardData();
  };
  ws.value.onmessage = (e: any) => {
    console.log('ws接收消息:',e.data);
    content.value = e.data
  };
  ws.value.onerror = (error: any) => {
    console.error('ws发生错误:', error);
  };
  ws.value.onclose = () => {
    console.log('ws关闭,10秒后重连!');
    setTimeout(() => {
      handleWebsocket();
    }, 1000);
  };
}

function listClipboardLogs() {
  list({
    page: page.value,
    pageSize: pageSize.value,
  }).then(response => {
    logs.value = response.data.list;
    total.value = response.data.total
  });
}

function addClipboardLog() {
  if (content.value === '') {
    ElMessage.info('当前内容为空!');
    return
  }
  add({content: content.value}).then(response => {
    if (response.success) {
      ElMessage.success(response.message);
      listClipboardLogs()
    } else {
      ElMessage.error(response.message);
    }
  });
}

function getClipboardData() {
  get().then(response => {
    if (response.success) {
      content.value = response.data.content;
    } else {
      ElMessage.error(response.message);
    }
  });
}

function updateClipboardData() {
  console.log(ws.value, ws.value.readyState, WebSocket.OPEN)
  if (ws.value && ws.value.readyState === WebSocket.OPEN) {
    ws.value.send(content.value)
    ElMessage.success({
      message: '更新成功',
      grouping: true,
    });
  } else {
    ElMessage.error('WebSocket 连接未打开，无法发送消息');
  }
}

function handleReset() {
  content.value = '';
}

function handleCopy(record: any) {
  if (record.content) {
    try {
      toClipboard(record.content);
      ElMessage.success('复制成功！');
    } catch (e) {
      ElMessage.error('复制失败！');
    }
  } else {
    try {
      toClipboard(content.value);
      ElMessage.success('复制成功！');
    } catch (e) {
      ElMessage.error('复制失败！');
    }
  }

}

function handleDelete(record: any) {
  del(record.id).then(response => {
    if (response.success) {
      listClipboardLogs();
      ElMessage.success(response.message);
    } else {
      ElMessage.error(response.message);
    }
  });
}

// region 处理分页
const handleSizeChange = (val: number) => {
  pageSize.value = val
  listClipboardLogs()
}

const handleCurrentChange = (val: number) => {
  page.value = val
  listClipboardLogs()
}
//endregion
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
</style>
