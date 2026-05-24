# q
回忆下这几种lustre锁的意义

char *ldlm_typename[] = {
    [LDLM_PLAIN] = "PLN",
    [LDLM_EXTENT] = "EXT",
    [LDLM_FLOCK] = "FLK",
    [LDLM_IBITS] = "IBT",
};
# a
`ldlm_typename[]` 映射的就是 LDLM 支持的 4 种锁资源类型，它们决定锁如何判断冲突、如何组织数据结构：

- `LDLM_PLAIN` → `"PLN"`：最基础的互斥/共享锁，常用于普通对象或者控制资源，没有额外的范围或位掩码概念。
- `LDLM_EXTENT` → `"EXT"`：字节范围锁（Byte Range Lock），结合 start/end 区间实现按文件“extent”粒度的读写并发控制。
- `LDLM_FLOCK` → `"FLK"`：POSIX flock/POSIX byte-range flock 兼容锁，带有进程 ID、起止偏移等 flock 语义。
- `LDLM_IBITS` → `"IBT"`：inode 位锁，对 inode 的不同属性位（例如权限、大小、布局）建立独立锁，常用于元数据操作精细化并发。

# q
ldlm 的主要是使用方式是啥
创建
加锁
+-引用
取消锁
跳过
# a
ldlm_interval_alloc() 这段其实只是 LDLM 内部在给“extent 锁”分配配套的区间节点，外层不需要关心细节。你如果只关注接口层“怎么用 LDLM 锁”，可以把各模块职责理解成下面这样：
面向使用者的接口（ldlm_request.c / ldlm_lock.c 提供）
ldlm_lock_create()：创建一把锁对象，传入 namespace、资源 ID（FID/对象标识）、锁类型（PLAIN/EXTENT/…）、请求模式，以及阻塞/完成/glimpse 回调。
ldlm_cli_enqueue() / ldlm_cli_enqueue_local()：向本地或远端 LDLM 发起加锁请求，成功后锁要么进入 granted，要么排队等待。
ldlm_lock_addref() / ldlm_lock_decref()：在持有锁期间增减引用；当引用清零且没有回调在途时，锁会进入 LRU 等待复用，或自动触发取消。
ldlm_cancel_lock() / ldlm_lock_decref_and_cancel()：显式取消锁，常用于关闭句柄、出错回滚。
ldlm_lock_match_with_skip() / ldlm_revalidate_lock_handle()：查找是否已有兼容锁、支持锁重用或恢复阶段的“重匹配”

因此，调用层只需要：
    确定 namespace 和资源 ID；
    选对锁类型和模式；
    调用 ldlm_lock_create + ldlm_cli_enqueue（或已有封装比如 osc_enqueue_lock、mdt_intent_lock）；
    在业务进入/退出时 addref/decref；
    处理必要的 AST 回调和取消即可。剩余的 slab 分配、树结构维护（包括你看到的 ldlm_interval_alloc()）都是 LDLM 内部为保证性能、正确性而封装的实现细节。

# q
解释下这个锁的用法？
/* lock types */
enum ldlm_mode {
	LCK_MINMODE	= 0,
	LCK_EX		= 1,
	LCK_PW		= 2,
	LCK_PR		= 4,
	LCK_CW		= 8,
	LCK_CR		= 16,
	LCK_NL		= 32,
	LCK_GROUP	= 64,
	LCK_COS		= 128,
	LCK_MAXMODE
};
# a
`enum ldlm_mode` 定义的是 LDLM 支持的锁模式，每个模式都是一个比特位，可以按位组合表示“请求/已经授予的访问权限”：

- `LCK_EX` (1)：独占锁。典型写锁，禁止其它任何访问。
    定义：最高级别的排他锁，同一资源同一时间仅允许一个持有者。
- `LCK_PW` (2)：保护写锁 (“Protected Write”)。写者可写，读者需持有相应的读锁才能访问；常用于 extent/data 写 I/O。
    定义：允许写操作，但禁止其他写操作，允许多个读操作。
- `LCK_PR` (4)：保护读锁 (“Protected Read”)。允许多个持有 PR 的读者并发，写者需 PR/PW 的组合做协调。
    定义：允许读操作，禁止写操作，允许多个读操作。
- `LCK_CW` (8)：并发写锁 (“Concurrent Write”)。允许多个写者并发写特定受约束资源（配合服务器端策略，如 OST dirty writeback）。
    定义：允许多个写操作并发执行，但需保证数据最终一致性。
- `LCK_CR` (16)：并发读锁 (“Concurrent Read”)。更宽松的读锁，允许与 CW 等并存，常见于只需弱一致性时。
    定义：允许多个读操作并发执行。
- `LCK_NL` (32)：空锁 (“No Lock”)。仅占位或者表示不需要资源，常用于取消、询问状态。
    定义：无实际锁语义，仅表示对资源的“兴趣”。
- `LCK_GROUP` (64)：组锁。为一组客户端 (group) 保留的特殊模式，多用于 DOM/COS 或者恢复场景，标记成组访问。
    定义：用于组级资源管理，如OST的条带化数据块。
- `LCK_COS` (128)：Commit-on-Sharing 力量的锁模式，用于共享写入时触发强制提交或转化。
    定义：强一致性锁，确保跨节点数据同步。

- `LCK_MINMODE`/`LCK_MAXMODE`：哨兵值，表示合法范围，实际不要使用。

这些模式由 LDLM 判断冲突矩阵决定互斥或兼容关系，调用方在 enqueue 时候指定请求模式，授予后可以在 `l_granted_mode` 里看到最终模式。

# q
lustre 有 空锁 并发读/写锁 保护读/写锁 独占锁 ，这些锁的冲突矩阵是啥
# a
```
      NL  CR  CW  PR  PW  EX
NL    1   1   1   1   1   1
CR    1   1   1   1   1   0
CW    1   1   1   0   0   0
PR    1   1   0   1   0   0
PW    1   1   0   0   0   0
EX    1   0   0   0   0   0
```

# q
lustre 命名空间有哪些
# a
```
enum ldlm_ns_type {
	LDLM_NS_TYPE_UNKNOWN = 0,	/**< invalid type */
	LDLM_NS_TYPE_MDC,		/**< MDC namespace */
	LDLM_NS_TYPE_MDT,		/**< MDT namespace */
	LDLM_NS_TYPE_OSC,		/**< OSC namespace */
	LDLM_NS_TYPE_OST,		/**< OST namespace */
	LDLM_NS_TYPE_MGC,		/**< MGC namespace */
	LDLM_NS_TYPE_MGT,		/**< MGT namespace */
};
```

