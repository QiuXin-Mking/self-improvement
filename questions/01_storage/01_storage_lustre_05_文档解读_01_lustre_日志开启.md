# q
Lustre 支持哪些调试日志类型？
# a
支持的日志类型包括：ioctl、neterror、warning、error、emerg、ha、config、console、lfsck。当前活跃的类型可通过 `lctl get_param debug` 查看。

# q
如何开启最详细的 Lustre 调试日志？
# a
使用命令 `lctl set_param debug=+all` 可将所有调试类型添加到当前掩码中，启用最详细的日志。

