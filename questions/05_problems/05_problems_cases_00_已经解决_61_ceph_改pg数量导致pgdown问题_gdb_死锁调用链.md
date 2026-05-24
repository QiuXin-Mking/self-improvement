# q
如何使用GDB定位Ceph OSD中线程死锁时锁的持有者？
# a
执行 `p ((pthread_mutex_t*)mutex_addr)->__data.__owner` 可获取互斥锁当前持有线程的LWP ID，再通过 `thread find <LWP>` 切换到该线程，查看其调用栈。例如：
```
(gdb) p ((pthread_mutex_t*)0x562a14951cf8)->__data.__owner
$6 = 107317
```
可知锁 0x562a14951cf8 被 LWP 107317 持有。

# q
Ceph BlueStore 发生死锁时，如何通过GDB调用栈判断形成了循环等待？
# a
检查所有阻塞在 `__lll_lock_wait` → `pthread_mutex_lock` 的线程，分别记录它们等待的锁地址和已持有的锁地址。如果出现如下循环：
- Thread 32 持有 `0x562a14951cf8`，等待 `0x562a15759018`
- Thread 14 持有 `0x562a15759018`，等待 `0x562a14951cf8`

则构成经典死锁。从堆栈进一步确认：
```
Thread 32:
  BlueStore::Onode::put() -> split_cache -> _split_collection
  持有锁: Onode锁 (0x562a14951cf8)
  等待锁: SharedBlob锁 (0x562a15759018)

Thread 14:
  BlueStore::SharedBlob::put() -> ExtentMap::~ExtentMap -> Onode::put()
  持有锁: SharedBlob锁 (0x562a14951cf8)
  等待锁: Onode锁 (0x562a15759018)
```

# q
此PG数量调整导致的pg down问题中，Ceph OSD死锁的典型根因是什么？
# a
根因是在PG分裂(`_split_collection`)过程中，两个操作获取锁的顺序不一致造成死锁：
- 一条路径先获取 Onode 锁，再获取 SharedBlob 锁（`split_cache` → `Onode::put`）
- 另一条路径先获取 SharedBlob 锁，再获取 Onode 锁（`SharedBlob::put` → `ExtentMap` 析构 → `Onode::put`）

当两个操作并发执行时形成循环等待，最终导致 `tp_osd_tp` 线程全部阻塞，PG进入down状态。

# q
解决此类死锁问题的标准流程是什么？
# a
1. 通过 GDB 分析死锁线程，找到争用的锁及其保护的数据结构。
2. 理清锁的获取顺序，确保所有代码路径以相同顺序获取多把锁（例如固定先 Onode 锁后 SharedBlob 锁，或反之）。
3. 修改源码中错误的锁顺序，可能将析构、引用释放等操作延迟到不持锁的上下文中。
4. 编译测试验证，并通过压力测试确保不再出现锁竞争导致的死锁。

