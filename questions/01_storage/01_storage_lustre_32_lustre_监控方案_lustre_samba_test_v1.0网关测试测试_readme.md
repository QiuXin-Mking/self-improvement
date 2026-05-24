# q
网关拨测脚本需要实现哪些核心功能？
# a
脚本需接收 `fs_id`、`nas_id` 和容器 IP 列表作为输入（通过 `/etc/lustre_mon/lustre_samba_monitor_data_input.json` 传入），在 `/monitor` 下为每个 IP 创建子目录，按 IP 挂载 Samba 或 NFS 共享点到对应子目录，为每个子目录启动拨测任务，并每分钟执行一次（频率可配置）。拨测指标包括：`open+sync` 文件元数据操作时延、拷贝大文件（默认 4 MB，支持修改）的读写吞吐，以及 mount/unmount 测试。拨测日志写入 `/var/log/nas_monitor/{NAS_ID}/{容器id}_{ip}/`，容器外对应路径为 `/var/log/manila/monitor/{NAS_ID}/{容器id}/`。

# q
如何向监控脚本传递 nas id 和容器 IP 等参数？
# a
通过配置文件 `/etc/lustre_mon/lustre_samba_monitor_data_input.json` 传入，该 JSON 结构包含 `file_system_id` 和 `share dockers` 数组，数组中每个元素包含 `share_docker_id` 和 `ip` 字段。

# q
拨测结果日志的目录结构是如何组织的？
# a
容器内日志路径为 `/var/log/nas_monitor/{NAS_ID}/{容器id}_{ip}/`，其中 `{NAS_ID}` 为传入的 NAS 标识，`{容器id}` 为 share_docker_id，`{ip}` 为容器 IP。容器外对应映射路径为 `/var/log/manila/monitor/{NAS_ID}/{容器id}/`。

