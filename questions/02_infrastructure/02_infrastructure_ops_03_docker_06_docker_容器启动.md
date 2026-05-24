# q
如何理解 `docker container run -d` 参数的作用？
# a
`-d`（或 `--detach`）让容器在**后台运行**（守护进程模式），启动后终端不会被容器占用，输出不会直接打印到控制台。

# q
`docker container run --restart always` 的 `--restart always` 重启策略是什么？
# a
表示容器在任何退出情况下（手动停止除外）**始终自动重启**，由 Docker 守护进程负责。适合生产环境保证服务可用性。

# q
`docker container run` 的 `--volume` 参数如何工作？下面命令中的挂载关系是什么？
```bash
--volume /etc/nginx/conf.d/http_nginx.conf:/etc/nginx/conf.d/http_nginx.conf
--volume /etc/ceph/ssl:/mnt
```
# a
`--volume` 将**宿主机路径**挂载到**容器内路径**。
- 第一条：将宿主机的 `/etc/nginx/conf.d/http_nginx.conf` 挂载到容器内的**同名路径**，直接覆盖容器内的默认配置。
- 第二条：将宿主机的 `/etc/ceph/ssl` **整个目录**挂载到容器内的 `/mnt`，容器可通过 `/mnt` 访问宿主机的 SSL 证书文件。

# q
`docker container run --network host` 的作用和适用场景是什么？
# a
`--network host` 让容器**与宿主机共享网络命名空间**，容器直接使用宿主机的 IP 和端口，**没有网络隔离**。适用于需要高性能网络或直接绑定宿主机端口的场景，例如代理服务（如 Nginx），避免 NAT 转换开销。

