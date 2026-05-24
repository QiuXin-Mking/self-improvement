# q
回忆下这几种lustre锁的意义

char *ldlm_typename[] = {
    [LDLM_PLAIN] = "PLN",
    [LDLM_EXTENT] = "EXT",
    [LDLM_FLOCK] = "FLK",
    [LDLM_IBITS] = "IBT",
};
# a
`ldlm_typename[]` 定义了 LDLM 的 4 种锁资源类型，决定冲突判定与数据组织：
- `LDLM_PLAIN` → PLN：基础互斥/共享锁，无范围或位掩码概念。
- `LDLM_EXTENT` → EXT：字节范围锁，按 start/end 区间控制文件 extent 粒度读写并发。
- `LDLM_FLOCK` → FLK：POSIX flock 语义锁，带进程 ID、起止偏移等。
- `LDLM_IBITS` → IBT：inode 位锁，对 inode 属性位（如权限、大小、布局）独立加锁。

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
`enum ldlm_mode` 定义 LDLM 的位掩码锁模式，每个比特位代表一种访问权限：
- `LCK_EX` (1)：独占写锁，禁止其他任何访问。
- `LCK_PW` (2)：保护写锁，允许多读，禁止其他写。
- `LCK_PR` (4)：保护读锁，允许多读，禁止写。
- `LCK_CW` (8)：并发写锁，允许多个写者并发写，需服务器端保证最终一致性。
- `LCK_CR` (16)：并发读锁，弱一致性读，可与 CW 等共存。
- `LCK_NL` (32)：空锁，仅占位或表示兴趣，无实际锁语义。
- `LCK_GROUP` (64)：组锁，为一组客户端保留，用于 DOM/COS 或恢复场景。
- `LCK_COS` (128)：Commit-on-Sharing 锁，共享写入时强制触发提交或转化。
- `LCK_MINMODE`/`LCK_MAXMODE`：哨兵值，不可实际使用。

# q
lustre 有 空锁 并发读/写锁 保护读/写锁 独占锁 ，这些锁的冲突矩阵是啥
# a
冲突矩阵（1 表示兼容，0 表示冲突）：
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
```c
enum ldlm_ns_type {
	LDLM_NS_TYPE_UNKNOWN = 0,  /* 无效类型 */
	LDLM_NS_TYPE_MDC,          /* MDC 客户端命名空间 */
	LDLM_NS_TYPE_MDT,          /* MDT 服务端命名空间 */
	LDLM_NS_TYPE_OSC,          /* OSC 客户端命名空间 */
	LDLM_NS_TYPE_OST,          /* OST 服务端命名空间 */
	LDLM_NS_TYPE_MGC,          /* MGC 客户端命名空间 */
	LDLM_NS_TYPE_MGT,          /* MGT 服务端命名空间 */
};
```
分别对应元数据客户端/服务端、对象存储客户端/服务端、管理客户端/服务端的锁命名空间。

