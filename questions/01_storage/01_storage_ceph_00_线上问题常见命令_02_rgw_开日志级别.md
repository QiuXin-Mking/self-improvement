# q
如何查询当前 RGW 实例的 `debug_rgw` 日志级别配置？
# a
使用 `ceph daemon` 命令查看指定 RGW 守护进程的运行配置，过滤 `debug_rgw` 相关项：

```bash
ceph daemon client.rgw.<id> config show | grep debug_rgw
```

示例输出：
```
"debug_rgw": "0/0",
"debug_rgw_sync": "1/5",
```

其中 `<id>` 替换为实际的 RGW 实例标识（如 `ees-scla-6-2`）。

# q
如何通过 `ceph tell` 命令临时动态调整 RGW 日志级别？
# a
使用 `ceph tell` 向指定 RGW 实例注入参数，修改 `debug_rgw` 的车载/日志级别（格式为 `日志级别/内存级别`）：

开启详细日志（级别 20）：
```bash
ceph tell client.rgw.<id> injectargs '--debug_rgw 20/20'
```

关闭日志（恢复默认 0）：
```bash
ceph tell client.rgw.<id> injectargs '--debug_rgw 0/0'
```

例如：
```bash
ceph tell client.rgw.ees-scla-6-2 injectargs '--debug_rgw 20/20'
```

# q
如何使用 `ceph daemon` 通过本地 admin socket 直接设置单个 RGW 的 `debug_rgw` 日志级别？
# a
通过指定 asok 文件路径直接修改运行中的 RGW 配置：

```bash
ceph daemon /var/run/ceph/client.rgw.<hostname>.asok config set debug_rgw <级别>
```

常用级别示例：
- 开启详细日志：`config set debug_rgw 20`
- 降低日志量：`config set debug_rgw 2` 或 `config set debug_rgw 1`

动态获取本机主机名的写法：
```bash
ceph daemon /var/run/ceph/client.rgw.$(hostname).asok config set debug_rgw 20
```

实际示例：
```bash
ceph daemon /var/run/ceph/client.rgw.ees-stxx-176-31-cxsj.asok config set debug_rgw 20
ceph daemon /var/run/ceph/client.rgw.ees-gdla-250-3.asok config set debug_rgw 2
```

# q
如何快速全局调整 MON 或 OSD 的调试日志开关？
# a
使用 `ceph tell` 配合通配符批量设置：

```bash
# 关闭 MON 的认证调试日志
ceph tell mon.* config set debug_auth 0

# 关闭 OSD 调试日志
ceph tell osd.* config set debug_osd 0

# 开启 OSD 详细日志
ceph tell osd.* config set debug_osd 20
```

