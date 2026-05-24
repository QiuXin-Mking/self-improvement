# q
如何查看Lustre目录或文件当前所在的MDT索引？
# a
使用 `lfs getdirstripe -m <目录>` 查看目录分布的MDT，使用 `lfs getstripe -m <文件>` 查看文件所在的MDT。例如：
```
lfs getdirstripe -m /mnt/lustre/test2
lfs getstripe -m /mnt/lustre/test2/Fio.25*.0
```
输出为整数，表示MDT编号（如0或1）。

# q
如何将Lustre目录（及其下的文件）迁移到指定的MDT？
# a
使用 `lfs migrate -m <MDT索引> <目录>` 命令。例如将 `/mnt/lustre/test2` 迁移到MDT0000：
```
lfs migrate -m 0 /mnt/lustre/test2
```
迁移后可通过 `lfs getstripe -m` 验证文件所在MDT已变为目标索引。

# q
在Lustre中如何限制某个OST上创建新对象，以便进行维护操作？
# a
将该OST对应的 `max_create_count` 参数设为0，有两种方式：
- 通过 `lctl` 设置：
  ```bash
  lctl set_param osp.nas_test-OST0010-osc-MDT*.max_create_count=0
  ```
- 直接写入 sysfs：
  ```bash
  echo 0 > /sys/fs/lustre/osp/nas_test-OST0010-osc*/max_create_count
  ```
完成后，客户端将不会在该OST上创建新对象。

# q
挂载Lustre文件系统客户端的基本命令是什么？
# a
使用 `mount -t lustre` 并指定MGS的网络地址和文件系统名称：
```bash
mount -t lustre 172.31.0.26@tcp:/nas_test /mnt/lustre2
```
其中 `172.31.0.26@tcp` 是MGS的NID，`nas_test` 是文件系统名称。

