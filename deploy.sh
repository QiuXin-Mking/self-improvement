#!/bin/bash

# 多租户艾宾浩斯遗忘曲线学习系统部署脚本 (中国地区优化版)
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

log_info "开始部署多租户艾宾浩斯遗忘曲线学习系统(中国优化版)..."

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

# 设置中国镜像源
setup_china_mirrors() {
    log_info "设置中国镜像源..."

    # 设置 Go 代理
    export GOPROXY=https://goproxy.cn,direct
    echo 'export GOPROXY=https://goproxy.cn,direct' >> ~/.bashrc

    # 检查是哪种 Linux 发行版并设置镜像源
    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        log_info "检测到 Debian/Ubuntu 系统，设置国内镜像源..."

        # 备份原始源列表
        if [ -f /etc/apt/sources.list ] && [ ! -f /etc/apt/sources.list.backup ]; then
            sudo cp /etc/apt/sources.list /etc/apt/sources.list.backup
            log_info "已备份原始源列表到 /etc/apt/sources.list.backup"
        fi

        # 使用阿里云镜像源 (根据系统版本选择)
        if [[ "$VER" =~ ^2[0-9]\.04$ ]]; then
            # Ubuntu 20.04
            sudo tee /etc/apt/sources.list > /dev/null << 'EOF'
deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
EOF
        elif [[ "$VER" =~ ^22\.[0-9][0-9]$ ]]; then
            # Ubuntu 22.04
            sudo tee /etc/apt/sources.list > /dev/null << 'EOF'
deb http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
EOF
        else
            # 其他版本使用较通用的配置
            sudo tee /etc/apt/sources.list > /dev/null << 'EOF'
deb http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs) main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs) main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ $(lsb_release -cs)-backports main restricted universe multiverse
EOF
        fi
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        log_info "检测到 CentOS/RHEL 系统，设置国内镜像源..."

        # 备份原始源配置
        if [ -f /etc/yum.repos.d/CentOS-Base.repo ] && [ ! -f /etc/yum.repos.d/CentOS-Base.repo.backup ]; then
            sudo cp /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
            log_info "已备份原始源配置到 /etc/yum.repos.d/CentOS-Base.repo.backup"
        fi

        # 使用阿里云镜像源 (适用于 CentOS 7)
        if [[ "$VER" == "7" ]]; then
            sudo yum install -y wget
            sudo mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
            sudo wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
        elif [[ "$VER" == "8" ]]; then
            sudo dnf -y install dnf-plugins-core
            sudo dnf config-manager --set-enabled PowerTools
            sudo dnf config-manager --add-repo=http://mirrors.aliyun.com/repo/Centos-vault8-extras.repo
        fi
    fi

    # 更新包列表
    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        sudo apt update
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        sudo yum makecache
    fi
}

# 检查并安装必要软件
install_dependencies() {
    log_info "安装依赖软件..."

    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
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

    # 安装或更新 Go
    GO_VERSION="1.21.5"
    log_info "从中国镜像下载 Go $GO_VERSION..."

    # 从 Go 官方国内镜像下载
    wget -O go.tar.gz "https://studygolang.com/dl/golang/go$GO_VERSION.linux-amd64.tar.gz"

    if [ $? -ne 0 ]; then
        log_warn "从 Go 中文网镜像下载失败，尝试使用官方下载地址..."
        # 如果上面的链接失效，可以使用官方地址配合代理
        wget -O go.tar.gz "https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz"
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

    # 使用国内镜像安装 Node.js
    log_info "使用国内镜像安装 Node.js..."

    if [[ "$OS" == "Ubuntu"* ]] || [[ "$OS" == "Debian"* ]]; then
        # 使用 NodeSource 国内镜像
        curl -fsSL https://registry.npmmirror.com/-/binary/node/latest-v18.x/node-v18.17.0-linux-x64.tar.xz -o nodejs.tar.xz
        sudo mkdir -p /opt/nodejs
        sudo tar -xf nodejs.tar.xz -C /opt/nodejs --strip-components=1
        sudo ln -sf /opt/nodejs/bin/node /usr/local/bin/node
        sudo ln -sf /opt/nodejs/bin/npm /usr/local/bin/npm
        sudo ln -sf /opt/nodejs/bin/npx /usr/local/bin/npx

        # 或者使用 Nodesource 国内镜像
        # curl -fsSL https://npmmirror.com/mirrors/setup_18.x | sudo -E bash -
        # sudo apt-get install -y nodejs
    elif [[ "$OS" == "CentOS"* ]] || [[ "$OS" == "Red Hat"* ]]; then
        # 对于 CentOS 使用国内镜像
        curl -fsSL https://npmmirror.com/mirrors/setup_18.x | sudo bash -
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

# 克隆或上传项目代码
setup_code() {
    log_info "设置项目代码..."

    cd "$APP_DIR"

    # 这里假定您已经将代码文件上传到了服务器
    # 如果使用 Git，则取消下面的注释并相应修改
    # git clone https://github.com/your-username/your-repository.git .

    # 验证必要的文件是否存在
    if [[ ! -f "$APP_DIR/web_server.go" || ! -f "$APP_DIR/main.go" ]]; then
        log_error "关键文件不存在。请确保已将项目代码上传到 $APP_DIR 目录。"
        exit 1
    fi

    log_info "项目代码已就位。"
}

# 设置环境变量
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
EOF

    chmod 600 "$APP_DIR/.env"
    sudo chown "$USER_NAME:$USER_NAME" "$APP_DIR/.env"
}

