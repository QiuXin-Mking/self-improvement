# q
为什么执行 `ceph osd crush reweight osd.3 1.0` 会报错 "device 'osd.3' does not appear in the crush map"？
# a
因为 OSD（如 osd.3）当前只存在于 OSD 树中，但未被加入到 CRUSH map 中。CRUSH map 是 Ceph 的数据分布映射表，只有加入其中的 OSD 才能参与数据分布和权重调整。报错 `Error ENOENT: device 'osd.3' does not appear in the crush map` 即表明在 CRUSH map 里找不到该设备。

# q
如何将不在 CRUSH map 中的 OSD 添加到 CRUSH map 中？
# a
使用 `ceph osd crush add` 命令，指定 OSD ID、权重、主机名和根节点。示例：
```
ceph osd crush add 3 1.0 host=cmp_trv-97-2 root=default
```
该命令会将 osd.3 以权重 1.0 添加到 CRUSH map，并归属到主机 `cmp_trv-97-2` 及根节点 `default` 下。添加后 OSD 树中该 OSD 及其主机的权重会相应更新。

# q
如何为已添加到 CRUSH map 的 OSD 设置设备类（device class）？
# a
使用 `ceph osd crush set-device-class` 命令。例如将 osd.4 设置为 SSD 类：
```
ceph osd crush set-device-class ssd osd.4
```
设置后，OSD 树中该 OSD 的 `CLASS` 列会显示为 `ssd`，从而参与对应存储类别的数据分布。

# q
如何将某个 OSD 的权重调整到与其他 OSD 一致（例如 18.23039）？
# a
在 OSD 已存在于 CRUSH map 的前提下，使用 `ceph osd crush reweight` 命令更新其权重。例如：
```
ceph osd crush reweight osd.3 18.23039
```
该命令会将 osd.3 在 CRUSH map 中的权重设置为 18.23039，使其与同主机其他 SSD OSD 的权重保持一致。

