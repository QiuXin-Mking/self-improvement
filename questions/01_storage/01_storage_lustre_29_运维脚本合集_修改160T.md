# q
如何缩容一个 Ceph RBD 镜像？
# a
使用 `rbd resize` 命令并指定 `--allow-shrink` 参数。例如，将 `lustre_ost00` 缩容至 4096G：
```bash
rbd resize rbd/lustre_ost00 --size 4096G --allow-shrink
```

# q
如何创建一个启用 layering 特性的 Ceph RBD 镜像？
# a
使用 `rbd create` 命令并加上 `--image-feature layering` 参数。例如，创建一个 4096G 的镜像：
```bash
rbd create --size 4096G rbd/lustre_ost20 --image-feature layering
```

# q
如何批量卸载并取消映射所有已挂载的 RBD 设备？
# a
可以执行以下循环脚本，先卸载设备再取消映射：
```bash
for dev in $(rbd showmapped | awk 'NR>1 {print $5}'); do
    umount $dev
    rbd unmap $dev
done
```

# q
如何使用 `mkfs.lustre` 创建一个 Lustre OST（对象存储目标）设备？
# a
使用 `--ost` 选项，并指定文件系统名、MGS 节点地址、索引号和设备路径：
```bash
mkfs.lustre --fsname=st_nas --ost --mgsnode=192.168.5.171@tcp --mgsnode=192.168.5.172@tcp --reformat --index=0 /dev/rbd1
```

