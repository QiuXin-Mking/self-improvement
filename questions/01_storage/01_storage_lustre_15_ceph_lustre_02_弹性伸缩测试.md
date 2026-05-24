# q
如何编译Lustre客户端rpm包？
# a
使用以下命令配置并生成RPM包：
```sh
./configure --disable-server --enable-client --enable-quota --with-linux=/home/chenlou/lustre_all_data/er-linux --with-zfs=no
make rpms
mv *.rpm ../client
```
生成的客户端RPM包最终存放在 `/home/chenlou/lustre_all_data/client` 目录。

# q
安装新内核后如何修改默认启动内核？
# a
使用 `grubby` 命令设置默认启动项：
```sh
grubby --set-default /boot/vmlinuz-4.18.0-553.5.1
```
查看当前默认内核：`grubby --default-kernel`  
查看所有已安装内核：`grubby --info=ALL`

# q
在Lustre弹性伸缩测试中，通常模拟哪些故障场景？
# a
- 单点故障：MDT 故障 + OSS 节点故障
- 单盘故障：OSS 节点故障
- 管理/元数据/对象存储服务切换：MGS、MDS、OSS 故障切换

