# q
如何检查系统服务是否存在丢日志的情况？
# a
使用以下命令查看 `systemd-journald` 状态，若输出中出现 `Suppressed ... messages from ...` 说明对应服务存在日志被抑制（丢失）的情况。
```bash
systemctl status systemd-journald
```

# q
`systemctl status systemd-journald` 输出中的 `Suppressed` 消息代表什么含义？
# a
表示 `systemd-journald` 因为日志产生速率过高或达到限制，主动抑制（丢弃）了来自指定服务的部分日志消息，即该服务存在丢日志。

