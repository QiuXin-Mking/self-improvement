# q
解决mon硬限制问题的标准配置调整命令是什么？
# a
执行以下命令将 `mon_max_pg_per_osd` 全局增大到 500：
```
ceph config set global mon_max_pg_per_osd 500
```
然后可用 `ceph daemon osd.xx config get mon_max_pg_per_osd` 验证配置是否生效。

# q
重启mon节点有什么风险？
# a
重启mon会导致服务断流，操作时应避开业务高峰并制定回退预案。

