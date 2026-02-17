#!/bin/bash

# 多租户艾宾浩斯遗忘曲线学习系统部署脚本 (国际版/海外优化版)
# 适用于香港等可以自由连接外部服务商的服务器环境
# 适用于 Ubuntu 20.04+ 或 Debian 10+

set -e  # 遇到错误时停止执行

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

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

# 检查是否为 root 用户
if [[ $EUID -eq 0 ]]; then
   log_error "此脚本不应以 root 用户身份运行。请使用普通用户运行，并确保该用户具有 sudo 权限。"
   exit 1
fi

# 检查是否安装了 sudo
if ! command -v sudo &> /dev/null; then
    log_error "sudo 命令未安装。请先安装 sudo。"
    exit 1
fi

log_info "开始部署多租户艾宾浩斯遗忘曲线学习系统(国际版)..."

# 定义变量
APP_DIR="/opt/spaced-repetition"
DATA_DIR="$APP_DIR/data"
QUESTIONS_DIR="$APP_DIR/questions"
USER_NAME="$(logname)"
BACKUP_DIR="/opt/backups"
DOMAIN_NAME=""

# 检查操作系统
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
else
    log_error "无法确定操作系统类型"
    exit 1
fi

log_info "检测到操作系统: $OS"

# 恢复原始镜像源配置（如果之前使用了中国镜像）
restore_original_sources() {
    log_info "检查是否需要恢复原始系统镜像源..."

    # 恢复 Ubuntu/Debian 原始源
    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        if [ -f /etc/apt/sources.list.backup ]; then
            log_info "发现中国镜像源备份，恢复原始源列表..."
            sudo cp /etc/apt/sources.list.backup /etc/apt/sources.list
            sudo apt update
            log_info "已恢复原始系统镜像源"
        else
            log_info "未找到中国镜像源备份，保持当前源设置"
        fi
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        if [ -f /etc/yum.repos.d/CentOS-Base.repo.backup ]; then
            log_info "发现中国镜像源备份，恢复原始源配置..."
            sudo cp /etc/yum.repos.d/CentOS-Base.repo.backup /etc/yum.repos.d/CentOS-Base.repo
            sudo yum makecache
            log_info "已恢复原始系统镜像源"
        else
            log_info "未找到中国镜像源备份，保持当前源设置"
        fi
    fi
}

# 配置全球/海外环境设置
setup_global_config() {
    log_info "配置全球/海外环境设置..."

    # 设置 Go 代理为官方默认（或不设置，直接访问）
    unset GOPROXY
    # 从 .bashrc 中删除中国 Go 代理设置
    sed -i '/goproxy.cn/d' ~/.bashrc
    log_info "已配置 Go 为默认代理模式"
}

# 检查并安装必要软件
install_dependencies() {
    log_info "安装依赖软件..."

    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        sudo apt update
        sudo apt install -y curl wget git build-essential sqlite3
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        sudo yum install -y epel-release
        sudo yum install -y curl wget git gcc make sqlite-devel
    else
        log_error "不支持的操作系统: $OS"
        exit 1
    fi
}

# 安装 Go
install_go() {
    log_info "检查 Go 是否已安装..."

    if command -v go &> /dev/null; then
        GO_VERSION=$(go version | cut -d ' ' -f 3 | sed 's/go//')
        MIN_VERSION="1.19"

        if [ "$(printf '%s\n' "$MIN_VERSION" "$GO_VERSION" | sort -V | head -n1)" = "$MIN_VERSION" ] || [ "$MIN_VERSION" = "$(printf '%s\n' "$MIN_VERSION" "$GO_VERSION" | sort -V | head -n1)" ]; then
            log_info "Go $GO_VERSION 已安装且满足要求"
        else
            log_info "Go 版本过低，需要安装 Go 1.19+"
            # 继续安装 Go
        fi
    else
        log_info "Go 未安装，正在安装..."
    fi

    # 安装或更新 Go 从官方源
    GO_VERSION="1.21.5"
    log_info "从官方源下载 Go $GO_VERSION..."

    wget -O go.tar.gz "https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz"

    if [ $? -ne 0 ]; then
        log_warn "从官方源下载失败，尝试使用备用源..."
        # 尝试使用 Go 中文网镜像作为最后备选
        wget -O go.tar.gz "https://studygolang.com/dl/golang/go$GO_VERSION.linux-amd64.tar.gz"
    fi

    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
}

