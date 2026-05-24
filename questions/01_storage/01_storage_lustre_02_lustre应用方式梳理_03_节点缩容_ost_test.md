# q
在Lustre缩容过程中，如何阻止在特定OST上继续创建新对象？
# a
将该OST对应的OSP `max_create_count` 参数设置为 0。可通过 sysfs 或 `lctl set_param` 实现：
```bash
echo 0 > /sys/fs/lustre/osp/<ost_name>-osc*/max_create_count
```
或
```bash
lctl set_param osp.<ost_name>-osc-MDT*.max_create_count=0
```

# q
如何将Lustre OST设置为非活跃（inactive）状态，使其停止处理请求？
# a
使用 `lctl set_param` 将对应OSP的 `active` 参数置 0：
```bash
lctl set_param osp.<ost_name>-osc-MDT*.active=0
```

# q
如何使用Lustre命令查找指定OST上的所有文件并将其迁移到其他OST？
# a
先用 `lfs find --ost <OST_UUID> <mountpoint>` 找出文件，再通过管道交给 `lfs_migrate -y` 强制迁移：
```bash
lfs find --ost nas_test-OST000a_UUID /mnt/lustre/test1 | lfs_migrate -y
```

# q
如何查看Lustre文件系统中所有OST及其当前状态？
# a
执行 `lfs osts <挂载点>`，输出将列出每个OST的索引、名称和状态（如 ACTIVE）：
```bash
lfs osts /mnt/lustre
```

