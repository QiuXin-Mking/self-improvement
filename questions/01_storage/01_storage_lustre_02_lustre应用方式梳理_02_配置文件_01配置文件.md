# q
/etc/modprobe.d/ 是什么
# a
它是 Linux 系统中用于配置内核模块加载行为的目录，包含通过 modprobe 工具读取的配置文件，可用于动态加载、卸载内核模块并设置模块参数。

# q
如何在 /etc/modprobe.d/ 中阻止某个内核模块自动加载
# a
使用黑名单配置：  
```
blacklist module_name
```

# q
如何在 /etc/modprobe.d/ 中为内核模块设置加载选项
# a
使用 options 指令：
```
options module_name option_name=value
```

# q
Lustre 中 OSD 级别的写缓存配置参数有哪些
# a
在 `/sys/fs/lustre/osd-ldiskfs/` 下，每个 OST 或 MDT 拥有 `read_cache_enable` 和 `writethrough_cache_enable` 参数，分别用于控制读缓存和透写缓存行为。

# q
Lustre 客户端配置 `options lnet networks=tcp(bond0.5)` 的含义是什么
# a
这是在 `/etc/modprobe.d/lustre.conf` 中为 LNet 内核模块设置的选项，表示 LNet 使用 TCP 网络类型并绑定在 bond0.5 网络接口上。

