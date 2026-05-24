# q
Ceph集群中出现“1 object unfound”时，如何快速定位该对象所在的PG？
# a
执行 `ceph health detail` 会详细列出丢失对象所在的池和 PG 编号。也可以对可疑 PG 执行：
```
ceph pg <pgid> list_unfound
```
直接查看该 PG 中所有未找到的对象。

# q
如何查找某个 RBD image 里是否包含特定的 data object（如 2a1098b6bf5cc5）？
# a
遍历池中所有 image（排除 config 前缀），并使用 `rbd info` 输出进行匹配：
```bash
for i in $(rbd ls -p vms | grep -v config); do
  echo $i
  rbd -p vms info $i | grep 2a1098b6bf5cc5
done
```
若命中，则该 image 使用了该对象。

# q
当 Ceph 集群因 unfound object 进入 error 状态且无法对外提供服务时，核心修复命令是什么？
# a
对受影响 PG 执行 revert 操作，让集群回退丢失的对象：
```
ceph pg <pgid> mark_unfound_lost revert
```
例如：
```
ceph pg 4.2e mark_unfound_lost revert
ceph pg 10.3ff mark_unfound_lost revert
```
执行后集群可以恢复正常 I/O 服务。

# q
如果本地有丢失对象的备份文件，如何手动将其重新写入 Ceph 集群？
# a
使用 `rados put` 命令：
```
rados -p <pool-name> put <object-name> <local-file>
```
该命令会将本地文件内容写入指定池中的对象，从而恢复数据。

# q
除了 `mark_unfound_lost revert`，还有哪些命令可用于调整 PG 的 OSD 映射以尝试找回丢失对象？
# a
可以使用 `pg-upmap-items` 重新映射 PG 内的 OSD：
```
ceph osd pg-upmap-items <pgid> <from-osd> <to-osd>
```
例如：
```
ceph osd pg-upmap-items 10.3ff 47 48
```
通过将 PG 映射到其他 OSD，可能访问到原本不可达的对象副本。

