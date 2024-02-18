#!/bin/bash

# 杀死具体的 Go 进程，根据程序名称或其他标识符来匹配
stop_program() {
    program_name="$1"
    process_ids=$(pgrep -f "program_name")
    # $(pgrep -o -f "$program_name")

    if [ -n "$process_ids" ]; then
        for pid in $process_ids; do
            # 使用 kill 命令终止该进程
            kill "$pid"
            echo "Process with PID $pid terminated."
        done
    else
        echo "$program_name is not running."
    fi
}

stop_cnpm(){
    # 使用 ps 命令查找正在运行的 cnpm run dev 进程的 PID
    pids=$(ps aux | grep "cnpm run dev" | grep -v grep | awk '{print $2}')

    # 检查是否找到了 PID
    if [ -z "$pids" ]; then
        echo "No 'cnpm run dev' process found."
    else
        for pid in $pids; do
            # 使用 kill 命令终止该进程
            kill "$pid"
            echo "Process with PID $pid terminated."
        done
    fi

    pid=$(lsof -ti tcp:8080)

    # 检查是否找到了进程
    if [ -n "$pid" ]; then
        # 终止进程
        echo "Killing process with PID: $pid"
        kill $pid
    else
        echo "No process found using port 8080."
    fi
}

# 停止 main.go
stop_program "main.go"

# 停止 starter.go
stop_program "starter"

stop_program "cnpm dev run"

# 停止 frontend.go
stop_cnpm