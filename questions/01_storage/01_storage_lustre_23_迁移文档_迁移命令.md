# q
如何将文件迁移到指定的 MDT 索引？
# a
使用 `lfs migrate -m` 命令，例如将 `/mnt/lustre/test2` 迁移到 MDT0：
```bash
lfs migrate -m 0 /mnt/lustre/test2
```

# q
如何将文件迁移到指定的 OST 或实现手动条带化？
# a
- 迁移到单个 OST：
  ```bash
  lfs migrate -o 2 /mnt/lustre/test2
  lfs migrate -o 1 /mnt/lustre/stg_ssd_2-102-174/Fio.0.0
  ```
- 手动指定多个 OST 实现条带化：
  ```bash
  lfs migrate -o 4,5,6,7 /mnt/lustre/stg_ssd_2-102-174/*
  lfs migrate -o 0,1 /mnt/lustre/net-101-171/*
  ```

# q
如何使用 `lfs_migrate` 进行随机离散迁移并避免阻塞？
# a
使用 `--non-block --non-direct -y` 参数，例如：
```bash
lfs_migrate --non-block --non-direct -y
```
若需针对特定 OST 上的文件进行迁移，结合 `lfs find`：
```bash
lfs find --ost nas-OST0005_UUID /mnt/lustre/ | lfs_migrate --non-block --non-direct -y
```

# q
如何查询文件的条带化布局信息？
# a
使用 `lfs getstripe` 命令：
```bash
lfs getstripe /mnt/lustre/stg_ssd_2-102-174/Fio.0.0
```

# q
如何解决 `lfs migrate` 报错 “Device or resource busy”？
# a
使用 `lfs find` 结合 `lfs_migrate`，并指定重试间隔（`-i`）与并发数（`-c`），例如：
```bash
lfs find --ost nas-OST0005_UUID /mnt/lustre/ | lfs_migrate --non-block --non-direct -i 10 -c 2 -y
lfs find --ost mj_nas-OST0000_UUID /mnt/lustre/ | lfs_migrate --non-block --non-direct -i 10 -c 2 -y
```

