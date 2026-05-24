# q
如何查看集群中所有的 CRUSH 规则列表？
# a
```bash
ceph osd crush rule list
```

# q
如何查看某条 CRUSH 规则的详细配置（如 `ssd_rule`）？
# a
```bash
ceph osd crush rule dump ssd_rule
```

# q
如何创建一条针对 SSD 设备的复制型 CRUSH 规则？
# a
首先确保 SSD OSD 已被标记为 `ssd` 设备类：
```bash
ceph osd crush set-device-class ssd <osd-id>
```
然后创建规则：
```bash
ceph osd crush rule create-replicated ssd_rule1 default host ssd
```
该规则表示从 `default` 根节点开始，挑选位于不同主机且设备类为 `ssd` 的 OSD 来存放副本。

# q
如何在 Ceph 中导出和反编译 CRUSH map 以便人工查看或编辑？
# a
```bash
# 导出二进制 CRUSH map
ceph osd getcrushmap -o crushmap.bin
# 反编译为文本
crushtool -d crushmap.bin -o crushmap.txt
```

# q
在 CRUSH 规则中，`step chooseleaf firstn 0 type host` 的含义是什么？
# a
该步骤表示选择叶子节点（OSD）来放置数据，选择粒度在主机（`host`）级别，确保副本分布到不同的主机上。`firstn 0` 中的 `0` 是一个占位符，运行时会被替换为存储池的实际副本数。

