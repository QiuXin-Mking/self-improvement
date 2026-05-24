# q
如何将 Ceph OSD 的 reweight 值修改为 1.0？
# a
使用 `ceph osd reweight osd.<ID> 1.0` 命令。例如，将 osd.33 的 reweight 恢复为 1.0：
```
ceph osd reweight osd.33 1.0
```

# q
如何验证 OSD 的 reweight 修改是否生效？
# a
执行 `ceph osd tree` 命令，检查输出中对应 OSD 的 `REWEIGHT` 列，确认值已变为 `1.00000`。
```
[root@ees-0-2 ~]# ceph osd tree
ID   CLASS  WEIGHT     TYPE NAME                  STATUS  REWEIGHT  PRI-AFF
-13          33.53662      host ees-0-5
33    hdd    3.67798          osd.33                 up   1.00000  1.00000
```

