# q
如何将 `start_services.sh` 脚本集成到容器镜像的默认启动命令中？
# a
将 `start_services.sh` 封装到镜像的默认启动命令里面，容器启动时自动执行该脚本。

# q
如何通过环境变量指定对象存储（如 S3）挂载到容器内的目录？
# a
通过设置环境变量的方式指定对象挂载容器目录，容器启动时会读取该环境变量来决定挂载路径。

# q
Jupyter Lab 启动时如何配置访问 Token？其默认值是什么？
# a
可以通过设置环境变量指定 Token 值，如果设置了则以设置的为准，未设置则使用默认 Token `3bv8uoqs855ukot7kr3c0s9nawj5b908srkb7tur89r248sn`。启动命令示例：
```
jupyter lab --ip=0.0.0.0 --port=8888 --allow-root --no-browser --NotebookApp.token="3bv8uoqs855ukot7kr3c0s9nawj5b908srkb7tur89r248sn" > /var/log/jupyter.log 2>&1 &
```

# q
容器镜像从本地导出到远程主机的完整流程是什么？
# a
流程为：`container` → `image` → `export` → `local tar` → `remote tar` → `import`，即将运行容器提交为镜像，导出为本地 tar 包，传输到远程后导入为镜像。

