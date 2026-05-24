# q
如何构建 node20-claude-frp 的 Docker 镜像？
# a
在 `修改部署和配置方案` 目录下执行：
```bash
cd "修改部署和配置方案"
docker build -t node20-claude-frp:latest .
```

# q
运行该容器时需要映射哪些端口？对应的作用是什么？
# a
端口映射如下：
- `-p 5173:5173`：宿主机 5173 → 容器 Vite 开发服务器
- `-p 7000:7000`：宿主机 7000 → 容器 FRP 端口
容器内部的 22 端口不需要直接映射到宿主机，它通过 frpc 内网穿透映射到远程服务器的 5555 端口。

# q
如何通过 SSH 连接到该容器？
# a
使用 frpc 内网穿透后的地址连接：
```bash
ssh root@115.190.235.149 -p 5555
```
密码为 `MKING5@0610mking`。

# q
如何在容器内检查 frpc 和 SSH 服务的运行状态？
# a
进入容器后执行以下命令：
```bash
# 查看 frpc 进程
ps aux | grep frpc

# 查看 SSH 服务状态
ps aux | grep sshd

# 查看所有相关服务
ps aux | grep -E "frpc|sshd"
```

# q
该容器中预配置的 Claude Code 环境变量有哪些？
# a
- 版本：2.1.19
- 已配置的环境变量：
  - `ANTHROPIC_AUTH_TOKEN`: sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6
  - `ANTHROPIC_BASE_URL`: https://turingai.plus
  - `ANTHROPIC_SMALL_FAST_MODEL`: claude-3-5-haiku-20241022
- npm 淘宝镜像源已配置加速下载，Claude Code CLI 已全局安装。

