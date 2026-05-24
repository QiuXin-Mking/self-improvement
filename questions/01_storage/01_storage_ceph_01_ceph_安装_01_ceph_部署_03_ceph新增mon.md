# q
在 Ceph 集群中新增 Monitor 节点时，如何配置 mon_host 参数？
# a
在 `/etc/ceph/` 目录下的配置文件中添加 `[mon]` 段，使用 `mon_host` 参数指定所有 Monitor 主机名，格式为 `mon_host = mon1,mon2,mon3`，主机名之间用英文逗号分隔，中间不能有空格。

# q
使用 ceph-deploy 新增 Monitor 的正确命令格式是什么？常见的错误写法有哪些？
# a
正确命令格式为 `ceph-deploy mon create <主机名列表>`，主机名以英文逗号分隔且不加空格，例如：
```
ceph-deploy mon create ees24,ees23
```
常见错误写法是在逗号后添加空格，如 `ceph-deploy mon create ees24, ees23`，这会导致创建失败。

# q
新增 Ceph Monitor 节点需要哪几个基本步骤？
# a
1. 进入 `/etc/ceph/` 目录；
2. 在配置文件中设置 `[mon]` 下的 `mon_host` 参数，列出所有 Monitor 主机名；
3. 执行 `ceph-deploy mon create <主机名列表>` 命令完成添加。

