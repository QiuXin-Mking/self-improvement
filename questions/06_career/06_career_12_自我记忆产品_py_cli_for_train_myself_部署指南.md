# q
部署完成后，如何初始化间隔重复训练系统的知识库？
# a
执行命令：
```bash
./train --init
```

# q
如何查看训练统计信息？
# a
执行命令：
```bash
./train --stats
```

# q
如何备份学习数据？
# a
复制 `data/learning_data.json` 文件，例如：
```bash
cp data/learning_data.json data/learning_data.json.bak
```
可配合 crontab 定时备份，如每天凌晨 2 点：
```bash
0 2 * * * cp /opt/spaced-repetition-train/data/learning_data.json /opt/spaced-repetition-train/data/learning_data.json.$(date +\%Y\%m\%d)
```

# q
使用 Docker 部署时，如何以交互模式运行训练并完成初始化？
# a
在项目目录下执行：
```bash
# 先构建镜像（如需）
docker-compose build

# 初始化知识库
docker-compose run --rm train --init

# 启动训练
docker-compose run --rm train
```

