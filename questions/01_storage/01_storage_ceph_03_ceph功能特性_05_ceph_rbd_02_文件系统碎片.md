# q
如何通过脚本模拟文件系统碎片？
# a
该脚本通过三种操作产生碎片化：1）创建20个大文件并立即删除，在文件系统中留下空洞；2）创建1000个小文件并多次追加写入使其分散；3）创建600个中型文件并随机删除，进一步制造不连续空间。脚本使用Python，目标目录为 `./fragment_test`。

# q
如何使用rbd命令创建Ceph RBD镜像的快照？
# a
使用命令 `rbd snap create <pool>/<image>`，例如 `rbd snap create rbd/cyy_nas` 为 rbd 池下的 cyy_nas 镜像创建快照。创建前需确保池已存在，如通过 `ceph osd pool create rbd 10` 创建名为 rbd、PG 数量为 10 的池。

# q
如何查看RBD镜像已有的快照列表？
# a
使用命令 `rbd snap ls <pool>/<image>`，例如 `rbd snap ls rbd/cyy_nas` 即可列出 cyy_nas 镜像的所有快照。

