# q
如何根据产品化规格和Ceph可用容量计算Lustre所需的OST总数及每节点分配？
# a
产品化规格每个OST固定为2TB。根据`ceph df`获取Ceph总可用容量（例如25TB），计算需创建的OST数量：
- 总OST数 = 总容量 ÷ 单OST容量 = 25T ÷ 2T = 12个
- 每节点分配的OST数 = 总OST数 ÷ 节点数 = 12 ÷ 3 = 4个
因此每个节点部署4个OST，共12个OST。

# q
在Lustre部署中，MDT容量如何根据现有环境比例进行估算？
# a
参考现有环境Lustre总容量与MDT总容量的比例。例如现有环境150T总容量对应250G MDT，则25T总容量对应MDT需求为：
25T × (250G ÷ 150T) ≈ 41.67G
再除以节点数（3）得每节点约13.67G，向上取整为20G。因此实际为每个MDT创建20GB的RBD。

# q
在三个节点上创建Lustre文件系统时，各节点分别需要执行哪些`mkfs.lustre`命令？请写出节点10.5.65.130的命令示例。
# a
节点10.5.65.130负责MGS、MDT0和OST0-3：
```bash
mkfs.lustre --fsname=nas --mgs --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp  /dev/rbd5
mkfs.lustre --fsname=nas --mdt --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp  --reformat --index=0  /dev/rbd0
mkfs.lustre --fsname=nas --ost --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp   --reformat --index=0  /dev/rbd1
mkfs.lustre --fsname=nas --ost --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp   --reformat --index=1  /dev/rbd2
mkfs.lustre --fsname=nas --ost --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp   --reformat --index=2  /dev/rbd3
mkfs.lustre --fsname=nas --ost --mgsnode=10.5.73.130@tcp --mgsnode=10.5.73.131@tcp   --reformat --index=3  /dev/rbd4
```
节点10.5.65.131负责MDT1和OST4-7，节点10.5.65.132负责MDT2和OST8-11，命令格式类似，仅索引和RBD设备号不同。

# q
切换Lustre内核时钟源为tsc需要修改哪些配置项并如何生效？
# a
修改 `/etc/default/grub` 中的 `GRUB_CMDLINE_LINUX` 行，添加 `tsc=reliable tsc=nowatchdog` 参数。然后执行：
```bash
grub2-mkconfig -o /boot/grub2/grub.cfg
```
重启后通过 `cat /sys/devices/system/clocksource/clocksource0/current_clocksource` 确认当前时钟源为 `tsc`。

