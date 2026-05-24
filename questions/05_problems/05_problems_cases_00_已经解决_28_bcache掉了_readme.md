# q
如何查看系统中所有bcache设备节点及其主次设备号？
# a
使用 `ls -l /dev/bcache*`，输出类似：
```
brw-rw---- 1 root disk 253,    0 Jan 10  2024 /dev/bcache0
brw-rw---- 1 root disk 253,  128 Jan 10  2024 /dev/bcache1
...
brw-rw---- 1 root disk 253, 2944 Jan 10  2024 /dev/bcache23
```
所有bcache设备的主设备号均为253，次设备号依次递增。

# q
如何查看系统中已注册的bcache缓存集UUID？
# a
执行 `ls /sys/fs/bcache/`，输出为一系列UUID目录，例如 `104da218-b9b7-41ef-9250-9a0c3effb9d4`，每个目录对应一个缓存集。

# q
当bcache设备丢失时，如何快速判断缓存集是否仍然被内核识别？
# a
检查 `/sys/fs/bcache/` 目录下的UUID列表，若某个缓存集对应的UUID目录存在，说明该缓存集仍被内核识别；若该目录消失，则表示缓存集已注销或缓存设备故障。

