# q
如何加载 LNET self-test 模块？
# a
执行命令 `modprobe lnet_selftest`。

# q
`lst new_session read_write` 命令的作用是什么？重复执行会发生什么？
# a
用于创建一个名为 read_write 的测试会话。如果会话已存在，再次执行会报错：`Reader error: 'new session creation failed' at 0` 和 `Failed to create session: Operation not supported`。

# q
`lst` 工具支持哪些常用命令？
# a
通过 `lst --list-commands` 可列出所有命令，包括：quit, exit, help, version, new_session, end_session, show_session, ping, add_group, del_group, update_group, list_group, stat, show_error, add_batch, run, stop, list_batch, query, add_test 等。

# q
如何运行并结束一个 LNET self-test 的 read session？
# a
创建会话后，使用 `lst run read_write` 运行测试，测试结束后执行 `lst end_session read_write` 关闭会话。

