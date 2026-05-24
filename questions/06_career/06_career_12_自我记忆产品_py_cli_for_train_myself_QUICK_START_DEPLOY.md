# q
如何使用PyInstaller打包并部署这个训练CLI工具？
# a
本地操作：
```bash
pip install pyinstaller
python build.py
```
然后将 `deploy/` 目录上传到服务器。服务器上执行：
```bash
cd /opt/spaced-repetition-train
chmod +x train
./train --init
./train
```

# q
Docker部署方式下如何构建镜像并运行训练？
# a
本地构建镜像：
```bash
docker build -t spaced-repetition-train .
```
上传整个项目目录（含 Dockerfile 和 docker-compose.yml）到服务器，然后使用 docker-compose 运行：
```bash
docker-compose up
# 或先初始化再训练
docker-compose run --rm train --init
docker-compose run --rm train
```

# q
直接部署（不打包）时需要上传哪些文件？如何初始化并启动训练？
# a
需要上传的文件包括：`train.py`、`parser.py`、`spaced_repetition.py`、`requirements.txt` 以及 `questions/` 目录。推荐使用部署脚本：
```bash
chmod +x deploy.sh
sudo ./deploy.sh
```
或手动执行：
```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python train.py --init
python train.py
```

# q
三种部署方式对应的常用命令是什么？
# a
无论哪种方式，初始化知识库和启动训练的命令：
```bash
./train --init       # 初始化知识库
./train              # 开始训练
./train --stats      # 查看统计
```
如果是 Python 直接运行则替换为 `python train.py`，Docker 方式为 `docker-compose run --rm train`。

