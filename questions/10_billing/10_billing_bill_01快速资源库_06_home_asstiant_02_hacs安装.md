# q
HACS集成的最新稳定版下载地址是什么？
# a
```bash
wget https://github.com/hacs/integration/releases/download/2.0.5/hacs.zip
```

# q
在 Home Assistant Docker 环境中，自定义组件（如 HACS）应放置在宿主机的哪个目录？
# a
根据容器的 Binds 映射，自定义组件目录为 `/mnt/config/custom_components`，需要在其中解压 HACS。

# q
如何查看 Home Assistant 容器的宿主机挂载路径？
# a
使用 `docker inspect 19ff9bb1626d` 查看输出的 "Binds" 字段。

# q
安装 HACS 后，如何重启 Home Assistant 容器使其生效？
# a
执行 `docker restart 19`（其中 19 是容器 ID 的前缀，也可用完整 ID 或容器名）。

