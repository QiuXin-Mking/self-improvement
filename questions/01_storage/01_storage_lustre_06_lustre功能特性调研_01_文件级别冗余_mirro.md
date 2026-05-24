# q
Lustre 文件级别镜像（file-level mirroring）的核心作用是什么？
# a
通过 `lfs mirror create` 命令为指定文件创建一个或多个冗余镜像副本。每个镜像副本可以独立配置条带参数（大小、数量）和 OST 池，实现文件级别的数据冗余，从而允许在故障域（如不同 OST 或池）中保护数据。

# q
`lfs mirror create` 命令的常用参数 `-N`、`-S`、`-c`、`-p` 分别表示什么含义？
# a
- `-N`：指定镜像副本的数量（若不指定则默认为 1）。
- `-S`：设置该副本的条带大小（单位：字节，如 `4M`）。
- `-c`：设置该副本的条带计数（即跨几个 OST 条带化）。
- `-p`：指定该副本使用的 OST 池名称（如 `flash`、`archive`）。

# q
使用 `lfs getstripe` 查看镜像文件时，`lcm_mirror_count` 和 `lcme_mirror_id` 字段分别代表什么？
# a
- `lcm_mirror_count`：文件当前拥有的镜像副本总数。
- `lcme_mirror_id`：当前布局条目所属的镜像副本编号（对应创建时的 -N 顺序）。

