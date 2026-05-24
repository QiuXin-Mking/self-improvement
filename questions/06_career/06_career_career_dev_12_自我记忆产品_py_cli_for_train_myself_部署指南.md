# q
使用PyInstaller打包部署时，本地打包需执行哪两个命令？
# a
```bash
pip install -r requirements-dev.txt
python build.py
```
打包完成后在 `dist/` 目录生成可执行文件（如 `train`），部署包自动创建于 `deploy/` 目录。

# q
Docker 容器化部署在服务器上如何初始化知识库？
# a
执行 `docker-compose run --rm train --init`。

# q
直接手动部署时，如何在服务器上创建虚拟环境并安装依赖？
# a
```bash
cd /opt/spaced-repetition-train
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```
随后执行 `python train.py --init` 初始化知识库。

