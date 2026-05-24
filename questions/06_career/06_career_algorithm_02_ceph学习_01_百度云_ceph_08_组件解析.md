# q
Ceph集群中的Monitor(ceph-mon)主要负责哪些功能？
# a
监控集群状态，守护进程和客户端验证，通常部署3-5个监视器。

# q
Manager组件在Ceph中的作用是什么？
# a
跟踪运行时的指标和Ceph的状态。

# q
Ceph中引入PG（Placement Groups）逻辑概念的目的是什么？
# a
将数据操作和磁盘操作实现隔离，一个PG包括多个OSD。

# q
CRUSH算法在Ceph中承担什么任务？
# a
将object对象确定到指定PG中，并将PG内的数据分散到不同OSD中。

