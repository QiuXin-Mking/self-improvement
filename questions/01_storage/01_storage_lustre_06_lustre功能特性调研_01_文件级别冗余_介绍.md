# q
Lustre中的文件级冗余（FIR）是什么？它有什么作用？
# a
文件级冗余（FIR）即多副本（mirror）功能，在Lustre 2.11版本后引入，用于缓解硬件故障。它通过为文件创建多个镜像副本，提升数据可靠性和读并行能力，但会在IO响应路径上引入额外开销。

# q
如何创建一个带多镜像和不同故障域的Lustre镜像文件？请给出命令示例。
# a
使用 `lfs mirror create` 命令。例如：
```
lfs mirror create -N -S 4M -c 2 -p flash -N -c -1 -p archive /mnt/testfs/file1
```
该命令创建两个镜像：第一个使用 flash 池，条带大小 4M，条带数 2；第二个使用 archive 池，条带数用 -1 表示使用池中的所有可用 OST。`-N` 后跟对应镜像的布局选项。

# q
在创建Lustre镜像文件前需要如何配置OST Pool？
# a
需要在 MGS 上创建指定的 OST 池并添加 OST。例如：
```
lctl pool_new lustre1.flash       # 创建池，lustre1 为文件系统名，flash 为池名
lctl pool_add flash OST0000       # 将 OST0000 加入池
lctl pool_list lustre1            # 查看池列表
```
然后可在 `lfs mirror create` 中用 `-p <pool>` 使用这些池。

# q
`lfs mirror` 提供了哪些子命令来管理镜像文件？
# a
`lfs mirror` 包含的子命令有：`create`（创建镜像文件）、`extend`（添加镜像）、`split`（分割镜像）、`read`（读取指定镜像内容）、`write`（写入指定镜像）、`copy`（复制镜像到其他镜像）、`resync`（重新同步）、`verify`（校验镜像）、`delete`（删除镜像）等。

