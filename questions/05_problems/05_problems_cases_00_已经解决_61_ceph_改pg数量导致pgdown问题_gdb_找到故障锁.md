# q
如何使用 GDB 定位 Ceph OSD 中因锁竞争导致的线程阻塞问题？
# a
1. 用 `info threads` 列出所有线程，找到大量卡在 `__lll_lock_wait` 的线程；
2. 执行 `thread <线程编号>` 切换到目标线程，`bt` 查看调用栈，锁定 `__gthread_mutex_lock` 处的 `__mutex` 地址（如 `0x562a14b16b80`）；
3. 用 `p ((pthread_mutex_t*)地址)->__data.__owner` 查出当前持有该锁的 LWP 号；
4. 再通过 `info threads` 找到该 LWP 对应的线程编号，重复以上步骤，即可还原锁依赖链，定位死锁或严重竞争点。

# q
案例中 Ceph OSD 出现多个 `safe_timer` 线程阻塞的典型锁竞争链条是什么？
# a
案例显示一条清晰的锁等待链：  
- **Thread 120** (`safe_timer`) 在 `MgrClient::update_daemon_health` 等待 `__mutex=0x562a14b16b80`，该锁被 LWP 107288 持有；  
- **Thread 61** (LWP 107288, `safe_timer`) 在 `OSD::get_perf_reports` → `PG::lock` 等待 `__mutex=0x562a1c7cf0b0`，该锁被 LWP 107317 持有；  
- **Thread 32** (LWP 107317, `tp_osd_tp`) 在 `BlueStore::Collection::split_cache` → `BlueStore::Onode::put` 等待 `__mutex=0x562a15759018`（递归锁），形成了跨模块的锁依赖竞争，最终导致多个线程阻塞。

