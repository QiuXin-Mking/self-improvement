# q
cosbench 读取 S3 数据时出现 “Unable to verify integrity of data download” 错误的典型根因是什么？
# a
典型根因是客户端（cosbench driver）计算的内容哈希值与 Amazon S3 返回的哈希值不一致，导致 MD5 完整性校验失败。这通常发生在 HTTPS 传输或客户端与服务端实现存在差异时，数据本身可能并未损坏。

# q
如何通过修改 cosbench 启动参数解决 “Unable to verify integrity of data download” 错误？
# a
在 cosbench 的 JVM 启动参数中增加 `-Dcom.amazonaws.services.s3.disableGetObjectMD5Validation=true`，关闭 S3 客户端的 MD5 校验。具体流程：

1. 执行 `stop-all.sh` 停止所有 cosbench 进程（controller 和 drivers）。
2. 编辑 `cosbench-start.sh`，在 Java 启动命令行中添加上述 JVM 属性。
3. 执行 `start-all.sh` 重启所有 cosbench 进程。

