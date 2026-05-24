# q
删除Ceph pool前需要设置什么配置参数？
# a
需要将mon的配置参数`mon_allow_pool_delete`设置为`true`。

# q
如何永久设置Ceph允许删除pool？
# a
使用ceph config命令持久化写入：
```bash
ceph config set mon mon_allow_pool_delete true
```

# q
如何在不重启服务的情况下临时允许删除pool？
# a
通过injectargs向所有mon注入运行时参数：
```bash
ceph tell mon.* injectargs '--mon-allow-pool-delete=true'
```

# q
如何查询当前mon是否已允许删除pool？
# a
```bash
ceph config get mon mon_allow_pool_delete
```
返回`true`则表示已允许。

