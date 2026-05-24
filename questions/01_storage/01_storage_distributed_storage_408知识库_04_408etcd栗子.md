# q
在Python项目中如何通过etcd3库连接etcd服务端？
# a
使用 `etcd3.client()` 方法，并指定服务端口。例如：
```python
etcd_cli = etcd3.client(port=23790)
```

# q
如何从etcd中获取指定键的值？
# a
使用客户端实例的 `get` 方法，传入键名。例如：
```python
state, _ = etcd_cli.get(key.FMT_UPGRADE_TASK_STATE)
```
该方法返回一个元组，第一个元素是键的值，第二个元素是元数据。

# q
项目中etcd键的命名格式常如何定义？
# a
通过定义常量来表示键的路径，例如：
```python
FMT_UPGRADE_CONTROLLER = "/upgrade/controller"
```
使用统一的格式常量可以确保键的一致性，避免硬编码。

