# q
如何使用 PyInstaller 打包 Spaced Repetition Train 应用？
# a
在项目根目录执行 `python build.py`，打包完成后会在 `dist/` 目录下生成平台对应的可执行文件（如 `dist/train`）以及一个包含可执行文件、示例问题和数据的 `deploy/` 目录。

# q
使用 Docker 部署时，如何初始化知识库并进行训练？
# a
在服务器上进入项目目录后，依次执行：
```bash
docker-compose build
docker-compose run --rm train --init
docker-compose run --rm train
```

# q
如何在 Linux 服务器上定期备份 Spaced Repetition Train 的学习数据？
# a
使用 crontab 设置定时任务，例如每天凌晨2点备份：
```bash
crontab -e
# 添加
0 2 * * * cp /opt/spaced-repetition-train/data/learning_data.json /opt/spaced-repetition-train/data/learning_data.json.$(date +\%Y\%m\%d)
```

# q
部署后终端出现中文乱码，应如何设置 Linux 环境变量？
# a
执行以下命令，确保终端支持 UTF-8 编码：
```bash
export LANG=zh_CN.UTF-8
export LC_ALL=zh_CN.UTF-8
```

