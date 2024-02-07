#!/bin/bash

# 杀死具体的 Go 进程，根据程序名称或其他标识符来匹配
stop_program() {
    program_name="$1"
    process_id=$(pgrep -o -f "$program_name")

    if [ -n "$process_id" ]; then
        echo "Stopping $program_name (PID: $process_id)..."
        kill "$process_id"
    else
        echo "$program_name is not running."
    fi
}

# 停止 main.go
stop_program "main.go"

# 停止 starter.go
stop_program "starter"

# 停止 frontend.go
stop_program "frontend"