# q
替换Mellanox mlx5_core内核模块并恢复服务的标准流程是什么？
# a
1. 卸载依赖模块（按顺序）：`modprobe -r mlx5_ib`、`modprobe -r mlxfw`、`modprobe -r pci_hyperv_intf`、`modprobe -r tls`、`modprobe -r mlx5_core`
2. 替换内核模块文件：`cp /root/mlx5_core.ko /lib/modules/4.18.0-3.2.1.er1.x86_64/kernel/drivers/net/ethernet/mellanox/mlx5/core/mlx5_core.ko`
3. 更新模块依赖：`depmod -a`
4. 重新加载模块：`modprobe mlx5_core`、`modprobe mlx5_ib`、`modprobe mlxfw`、`modprobe pci_hyperv_intf`、`modprobe tls`（若涉及Lustre还需加载lustre）
5. 重启网络：`systemctl restart network`
6. 执行网络初始化脚本：`sh /usr/bin/initnet_for_reboot.sh`
7. 重新挂载Lustre：执行`lustre_mount.sh`并检查`df`状态

# q
如何从日志或系统状态定位mlx5_core模块替换后Lustre挂载超时的问题？
# a
- 观察`df`命令是否返回 `df: ‘/mnt/lustre’: Connection timed out`，表明Lustre客户端无法连接服务器
- 检查模块加载状态：确认`mlx5_core`、`mlx5_ib`、`lustre`等已正确加载
- 使用`lnetctl net show`查看LNet网络配置是否丢失或异常
- 若网络配置缺失，重新添加TCP网络接口：`lnetctl net del --net tcp`（如存在则删除），`lnetctl net add --net tcp --if bond0.5`
- 执行上述步骤后重新挂载Lustre并验证`df`不再超时

# q
卸载mlx5_core模块前必须遵循的依赖卸载顺序是什么？
# a
按依赖关系反向卸载：
```
modprobe -r mlx5_ib
modprobe -r mlxfw
modprobe -r pci_hyperv_intf
modprobe -r tls
modprobe -r mlx5_core
```
此顺序确保先移除依赖模块，最后移除核心模块。

# q
替换内核模块文件后如何让系统立即识别新模块？
# a
拷贝新`.ko`文件到模块目录后，执行`depmod -a`命令更新模块依赖信息，然后使用`modprobe <模块名>`加载新模块，无需重启系统。

