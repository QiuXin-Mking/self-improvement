# q
如何创建一个指定大小的RBD镜像？
# a
使用 `rbd create` 命令，指定镜像名、大小和存储池：
```bash
rbd create <image-name> --size <size> --pool <pool-name>
```
示例：
```bash
rbd create qx_image --size 1G --pool pool1
```

# q
如何查看某个Ceph存储池中有多少个RBD镜像？
# a
使用 `rbd -p <pool-name> ls` 命令：
```bash
rbd -p pool1 ls
```

# q
如何将RBD镜像映射为本地块设备？
# a
使用 `rbd device map` 命令，指定存储池和镜像名：
```bash
rbd device map pool1/qx_image
```
命令执行后会返回一个特定的设备路径（如 `/dev/rbd0`）。

# q
如何取消RBD块设备的映射？
# a
使用 `rbd unmap` 命令，指定存储池和镜像名：
```bash
rbd unmap pool1/qx_image
```

