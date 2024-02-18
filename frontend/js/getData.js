// 指定后端的 IP 地址和端口号
const backendIpAddress = "60.205.179.8";  // 替换为实际的后端 IP 地址
const backendPort = "6301";  // 替换为实际的后端端口号
// 准备一次加载的项目数
const itemsPerLoad = 10;

let pn = 0;
let rn = itemsPerLoad;

function fetchData() {
    if (pn >= 300){
        return
    }

    // 使用 Fetch API 进行 AJAX 请求
    fetch(`http://${backendIpAddress}:${backendPort}/getMovieData?limit=${pn},${rn}&orderType=asc`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json", // 设置请求头
        },
        credentials: "include", // 允许发送凭据（例如，跨域请求时发送 cookies）
    }).then(response => {
        // console.error("response.status : " + response.status); // 添加调试信息
        if (response.status === 204) {
            // 处理空响应：隐藏按钮并返回一个空数组
            // loadMoreButton.style.display = "none"; // 隐藏按钮
            return [];
        } else if (response.ok) {
            return response.json(); // 将非空响应解析为 JSON
        } else {
            // 处理响应不成功的情况
            // loadMoreButton.style.display = "none"; // 隐藏按钮
            throw new Error(`Request failed with status: ${response.status}`);
        }
    }).then(movieData => {
        // 检查响应数据是否为空
        if (movieData.length === 0) {
            // loadMoreButton.style.display = "none"; // 隐藏按钮
        } else {
            pn += itemsPerLoad;
            return movieData;
            
        }
    }).catch(error => {
        console.error("Error loading data: " + error);
    });


}
    
    
            // 监听滚动事件以判断是否滚动到底部
