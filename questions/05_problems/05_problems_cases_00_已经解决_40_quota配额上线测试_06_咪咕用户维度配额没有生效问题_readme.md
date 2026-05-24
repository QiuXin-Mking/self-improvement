# q
如何为Ceph RGW用户设置并启用用户维度的配额（对象数和大小）？
# a
1. 设置配额值：  
```  
radosgw-admin quota set --uid=<用户ID> --quota-scope=user --max-objects=<对象数> --max-size=<字节数>  
```  
2. 启用配额：  
```  
radosgw-admin quota enable --uid=<用户ID> --quota-scope=user  
```  
**注意**：必须先 set 后 enable，且 `--quota-scope` 必须指定为 `user` 才能对用户维度生效。

# q
如何检查RGW用户配额是否已生效？
# a
使用以下任一命令查看配额状态：  
- `radosgw-admin user info --uid=<用户ID>`（在输出中查找 quota 相关字段）  
- `radosgw-admin quota get --uid=<用户ID>`（直接显示配额设置）  
示例：  
```
radosgw-admin user info --uid=testuser1
radosgw-admin quota get --uid=testuser1
```

# q
RGW用户配额设置后未生效，最常见的根因是什么？
# a
最常见的原因是只执行了 `quota set` 而未执行 `quota enable`，导致配额未被激活。  
另外，`quota set` 时若 `--quota-scope` 误设为 `bucket` 而非 `user`，配额将只作用于桶级别而非用户级别。排查步骤：  
1. 确认是否已执行 `quota enable`。  
2. 确认 `--quota-scope` 是否为 `user`。  
3. 用 `quota get` 检查当前配额配置。  
4. 若仍无效，开启 RGW debug 日志定位。

# q
如何通过管理套接字动态调整RGW的调试日志级别，以排查配额不生效问题？
# a
使用 ceph tell 或 admin socket 命令：  
```
ceph --admin-daemon /var/run/ceph/client.rgw.<实例名>.asok config set debug_rgw 20
```  
完成后关闭调试：  
```
ceph --admin-daemon /var/run/ceph/client.rgw.<实例名>.asok config set debug_rgw 0
```  
日志中会记录配额检查过程，帮助分析未生效原因。

