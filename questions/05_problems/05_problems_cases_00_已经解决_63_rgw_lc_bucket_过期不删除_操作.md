# q
如何手动触发RGW生命周期处理以排查bucket过期不删除问题？
# a
使用 `radosgw-admin lc process` 命令手动触发生命周期处理。

# q
在手动触发RGW生命周期处理时如何开启详细调试日志以收集诊断信息？
# a
在命令前设置环境变量 `CEPH_ARGS="--debug-rgw=20"`，并将标准输出和标准错误重定向到日志文件。示例命令：
```bash
CEPH_ARGS="--debug-rgw=20" radosgw-admin lc process > /home/qiuxin/lc_process_$(date +%s).log 2>&1
```

