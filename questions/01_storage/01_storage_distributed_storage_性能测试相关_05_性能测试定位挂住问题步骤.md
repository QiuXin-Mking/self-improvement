# q
如何从engine-fs/ocache日志判断集群是否出现IO挂住问题？
# a
进入 `/engine-fs/ocache/` 目录，查看对应模块的 `recv_io_req` 和 `send_io_rsp` 文件：
```sh
cd /engine-fs/ocache/
cat 8/recv_io_req ; cat 8/send_io_rsp
```
如果 `recv io req` 的计数持续增长而 `send io rsp` 计数停滞或两者差距不断增大，说明存在未响应的IO请求，很可能是挂住问题。示例输出中 `osd` 模块 recv 为 27798854，send 为 27797594，两者差值较小，若差值异常扩大需重点关注。

# q
分布式存储挂住问题需要查看哪些关键日志和配置文件？
# a
1. 检查每个节点的核心转储位置：`/core`
2. 查看SPDK及引擎日志：`/var/log/mdbs/` 下的 `spdk.log`、`engine.log`、`python_app.log`
3. 检查engine-fs代理信息：`/engine-fs/proxy/1/proxy_info`
4. 进入 `/engine-fs/ocache/` 目录，执行 `cat <模块ID>/recv_io_req; cat <模块ID>/send_io_rsp` 对比IO请求与响应计数，定位挂住模块

