# q
如何将 raw 格式的磁盘镜像转换为 qcow2 格式？
# a
使用 `qemu-img convert` 命令：
```bash
qemu-img convert -f raw -O qcow2 源文件.raw 目标文件.qcow2
```

# q
如何将 qcow2 格式的磁盘镜像转换为 raw 格式？
# a
使用 `qemu-img convert` 命令：
```bash
qemu-img convert -f qcow2 -O raw 源文件.qcow2 目标文件.raw
```