# 安装 Node.js (仅当需要重新构建前端时才需要)
install_nodejs() {
    log_info "检查 Node.js 是否已安装..."

    if command -v node &> /dev/null; then
        NODE_VERSION=$(node -v | sed 's/v//')
        MIN_VERSION="16.0.0"

        if [ "$(printf '%s\n' "$MIN_VERSION" "$NODE_VERSION" | sort -V | head -n1)" = "$MIN_VERSION" ] || [ "$MIN_VERSION" = "$(printf '%s\n' "$MIN_VERSION" "$NODE_VERSION" | sort -V | head -n1)" ]; then
            log_info "Node.js $NODE_VERSION 已安装且满足要求"
        else
            log_info "Node.js 版本过低，正在更新..."
        fi
    else
        log_info "Node.js 未安装，正在安装..."
    fi

    # 从官方源安装 Node.js
    log_info "从官方源安装 Node.js..."

    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        # 使用 Nodesource 官方安装脚本
        curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
        sudo apt-get install -y nodejs
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        # 对于 CentOS 使用官方安装方式
        curl -fsSL https://rpm.nodesource.com/setup_18.x | sudo bash -
        sudo yum install -y nodejs
    fi
}

# 创建应用目录
create_directories() {
    log_info "创建应用目录..."

    sudo mkdir -p "$APP_DIR"
    sudo mkdir -p "$DATA_DIR"
    sudo mkdir -p "$QUESTIONS_DIR"
    sudo mkdir -p "$BACKUP_DIR"
    sudo chown -R "$USER_NAME:$USER_NAME" "$APP_DIR"
}

# 检查现有代码并更新/继续部署
check_and_update_code() {
    log_info "检查现有代码并准备更新部署..."

    if [ -d "$APP_DIR" ]; then
        log_info "检测到已存在的部署目录，检查是否为有效项目..."

        if [[ ! -f "$APP_DIR/web_server.go" || ! -f "$APP_DIR/main.go" ]]; then
            log_error "关键文件不存在。请确保已将项目代码上传到 $APP_DIR 目录。"
            exit 1
        fi

        log_info "确认项目代码完整性，准备继续部署..."

        # 检查是否有 Git 存储库，如果有则拉取最新代码
        if [ -d "$APP_DIR/.git" ]; then
            log_info "检测到 Git 存储库，尝试拉取最新代码..."
            cd "$APP_DIR"
            git fetch
            git pull origin main
            log_info "代码已更新到最新版本"
        fi
    else
        log_error "未找到现有部署目录。此脚本旨在更新已部署的环境配置。请先使用 deploy.sh 部署或手动上传项目代码。"
        exit 1
    fi
}

# 设置环境变量（国际优化版）
setup_environment() {
    log_info "设置环境变量..."

    cat > "$APP_DIR/.env" << 'EOF'
# 服务器配置
PORT=5000
DATABASE_PATH=/opt/spaced-repetition/data/app.db

# 安全配置 - 请务必更换为随机字符串
JWT_SECRET=CHANGE_THIS_TO_A_RANDOM_SECRET_STRING_FOR_PRODUCTION

# 日志级别
LOG_LEVEL=info

# 国际化相关配置
LANGUAGE=en
TIMEZONE=Asia/Hong_Kong
EOF

    chmod 600 "$APP_DIR/.env"
    sudo chown "$USER_NAME:$USER_NAME" "$APP_DIR/.env"
}

