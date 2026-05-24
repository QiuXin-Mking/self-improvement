# q
替换Mellanox mlx5内核模块前如何安全卸载原有驱动？
# a
```bash
modprobe -r mlx5_ib
modprobe -r mlx5_core
```
按顺序卸载上层IB模块，再卸载core模块，避免依赖冲突。

# q
替换mlx5_core.ko的具体文件操作和MD5校验方法是什么？
# a
```bash
# 备份原始模块并记录MD5
md5sum /lib/modules/4.18.0-3.2.1.er1.x86_64/kernel/drivers/net/ethernet/mellanox/mlx5/core/mlx5_core.ko > /home/qiuxin/original_md5

# 记录新模块MD5
md5sum /root/mlx5_core.ko > /home/qiuxin/321_md5

# 替换内核模块文件
cp /root/mlx5_core.ko /lib/modules/4.18.0-3.2.1.er1.x86_64/kernel/drivers/net/ethernet/mellanox/mlx5/core/mlx5_core.ko
```
之后通过 `cat /home/qiuxin/original_md5` 和 `cat /home/qiuxin/321_md5` 对比MD5值，确认替换生效且版本正确。

# q
替换模块后如何加载新驱动并恢复网络？
# a
```bash
# 加载模块
modprobe mlx5_core
modprobe mlx5_ib

# 重启网络服务
systemctl restart network

# 重新初始化IP配置
sh /usr/bin/initnet_for_reboot.sh
```
依次加载core和ib模块，重启网络使配置生效，再运行自定义脚本来添加业务IP地址。

# q
如何验证替换后的mlx5内核模块版本是否符合预期（如3.3.1）？
# a
```bash
modinfo mlx5_core | grep 3.3.1
```
使用modinfo查询模块信息，通过grep过滤出版本号，确认是否包含目标版本标识。

# q
在Lustre环境中替换mlx5内核模块的完整操作流程是什么？
# a
1. 卸载现有模块：`modprobe -r mlx5_ib`、`modprobe -r mlx5_core`
2. 备份原始模块并记录MD5，将新模块复制到内核目录
3. 加载新模块：`modprobe mlx5_core`、`modprobe mlx5_ib`
4. 重启网络并初始化IP：`systemctl restart network`，执行`initnet_for_reboot.sh`
5. 验证模块版本：`modinfo mlx5_core | grep 3.3.1`
（额外涉及Lustre调优命令如 `tunefs.lustre ... --dryrun` 用于检查配置，相关节点IP：10.176.104.139-142）

