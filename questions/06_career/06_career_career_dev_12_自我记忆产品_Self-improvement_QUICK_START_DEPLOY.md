# q
PyInstaller打包部署的最简单流程是什么？
# a
本地打包：`pip install pyinstaller` 后执行 `python build.py`，然后将 `deploy/` 目录上传到服务器。  
服务器运行：
```bash
cd /opt/spaced-repetition-train
chmod +x train
./train --init
./train
```

# q
使用Docker部署时，本地构建和服务器运行的命令分别是什么？
# a
本地构建：`docker build -t spaced-repetition-train .`  
服务器运行：`docker-compose up`，或分步执行 `docker-compose run --rm train --init` 和 `docker-compose run --rm train`。

# q
直接部署的两种方式及其命令是什么？
# a
自动部署（推荐）：
```bash
chmod +x deploy.sh
sudo ./deploy.sh
```
手动部署：
```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python train.py --init
python train.py
```

# q
三种部署方式分别需要上传哪些文件到服务器？
# a
- PyInstaller方式：`deploy/` 整个目录
- Docker方式：整个项目目录（包括 Dockerfile 和 docker-compose.yml）
- 直接部署方式：`train.py`、`parser.py`、`spaced_repetition.py`、`requirements.txt` 和 `questions/` 目录

# q
初始化知识库和查看训练统计的常用命令是什么？
# a
初始化：`./train --init`（或 `python train.py --init`）  
开始训练：`./train`  
查看统计：`./train --stats`

