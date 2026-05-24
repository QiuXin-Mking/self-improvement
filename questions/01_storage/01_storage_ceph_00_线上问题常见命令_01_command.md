# q
如何查看Ceph集群的整体健康状态、数据分布及监控状态？
# a
使用 `ceph -s` 命令：
```bash
ceph -s
```

# q
如何导出所有Placement Group (PG)的详细信息？
# a
使用 `ceph pg dump` 命令（输出较多，可配合 `grep` 过滤）：
```bash
ceph pg dump
```

# q
如何在不重启OSD的情况下动态调整其调试日志级别？
# a
使用 `ceph tell` 方式注入参数（推荐）：
```bash
ceph tell osd.13 injectargs '--debug_osd 20'
```
或通过 Admin Socket 直接设置：
```bash
ceph daemon /var/run/ceph/ceph-osd.13.asok config set debug_osd 20
```

# q
如何设置Ceph集群的回填满阈值，控制空间占用过高时停止回填？
# a
设置 `backfillfull-ratio`：
```bash
ceph osd set-backfillfull-ratio 0.93
```
查看当前满阈值相关配置：
```bash
ceph osd dump | grep back | grep full | grep ratio
```

# q
如何从大日志文件中按行号范围提取内容片段？
# a
使用 `sed` 按指定行号范围提取：
```bash
sed -n '2077733,2079733p' ceph-osd.0.log-20250826 > ceph-osd.0.log-20250826_1000
```
若日志为 gzip 压缩文件，可结合 `zcat` 提取：
```bash
zcat ceph-osd.6.log-20250904.gz | sed -n '27600,29600p'
```

