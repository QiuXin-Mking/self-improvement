#!/bin/bash

# 日志管理脚本 - 查看和管理前后端日志

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LOGS_DIR="$SCRIPT_DIR/logs"
BACKEND_LOG_LINK="$LOGS_DIR/backend.log"
FRONTEND_LOG_LINK="$LOGS_DIR/frontend.log"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 检查日志目录是否存在
if [ ! -d "$LOGS_DIR" ]; then
    echo -e "${RED}日志目录不存在: $LOGS_DIR${NC}"
    echo "请先运行 ./start.sh 启动服务"
    exit 1
fi

# 显示使用说明
show_usage() {
    cat << EOF
日志管理脚本

用法: ./logs.sh [选项] [服务]

选项:
    -h, --help              显示此帮助信息
    -l, --list              列出所有日志文件
    -f, --follow [service]  实时查看日志 (服务: backend/frontend/all)
    -t, --tail [service]    查看最后 50 行日志 (服务: backend/frontend/all)
    -e, --error [service]   查看错误日志 (服务: backend/frontend/all)
    -c, --clean             清理超过 7 天的旧日志
    -d, --delete-all        删除所有日志文件

服务:
    backend                后端服务日志
    frontend               前端服务日志
    all                   所有日志 (默认)

示例:
    ./logs.sh -l                    列出所有日志文件
    ./logs.sh -f                    实时查看所有日志
    ./logs.sh -f backend            实时查看后端日志
    ./logs.sh -t                    查看所有日志最后 50 行
    ./logs.sh -e backend            查看后端错误日志
    ./logs.sh -c                    清理 7 天前的旧日志
    ./logs.sh -d                    删除所有日志文件
EOF
}

