# q
ceph 是应用了 Paxos、Raft 哪一个一致性协议。
# a
Monitor 节点通过 Paxos 协议维护集群的元数据（如 OSD 状态、CRUSH 映射等），确保所有节点对集群状态的视图一致。每个 Monitor 节点都可以接受写请求，通过 Paxos 的多数派投票机制达成共识

Paxos：用于元数据管理（Monitor、MDS）的强一致性场景。
Raft-like 日志复制：用于数据存储层（OSD、PG）的高性能一致性保障。

# q
简述 ceph 写io的流程， 是如何实现三副本一致性的，是写完三副本后返回，还是写完单独副本后返回。
# a
同步写入：Primary 等待所有副本确认（CEPH_OSD_FLAG_ONDISK）后才返回。
确认标志：副本必须返回 CEPH_OSD_FLAG_ONDISK，表示数据已落盘。
等待集合：waiting_for_commit 包含所有副本（包括 Primary），全部清空后才完成。

写时序图如下
```
Client                    Primary OSD              Replica OSD 1        Replica OSD 2
  |                          |                          |                    |
  |--- Write Request ------->|                          |                    |
  |                          |                          |                    |
  |                          |--- Write to Local ------>|                    |
  |                          |--- Send to Replica 1 ------------------------>|
  |                          |--- Send to Replica 2 ------------------------>|
  |                          |                          |                    |
  |                          |<-- Local Commit ---------|                    |
  |                          |<-- Replica 1 Commit --------------------------|
  |                          |<-- Replica 2 Commit --------------------------|
  |                          |                          |                    |
  |                          | (waiting_for_commit.empty() == true)          |
  |<-- ACK (all committed) --|                          |                    |
```

# q
回忆下 ceph 写操作的和回调的流程
# a
1. 写操作流程（PrimaryLogPG::execute_ctx）
```cpp
// src/osd/PrimaryLogPG.cc:4012-4046
// 注册 commit 回调，只有在所有副本都 commit 后才会发送回复
ctx->register_on_commit([m, ctx, this](){
  if (ctx->op && !ctx->sent_reply) {
    MOSDOpReply *reply = ctx->reply;
    reply->add_flags(CEPH_OSD_FLAG_ACK | CEPH_OSD_FLAG_ONDISK);
    osd->send_message_osd_client(reply, m->get_connection());
    ctx->sent_reply = true;
  }
});

// 发送副本写入请求
issue_repop(repop, ctx);
```

2. 副本写入和确认（ReplicatedBackend::submit_transaction）
```cpp
// src/osd/ReplicatedBackend.cc:447-528
void ReplicatedBackend::submit_transaction(...) {
  // 创建 InProgressOp，记录需要等待的副本
  op.waiting_for_commit.insert(
    parent->get_acting_recovery_backfill_shards().begin(),
    parent->get_acting_recovery_backfill_shards().end());
  
  // 发送到所有副本
  issue_op(soid, at_version, tid, ...);
  
  // 本地也写入
  parent->queue_transactions(tls, op.op);
}
```

3. 等待所有副本确认（ReplicatedBackend::do_repop_reply）
```cpp
// src/osd/ReplicatedBackend.cc:558-612
void ReplicatedBackend::do_repop_reply(OpRequestRef op) {
  auto r = op->get_req<MOSDRepOpReply>();
  
  // 如果副本确认已写入磁盘（CEPH_OSD_FLAG_ONDISK）
  if (r->ack_type & CEPH_OSD_FLAG_ONDISK) {
    // 从等待列表中移除该副本
    ip_op.waiting_for_commit.erase(from);
  }
  
  // 关键：只有当所有副本都确认后，才调用 on_commit
  if (ip_op.waiting_for_commit.empty() && ip_op.on_commit) {
    ip_op.on_commit->complete(0);  // 触发 Primary 的回调
    in_progress_ops.erase(iter);
  }
}
```

# q
简单介绍下 rocksDB, RocksDB  相对于 LevelDB的差异有哪些，又是如何改进的
# a
RocksDB 支持多线程压缩（compaction）和多线程 memtable 写入，极大提升了并发写入能力，而 LevelDB 多为单线程操作
RocksDB 针对 LevelDB 的写放大问题做了大量优化，提高了高并发下的写入吞吐，很适合大规模数据写入场景
RocksDB 专为 SSD 优化，能处理更大的数据量和更高的 IOPS，适合大数据与分布式环境
RocksDB 由 Facebook 主导开发，社区活跃，解决了 LevelDB 主动限制写入等实际问题，并持续增加功能

