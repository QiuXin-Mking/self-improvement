# q
网关拨测脚本按功能划分为哪些核心模块？
# a
划分为 7 个模块：
1. 参数解析与初始化模块
2. 目录管理模块
3. 挂载管理模块
4. 拨测任务管理模块
5. 具体拨测功能模块
6. 日志与监控模块
7. 主流程控制模块

每个模块由若干具体函数实现（如 `parse_arguments`、`mount_shared_directory`、`probe_task_loop`、`test_open_sync_latency`、`log_probe_result`、`health_check` 等）。

# q
参数解析函数 `parse_arguments` 支持哪些选项，校验逻辑是什么？
# a
支持的选项：
- `--nas-id`：设置 NAS 标识
- `--ips`：逗号分隔的容器 IP 列表，解析为数组
- `--file-size`：测试文件大小，默认 4194304（4 MB）
- `--frequency`：执行频率（秒），默认 60
- `--help`：显示用法

校验逻辑：若 `NAS_ID` 为空或 `CONTAINER_IPS` 数组长度为 0，则调用 `handle_error` 终止脚本。

# q
脚本如何实现多个容器 IP 的并发拨测任务管理？
# a
通过 `start_probe_tasks` 函数：遍历 `CONTAINER_IPS`，为每个 IP 调用 `probe_task_loop` 并将其放在后台执行（`&`），并将子进程 PID 收集到全局数组 `PROBE_PIDS` 中。
`probe_task_loop` 内是一个无限循环，按 `FREQUENCY` 间隔重复执行 `run_all_probes`，直到接收到退出信号。
主函数最后调用 `wait` 等待所有后台 PID 完成，配合 `wait_for_completion` 检测异常退出。

# q
脚本中的错误处理和清理机制是如何设计的？
# a
定义了统一的错误处理函数 `handle_error`：接收错误消息，打印错误信息并返回退出码，然后调用 `cleanup_and_exit`。
`cleanup_and_exit` 负责停止所有拨测任务（`stop_probe_tasks`）、卸载已挂载的共享目录，最后以指定退出码退出脚本。
所有关键失败路径均通过 `handle_error` 触发，保证资源清理一致。

