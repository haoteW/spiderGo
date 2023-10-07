

document.addEventListener("DOMContentLoaded", function () {
    const content = document.getElementById("content");
    const loadMoreButton = document.getElementById("loadMore");
    let itemCounter = 2; // 从第三个 item 开始加载

    // 异步加载更多内容的函数
    function loadMoreItems() {
        // 模拟异步请求，例如使用 Fetch API 请求数据
        setTimeout(function () {
            const newItem = document.createElement("div");
            newItem.classList.add("item");
            newItem.textContent = "Item " + itemCounter;
            content.appendChild(newItem);
            itemCounter++;
        }, 1000); // 模拟延迟 1 秒

        loadMoreButton.disabled = true; // 防止多次点击
        loadMoreButton.textContent = "Loading...";

        loadMoreItems(); // 初始加载

        loadMoreButton.addEventListener("click", function () {
            loadMoreItems();
            // 恢复按钮状态
            loadMoreButton.disabled = false;
            loadMoreButton.textContent = "Load More";
        });
    }
});
