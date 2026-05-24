# q
在这个Lustre部署规划中，如何根据预期文件数量计算MDT所需容量？
# a
基于单个inode大小计算：每个inode约为2.4 KiB（1024+1536），若预期1亿个文件，则MDT容量需求约为 100,000,000 × 2.4 / 1024 / 1024 ≈ 228.88 GiB，向上取整约需250 GiB。实际配置中创建了5个50 GiB的MDT，合计250 GiB。

# q
如何使用Ceph RBD为Lustre创建和管理块设备？
# a
使用 `rbd create` 命令创建指定大小的RBD映像，并启用 `layering` 特性；之后在各节点用 `sudo rbd map rbd/<映像名> --id admin` 映射为本地块设备；调整大小使用 `rbd resize`。例如：
```bash
rbd create --size 4860G rbd/lustre_ost00 --image-feature layering
sudo rbd map rbd/lustre_ost00 --id admin
```

# q
部署Lustre之前需要完成哪些系统层面的准备工作？
# a
1. **内核替换**：安装定制内核RPM包并重启（`rpm -ivh --nodeps --force ./*.rpm`）。  
2. **时钟源调整**：在 `/etc/default/grub` 的 `GRUB_CMDLINE_LINUX` 中添加 `tsc=reliable tsc=nowatchdog`，然后执行 `grub2-mkconfig -o /boot/grub2/grub.cfg` 并重启。  
3. **本地软件仓库配置**：用 `createrepo` 创建本地YUM仓库，添加 `lustre.repo`，安装依赖包（如 `libmount-devel`, `libyaml-devel`, `e2fsprogs-devel`, `e2fsprogs`）。  
4. **安装Lustre软件**：依次安装客户端和服务端的RPM包（`cd /home/client; rpm -ivh --nodeps --force ./*.rpm` 和 `/home/server/` 下的同样操作）。

# q
如何配置Lustre使用指定的网络接口？
# a
在 `/etc/modprobe.d/lustre.conf` 中配置 `lnet` 模块参数，指定网络类型和接口，例如：
```
options lnet networks=tcp(bond0.5)
```
这表示Lustre将使用TCP网络，绑定在 `bond0.5` 接口上。

# q
mkfs.lustre命令中各个关键参数的作用是什么？
# a
`mkfs.lustre` 用于创建Lustre文件系统组件，关键参数：
- `--fsname=st_nas`：设定文件系统名称。
- `--mgs`：指定该设备为管理服务（MGS）。
- `--mdt` / `--ost`：分别表示元数据目标（MDT）或对象存储目标（OST）。
- `--mgsnode=192.168.5.171@tcp`：指定MGS节点的网络地址（LNet NID），可多个实现高可用。
- `--reformat`：强制重新格式化已有设备。
- `--index=0`：组件在集群中的索引，MDT和OST从0开始编号。

