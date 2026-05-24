# q
如何配置 Docker 守护进程使用阿里云容器镜像加速？
# a
编辑 `/etc/docker/daemon.json`，添加 `"registry-mirrors"` 字段并指定加速地址，示例内容：
```json
{
  "dns": ["8.8.8.8", "8.8.4.4"],
  "registry-mirrors": ["https://1umjrwoj.mirror.aliyuncs.com"]
}
```

# q
修改 Docker 守护进程配置文件后，如何使更改生效？
# a
依次执行：
```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```

