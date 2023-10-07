package main

import (
	"net/http"
)

func main() {
	// 配置路由
	http.HandleFunc("/", ServeHTML)

	// 启动 HTTP 服务器并指定监听端口
	http.ListenAndServe(":6300", nil)
}

// ServeHTML 处理请求并提供 HTML 文件
func ServeHTML(w http.ResponseWriter, r *http.Request) {
	// 打开并读取 HTML 文件
	http.ServeFile(w, r, "frontend/template.html") // 替换为实际的 HTML 文件路径
}
