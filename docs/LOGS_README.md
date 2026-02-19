# 服务启动与日志管理

## 概述

本项目提供了完整的服务启动、停止和日志管理脚本，帮助您更方便地管理前后端服务。

## 脚本说明

### 1. start.sh - 启动服务

启动前后端服务并记录日志。

**功能特点：**
- ✅ 自动创建日志目录 (`logs/`)
- ✅ 每次启动创建带时间戳的日志文件
- ✅ 创建日志软链接方便访问最新日志
- ✅ 自动检测服务是否已运行
- ✅ 启动失败时提供清晰的错误提示

**日志文件命名：**
```
logs/
├── backend_20240217_150000.log  # 后端日志（带时间戳）
├── backend.log                  # 后端日志软链接（指向最新日志）
├── frontend_20240217_150000.log # 前端日志（带时间戳）
└── frontend.log                 # 前端日志软链接（指向最新日志）
```

**使用方法：**
```bash
# 启动服务（默认端口：后端 8000，前端 3000）
./start.sh

# 指定后端端口
BACKEND_PORT=9000 ./start.sh

# 指定前端端口
FRONTEND_PORT=4000 ./start.sh

# 指定 JWT secret
JWT_SECRET="my-secret" ./start.sh
```

**输出信息：**
启动成功后会显示：
- 服务端口
- 日志文件位置
- 实时查看日志的命令
- 停止服务的命令

---

### 2. stop.sh - 停止服务

停止前后端服务并记录停止时间。

**功能特点：**
- ✅ 优雅停止服务进程
- ✅ 清理残留进程
- ✅ 记录停止时间到日志
- ✅ 显示日志文件信息
- ✅ 提供清理旧日志的命令

**使用方法：**
```bash
./stop.sh
```

**停止过程：**
1. 尝试通过 PID 文件停止进程
2. 查找并停止所有相关进程
3. 强制停止残留进程
4. 记录停止时间到日志

---

### 3. logs.sh - 日志管理

提供丰富的日志查看和管理功能。

**使用方法：**

#### 查看帮助
```bash
./logs.sh -h
```

#### 列出所有日志文件
```bash
./logs.sh -l
```
输出示例：
```
[最新] backend.log  (大小: 2.5MB, 修改: 2024-02-17 15:30:00)
      backend_20240217_140000.log  (大小: 1.2MB, 修改: 2024-02-17 14:00:00)
[最新] frontend.log  (大小: 800KB, 修改: 2024-02-17 15:30:00)
      frontend_20240217_140000.log  (大小: 400KB, 修改: 2024-02-17 14:00:00)
```

#### 实时查看日志
```bash
# 实时查看所有日志
./logs.sh -f

# 实时查看后端日志
./logs.sh -f backend

# 实时查看前端日志
./logs.sh -f frontend
```

#### 查看最后 N 行日志
```bash
# 查看所有日志最后 50 行（默认）
./logs.sh -t

# 查看后端日志最后 100 行
./logs.sh -t backend 100

# 查看前端日志最后 20 行
./logs.sh -t frontend 20
```

#### 查看错误日志
```bash
# 查看所有错误日志
./logs.sh -e

# 查看后端错误日志
./logs.sh -e backend

# 查看前端错误日志
./logs.sh -e frontend
```

#### 清理旧日志
```bash
# 清理 7 天前的旧日志（默认）
./logs.sh -c

# 清理 30 天前的旧日志
./logs.sh -c 30
```

#### 删除所有日志
```bash
./logs.sh -d
```
⚠️ 警告：此操作将删除所有日志文件，需要确认。

---

## 日志格式

### 时间戳格式
所有日志都包含精确的时间戳：
```
[2024-02-17 15:30:00] 启动后端服务...
[2024-02-17 15:30:05] 后端服务启动成功 (PID: 12345)
[2024-02-17 15:30:10] 收到请求: GET /api/stats
```

