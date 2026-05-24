# q
如何使用 PyInstaller 打包并部署项目到服务器？
# a
1. 本地安装依赖：`pip install -r requirements-dev.txt`
2. 执行打包：`python build.py`，生成 `dist/` 下的可执行文件（如 `train`、`train.exe`）
3. 自动创建的 `deploy/` 目录包含可执行文件、`questions/`、`data/` 等
4. 上传部署包：`scp -r deploy/* user@server:/opt/spaced-repetition-train/`
5. SSH 登录后设置权限：`chmod +x train`
6. 初始化知识库：`./train --init`
7. 使用：`./train`（训练）或 `./train --stats`（查看统计）

# q
在 Docker 部署方式中，如何构建镜像、初始化知识库和运行交互式容器？
# a
1. 本地构建镜像：`docker build -t spaced-repetition-train .`
2. 上传项目：`scp -r . user@server:/opt/spaced-repetition-train/`
3. 服务器上构建和运行：
   - `docker-compose build`
   - 初始化：`docker-compose run --rm train --init`
   - 交互训练：`docker-compose run --rm train`
   - 查看统计：`docker-compose run --rm train --stats`
4. 后台运行（可选）：`docker-compose up -d`，查看日志 `docker-compose logs -f`

# q
直接部署时，如何手动创建虚拟环境并安装依赖？
# a
```bash
cd /opt/spaced-repetition-train
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```
然后初始化：`python train.py --init`

# q
如何为学习数据进行自动备份？给出 crontab 定时备份命令。
# a
编辑 crontab：`crontab -e`
添加每日凌晨2点备份一行：
```
0 2 * * * cp /opt/spaced-repetition-train/data/learning_data.json /opt/spaced-repetition-train/data/learning_data.json.$(date +\%Y\%m\%d)
```
此命令会在 `data/` 目录下生成带日期的备份文件。

# q
三种部署方式各有哪些优点和适用场景？
# a
- **PyInstaller打包**：无需服务器安装Python，单文件部署简单，适合快速、资源受限的部署
- **Docker容器化**：环境隔离，易版本管理和回滚，适合已有Docker环境的多服务器部署
- **直接部署**：完全控制环境，性能最优（无容器开销），适合需要自定义配置的开发测试场景

