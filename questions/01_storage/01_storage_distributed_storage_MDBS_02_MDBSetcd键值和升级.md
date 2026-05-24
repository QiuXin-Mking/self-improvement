# q
重新部署（执行 mdbs_deploy.sh）与升级部署（mcli 内 upgrade 命令）对用户数据的处理有何区别？
# a
重新部署会丢失用户数据，升级部署不会丢失用户数据。

# q
集群升级准备阶段需要暂停哪些后台任务？
# a
先暂停后台空间均衡任务，再暂停后台 volume 的快照策略任务。

# q
在升级流程中，`/id/ocache` 键值的作用是什么？
# a
用于判断当前版本是否已完成相关步骤：如果 etcd 中存在该键，则跳过后续操作；如果不存在，则需要查询所有 ocache 相关键，取最大值写入 `/id/ocache`。

# q
使用 `etcd_get` 获取键值时的正确写法及返回值类型是什么？举例说明。
# a
`etcd_get` 返回一个元组 `(value, metadata)`，应使用解包方式接收，如：
```python
o, b = etcd_get(key=key.CDIKS_ID_KEY)
```
其中 `o` 是键值（`bytes` 或 `None`），`b` 是元数据对象（`KVMetadata` 或 `None`）。直接赋值给单个变量会导致拿到整个元组。

