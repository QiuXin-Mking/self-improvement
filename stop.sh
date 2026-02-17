#!/bin/bash

# 停止 spaced repetition 服务（前端 + 后端）

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BACKEND_PID_FILE="$SCRIPT_DIR/.web_server.pid"
FRONTEND_PID_FILE="$SCRIPT_DIR/.frontend.pid"
LOGS_DIR="$SCRIPT_DIR/logs"
BACKEND_LOG_LINK="$LOGS_DIR/backend.log"
FRONTEND_LOG_LINK="$LOGS_DIR/frontend.log"

echo "=================================="
echo "正在停止服务..."
echo "=================================="
echo ""

# 记录停止时间到日志
if [ -f "$BACKEND_LOG_LINK" ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 正在停止服务..." >> "$BACKEND_LOG_LINK"
fi
if [ -f "$FRONTEND_LOG_LINK" ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 正在停止服务..." >> "$FRONTEND_LOG_LINK"
fi

# 停止后端
echo "--- 后端服务 ---"
STOPPED_BACKEND=false
if [ -f "$BACKEND_PID_FILE" ]; then
    PID=$(cat "$BACKEND_PID_FILE")
    if ps -p "$PID" > /dev/null 2>&1; then
        kill "$PID" 2>/dev/null
        echo "已通过 PID 文件停止进程 (PID: $PID)"
        STOPPED_BACKEND=true
    fi
    rm "$BACKEND_PID_FILE"
fi

# 查找并停止所有 web_server 相关进程
PIDS=$(pgrep -f "web_server" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep")
if [ -n "$PIDS" ]; then
    for pid in $PIDS; do
        echo "停止后端进程: $pid"
        kill "$pid" 2>/dev/null
        STOPPED_BACKEND=true
    done
fi

# 强制停止残留的后端进程
sleep 1
PIDS=$(pgrep -f "web_server" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep")
if [ -n "$PIDS" ]; then
    for pid in $PIDS; do
        kill -9 "$pid" 2>/dev/null
        echo "强制停止后端进程: $pid"
        STOPPED_BACKEND=true
    done
fi

if [ "$STOPPED_BACKEND" = true ] && [ -f "$BACKEND_LOG_LINK" ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 后端服务已停止" >> "$BACKEND_LOG_LINK"
fi

# 停止前端
echo ""
echo "--- 前端服务 ---"
STOPPED_FRONTEND=false
if [ -f "$FRONTEND_PID_FILE" ]; then
    PID=$(cat "$FRONTEND_PID_FILE")
    if ps -p "$PID" > /dev/null 2>&1; then
        kill "$PID" 2>/dev/null
        echo "已通过 PID 文件停止进程 (PID: $PID)"
        STOPPED_FRONTEND=true
    fi
    rm "$FRONTEND_PID_FILE"
fi

# 查找并停止所有 vite 相关进程
PIDS=$(pgrep -f "vite" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep")
if [ -n "$PIDS" ]; then
    for pid in $PIDS; do
        echo "停止前端进程: $pid"
        kill "$pid" 2>/dev/null
        STOPPED_FRONTEND=true
    done
fi

# 强制停止残留的前端进程
sleep 1
PIDS=$(pgrep -f "vite" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep")
if [ -n "$PIDS" ]; then
    for pid in $PIDS; do
        kill -9 "$pid" 2>/dev/null
        echo "强制停止前端进程: $pid"
        STOPPED_FRONTEND=true
    done
fi

# 清理 esbuild 进程（Vite 的构建工具）
PIDS=$(pgrep -f "esbuild" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep")
if [ -n "$PIDS" ]; then
    for pid in $PIDS; do
        kill "$pid" 2>/dev/null
    done
fi

if [ "$STOPPED_FRONTEND" = true ] && [ -f "$FRONTEND_LOG_LINK" ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 前端服务已停止" >> "$FRONTEND_LOG_LINK"
fi

echo ""
echo "=================================="
echo "完成！所有服务已停止"
echo "=================================="
echo ""

# 显示日志信息
if [ -d "$LOGS_DIR" ]; then
    LOG_COUNT=$(ls -1 "$LOGS_DIR"/*.log 2>/dev/null | wc -l)
    echo "日志文件: $LOG_COUNT 个"
    echo "日志目录: $LOGS_DIR"
    echo ""
    echo "查看最新日志:"
    echo "  后端: tail -20 $BACKEND_LOG_LINK"
    echo "  前端: tail -20 $FRONTEND_LOG_LINK"
    echo ""
    echo "清理旧日志:"
    echo "  find $LOGS_DIR -name '*.log' -mtime +7 -delete"
fi
