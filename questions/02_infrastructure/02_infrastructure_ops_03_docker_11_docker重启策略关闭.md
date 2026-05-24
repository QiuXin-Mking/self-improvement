# q
如何使用Docker命令修改容器的重启策略为“不自动重启”？
# a
使用 `docker update --restart=no <container_id_or_name>` 命令，将指定容器的重启策略设置为 `no`（即不会自动重启）。

# q
如何查看Docker容器当前配置的重启策略？
# a
通过 `docker inspect <container_id_or_name> | grep -i restart` 命令从容器详细信息的输出中过滤与重启相关的行，即可查看重启策略及参数。

