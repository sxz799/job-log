import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

export default (({ mode }) => {

  return defineConfig({
    plugins: [vue()],
    server: {
      port: 6060,
      proxy: {
        '/prod-api': {
          target: "http://127.0.0.1:4000",
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/prod-api/, '')
        },
      }
    },

  })
})
