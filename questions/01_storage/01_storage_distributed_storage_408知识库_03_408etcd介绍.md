# q
etcd是什么？
# a
etcd是一个分布式键值对存储系统，由CoreOS开发，内部采用Raft协议作为一致性算法，用于可靠、快速地保存关键数据并提供访问。它常用于配置管理、服务发现和分布式协调，支持分布式锁、leader选举等。etcd默认使用2379端口提供HTTP API服务，使用2380端口与peer通信，生产环境推荐奇数节点集群部署。

# q
etcd的核心工作原理是什么？
# a
etcd架构主要分为四个部分：HTTP Server处理用户API请求及节点间同步与心跳；Store负责各类功能的事务处理，如数据索引、状态变更、事件处理等；Raft是强一致性算法的具体实现，属于etcd核心；WAL（预写式日志）用于持久化存储数据，所有数据提交前都会记录日志，Snapshot用于防止数据过多。用户请求经HTTP Server转发给Store，若涉及节点修改则交给Raft进行状态变更和日志记录，再同步到其他etcd节点确认数据提交，最后完成提交并再次同步。

# q
etcd中Raft相关的核心概念有哪些？
# a
主要包括：Raft（保证分布式系统强一致性的算法），Node（一个Raft状态机实例），Leader（通过竞选产生、处理所有数据提交的节点），Follower（竞选失败的从属节点），Candidate（Follower超时未收到Leader心跳时转变的候选者），Term（某个节点成为Leader到下一次竞选的时间周期），Index（数据项编号，与Term共同用于定位数据）。此外还有WAL（预写式日志）、Snapshot（快照）等持久化机制。

# q
etcd常见应用场景有哪些？
# a
服务发现（Service Discovery）、消息发布与订阅、负载均衡、分布式通知与协调、分布式锁与分布式队列、集群监控与Leader竞选。

# q
本项目代码中如何通过etcd获取指定nid的IP地址？
# a
通过etcd_get读取对应key的值，将字节解码并解析为JSON，若存在"ipaddr"字段则返回该IP，否则返回空字符串。示例代码：
```python
def get_ip_by_nid(nid):
    try:
        val, _ = etcd_get(key=key.node_key(nid))
        o = json.loads(bytes.decode(val))
        if "ipaddr" in o:
            return o["ipaddr"]
    except Exception as e:
        pass
    return
```

