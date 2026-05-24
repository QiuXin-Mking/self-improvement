# q
如何通过HTTP请求获取新的OSD ID？
# a
实例化 `Osdid` 对象（传入 master IP 和认证 token），调用 `get_osdid()` 方法发送 GET 请求到 `/osds/osdid`；服务端从 etcd 读取当前 OID，自增后写回并返回新 OID。

# q
`Osdid` 类是如何定义的？
# a
`Osdid` 继承 `NodeAPIObject`，初始化时调用 `super().__init__(client, "/osds/osdid", token=token)` 指定请求路径；`get_osdid` 方法通过 `self.http_get(url="/osds/osdid")` 发送 GET 请求获取新 OID。

# q
服务端处理 `/osds/osdid` GET 请求时，如何生成新的 OSD ID？
# a
从 etcd 读取 `OID_KEY` 的当前值，转为整数后加 1，再通过 `etcd_put` 更新回 etcd；若更新失败抛出 `ServiceError`，成功则返回新 OID。

