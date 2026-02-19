# 快速部署参考

## 🚀 两种部署方式快速选择

### 方式一：Go二进制打包（最简单）

**本地操作：**
```bash
# 1. 确保已安装Go
go version

# 2. 构建可执行文件
go build -o bin/train main.go
go build -o bin/web_app web_server.go

# 3. 创建部署目录
mkdir -p deploy
cp bin/train bin/web_app deploy/
cp -r questions/ deploy/
mkdir -p deploy/data
```

**服务器操作：**
```bash
cd /opt/spaced-repetition-go
chmod +x train web_app
./train --init
./train
```

---

### 方式二：直接部署

**服务器操作：**
```bash
# 手动部署
# 1. 安装Go
wget https://golang.org/dl/go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 2. 上传源代码
# 上传 main.go, web_server.go, internal/, go.mod, go.sum 等文件

# 3. 安装依赖
go mod tidy

# 4. 运行
go run main.go --init
go run main.go
```

---

## 📦 需要上传到服务器的文件

**Go二进制方式：**
- `train` (或 `train.exe`)
- `web_app` (或 `web_app.exe`)
- `questions/` 目录
- `data/` 目录（数据存储）

**直接部署方式：**
- `main.go`
- `web_server.go`
- `internal/` 目录
- `go.mod`
- `go.sum`
- `questions/` 目录（问题文件）

---

## ⚡ 常用命令

```bash
# 初始化知识库
./train --init

# 开始训练
./train

# 查看统计
./train --stats

# 启动Web服务
./web_app
```

---

详细说明请查看 `部署指南.md`

