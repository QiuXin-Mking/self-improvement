# q
在Lustre缩容MDT时，如何将MDT上的元数据迁移到其他MDT？
# a
使用命令 `lfs migrate -m <目标MDT索引列表> /mnt/lustre/*`，例如将MDT0的元数据迁移到MDT1、MDT2、MDT3：
```bash
lfs migrate -m 1,2,3 /mnt/lustre/*
```

# q
缩容OST时，如何禁止在待移除的OST上创建新文件？
# a
在所有节点上，通过 `lctl dl` 找出目标OST对应的 `osp` 实例名称（如 `osp.nas_test-OST0000-osc-MDT0000`），然后设置其 `max_create_count=0`：
```bash
lctl set_param osp.nas_test-OST0000-osc-MDT0000.max_create_count=0
```

# q
如何将待缩容OST上的所有对象数据迁移到其他OST？
# a
使用 `lfs find` 筛选目标OST上的文件，然后通过 `lfs_migrate` 迁移：
```bash
lfs find /mnt/lustre/ -obd nas_test-OST0000 | lfs_migrate -y
```
该操作耗时长，且因之前已禁止新数据写入，迁移不会再将数据写回原OST。