### 日志级别
- `INFO`：一般信息
- `ERROR`：错误信息
- `WARN`：警告信息
- `DEBUG`：调试信息

---

## 常见使用场景

### 场景 1：启动服务并查看实时日志
```bash
# 1. 启动服务
./start.sh

# 2. 查看后端实时日志
./logs.sh -f backend
```

### 场景 2：排查启动问题
```bash
# 1. 查看错误日志
./logs.sh -e backend

# 2. 查看完整日志
./logs.sh -t backend 100
```

### 场景 3：定期清理日志
```bash
# 每周清理一次 7 天前的旧日志
./logs.sh -c 7
```

### 场景 4：查看历史日志
```bash
# 1. 列出所有日志文件
./logs.sh -l

# 2. 查看特定日志文件
tail -f logs/backend_20240217_140000.log
```

---

## 环境变量

可以通过环境变量自定义配置：

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `BACKEND_PORT` | 8000 | 后端服务端口 |
| `FRONTEND_PORT` | 3000 | 前端服务端口 |
| `JWT_SECRET` | my-secret-key | JWT 加密密钥 |

**使用示例：**
```bash
# 启动时指定端口
BACKEND_PORT=9000 FRONTEND_PORT=4000 ./start.sh

# 指定 JWT 密钥
JWT_SECRET="my-secure-secret-key" ./start.sh
```

---

## 故障排查

### 问题 1：服务启动失败

**检查步骤：**
```bash
# 1. 查看错误日志
./logs.sh -e

# 2. 查看完整日志
./logs.sh -t

# 3. 检查端口是否被占用
lsof -i :8000  # 后端
lsof -i :3000  # 前端
```

### 问题 2：日志文件不存在

**可能原因：**
- 服务未启动
- 日志目录权限问题

**解决方案：**
```bash
# 1. 确保日志目录存在
mkdir -p logs

# 2. 检查目录权限
ls -la logs

# 3. 重新启动服务
./stop.sh
./start.sh
```

### 问题 3：日志文件过大

**解决方案：**
```bash
# 清理旧日志
./logs.sh -c

# 或者删除所有日志后重新启动
./logs.sh -d
./start.sh
```

---

## 最佳实践

1. **定期清理日志**
   - 建议每周清理一次 7 天前的旧日志
   - 可以设置 crontab 自动清理

2. **使用软链接访问最新日志**
   - 直接访问 `logs/backend.log` 和 `logs/frontend.log`
   - 无需查找带时间戳的日志文件

3. **查看错误优先**
   - 遇到问题先使用 `./logs.sh -e` 查看错误日志
   - 节省排查时间

4. **实时监控关键服务**
   - 使用 `./logs.sh -f backend` 实时监控后端
   - 快速发现和定位问题

---

## 进阶配置

### 设置自动日志清理（crontab）

```bash
# 编辑 crontab
crontab -e

# 添加以下行（每周日凌晨 2 点清理 7 天前的日志）
0 2 * * 0 /path/to/Self-improvement/logs.sh -c 7
```

### 设置日志轮转（logrotate）

如果需要更专业的日志管理，可以使用 logrotate：

1. 创建配置文件 `/etc/logrotate.d/spaced-repetition`：
```
/path/to/Self-improvement/logs/*.log {
    daily
    rotate 7
    compress
    delaycompress
    notifempty
    create 0644 qiuxin staff
    missingok
}
```

2. 测试配置：
```bash
logrotate -d /etc/logrotate.d/spaced-repetition
```

---

## 总结

| 脚本 | 功能 | 关键命令 |
|------|------|----------|
| `start.sh` | 启动前后端服务 | `./start.sh` |
| `stop.sh` | 停止前后端服务 | `./stop.sh` |
| `logs.sh` | 日志管理 | `./logs.sh -l` (列表) |
| | | `./logs.sh -f` (实时) |
| | | `./logs.sh -e` (错误) |
| | | `./logs.sh -c` (清理) |

如有问题，请先查看日志文件获取详细信息。
