# q
如何通过 `lfs migrate` 提升 AI 训练平台加载 500GB 数据集的读取带宽？
# a
将数据集以 16 条带、条带大小 4M 分散到指定的 SSD OST 上：
```bash
lfs migrate -c 16 -S 4M -o 0,2,4,6,8,10,12,14 /ai_dataset
```
效果：读取带宽从 2GB/s 提升至 25GB/s，训练迭代时间缩短 40%。

# q
如何解决视频编辑中单 HDD OST 带宽饱和（200MB/s）导致的卡顿问题？
# a
将视频项目文件迁移至 SSD 池，并使用 4 条带、1MB 条带大小：
```bash
lfs migrate -c 4 -S 1M --pool ssd_pool /video_project
```
效果：单文件读取速度达 800MB/s，实时编辑流畅度显著提升。

# q
在 Lustre 中如何对金融交易日志实现存储分层（热数据使用 SSD，冷数据归档到 HDD）？
# a
利用 `lfs migrate` 指定 OST 列表和迁移模式：
- 将热数据（7 天内高频访问）迁移至 SSD OST：
```bash
lfs migrate -o 1,3,5 --block /transaction_logs
```
- 30 天后将数据迁移至 HDD OST 作为归档（非阻塞模式）：
```bash
lfs migrate -o 11,13,15 --non-block /transaction_logs
```

# q
针对不同场景，`lfs migrate` 的参数配置最佳实践有哪些？
# a
| 场景需求 | 推荐参数组合 | 注意事项 |
|---------|-------------|---------|
| 百万级小文件目录 | `-c 4 -H fnv_1a_64` | 避免对小文件条带化 |
| TB 级大文件读写 | `-c 8 -S 4M` | 条带大小需匹配 I/O 块大小 |
| 存储分层/资源隔离 | `-o <ost_list> --pool <name>` | 提前规划 OST 分组策略 |
| 业务连续性要求高 | `--non-block --non-direct` | 迁移后需校验数据一致性 |

