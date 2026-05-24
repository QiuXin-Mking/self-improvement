# q
Docker 容器的启动行为由镜像中的哪两个指令共同决定？
# a
CMD 和 ENTRYPOINT 指令共同决定。

# q
如何检查 Docker 容器内的启动命令？
# a
```bash
docker inspect <container_name> --format '{{.Config.Entrypoint}} {{.Config.Cmd}}'
```
例如：
```bash
docker inspect vue3-dev --format '{{.Config.Entrypoint}} {{.Config.Cmd}}'
```

# q
运行 `docker inspect vue3-dev --format '{{.Config.Entrypoint}} {{.Config.Cmd}}'` 得到 `[docker-entrypoint.sh] [bash]`，这代表什么含义？
# a
- `Entrypoint`：主命令，是容器的启动入口，即要运行的核心程序或脚本，不可随意替换。
- `Cmd`：参数，是传递给 Entrypoint 的运行参数，可以灵活修改或覆盖。

# q
如何在已有容器 `vue3-dev` 中配置自启动 FRP 并保留交互式 Shell？（约束：必须基于现有容器，不能仅用镜像）
# a
已知该容器的入口点为 `/docker-entrypoint.sh`，可通过以下脚本实现：
```sh
#!/bin/bash
# 执行原入口点初始化
/docker-entrypoint.sh bash -c "
  # 静默启动 FRP（完全丢弃日志）
  /app/frp_0.63.0_linux_amd64/frpc -c /app/frp_0.63.0_linux_amd64/frpc.toml > /dev/null 2>&1 &
  # 保留交互式 Shell
  exec bash
"
```

