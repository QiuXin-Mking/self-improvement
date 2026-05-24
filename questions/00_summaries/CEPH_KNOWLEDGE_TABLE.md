# q
Ceph 的核心组件有哪些？各自的作用、默认端口及关键命令是什么？
# a
Ceph 的核心组件及概要信息如下：

| 组件 | 全称 | 作用 | 端口 | 关键命令 |
|------|------|------|------|----------|
| **MON** | Monitor | 集群状态监控 | 6789 | `ceph mon stat` |
| **OSD** | Object Storage Daemon | 数据存储/恢复 | 6800+ | `ceph osd stat` |
| **RGW** | RADOS Gateway | S3/Swift 接口 | 7480/8080 | `radosgw-admin` |
| **MDS** | Metadata Server | CephFS 元数据 | 6801+ | `ceph mds stat` |
| **MGR** | Manager | 监控指标收集 | 7000+ | `ceph mgr stat` |

# q
在 Ceph OSD 状态矩阵中，UP 和 IN 的各种组合分别代表什么含义？
# a
OSD 的 UP 与 IN 组合的状态含义及处理方案：

| 状态 | UP | IN | 含义 | 处理方案 |
|------|-----|-----|------|----------|
| **正常** | ✓ | ✓ | OSD 正常运行 | - |
| **DOWN** | ✗ | ✓ | OSD 进程停止 | 重启 OSD |
| **OUT** | ✗ | ✗ | OSD 被移除 | `ceph osd in <id>` |
| **UP/DOWN** | ✓ | ✗ | OSD 运行但不在 CRUSH | `ceph osd in <id>` |
| **DNE** | - | - | OSD 不存在 | 重建或清理 |

# q
OSD 核心参数 `osd_op_num_shards` 的作用及其在不同磁盘类型下的调整建议是什么？
# a
`osd_op_num_shards` 是 OSD 的并发分片数，默认值为 5。调整建议：
- HDD：8-16
- SSD：16-32

该参数用于控制并发操作的分片数量，提升 OSD 的处理并行度。

# q
Ceph RGW 支持哪些认证方式？各自对应的配置项或文档参考是什么？
# a
RGW 支持的认证方式如下：

| 认证方式 | 场景 | 配置项 | 文档/参考 |
|----------|------|--------|------|
| AK/SK | S3 API | `rgw_s3_auth_use_keystone` | [021_认证相关] |
| Keystone | OpenStack 集成 | `rgw_keystone_url` | [044_openstack] |
| IAM Policy | 细粒度授权 | `rgw_enable_apis` | [Bucket Policy] |
| STS | 临时凭证 | `rgw_sts_key` | [STS 功能] |

# q
当 Ceph PG 出现 unfound 或 inconsistent 时，使用什么命令进行修复？
# a
使用 `ceph pg repair <pg_id>` 命令进行 PG 修复。

相关辅助命令：
```bash
# 查看 PG 状态
ceph pg stat
ceph pg dump | grep <pg_id>

# 查询 PG 位置
ceph pg map <pg_id>
```