# 安装 Go 依赖并重新构建
build_application() {
    log_info "安装 Go 依赖并重新构建应用..."

    cd "$APP_DIR"

    # 确保使用默认 Go 代理（能够访问 golang.org 的环境）
    unset GOPROXY
    go env -w GOPROXY=direct

    # 安装依赖
    go mod tidy

    # 构建应用
    log_info "重新构建 Web 服务器应用..."
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bin/web_app web_server.go

    log_info "重新构建 CLI 应用..."
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bin/train main.go

    log_info "应用重建完成。"
}

# 重建前端 (如果需要)
rebuild_frontend() {
    if [ -d "$APP_DIR/frontend" ]; then
        log_info "重建前端应用..."

        cd "$APP_DIR/frontend"

        # 使用官方 npm 源
        npm config set registry https://registry.npmjs.org/
        npm config delete disturl

        # 安装前端依赖
        npm install

        # 重建前端
        npm run build

        cd "$APP_DIR"
        log_info "前端重建完成。"
    else
        log_warn "前端目录不存在，跳过前端重建。"
    fi
}

# 重新初始化数据库（如果需要）
reinitialize_database() {
    log_info "检查数据库状态..."

    # 确保数据目录存在
    mkdir -p "$DATA_DIR"

    if [ -f "$DATA_DIR/app.db" ]; then
        log_info "数据库文件已存在，跳过初始化。如需重新初始化，请手动删除数据库文件。"
    else
        log_info "数据库文件不存在，执行初始化..."
        if [ -f "$APP_DIR/migrations/001_initial_schema.sql" ]; then
            sqlite3 "$DATA_DIR/app.db" < "$APP_DIR/migrations/001_initial_schema.sql"
            log_info "数据库初始化完成。"
        else
            log_warn "未找到数据库初始化脚本，跳过数据库初始化。"
        fi
    fi
}

# 重启服务（由于环境配置改变）
restart_service() {
    log_info "重新配置和重启服务..."

    # 检查服务是否已存在
    if sudo systemctl is-active --quiet spaced-repetition; then
        log_info "停止现有服务..."
        sudo systemctl stop spaced-repetition
    fi

    # 重新加载 systemd 配置
    sudo systemctl daemon-reload

    # 重新启用并启动服务
    sudo systemctl enable spaced-repetition
    sudo systemctl start spaced-repetition

    # 等待服务启动
    sleep 5

    if sudo systemctl is-active --quiet spaced-repetition; then
        log_info "服务重新启动成功！"
    else
        log_error "服务启动失败。请检查日志。"
        sudo journalctl -u spaced-repetition --no-pager -l
        exit 1
    fi
}

