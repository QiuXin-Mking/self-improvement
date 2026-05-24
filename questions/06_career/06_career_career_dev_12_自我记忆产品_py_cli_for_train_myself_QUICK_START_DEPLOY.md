# q
使用 PyInstaller 打包本项目最简单的步骤是什么？
# a
先在本地执行 `python build.py` 生成 `deploy/` 目录，然后将整个目录上传到服务器，最后在服务器上运行 `chmod +x train && ./train --init && ./train`。

# q
如何用 Docker 部署本项目并初始化知识库？
# a
在本地构建镜像 `docker build -t spaced-repetition-train .`，上传项目后，在服务器使用 `docker-compose run --rm train --init` 初始化，然后通过 `docker-compose up` 或 `docker-compose run --rm train` 启动训练。

# q
直接部署本项目时，服务器上手动执行的完整命令顺序是什么？
# a
```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python train.py --init
python train.py
```

# q
用 PyInstaller 打包后，服务器的初始化与训练命令是什么？
# a
初始化：`./train --init`；开始训练：`./train`；查看统计：`./train --stats`。

