#!/bin/bash

# 启动 spaced repetition 服务（前端 + 后端）

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
FRONTEND_DIR="$SCRIPT_DIR/frontend"
BACKEND_PID_FILE="$SCRIPT_DIR/.web_server.pid"
FRONTEND_PID_FILE="$SCRIPT_DIR/.frontend.pid"
LOGS_DIR="$SCRIPT_DIR/logs"

# 创建日志目录
mkdir -p "$LOGS_DIR"

# 日志文件命名（带时间戳）
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKEND_LOG_FILE="$LOGS_DIR/backend_${TIMESTAMP}.log"
FRONTEND_LOG_FILE="$LOGS_DIR/frontend_${TIMESTAMP}.log"
BACKEND_LOG_LINK="$LOGS_DIR/backend.log"
FRONTEND_LOG_LINK="$LOGS_DIR/frontend.log"

# 默认配置（前端端口根据 vite.config.ts 设置为 3000）
BACKEND_PORT="${BACKEND_PORT:-8000}"
FRONTEND_PORT="${FRONTEND_PORT:-3000}"
JWT_SECRET="${JWT_SECRET:-my-secret-key}"

# 检查服务是否已运行
if pgrep -f "web_server" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep" > /dev/null; then
    echo "后端服务已在运行中"
    echo "如需重启，请先运行 ./stop.sh"
    exit 1
fi

if pgrep -f "vite" | grep -v "\.sh$" | grep -v "stop.sh" | grep -v "start.sh" | grep -v "grep" > /dev/null; then
    echo "前端服务已在运行中"
    echo "如需重启，请先运行 ./stop.sh"
    exit 1
fi

echo "=================================="
echo "正在启动服务..."
echo "=================================="
echo ""
echo "--- 后端配置 ---"
echo "  端口: $BACKEND_PORT"
echo "  日志: $BACKEND_LOG_FILE"
echo ""
echo "--- 前端配置 ---"
echo "  端口: $FRONTEND_PORT"
echo "  日志: $FRONTEND_LOG_FILE"
echo ""

# 启动后端服务
cd "$SCRIPT_DIR" || exit 1
echo "[$(date '+%Y-%m-%d %H:%M:%S')] 启动后端服务..." >> "$BACKEND_LOG_FILE"
PORT="$BACKEND_PORT" JWT_SECRET="$JWT_SECRET" nohup go run web_server.go >> "$BACKEND_LOG_FILE" 2>&1 &
BACKEND_PID=$!
echo "$BACKEND_PID" > "$BACKEND_PID_FILE"

# 创建日志软链接（方便访问最新日志）
ln -sf "$(basename "$BACKEND_LOG_FILE")" "$BACKEND_LOG_LINK"

# 启动前端服务
cd "$FRONTEND_DIR" || exit 1
echo "[$(date '+%Y-%m-%d %H:%M:%S')] 启动前端服务..." >> "$FRONTEND_LOG_FILE"
nohup npm run dev >> "$FRONTEND_LOG_FILE" 2>&1 &
FRONTEND_PID=$!
echo "$FRONTEND_PID" > "$FRONTEND_PID_FILE"

# 创建日志软链接（方便访问最新日志）
ln -sf "$(basename "$FRONTEND_LOG_FILE")" "$FRONTEND_LOG_LINK"

# 等待服务启动
sleep 3

# 检查后端服务
cd "$SCRIPT_DIR" || exit 1
echo ""
echo "--- 检查后端服务 ---"
if ps -p "$BACKEND_PID" > /dev/null; then
    echo "✓ 后端启动成功 (PID: $BACKEND_PID)"
    echo "  访问: http://localhost:$BACKEND_PORT"
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 后端服务启动成功 (PID: $BACKEND_PID)" >> "$BACKEND_LOG_FILE"
else
    echo "✗ 后端启动失败，请查看日志: $BACKEND_LOG_FILE"
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 后端服务启动失败" >> "$BACKEND_LOG_FILE"
    rm "$BACKEND_PID_FILE" 2>/dev/null
    exit 1
fi

# 检查前端服务
echo ""
echo "--- 检查前端服务 ---"
if ps -p "$FRONTEND_PID" > /dev/null; then
    echo "✓ 前端启动成功 (PID: $FRONTEND_PID)"
    echo "  访问: http://localhost:$FRONTEND_PORT"
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 前端服务启动成功 (PID: $FRONTEND_PID)" >> "$FRONTEND_LOG_FILE"
else
    echo "✗ 前端启动失败，请查看日志: $FRONTEND_LOG_FILE"
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] 前端服务启动失败" >> "$FRONTEND_LOG_FILE"
    rm "$FRONTEND_PID_FILE" 2>/dev/null
    exit 1
fi

echo ""
echo "=================================="
echo "所有服务启动成功！"
echo "=================================="
echo ""
echo "查看实时日志:"
echo "  后端: tail -f $BACKEND_LOG_LINK"
echo "  前端: tail -f $FRONTEND_LOG_LINK"
echo ""
echo "查看历史日志:"
echo "  ls -lh $LOGS_DIR/"
echo ""
echo "停止服务: ./stop.sh"
