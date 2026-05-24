# q
如何通过GDB线程堆栈定位Ceph OSD发生死锁导致的PG down问题？
# a
使用`gdb -p <osd_pid>`附加进程，执行`thread apply all bt`。若多个线程卡在`__lll_lock_wait`，且等待的锁涉及PG操作和BlueStore存储操作（如`PG::lock`、`BlueStore::OnodeSpace::lookup`），则表明存在锁竞争或死锁。例如：
```
Thread 4 在 PG::lock(bool) const () at PG.cc:235 等待锁；
Thread 6 在 BlueStore::OnodeSpace::lookup(ghobject_t const&) 等待递归互斥锁。
```
此类阻塞会导致PG无法处理请求而down。

# q
Ceph修改PG数量后PG down，在GDB中看到`PG::lock`和`BlueStore::OnodeSpace::lookup`同时阻塞的典型根因是什么？
# a
根因是PG数量变化引发大量PG并发操作（如分裂、迁移、peering），这些操作同时争用PG级别锁和BlueStore内部的`OnodeSpace`递归互斥锁。当Worker线程分别持有一种锁并等待对方释放时形成死锁，导致处理链中断，PG被标记为down。

