<template>

    <h1>Todo App</h1>
    <el-button type="primary" @click="listTodos">Primary</el-button>

</template>

<script lang="ts">
import { defineComponent } from 'vue'
import api from "../api";

export default defineComponent({
  data() {
    return {
      newTodo: '',
      todos: [] as string[]
    }
  },
  methods: {
    listTodos() {
      api.list().then(response=>{
        console.log(response)
      })
    },
  }
})
</script>


