# q
Ceph 15集群中OSD在backfill时必现ceph_assert失败的典型根因是什么？
# a
社区已知bug（https://tracker.ceph.com/issues/56764），位于`SnapSet::get_clone_bytes(snapid_t) const`中处理克隆快照字节时触发断言失败，导致PG（如4.112）down。该bug于2022年确认，2024年11月分配，至今处于In Progress状态，暂无官方修复方案。

# q
如何从日志定位OSD backfill中由克隆故障引发的断言失败问题？
# a
通过故障OSD（如osd.10）的crash堆栈定位，关键调用路径：
```
SnapSet::get_clone_bytes(snapid_t) const
PrimaryLogPG::add_object_context_to_pg_stat(std::shared_ptr<ObjectContext>, pg_stat_t*)
PrimaryLogPG::recover_backfill(unsigned long, ThreadPool::TPHandle&, bool*)
PrimaryLogPG::start_recovery_ops(unsigned long, ThreadPool::TPHandle&, unsigned long*)
OSD::do_recovery(PG*, unsigned int, unsigned long, ThreadPool::TPHandle&)
ceph::osd::scheduler::PGRecovery::run(OSD*, OSDShard*, boost::intrusive_ptr<PG>&, ThreadPool::TPHandle&)
OSD::ShardedOpWQ::_process(unsigned int, ceph::heartbeat_handle_d*)
ShardedThreadPool::shardedthreadpool_worker(unsigned int)
ShardedThreadPool::WorkThreadSharded::entry()
```
堆栈表明：ShardedOpWQ线程执行recovery任务→backfill阶段调用`add_object_context_to_pg_stat`→`SnapSet::get_clone_bytes`时触发断言，导致osd crash和pg down。

# q
面对该ceph_assert故障，目前的处理建议是什么？
# a
社区尚未发布补丁（状态In Progress），当前无直接代码修复。可采取的临时规避思路：尽量避免在受影响集群上主动触发backfill（如调整recovery/backfill限速、维护窗口外操作），或关注社区补丁进展后升级。对于已触发宕机的OSD，手动拉起OSD会重复触发，需协调业务切换或等待修复。