# 列出所有日志文件
list_logs() {
    echo -e "${GREEN}=========================================="
    echo "          日志文件列表"
    echo "==========================================${NC}"
    echo ""

    if [ -z "$(ls -A "$LOGS_DIR"/*.log 2>/dev/null)" ]; then
        echo -e "${YELLOW}没有找到日志文件${NC}"
        return
    fi

    # 获取所有日志文件并按时间排序
    for log_file in $(ls -t "$LOGS_DIR"/*.log 2>/dev/null); do
        filename=$(basename "$log_file")
        size=$(du -h "$log_file" | cut -f1)
        mtime=$(stat -f "%Sm" -t "%Y-%m-%d %H:%M:%S" "$log_file" 2>/dev/null || stat -c "%y" "$log_file" 2>/dev/null)

        if [ "$filename" = "backend.log" ] || [ "$filename" = "frontend.log" ]; then
            echo -e "${BLUE}[最新]${NC} $filename  (大小: $size, 修改: $mtime)"
        else
            echo "      $filename  (大小: $size, 修改: $mtime)"
        fi
    done

    echo ""
    echo -e "日志目录: $LOGS_DIR"
}

# 实时查看日志
follow_logs() {
    local service=$1

    if [ "$service" = "backend" ]; then
        if [ -f "$BACKEND_LOG_LINK" ]; then
            echo -e "${GREEN}实时查看后端日志 (按 Ctrl+C 退出)...${NC}"
            echo ""
            tail -f "$BACKEND_LOG_LINK"
        else
            echo -e "${RED}后端日志文件不存在${NC}"
        fi
    elif [ "$service" = "frontend" ]; then
        if [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${GREEN}实时查看前端日志 (按 Ctrl+C 退出)...${NC}"
            echo ""
            tail -f "$FRONTEND_LOG_LINK"
        else
            echo -e "${RED}前端日志文件不存在${NC}"
        fi
    else
        # 同时查看前后端日志
        if [ -f "$BACKEND_LOG_LINK" ] && [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${GREEN}实时查看所有日志 (按 Ctrl+C 退出)...${NC}"
            echo ""
            tail -f "$BACKEND_LOG_LINK" "$FRONTEND_LOG_LINK"
        else
            echo -e "${RED}日志文件不存在${NC}"
        fi
    fi
}

# 查看日志最后几行
tail_logs() {
    local service=$1
    local lines=${2:-50}

    if [ "$service" = "backend" ]; then
        if [ -f "$BACKEND_LOG_LINK" ]; then
            echo -e "${GREEN}后端日志最后 $lines 行:${NC}"
            echo ""
            tail -n "$lines" "$BACKEND_LOG_LINK"
        else
            echo -e "${RED}后端日志文件不存在${NC}"
        fi
    elif [ "$service" = "frontend" ]; then
        if [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${GREEN}前端日志最后 $lines 行:${NC}"
            echo ""
            tail -n "$lines" "$FRONTEND_LOG_LINK"
        else
            echo -e "${RED}前端日志文件不存在${NC}"
        fi
    else
        # 同时显示前后端日志
        if [ -f "$BACKEND_LOG_LINK" ] && [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${GREEN}所有日志最后 $lines 行:${NC}"
            echo ""
            echo -e "${BLUE}=== 后端日志 ===${NC}"
            tail -n $((lines / 2)) "$BACKEND_LOG_LINK"
            echo ""
            echo -e "${BLUE}=== 前端日志 ===${NC}"
            tail -n $((lines / 2)) "$FRONTEND_LOG_LINK"
        else
            echo -e "${RED}日志文件不存在${NC}"
        fi
    fi
}

# 查看错误日志
error_logs() {
    local service=$1

    if [ "$service" = "backend" ]; then
        if [ -f "$BACKEND_LOG_LINK" ]; then
            echo -e "${RED}后端错误日志:${NC}"
            echo ""
            grep -i "error\|fail\|fatal\|panic\|exception" "$BACKEND_LOG_LINK" | tail -20
        else
            echo -e "${RED}后端日志文件不存在${NC}"
        fi
    elif [ "$service" = "frontend" ]; then
        if [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${RED}前端错误日志:${NC}"
            echo ""
            grep -i "error\|fail\|fatal\|panic\|exception" "$FRONTEND_LOG_LINK" | tail -20
        else
            echo -e "${RED}前端日志文件不存在${NC}"
        fi
    else
        # 同时显示前后端错误日志
        if [ -f "$BACKEND_LOG_LINK" ] && [ -f "$FRONTEND_LOG_LINK" ]; then
            echo -e "${RED}所有错误日志:${NC}"
            echo ""
            echo -e "${BLUE}=== 后端错误 ===${NC}"
            grep -i "error\|fail\|fatal\|panic\|exception" "$BACKEND_LOG_LINK" | tail -10
            echo ""
            echo -e "${BLUE}=== 前端错误 ===${NC}"
            grep -i "error\|fail\|fatal\|panic\|exception" "$FRONTEND_LOG_LINK" | tail -10
        else
            echo -e "${RED}日志文件不存在${NC}"
        fi
    fi
}

# 清理旧日志
clean_logs() {
    local days=${1:-7}

    echo -e "${YELLOW}清理超过 $days 天的旧日志...${NC}"

    # 查找并删除旧日志文件
    local count=$(find "$LOGS_DIR" -name "*.log" -type f -mtime +$days 2>/dev/null | wc -l)

    if [ "$count" -gt 0 ]; then
        find "$LOGS_DIR" -name "*.log" -type f -mtime +$days -delete
        echo -e "${GREEN}已删除 $count 个旧日志文件${NC}"
    else
        echo -e "${GREEN}没有需要清理的日志文件${NC}"
    fi
}

# 删除所有日志
delete_all_logs() {
    echo -e "${RED}警告: 即将删除所有日志文件！${NC}"
    read -p "确定要删除吗？(yes/no): " confirm

    if [ "$confirm" = "yes" ]; then
        rm -f "$LOGS_DIR"/*.log
        echo -e "${GREEN}所有日志文件已删除${NC}"
    else
        echo -e "${YELLOW}操作已取消${NC}"
    fi
}

# 主逻辑
case "${1:-}" in
    -h|--help)
        show_usage
        ;;
    -l|--list)
        list_logs
        ;;
    -f|--follow)
        follow_logs "${2:-all}"
        ;;
    -t|--tail)
        tail_logs "${2:-all}" "${3:-50}"
        ;;
    -e|--error)
        error_logs "${2:-all}"
        ;;
    -c|--clean)
        clean_logs "${2:-7}"
        ;;
    -d|--delete-all)
        delete_all_logs
        ;;
    *)
        show_usage
        ;;
esac
