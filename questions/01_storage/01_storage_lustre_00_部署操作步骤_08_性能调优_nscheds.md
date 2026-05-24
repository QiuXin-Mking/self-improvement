# q
如何查看 Lustre ksocklnd 模块的当前 nscheds 参数值？
# a
使用以下命令之一：
```bash
cat /sys/module/ksocklnd/parameters/nscheds
```
或
```bash
sudo systool -v -m ksocklnd | grep nscheds
```

# q
如何修改 Lustre ksocklnd 模块的 nscheds 参数？
# a
编辑配置文件 `/etc/modprobe.d/ksocklnd.conf`，添加或修改如下行：
```bash
options ksocklnd nscheds=<desired_value>
```
例如 `options ksocklnd nscheds=12`。

# q
修改 ksocklnd 的 nscheds 参数后，如何使其生效？
# a
需要卸载 Lustre 并重新加载模块：
1. 卸载所有 Lustre 挂载点（如 `umount /data/lustre_ost_4` 等）
2. 卸载 Lustre 内核模块：`lustre_rmmod`
3. 确保 ksocklnd 未被加载：`lsmod | grep ksocklnd`
4. 重新挂载 Lustre 服务时模块会自动加载新参数。

