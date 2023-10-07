package main

import (
	"encoding/json"
	"good-spider/persist"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// 添加处理程序到 mux
	mux.HandleFunc("/getMovieData", func(w http.ResponseWriter, r *http.Request) {
		// 添加 CORS 头部以允许跨域请求
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) //允许访问所有域
		// 必须，设置服务器支持的所有跨域请求的方法
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
		w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			// 如果是预检请求（OPTIONS 请求），则只返回 CORS 头部，不处理实际请求
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 获取查询参数
		queryParams := r.URL.Query()

		// 获取特定参数的值
		limit := queryParams.Get("limit")
		orderType := queryParams.Get("orderType")

		movieData := persist.Search(limit, orderType)

		if len(movieData) == 0 {
			http.Error(w, "there is no more data", http.StatusNoContent)
			return
		}
		// 将数据编码为 JSON 并发送回客户端
		jsonData, err := json.Marshal(movieData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}) // 示例：添加处理程序到 /api/data 路径

	// 启动 HTTP 服务器并使用 mux 作为多路复用器
	http.ListenAndServe(":8080", mux)
}
