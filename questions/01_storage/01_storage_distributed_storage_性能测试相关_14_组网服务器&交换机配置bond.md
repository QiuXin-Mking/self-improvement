# q
如何配置Linux服务器的网卡bond（聚合）？需要修改哪些配置文件？
# a
需要修改 `/etc/sysconfig/network-scripts/` 下的配置文件：
1. **修改实体网卡配置文件**（如 `ifcfg-enp175s0f1`、`ifcfg-enp175s0f3`），设置为 slave 模式：
   ```
   TYPE=Ethernet
   BOOTPROTO=none
   NAME=enp175s0f1
   DEVICE=enp175s0f1
   ONBOOT=yes
   MASTER=bond4hci
   SLAVE=yes
   MTU=9000
   ```
   其中 `MASTER` 对应该网卡所属的 bond 虚拟设备名，`SLAVE=yes` 表示作为成员接口。
2. **创建 bond 主配置文件**（如 `ifcfg-bond4hci`）：
   ```
   DEVICE=bond4hci
   NAME=bond4hci
   BONDING_MASTER=yes
   TYPE=Bond
   IPADDR=172.250.101.190
   NETMASK=255.255.0.0
   ONBOOT=yes
   BOOTPROTO=static
   BONDING_OPTS="mode=4 miimon=100 lacp_rate=fast arp_validate=0"
   MTU=9000
   ```
   注意：`BONDING_OPTS` 中的 `mode=4` 表示 LACP（802.3ad）动态聚合模式。

# q
bond 配置中 `BONDING_OPTS="mode=4 miimon=100 lacp_rate=fast arp_validate=0"` 各参数的含义是什么？
# a
- `mode=4`：绑定模式为 IEEE 802.3ad 动态链路聚合（LACP），需要交换机支持 LACP。
- `miimon=100`：链路监控频率，每 100 毫秒检查一次链路状态。
- `lacp_rate=fast`：LACP 的报文发送速率设为快速（每秒发送一次 LACPDU）。
- `arp_validate=0`：关闭 ARP 监控验证（不通过 ARP 检测链路故障）。

# q
交换机 M-LAG（多机箱链路聚合）的配置包括哪些关键步骤？
# a
以两台交换机实现 M-LAG 为例，关键步骤：
1. **配置 peer-link**：使用高速接口（如 40G 口 eth-0-49、50）创建聚合口并设为 trunk，允许所有 VLAN 通过，然后部署为 MLAG 的 peer-link。
   ```bash
   interface range eth-0-49 - 50
   no shutdown
   switchport mode trunk
   switchport trunk allowed vlan all
   static-channel-group 49
   exit
   interface agg 49
   spanning-tree port disable
   exit
   mlag configuration
   peer-link agg 49
   ```
2. **配置三层接口用于 peer-address**：创建一个专用 VLAN（如 4094），配置 /24 的 IP，并在 MLAG 配置中指定对端地址。
   ```bash
   vlan 4094
   interface vlan 4094
   ip address 10.10.0.1/24   # 交换机A
   mlag configuration
   peer-address 10.10.0.2
   ```
   交换机 B 配置对称（地址为 10.10.0.2，peer-address 指 10.10.0.1）。
3. **将业务接口加入 M-LAG 组**：在交换机上配置接入端口（如 eth-0-1、eth-0-2），加入聚合口并指定 `channel-group`，然后对聚合接口启用 `mlag <id>`。
   ```bash
   interface eth-0-1
   switchport access vlan 10
   channel-group 1 mode active
   exit
   interface agg 1
   mlag 1
   ```
   两台交换机上的同一 MLAG ID 形成一组跨设备聚合。

# q
如何验证 Linux bond 是否实现了链路冗余？
# a
基本测试方法：
1. 使用 `ping` 从另一主机持续 ping 该 bond IP，如 `ping 172.22.101.190`。
2. 在 bond 所在服务器上依次执行：
   ```bash
   ifconfig <物理网卡名> down   # 关闭其中一个成员接口
   ```
   观察 ping 是否中断，然后执行 `ifconfig <物理网卡名> up` 恢复。
3. 交替 down/up 不同的成员网卡，若 ping 始终不丢包，则说明 bond 冗余正常。
   也可通过 `ip a` 查看网卡状态，或观察交换机/网卡 LED 灯亮灭来辅助判断链路状态。

