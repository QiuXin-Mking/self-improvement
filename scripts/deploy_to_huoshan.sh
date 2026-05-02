#!/bin/bash
# 部署脚本 - 在服务器上编译并部署，最终只保留二进制文件

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

# 配置
SERVER="huoshan"
SERVER_PATH="/root/spaced-repetition"
TEMP_BUILD_DIR="/tmp/spaced-repetition-build"

log_info "========================================="
log_info "部署到 $SERVER 服务器"
log_info "========================================="
echo ""

# 本地编译前端
log_info "========================================="
log_info "本地编译前端..."
log_info "========================================="
cd frontend && npm run build && cd ..

# 准备上传包（不含源码）
log_info "========================================="
log_info "准备上传包..."
log_info "========================================="

DEPLOY_DIR="deploy_temp"
rm -rf "$DEPLOY_DIR"
mkdir -p "$DEPLOY_DIR"

# 只上传必要的文件
mkdir -p "$DEPLOY_DIR/internal"

# 仅上传 web_server.go（cli_server.go 同样是 package main + main()，一起编译会冲突）
cp web_server.go "$DEPLOY_DIR/"
cp go.mod go.sum "$DEPLOY_DIR/"
cp -r internal/* "$DEPLOY_DIR/internal/"

# 上传前端编译产物
mkdir -p "$DEPLOY_DIR/frontend/dist"
cp -r frontend/dist/* "$DEPLOY_DIR/frontend/dist/"

# 创建服务器端编译和部署脚本
cat > "$DEPLOY_DIR/build_and_deploy.sh" << 'BEOF'
#!/bin/bash
set -e

BUILD_DIR="$(pwd)"
DEPLOY_PATH="/root/spaced-repetition"

echo "========================================="
echo "服务器端编译和部署"
echo "========================================="

# 安装 Go（如果未安装）
if ! command -v go &> /dev/null; then
    echo "安装 Go..."
    apt-get update -qq && apt-get install -y -qq golang-1.22-go 2>/dev/null
    rm -rf /usr/local/go
    ln -sf /usr/lib/go-1.22 /usr/local/go
fi
export PATH=/usr/local/go/bin:$PATH
GO=/usr/local/go/bin/go

# 停止现有服务
if [ -f "$DEPLOY_PATH/stop.sh" ]; then
    bash "$DEPLOY_PATH/stop.sh" 2>/dev/null || true
fi

# 保留用户数据（数据库 + .env，部署时不清除）
if [ -d "$DEPLOY_PATH/data" ]; then
    echo "保留用户数据..."
    cp -r "$DEPLOY_PATH/data" /tmp/spaced-repetition-data-backup
fi
if [ -f "$DEPLOY_PATH/.env" ]; then
    cp "$DEPLOY_PATH/.env" /tmp/spaced-repetition-env-backup
fi

# 备份当前版本
if [ -d "$DEPLOY_PATH" ] && [ "$(ls -A $DEPLOY_PATH 2>/dev/null)" ]; then
    echo "备份当前版本..."
    cp -r "$DEPLOY_PATH" "$DEPLOY_PATH.backup.$(date +%Y%m%d_%H%M%S)" 2>/dev/null || true
fi

# 编译 Go 程序（当前目录已包含源码，直接编译）
echo "编译 Go 程序..."
cd "$BUILD_DIR"
export GOPROXY=https://goproxy.cn,direct
export PATH=/usr/local/go/bin:$PATH
$GO mod download 2>/dev/null || true
$GO build -o web_server . || {
    echo "编译失败，查看错误..."
    exit 1
}

# 创建最终部署目录
rm -rf "$DEPLOY_PATH"
mkdir -p "$DEPLOY_PATH/bin"
mkdir -p "$DEPLOY_PATH/data"
mkdir -p "$DEPLOY_PATH/questions"

# 恢复用户数据
if [ -d /tmp/spaced-repetition-data-backup ]; then
    cp -r /tmp/spaced-repetition-data-backup/* "$DEPLOY_PATH/data/"
    rm -rf /tmp/spaced-repetition-data-backup
fi
if [ -f /tmp/spaced-repetition-env-backup ]; then
    cp /tmp/spaced-repetition-env-backup "$DEPLOY_PATH/.env"
    rm -f /tmp/spaced-repetition-env-backup
fi

# 复制编译产物（不含源码）
cp web_server "$DEPLOY_PATH/bin/web_server"
cp -r frontend/dist "$DEPLOY_PATH/"

# 创建启动脚本
cat > "$DEPLOY_PATH/start.sh" << 'EOF'
#!/bin/bash
APP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PID_FILE="$APP_DIR/.app.pid"
LOG_FILE="$APP_DIR/app.log"

# 加载环境变量
if [ -f "$APP_DIR/.env" ]; then
    set -a
    source "$APP_DIR/.env"
    set +a
fi

if [ -f "$PID_FILE" ]; then
    OLD_PID=$(cat "$PID_FILE")
    if ps -p "$OLD_PID" > /dev/null 2>&1; then
        kill "$OLD_PID" 2>/dev/null || true
        sleep 2
        kill -9 "$OLD_PID" 2>/dev/null || true
    fi
    rm -f "$PID_FILE"
fi

cd "$APP_DIR"
nohup ./bin/web_server > "$LOG_FILE" 2>&1 &
echo $! > "$PID_FILE"
sleep 2

if ps -p $(cat "$PID_FILE") > /dev/null; then
    echo "服务启动成功 (PID: $(cat $PID_FILE))"
else
    echo "服务启动失败，请查看日志: $LOG_FILE"
    exit 1
fi
EOF
chmod +x "$DEPLOY_PATH/start.sh"

# 创建停止脚本
cat > "$DEPLOY_PATH/stop.sh" << 'EOF'
#!/bin/bash
APP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PID_FILE="$APP_DIR/.app.pid"

if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    if ps -p "$PID" > /dev/null 2>&1; then
        kill "$PID" 2>/dev/null || true
        sleep 2
        kill -9 "$PID" 2>/dev/null || true
    fi
    rm -f "$PID_FILE"
    echo "服务已停止"
fi
EOF
chmod +x "$DEPLOY_PATH/stop.sh"

# 创建 .env.example
cat > "$DEPLOY_PATH/.env.example" << 'EOF'
PORT=4430
DATABASE_PATH=data/app.db
JWT_SECRET=change-this-to-a-random-secret-string-for-production
LOG_LEVEL=info
EOF

# 创建 .env（如果不存在）
if [ ! -f "$DEPLOY_PATH/.env" ]; then
    cp "$DEPLOY_PATH/.env.example" "$DEPLOY_PATH/.env"
fi

# 设置权限
chmod +x "$DEPLOY_PATH/bin/web_server"

# 启动服务
echo "启动服务..."
bash "$DEPLOY_PATH/start.sh"

echo "========================================="
echo "部署完成！"
echo "========================================="
BEOF
chmod +x "$DEPLOY_DIR/build_and_deploy.sh"

# 打包
tar -czf deploy-package.tar.gz -C "$DEPLOY_DIR" .

# 上传并执行
log_info "========================================="
log_info "上传并部署..."
log_info "========================================="

ssh "$SERVER" "mkdir -p $TEMP_BUILD_DIR"
scp deploy-package.tar.gz "$SERVER:$TEMP_BUILD_DIR/"

ssh "$SERVER" "
cd $TEMP_BUILD_DIR && \
tar -xzf deploy-package.tar.gz && \
bash build_and_deploy.sh && \
rm -f deploy-package.tar.gz
"

# 清理本地
rm -rf "$DEPLOY_DIR" deploy-package.tar.gz

log_info "========================================="
log_info "部署完成！"
log_info "========================================="
echo ""
echo "查看服务:"
echo "  ssh $SERVER 'cd $SERVER_PATH && cat .app.pid'"
echo "  ssh $SERVER 'cd $SERVER_PATH && tail -f app.log'"
echo ""
