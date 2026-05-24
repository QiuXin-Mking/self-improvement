# q
如何构建名为 `node20-claude-frp` 的 Docker 镜像？
# a
在 `修改部署和配置方案` 目录下执行：
```bash
docker build -t node20-claude-frp:latest .
```

# q
启动该容器的 `docker run` 命令是什么，并简述其端口和卷映射？
# a
```bash
docker run -itd \
  --name claude-dev \
  -p 5173:5173 \
  -p 7000:7000 \
  -v "C:\Users\Administrator\Desktop\vue3-demo:/app" \
  -v "D:\qiuxin_aliyun_back\:/qiuxin_aliyun_back" \
  -w /app \
  node20-claude-frp:latest
```
- 端口映射：将容器的 5173（Vite 开发服务器）和 7000（FRP 端口）分别映射到宿主机的 5173 和 7000。
- 卷映射：将 Windows 主机路径 `C:\Users\Administrator\Desktop\vue3-demo` 挂载为容器工作目录 `/app`，将 `D:\qiuxin_aliyun_back\` 挂载为 `/qiuxin_aliyun_back`。

# q
为什么容器 22 端口未映射到宿主机，以及如何通过 SSH 连接到该容器？
# a
容器内的 SSH 服务（22 端口）通过 frpc 内网穿透直接映射到远程服务器 `115.190.235.149` 的 `5555` 端口，因此无需在宿主机映射。连接命令：
```bash
ssh root@115.190.235.149 -p 5555
```
密码：`MKING5@0610mking`

# q
进入容器后如何查看 frpc 和 SSH 服务的运行状态？
# a
在容器内执行以下命令：
- 查看 frpc 进程：`ps aux | grep frpc`
- 查看 SSH 服务：`ps aux | grep sshd`
- 一键查看两个服务：`ps aux | grep -E "frpc|sshd"`

# q
容器中 Claude Code 的 API Key 和 Base URL 是如何配置的？
# a
通过环境变量设置：
- `ANTHROPIC_AUTH_TOKEN`: `sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6`
- `ANTHROPIC_BASE_URL`: `https://turingai.plus`
- `ANTHROPIC_SMALL_FAST_MODEL`: `claude-3-5-haiku-20241022`

