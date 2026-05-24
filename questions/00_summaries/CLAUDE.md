# q
这个个人技术知识库主要涵盖哪些技术领域？
# a
主要涵盖 Ceph 分布式存储、对象存储、运维经验、嵌入式开发和问题排查记录，已按 10 域重组：01_storage（存储核心）、02_infrastructure（基础设施）、03_languages（编程语言）、04_linux（Linux）、05_problems（问题案例库）、06_career（职场成长）、07_projects（实战项目）、08_tools（工具技巧）、09_ai（AI相关）、10_billing（计费/接单项目）。

# q
如何快速导航这个个人技术知识库？
# a
先看 README.md 了解全貌，参考 5 个 `_summary_*.md` 领域摘要文件（如 `_summary_storage.md`），使用 CEPH_KNOWLEDGE_TABLE.md、PROBLEM_CASES_TABLE.md 和 COMMANDS_TABLE.md 三个知识表格进行网状索引。

# q
在 Ceph 集群中，常用的运维命令有哪些？
# a
查看集群状态：```ceph -s```；查看 OSD 层级：```ceph osd tree```；管理 noout 标志：```ceph osd set noout```、```ceph osd unset noout```；停止指定 OSD：```systemctl stop ceph-osd@<id>```；动态调整 OSD 日志级别：```ceph daemon osd.<id> config set debug_osd 20```；转储历史操作：```ceph daemon osd.<id> dump_historic_ops```。

# q
问题排查时，如何动态调整 Ceph RGW 的日志级别？
# a
使用命令：```ceph daemon /var/run/ceph/client.rgw.<id>.asok config set debug_rgw 20```

# q
这个知识库中记录的技术栈包括哪些关键组件？
# a
存储系统：Ceph (RGW/RBD/OSD/MON/MDS)、Lustre、BlueStore、RocksDB；语言：C、C++、Python、Go、Shell；工具：Ansible、Docker、Gerrit、Git、VSCode、Claude Code；协议：S3、HTTP、NFS、SMB、TCP/IP；硬件：ESP32、STM32、树莓派、SSD/HDD、bcache、LVM。

