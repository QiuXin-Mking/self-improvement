# q
class_export_put 函数的核心作用是什么？
# a
负责将一个 export 会话对象的引用计数减 1。当引用计数归零时，根据情况直接销毁 export 对象并回收 obd 设备，或将对象放入“僵尸队列”异步销毁。它是 Lustre 内核端引用计数与资源管理的典型实现。

# q
日志 `PUTting export 00000000d094ce03 : new refcount 4` 表示什么？
# a
表示 class_export_put 被调用，export 对象 `00000000d094ce03` 的引用计数从 5 减为 4，即当前仍有 4 个持有者。

# q
export 对象引用计数减为 0 后的可能处理路径是什么？
# a
如果调用者是唯一持有者，则直接彻底销毁 export 对象并回收 obd 设备；否则对象会被放入僵尸队列，由异步机制负责最终销毁。

