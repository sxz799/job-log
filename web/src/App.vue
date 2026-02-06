<template>
  <div :class="isDark ? 'dark' : ''">
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100 transition-colors duration-300 font-sans pb-20">
      
      <!-- Header -->
      <header class="sticky top-0 z-50 bg-white/80 dark:bg-gray-800/80 backdrop-blur-md shadow-sm">
        <div class="max-w-md mx-auto px-4 h-14 flex items-center justify-between">
          <h1 class="text-lg font-bold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
            Job Log
          </h1>
          <button 
            @click="toggleDark"
            class="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
          >
            <el-icon :size="20">
              <component :is="isDark ? 'Sunny' : 'Moon'" />
            </el-icon>
          </button>
        </div>
      </header>

      <!-- Main Content -->
      <main class="max-w-md mx-auto px-4 pt-4">
        <transition name="fade" mode="out-in">
          <keep-alive>
            <component :is="activeComponent" />
          </keep-alive>
        </transition>
      </main>

      <!-- Bottom Navigation -->
      <nav class="fixed bottom-0 left-0 right-0 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 pb-safe">
        <div class="max-w-md mx-auto flex justify-around items-center h-16">
          <button 
            v-for="tab in tabs" 
            :key="tab.name"
            @click="activeName = tab.name"
            class="flex-1 flex flex-col items-center justify-center h-full space-y-1 relative"
            :class="activeName === tab.name ? 'text-blue-600 dark:text-blue-400' : 'text-gray-500 dark:text-gray-400'"
          >
            <div class="relative p-1 rounded-xl transition-all duration-300"
                 :class="activeName === tab.name ? 'bg-blue-50 dark:bg-blue-900/30' : ''">
              <el-icon :size="24">
                <component :is="tab.icon" />
              </el-icon>
            </div>
            <span class="text-xs font-medium">{{ tab.label }}</span>
            
            <!-- Active Indicator -->
            <span 
              v-if="activeName === tab.name"
              class="absolute top-0 w-8 h-1 bg-blue-600 dark:bg-blue-400 rounded-b-full shadow-[0_0_8px_rgba(37,99,235,0.5)]"
            ></span>
          </button>
        </div>
      </nav>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import TodoList from './components/TodoList.vue'
import Clipboard from "./components/Clipboard.vue"
import { DocumentCopy, List, Moon, Sunny } from '@element-plus/icons-vue'

// Tabs Configuration
const tabs = [
  { name: 'Clipboard', label: '剪切板', icon: DocumentCopy, component: Clipboard },
  { name: 'TodoList', label: '待办事项', icon: List, component: TodoList }
]

const activeName = ref('Clipboard')
const activeComponent = computed(() => {
  return tabs.find(t => t.name === activeName.value)?.component || Clipboard
})

// Dark Mode
const isDark = ref(false)
const toggleDark = () => {
  isDark.value = !isDark.value
  if (isDark.value) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
}

// Initialize Dark Mode based on system preference
onMounted(() => {
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
})
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Safe area padding for iPhone X+ */
.pb-safe {
  padding-bottom: env(safe-area-inset-bottom);
}
</style>
