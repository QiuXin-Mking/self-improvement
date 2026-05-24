# q
如何使用ceph osd tree命令快速定位状态为down的OSD及其所在主机？
# a
执行`ceph osd tree`命令，查看`UP/DOWN`列。例如在以下输出中，osd.22 标记为 `down`，且位于主机 `stg-24-130` 下：
```
-4  139.61993     host stg-24-130
...
22   11.63499         osd.22         down  1.00000          1.00000
```

# q
在给出的ceph osd tree输出中，哪个OSD处于down状态？它属于哪台物理主机？
# a
处于down状态的OSD是osd.22，它属于主机stg-24-130。

