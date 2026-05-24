# q
`mdbs_bond4_configure_backend.sh` 脚本有什么功能，需要哪些必选参数？
# a
该脚本用于一次性完成 bond4 模式的配置、IP 配置、网关配置，并默认将 MTU 设置为 9000。所有入参均为必选：
- `-b` bond 虚拟网卡名称
- `-i` IP 地址
- `-f` 第一个实体网卡名称
- `-l` 第二个实体网卡名称
- `-m` 子网掩码（网关）  
示例：`bash mdbs_bond4_configure_backend.sh -b bond4hci -i 172.22.251.104 -f p1p1 -l p1p2 -m 255.255.0.0`

# q
当 `mdbs_bond4_configure_backend.sh` 执行出错时，如何回滚配置？需要哪些参数？
# a
使用回滚脚本 `mdbs_bond4_configure_backend_recover.sh`，只需指定两个实体网卡名称：
- `-f` 第一个实体网卡名称
- `-l` 第二个实体网卡名称  
示例：`bash mdbs_bond4_configure_backend_recover.sh -f p1p1 -l p1p2`

# q
如果只需要配置 bond 聚合接口但不设置 IP 和 MTU，应使用哪个脚本？如何验证实体网卡是否存在？
# a
使用 `mdbs_bond4_configure_only.sh`，参数为：
- `-b` bond 虚拟网卡名称
- `-f` 第一个实体网卡名称
- `-l` 第二个实体网卡名称  
验证网卡存在的命令：`ip addr show dev <网卡名>`。若网卡不存在，脚本会退出。

# q
在已配置 bond4 接口的基础上，后续如何单独配置 IP 和 MTU？各自的回滚脚本如何使用？
# a
- 配置 IP（需在 `mdbs_bond4_configure_only.sh` 基础上）：  
  `mdbs_bond4_configure_ip.sh -b bond4hci -i 172.22.250.101 -m 255.255.0.0`  
  回滚：`mdbs_bond4_configure_ip_recover.sh -b bond4hci`
- 配置 MTU（需在 `mdbs_bond4_configure_only.sh` 或 `_backend.sh` 基础上）：  
  `mdbs_bond4_configure_mtu.sh -b bond4hci -m 9000`  
  回滚：`mdbs_bond4_configure_mtu_recover.sh -b bond4hci`