# 安装 Go 依赖并构建
build_application() {
    log_info "安装 Go 依赖并构建应用..."

    cd "$APP_DIR"

    # 确保使用中国 Go 代理
    export GOPROXY=https://goproxy.cn,direct

    # 安装依赖
    go mod tidy

    # 构建应用
    log_info "构建 Web 服务器应用..."
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bin/web_app web_server.go

    log_info "构建 CLI 应用..."
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bin/train main.go

    log_info "应用构建完成。"
}

# 构建前端 (如果需要)
build_frontend() {
    if [ -d "$APP_DIR/frontend" ]; then
        log_info "构建前端应用..."

        cd "$APP_DIR/frontend"

        # 设置 npm 国内镜像源
        npm config set registry https://registry.npmmirror.com/
        npm config set disturl https://npmmirror.com/mirrors/node/

        # 安装前端依赖
        npm install

        # 构建前端
        npm run build

        cd "$APP_DIR"
        log_info "前端构建完成。"
    else
        log_warn "前端目录不存在，跳过前端构建。"
    fi
}

# 初始化数据库
initialize_database() {
    log_info "初始化数据库..."

    # 确保数据目录存在
    mkdir -p "$DATA_DIR"

    # 检查是否已存在数据库文件
    if [ -f "$DATA_DIR/app.db" ]; then
        log_warn "数据库文件已存在，跳过初始化。"
    else
        # 创建表结构
        sqlite3 "$DATA_DIR/app.db" < migrations/001_initial_schema.sql
        log_info "数据库初始化完成。"
    fi
}

# 创建 systemd 服务文件
create_systemd_service() {
    log_info "创建 systemd 服务文件..."

    sudo tee /etc/systemd/system/spaced-repetition.service << 'EOF'
[Unit]
Description=Spaced Repetition Multi-Tenant Application
After=network.target

[Service]
Type=simple
User=www-data
Group=www-data
WorkingDirectory=/opt/spaced-repetition
EnvironmentFile=/opt/spaced-repetition/.env
ExecStart=/opt/spaced-repetition/bin/web_app
Restart=always
RestartSec=10

# 确保数据目录有正确权限
ExecStartPre=/bin/mkdir -p /opt/spaced-repetition/data
ExecStartPre=/bin/chown -R www-data:www-data /opt/spaced-repetition/data

[Install]
WantedBy=multi-user.target
EOF
}

# 设置应用权限
set_permissions() {
    log_info "设置应用权限..."

    sudo chown -R www-data:www-data "$APP_DIR"
    sudo chmod +x "$APP_DIR/bin/web_app"
    sudo chmod +x "$APP_DIR/bin/train"
}

# 启动服务
start_service() {
    log_info "启动服务..."

    sudo systemctl daemon-reload
    sudo systemctl enable spaced-repetition
    sudo systemctl start spaced-repetition

    # 等待服务启动
    sleep 5

    if sudo systemctl is-active --quiet spaced-repetition; then
        log_info "服务启动成功！"
    else
        log_error "服务启动失败。请检查日志。"
        sudo journalctl -u spaced-repetition --no-pager -l
        exit 1
    fi
}

# 配置 Nginx (可选)
configure_nginx() {
    log_info "询问是否需要配置 Nginx..."

    read -p "是否需要配置 Nginx 反向代理？(y/n): " -n 1 -r
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

        sudo ln -s /etc/nginx/sites-available/spaced-repetition /etc/nginx/sites-enabled/

        # 测试 Nginx 配置
        sudo nginx -t

        # 重新加载 Nginx
        sudo systemctl reload nginx

        log_info "Nginx 配置完成。"

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

# 创建备份脚本
create_backup_script() {
    log_info "创建备份脚本..."

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

    # 添加定时备份 (每天凌晨2点)
    echo "0 2 * * * root /opt/backup-script.sh" | sudo tee -a /etc/crontab

    log_info "备份脚本创建完成。"
}

# 显示完成信息
show_completion_message() {
    log_info "部署完成！"
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
    echo "下一步操作建议："
    echo "1. 编辑 $APP_DIR/.env 文件并设置安全的 JWT_SECRET"
    echo "2. 上传您的 .md 问题文件到 $QUESTIONS_DIR 目录"
    echo "3. 访问应用并创建初始用户账户"
    echo "4. 定期备份数据 ($BACKUP_DIR 目录)"
    echo
    log_info "请妥善保管服务器访问凭据和 JWT 密钥。"
}

# 主函数
main() {
    log_info "开始自动化部署(中国优化版)..."

    setup_china_mirrors
    install_dependencies
    install_go
    install_nodejs  # 如果不需要重新构建前端，可以注释掉这一行
    create_directories
    setup_code
    setup_environment
    build_application
    build_frontend  # 如果不需要重新构建前端，可以注释掉这一行
    initialize_database
    create_systemd_service
    set_permissions
    start_service
    configure_nginx
    create_backup_script
    show_completion_message

    log_info "部署过程完成！"
}

# 执行主函数
main "$@"
