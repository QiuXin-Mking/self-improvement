# q
MDBS 中追踪消息内存分配的调试套路是什么？
# a
通过声明一个全局指针变量（如 `void *dbg_msg_rsp = NULL`），在消息分配代码中将其赋值为分配得到的消息指针（如 `dbg_msg_rsp = dbg_msg_mdlog`），然后利用 `DEFINE_FUNC_RETURN` 和 `MAKE_FUNC_RETURN` 宏定义桩函数。`DEFINE_FUNC_RETURN(__msg_alloc_tracing, void*, dbg_msg_rsp)` 会定义一个返回全局变量的桩函数，而 `MAKE_FUNC_RETURN(__msg_alloc_tracing, dbg_msg_rsp,1)` 用于桩的触发或设置返回值，从而让外部能捕获到每次消息分配的具体指针。

# q
`DEFINE_FUNC_RETURN` 宏在消息分配追踪中起什么作用？
# a
它定义了一个桩函数（如 `__msg_alloc_tracing`），其返回类型为声明时的类型（示例中为 `void*`），该函数返回对应的全局变量（如 `dbg_msg_rsp`）。这样，无需修改原始函数接口，即可在外部获取最近分配的消息指针。

# q
`MAKE_FUNC_RETURN` 宏的典型用法是什么？
# a
`MAKE_FUNC_RETURN(__msg_alloc_tracing, dbg_msg_rsp, 1)` 通常用于在消息分配完成后，设置桩函数的行为（如强制返回值或标记已追踪），从而激活一次追踪记录。第三个参数 `1` 可能表示一次追踪标志或返回值状态，确保每次分配都能被外部同步获取。

