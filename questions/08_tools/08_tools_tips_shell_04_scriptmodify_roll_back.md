# q
这个脚本的主要功能是什么？
# a
该脚本用于修改当前节点所有 proxy 的 `proxy_rollback` 参数，将其统一设置为 true 或 false。

# q
如何使用该脚本将所有 proxy 的回滚标志设置为 true？
# a
```bash
modify_roll_back -b 1
```

# q
脚本安装时需要做哪些准备？
# a
1. 将脚本放在 `/usr/bin` 目录下，以便通过 `modify_roll_back` 命令直接调用。
2. 赋予脚本可执行权限：`chmod +x modify_roll_back`。

# q
脚本实际修改的配置文件路径及写入的值是什么？
# a
脚本遍历 `/engine-fs/proxy` 下的每个子目录，将传入的 `Bool_value`（0 或 1）写入每个子目录中的 `proxy_rollback` 文件。路径格式为 `/engine-fs/proxy/*/proxy_rollback`。

