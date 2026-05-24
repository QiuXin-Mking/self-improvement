# q
修改Lustre节点IP地址后集群全面挂壁（所有节点无法通信）的典型根因是什么？
# a
LNet的NID配置未随IP变化更新，导致节点间通信仍使用旧IP地址。执行 `lctl list_nids` 可发现显示的NID（如 `172.31.0.26@tcp`）与当前网络接口的实际IP不一致，所有RPC请求无法路由到正确节点。

# q
如何从命令输出快速定位该问题的直接原因？
# a
对比 `lctl list_nids` 显示的NID与 `ip addr show` 输出的接口IP。例如：
```bash
[root@lustre_back1 ~]# lctl list_nids
172.31.0.26@tcp

# 而实际接口配置可能是错误的子网或IP：
2: eth0: ...
    inet 172.31.0.26/25 scope global eth0
```
若两者不一致（如掩码错误或IP已变更），即表明LNet仍在使用旧配置，所有节点因此无法通信。

# q
解决此类问题的标准恢复流程是什么？
# a
1. 将IP配置恢复为与LNet NID完全匹配的地址和掩码，如：
   ```bash
   sudo ip addr add 172.31.0.26/24 dev eth0
   sudo ip addr del 172.31.0.26/25 dev eth0
   ```
2. 重启LNet或重新加载Lustre模块使LNet重新识别网络接口（如 `lctl network unconfigure && lctl network configure` 或重启节点）。
3. 若需永久修改IP，必须同步更新Lustre配置文件中的networks参数（如 `/etc/modprobe.d/lustre.conf` 的 `options lnet networks=tcp0(eth0)`），再重新生成NID。

