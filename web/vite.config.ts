import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default (({ mode }) => {

  return defineConfig({
    plugins: [vue()],
    server: {
      port: 6060,
      proxy: {
        '/dev-api': {
          target: "https://ddns.sxz799.xyz:3000/",
          //target: "http://127.0.0.1:3000/",
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/dev-api/, '')
        },
      }
    },
  })
})
