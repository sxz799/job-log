import { defineConfig ,loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'



export default (({ mode }) => {
  const env = loadEnv(mode, process.cwd())
  return defineConfig({
    base: env.VITE_BASE_PATH,
    plugins: [vue()],
    server: {
      port: 6060,
      proxy: {
        '/dev-api': {
          target: "https://job-log.yumu799.cyou:8443/",
          // target: "http://127.0.0.1:3000/",
          changeOrigin: true,
          ws: true,
          rewrite: (path) => path.replace(/^\/dev-api/, '')
        },
      }
    },
  })
})