# 配置 Nginx (可选)
configure_nginx() {
    log_info "询问是否需要重新配置 Nginx..."

    read -p "是否需要重新配置 Nginx 反向代理？(y/n): " -n 1 -r
    echo

    if [[ $REPLY =~ ^[Yy]$ ]]; then
        # 检查 Nginx 是否已安装
        if ! command -v nginx &> /dev/null; then
            log_info "安装 Nginx..."

            if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
                sudo apt install -y nginx
            elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
                sudo yum install -y nginx
            fi
        fi

        # 询问域名
        read -p "请输入您的域名 (例如 example.com): " DOMAIN_NAME

        # 创建 Nginx 配置
        sudo tee /etc/nginx/sites-available/spaced-repetition << EOF
server {
    listen 80;
    server_name $DOMAIN_NAME;

    location / {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;

        # 增加超时时间，避免大文件上传超时
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 安全头配置
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Content-Type-Options "nosniff" always;
}
EOF

        sudo rm -f /etc/nginx/sites-enabled/spaced-repetition
        sudo ln -s /etc/nginx/sites-available/spaced-repetition /etc/nginx/sites-enabled/

        # 测试 Nginx 配置
        sudo nginx -t

        # 重新加载 Nginx
        sudo systemctl reload nginx

        log_info "Nginx 重新配置完成。"

        # 询问是否需要 SSL 证书
        read -p "是否需要配置 SSL 证书？(y/n): " -n 1 -r
        echo

        if [[ $REPLY =~ ^[Yy]$ ]]; then
            setup_ssl_cert
        fi
    fi
}

# 配置 SSL 证书
setup_ssl_cert() {
    log_info "配置 SSL 证书..."

    # 检查 Certbot 是否已安装
    if ! command -v certbot &> /dev/null; then
        log_info "安装 Certbot..."

        if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
            sudo apt install -y certbot python3-certbot-nginx
        elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
            sudo yum install -y certbot python3-certbot-nginx
        fi
    fi

    # 获取 SSL 证书
    sudo certbot --nginx -d "$DOMAIN_NAME"

    log_info "SSL 证书配置完成。"
}

# 更新备份策略（如果需要）
update_backup_script() {
    log_info "更新备份脚本配置..."

    sudo tee /opt/backup-script.sh << 'EOF'
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/opt/backups"
APP_DIR="/opt/spaced-repetition"

mkdir -p $BACKUP_DIR

# 备份数据库
cp $APP_DIR/data/app.db $BACKUP_DIR/app_backup_$DATE.db

# 清理7天前的备份
find $BACKUP_DIR -name "app_backup_*.db" -mtime +7 -delete

echo "Backup completed at $DATE"
EOF

    sudo chmod +x /opt/backup-script.sh

    # 检查是否已有备份任务，如果没有则添加
    if ! crontab -l 2>/dev/null | grep -q "backup-script"; then
        # 添加定时备份 (每天凌晨2点)
        echo "0 2 * * * /opt/backup-script.sh" | crontab -
        log_info "新的备份任务已添加到 crontab。"
    else
        log_info "备份任务已在 crontab 中存在。"
    fi

    log_info "备份脚本更新完成。"
}

# 显示完成信息
show_completion_message() {
    log_info "环境更新和重新部署完成！"
    echo
    echo "以下是重要信息："
    echo "- 应用将在 http://localhost:5000 或 http://your-server-ip:5000 运行"
    if [ -n "$DOMAIN_NAME" ]; then
        echo "- 通过域名访问: https://$DOMAIN_NAME (如果配置了 SSL)"
    fi
    echo "- 服务状态: sudo systemctl status spaced-repetition"
    echo "- 查看日志: sudo journalctl -u spaced-repetition -f"
    echo "- 配置文件: $APP_DIR/.env (请务必更新 JWT_SECRET)"
    echo
    echo "已执行的环境更改："
    echo "1. 恢复了系统镜像源为默认设置"
    echo "2. 移除了中国地区的 Go 代理设置"
    echo "3. 更新了应用配置以适应国际网络环境"
    echo "4. 重建了应用二进制文件"
    echo "5. 重新启动了服务"
    echo
    log_info "请务必更新 $APP_DIR/.env 中的 JWT_SECRET 为安全的随机字符串。"
}

# 主函数
main() {
    log_info "开始更新部署环境配置(国际版)..."

    # 恢复原始镜像源
    restore_original_sources

    # 配置全球环境
    setup_global_config

    # 检查并更新现有代码
    check_and_update_code

    # 重新设置环境变量
    setup_environment

    # 重新构建应用程序
    build_application

    # 重建前端（如果存在）
    rebuild_frontend

    # 重新初始化数据库（如果需要）
    reinitialize_database

    # 更新备份脚本
    update_backup_script

    # 重启服务以应用新配置
    restart_service

    # 重新配置 Nginx（可选）
    configure_nginx

    # 显示完成信息
    show_completion_message

    log_info "环境更新过程完成！"
}

# 执行主函数
main "$@"