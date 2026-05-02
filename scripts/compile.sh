#!/bin/bash
# 同时编译 Go 和前端代码

set -e  # 遇到错误立即退出

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

echo "========================================="
echo "开始清理旧的编译内容..."
echo "========================================="

# 删除旧的 Go 编译产物
if [ -f "bin/web_server" ]; then
    echo "删除 bin/web_server"
    rm -f bin/web_server
fi

# 删除旧的前端编译产物
if [ -d "frontend/dist" ]; then
    echo "删除 frontend/dist"
    rm -rf frontend/dist
fi

# 如果 bin 目录为空，删除它
if [ -d "bin" ] && [ -z "$(ls -A bin 2>/dev/null)" ]; then
    echo "删除空的 bin 目录"
    rmdir bin
fi

echo "清理完成!"
echo ""

echo "========================================="
echo "开始编译 Go 代码..."
echo "========================================="

# 检查是否需要交叉编译为 Linux
if [ "$BUILD_TARGET" = "linux" ]; then
    echo "编译目标: Linux (GOOS=linux GOARCH=amd64)"
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build -ldflags="-linkmode external -extldflags '-static'" -o bin/web_server web_server.go 2>/dev/null || \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags sqlite_omit_load_extension -o bin/web_server web_server.go
else
    echo "编译目标: 本地平台"
    go build -o bin/web_server web_server.go
fi
echo "Go 编译完成: bin/web_server"
echo ""

echo "========================================="
echo "开始编译前端代码..."
echo "========================================="
cd frontend
npm run build
cd ..
echo "前端编译完成: frontend/dist"
echo ""

echo "========================================="
echo "编译完成!"
echo "========================================="
echo "Go 二进制: bin/web_server"
echo "前端产物: frontend/dist"
