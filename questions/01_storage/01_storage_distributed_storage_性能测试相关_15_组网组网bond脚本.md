# q
配置Bond模式4时，脚本中BONDING_OPTS参数的具体值是什么？各选项的含义是什么？
# a
BONDING_OPTS 的配置为：
```
BONDING_OPTS="mode=4 miimon=100 lacp_rate=fast arp_validate=0"
```
- `mode=4`：指定 Bond 模式为 IEEE 802.3ad 动态链路聚合（LACP）。
- `miimon=100`：每 100 毫秒检测一次链路状态。
- `lacp_rate=fast`：LACP 报文发送速率为快速（每秒发送一次，而非默认的每 30 秒一次）。
- `arp_validate=0`：禁用 ARP 监测验证。

# q
脚本 1.0 和 2.0 在 MTU 配置上有何区别？
# a
1.0 版本为 Bond 接口和实体从网卡配置的 MTU 均为 1500（默认以太网 MTU），2.0 版本则将 MTU 统一配置为 9000，即支持巨型帧，通常用于高性能存储网络以减少传输开销。

# q
在脚本中，配置实体网卡成为 Bond 从设备时需要添加哪些关键参数？
# a
通过 `sed` 追加到实体网卡配置文件的参数如下：
- `TYPE=Ethernet`
- `BOOTPROTO=none`
- `ONBOOT=yes`
- `SLAVE=yes`
- `MTU=1500`（1.0）或 `MTU=9000`（2.0）
- `MASTER=<bondhci>`（指定所属的 Bond 接口名称）
此外，会从原备份文件中保留 `UUID`、`NAME`、`DEVICE` 等标识信息。

# q
脚本在执行修改前如何避免丢失原始网卡配置？
# a
先将原网卡配置文件重命名为 `.bak` 备份，如 `mv ifcfg-${former_card} ifcfg-${former_card}.bak`，然后从备份文件中提取 `UUID`、`NAME`、`DEVICE` 字段写入新配置文件，再追加 Bond 从设备所需的参数，从而保留设备唯一标识并允许随时恢复原配置。

# q
脚本 2.0 支持通过命令行参数进行配置，其用法格式是什么？
# a
使用 `getopts` 解析参数，用法示例：
```
./script -b <bond名称> -i <IP地址> -f <第一张实体网卡> -l <第二张实体网卡> -m <子网掩码>
```
如：
```
./t3 -b bond4hci1 -i 172.22.251.105 -f p1p1 -l p1p2 -m 255.255.0.0
```

