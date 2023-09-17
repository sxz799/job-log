

<template>
  <el-row :gutter="24">

    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <el-tabs v-model="activeName" class="demo-tabs">
        <el-tab-pane label="工作日志" name="Clipboard">  <Clipboard msg="Clipboard" /></el-tab-pane>
        <el-tab-pane label="待办事项" name="TodoList">  <TodoList msg="TodoList" /></el-tab-pane>
      </el-tabs>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>


</template>

<script setup lang="ts">
import TodoList from './components/TodoList.vue'
import Clipboard from "./components/Clipboard.vue";
import {ref} from "vue";

const activeName = ref('Clipboard')
</script>

<style scoped>

</style>
