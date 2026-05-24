# q
快照创建的API接口有哪些？
# a
快照创建可通过以下接口调用：
```python
cluster.client.SnapApi.create
cluster.api.Snap.post
cluster.api.Snap.create
```

# q
卷的proxy master信息在etcd中的存储路径是什么？
# a
卷的proxy master信息存储在etcd的键 `/volumes/%s/proxy/master` 中（`%s` 为卷标识）。

# q
如何从etcd中读取并解析JSON格式的数据？
# a
使用 `etcd_cli.get` 获取键值对（返回 bytes 类型），然后通过以下步骤解析：
```python
# 确保有值，然后解码并转换为字典
ref_ds = etcd_cli.get(key)
ref_obj = json.loads(bytes.decode(ref_ds))
```
注意：`bytes.decode` 必须保证入参有值。

# q
DS（数据存储）查询时使用的etcd键格式是什么？
# a
DS查询的键格式定义为：
```python
FMT_DS_KEY = "/ds/%s"
```
通过该模板生成具体键进行查询。

# q
`etcd_cli.transaction` 的作用是什么？
# a
`etcd_cli.transaction` 用于向 etcd 中执行事务性操作，例如放入一个键值对。具体使用示例需参考相关文档或示例代码。

