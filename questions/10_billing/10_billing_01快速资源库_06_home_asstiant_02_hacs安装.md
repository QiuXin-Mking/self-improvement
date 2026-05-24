# q
HACS 的 GitHub 官方发布页下载链接是什么？
# a
https://github.com/hacs/integration/releases

# q
在 Home Assistant 容器环境下，安装 HACS 需要将下载的 hacs.zip 解压到哪个目录？
# a
/mnt/config/custom_components/ （即容器内映射的 config/custom_components，执行 `unzip hacs.zip -d hacs` 后形成 `hacs` 子目录）

# q
在 Home Assistant 中启用 HACS 集成后，如何重启服务？
# a
通过 `docker restart <容器ID>` 重启 Home Assistant 容器（如示例中 `docker restart 19`）

# q
查看 Home Assistant Docker 容器的绑定挂载路径（Binds）可以使用什么命令？
# a
`docker inspect 19ff9bb1626d`（其中 19ff9bb1626d 为容器 ID）并查看输出中的 "Binds" 部分

