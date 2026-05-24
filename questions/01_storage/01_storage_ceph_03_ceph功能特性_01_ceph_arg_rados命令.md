# q
如何在使用 rados 命令时通过 CEPH_ARGS 启用调试输出？
# a
在命令前设置 `CEPH_ARGS` 环境变量，直接在同一行执行 rados 操作，例如：
```bash
CEPH_ARGS="--debug_rados 20 --debug_ms 1" rados -p <pool> stat <object>
```
其中 `<pool>` 替换为实际的存储池名称，`<object>` 替换为对象名称。

# q
CEPH_ARGS 环境变量的作用是什么？
# a
用于在执行 rados 命令时临时添加调试或运行参数（如 `--debug_rados 20 --debug_ms 1`），从而在不修改全局配置的情况下启用 Ceph 的调试日志输出。

