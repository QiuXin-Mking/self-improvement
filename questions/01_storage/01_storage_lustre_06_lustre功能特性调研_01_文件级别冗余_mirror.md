# q
Lustre中如何为文件添加镜像副本以实现文件级别冗余？
# a
使用 `lfs mirror extend` 命令，例如：
```bash
lfs mirror extend -N -S 8M -c -1 -p archive /path/to/file
```
参数含义：
- `-N`：创建新镜像（不指定镜像ID，自动分配）
- `-S 8M`：设置条带大小为 8MB
- `-c -1`：使用所有可用 OST
- `-p archive`：指定存储池为 archive
此外，也可以使用 `lfs mirror create` 命令以类似方式添加镜像副本，并支持分区域布局（如 `-E 4M -p flash -E eof -N2 -p archive`）。

# q
如何查看文件的镜像布局信息？
# a
使用 `lfs getstripe <file>` 命令。输出中关键字段：
- `lcm_mirror_count`：镜像副本总数
- `lcm_entry_count`：镜像条目数
- 每个镜像条目包含 `lcme_mirror_id`（镜像ID）、`lcme_flags`（如 `init` 表示初始状态）、扩展区间以及各自的条带布局（`lmm_stripe_count`、`lmm_stripe_size`、`lmm_pool` 等）。

# q
如何使用 `lfs find` 查找具有特定镜像数量的文件？
# a
```bash
lfs find --mirror-count +1 --type f /mnt/lustre
```
- `--mirror-count +1`：查找镜像数大于1的文件
- `--mirror-count +2`：镜像数大于2，以此类推
- `--type f`：仅查找普通文件
- 还可结合 `--mirror-state` 过滤镜像状态（如 `--mirror-state=^ro` 排除只读镜像）。

# q
`lfs mirror resync` 命令的作用是什么？
# a
`lfs mirror resync <file>` 用于手动触发镜像副本之间的数据重新同步，确保各个镜像的数据一致。这在镜像副本因故障出现过不一致时用于恢复冗余。

# q
`lfs mirror split` 命令在什么情况下会失败？
# a
当目标文件只有一个镜像组件（`lcm_entry_count: 1`）时执行 `lfs mirror split` 会报错，例如：
```
error lfs mirror split: file '...' has 1 component, cannot split
```
因为 split 需要存在至少两个镜像组件才能将其中一个分离为独立文件。

