<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <el-button size="small" type="success" @click="addClipboardLog" :icon="Plus" round>记录</el-button>
        <el-button-group>
          <el-button size="small" type="primary" @click="handleCopy" :icon="CopyDocument" round>复制</el-button>
          <el-button size="small" type="danger" @click="handleReset" :icon="CloseBold" round>清空</el-button>
        </el-button-group>
      </div>
    </template>
    <div class="card-body">
      <el-input
          v-model="content"
          :autosize="{ minRows: 8, maxRows: 45 }"
          type="textarea"
          @change="updateClipboardData"
          placeholder="请输入内容"
      />
    </div>
  </el-card>
  <br>
  <el-table border :data="logs">
    <el-table-column fixed="left" width="90" prop="create_at" label="日期">
      <template #default="{ row }">
        <div v-html="row.create_at.substring(2).replace(' ', '<br/>')"></div>
      </template>
    </el-table-column>
    <el-table-column prop="content" label="内容">
      <template #default="{ row }">
        <div v-html="row.content.replace(/\n/g, '<br/>')"></div>
      </template>
    </el-table-column>
    <el-table-column fixed="right" :warp="false" width="110" label="操作">
      <template #default="scope">

        <el-button type="primary" :icon="CopyDocument" circle @click="handleCopy(scope.row)"></el-button>
        <el-popconfirm
            confirm-button-text="是"
            cancel-button-text="否"
            :icon="InfoFilled"
            title="确定要删除吗?"
            @confirm="handleDelete(scope.row)"
        >
          <template #reference>
            <el-button type="danger" :icon="Delete" circle></el-button>
          </template>
        </el-popconfirm>

      </template>
    </el-table-column>
  </el-table>
  <br>
  <el-pagination
      size="small"
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
  getClipboardData();
  handleWebsocket();
  listClipboardLogs();
});


function handleWebsocket() {
  const url = import.meta.env.VITE_BASE_URL + '/api/clipboard/ws'
  ws.value = new WebSocket(url);
  ws.value.onopen = () => {
    console.log('ws connected');
  };
  ws.value.onmessage = (e) => {
    console.log('ws message', e.data);
    content.value = e.data
  };
  ws.value.onerror = (error) => {
    console.error('WebSocket error:', error);
  };
  ws.value.onclose = (event) => {
    console.log('WebSocket connection closed:', event);
    // 可选：实现自动重连逻辑
    setTimeout(() => {
      console.log('Reconnecting...');
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
    ElMessage.success('更新成功');
  } else {
    ElMessage.error('WebSocket 连接未打开，无法发送消息');
  }
}

function handleReset() {
  content.value = '';
}

function handleCopy(record: any) {
  try {
    toClipboard(record.content);
    ElMessage.success('复制成功！');
  } catch (e) {
    ElMessage.error('复制失败！');
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


const handleSizeChange = (val: number) => {
  pageSize.value = val
  listClipboardLogs()
}

const handleCurrentChange = (val: number) => {
  page.value = val
  listClipboardLogs()
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
