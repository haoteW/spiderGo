import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:{
    proxy:{
      '/getMovieData':{//获取路径中包含了/api的请求
          target:'http://127.0.0.1:6301',//后台服务所在的源
          changeOrigin:true,//修改源
          rewrite:(path)=>path.replace(/^\/getMovieData/,'getMovieData'),///api替换为''
          secure: false, // 接受运行在https上，默认不接受  
        }
    },
    host:'0.0.0.0',
    port:8080,
    https:false
  },
})
