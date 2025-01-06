

<template>
  <el-row>

    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <el-tabs :stretch="true" v-model="activeName">
        <el-tab-pane label="剪切板" name="Clipboard">  <Clipboard  /></el-tab-pane>
        <el-tab-pane label="待办事项" name="TodoList">  <TodoList  /></el-tab-pane>
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
