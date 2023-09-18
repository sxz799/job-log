import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default (({ mode }) => {

  return defineConfig({
    base:'/dist/',
    plugins: [vue()],
    server: {
      port: 6060,
      proxy: {
        '/dev-api': {
          target: "https://ddns.sxz799.xyz:3000",
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/dev-api/, '')
        },
      }
    },

  })
})
