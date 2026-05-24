# q
在Home Assistant的Docker部署中，安装HACS集成需要先创建哪些目录？
# a
需要在 Home Assistant 的配置目录（如 `/mnt/config`）下创建 `www` 和 `custom_components` 两个文件夹。

# q
如何通过命令行下载并安装 HACS 2.0.5 版本？
# a
进入 `custom_components` 目录，然后执行：
```
wget https://github.com/hacs/integration/releases/download/2.0.5/hacs.zip
unzip hacs.zip -d hacs
```

# q
HACS 集成文件安装完成后，如何重启通过 Docker 运行的 Home Assistant 以使其生效？
# a
使用 `docker restart` 命令重启对应的容器，例如：
```
docker restart 19ff9bb1626d
```
或使用容器名称：
```
docker restart home-assistant
```

