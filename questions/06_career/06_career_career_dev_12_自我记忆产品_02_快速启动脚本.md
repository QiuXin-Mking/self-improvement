# q
启动 vue3-dev 开发容器的脚本中包含怎样的 Docker 引擎就绪检测逻辑？
# a
脚本通过循环执行 `docker info` 检测 Docker 是否可用，超时时间为 60 秒，每 3 秒重试一次；超时后将提示错误并退出（`exit 1`）。示例代码如下：
```bash
timeout=60
waited=0
while ! docker info > /dev/null 2>&1; do
    if [ $waited -ge $timeout ]; then
        echo "等待超时，Docker 未能启动"
        exit 1
    fi
    echo "Docker 尚未就绪，等待 3 秒..."
    sleep 3
    waited=$((waited + 3))
done
```

# q
创建 vue3-dev 容器的完整 `docker run` 命令是什么？各参数含义如何？
# a
创建命令：
```bash
docker run -itd \
  --name vue3-dev \
  -p 5173:5173 \
  -v "C:\Users\Administrator\Desktop\vue3-demo:/app" \
  -v "D:\qiuxin_aliyun_back\04_职场修炼\12_自我记忆产品\Self-improvement:/Self-improvement" \
  -w /app \
  node:20 \
  bash
```
- `-itd`：交互式、分配终端、后台运行
- `--name vue3-dev`：指定容器名称
- `-p 5173:5173`：将宿主机 5173 端口映射到容器 5173（Vite 默认端口）
- 第一个 `-v`：挂载宿主机 `C:\...\vue3-demo` 到容器 `/app`
- 第二个 `-v`：挂载宿主机 `D:\...\Self-improvement` 到容器 `/Self-improvement`
- `-w /app`：设置容器工作目录为 `/app`
- `node:20`：使用 Node.js 20 镜像
- `bash`：容器启动后执行的命令（交互式 bash）

# q
在 Windows 下启动 Docker Desktop 时，路径含空格如何处理？
# a
在脚本中使用带空格的路径时，需要用引号括起来，例如：
```bash
"/drives/C/Program Files/Docker/Docker/Docker Desktop.exe" &
```

