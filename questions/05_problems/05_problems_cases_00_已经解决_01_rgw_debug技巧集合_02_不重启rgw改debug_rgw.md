# q
如何在不重启RGW守护进程的情况下动态调整debug_rgw日志级别？
# a
使用 `ceph daemon` 命令与RGW的admin socket通信，直接设置配置项。例如将日志级别调整为20：
```bash
ceph daemon /var/run/ceph/client.rgw.$(hostname).asok config set debug_rgw 20
```
调整后立即生效，无需重启服务。也可以指定完整的asok路径，如：
```bash
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2-cxsj.asok config set debug_rgw 20
```
恢复默认日志级别（0）时，同样方式设置为0：
```bash
ceph daemon /var/run/ceph/client.rgw.ees-gdla-250-2.asok config set debug_rgw 0
```

# q
如何查看RGW当前生效的debug_rgw日志级别？
# a
使用 `config get` 命令通过admin socket查询：
```bash
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-2-cxsj.asok config get debug_rgw
```
返回类似 `20` 的数值，表示当前debug级别。

