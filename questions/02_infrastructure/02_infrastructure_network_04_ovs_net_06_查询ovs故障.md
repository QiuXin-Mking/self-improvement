# q
如何查询 OVS 中 br-tun 桥的内核网络接口状态？
# a
使用 `ip link show br-tun` 命令。该命令会显示该网桥接口的索引号、MTU、状态标志（如 UP、LOWER_UP）等信息。

# q
如何列出连接到 br-tun 桥的所有 OVS 端口？
# a
使用命令：
```
ovs-vsctl list-ports br-tun
```
输出会列出所有端口名，例如隧道端口 `tun-2d6a034e-b2b5-11e9-8216-6c0b84af2407` 以及 veth 端口 `veth_tun_clu` 等。

# q
如何查看 OVS 中某个隧道接口的对端连接配置（options 字段）？
# a
使用 `ovs-vsctl list interface <接口名> | grep options`。例如：
```
ovs-vsctl list interface tun-2d6a034e-b2b5-11e9-8216-6c0b84af2407 | grep options
```
options 字段通常包含对端 IP、隧道类型等关键信息，用于排查隧道连通性问题。

