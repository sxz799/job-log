<template>



      <el-card class="box-card">
        <template #header>
          <div class="card-header">

            <el-button size="small" type="success" @click="addClipboardData" :icon="Plus" round>记录</el-button>

            <el-button-group>
              <el-button size="small" type="primary" @click="copy" :icon="CopyDocument" round>复制</el-button>

              <el-button size="small" type="danger" @click="reset" :icon="CloseBold" round>清空</el-button>
            </el-button-group>
          </div>
        </template>
        <div class="card-body">
          <el-input
              v-model="content"
              :autosize="{ minRows: 8, maxRows: 45 }"
              type="textarea"
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
        <el-table-column fixed="right" :warp="false" width="105" label="操作">
          <template #default="scope">

              <el-button type="primary"  :icon="CopyDocument" circle @click="handleCopy(scope.row)"></el-button>
              <el-popconfirm
                  confirm-button-text="是"
                  cancel-button-text="否"
                  :icon="InfoFilled"
                  title="确定要删除吗?"
                  @confirm="handleDelete(scope.row)"
              >
                <template #reference>
                  <el-button type="danger" :icon="Delete" circle ></el-button>
                </template>
              </el-popconfirm>

          </template>
        </el-table-column>
      </el-table>
      <br>
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



</template>

<script setup lang="ts">
import {ref, onMounted, onBeforeUnmount} from 'vue';
import {ElMessage} from 'element-plus';
import useClipboard from 'vue-clipboard3';
import {CloseBold, CopyDocument, Delete, InfoFilled, Plus} from '@element-plus/icons-vue';
import {add,list, get, update, del} from '../api/clipboard';

const {toClipboard} = useClipboard();

const content = ref('');
const lastContent = ref('');
const logs = ref([]);
let timer: NodeJS.Timeout | null = null;
const page = ref(1)
const total = ref(0)
const pageSize = ref(5)

onMounted(() => {
  getClipboardData();
  listClipboardLogs();
  timer = setInterval(refreshData, 3000);
});

onBeforeUnmount(() => {
  if (timer) {
    clearInterval(timer);
  }
});

function listClipboardLogs() {
  list({
    page: page.value,
    pageSize: pageSize.value,
  }).then(response => {
    logs.value = response.data.list;
    total.value = response.data.total
  });
}

function addClipboardData() {
  if(content.value===''){
    ElMessage.info('当前内容为空!');
    return
  }
  add({content:content.value}).then(response => {
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
      if (response.data.content !== lastContent.value) {
        ElMessage.success(response.message);
        lastContent.value = response.data.content
      }
      content.value = response.data.content;
    } else {
      ElMessage.error(response.message);
    }
  });
}

function updateClipboardData() {
  update({content: content.value}).then(response => {
    if (response.success) {
      lastContent.value = content.value;
      ElMessage.success(response.message);
    } else {
      ElMessage.error(response.message);
    }
  });
}

function reset() {
  content.value = '';
}

function handleCopy(record) {
  try {
    toClipboard(record.content);
    ElMessage.success('复制成功！');
  } catch (e) {
    ElMessage.error('复制失败！');
  }
}

function handleDelete(record) {
  del(record.id).then(response => {
    if (response.success) {
      listClipboardLogs();
      ElMessage.success(response.message);
    } else {
      ElMessage.error(response.message);
    }
  });
}

function copy() {
  try {
    toClipboard(content.value);
    ElMessage.success('复制成功！');
  } catch (e) {
    ElMessage.error('复制失败！');
  }
}

function refreshData() {
  if (lastContent.value === content.value) {
    getClipboardData();
  } else if (content.value !== '') {
    updateClipboardData();
  }
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
