# q
mdbs_bond4_configure_backend.sh 脚本的核心功能是什么，需要哪些必选参数？
# a
该脚本用于一次性完成 bond4 模式网卡的完整配置，包括创建 bond 虚拟网卡、绑定两个物理网卡、配置 IP 地址、子网掩码，并默认将 MTU 设置为 9000。  
所有参数均为必选，具体如下：
- `-b bondname`：bond 虚拟网卡名称
- `-i ip`：IP 地址
- `-f former net card name`：第一个实体网卡名称
- `-l last net card name`：第二个实体网卡名称
- `-m mask`：子网掩码

输入不存在的网卡名称会导致脚本退出。

# q
如何回滚 mdbs_bond4_configure_backend.sh 的配置操作？
# a
执行对应的回滚脚本 `recover_mdbs_bond4_configure_backend.sh`，需提供两个实体网卡名称参数：
```bash
bash recover_mdbs_bond4_configure_backend.sh -f <第一个实体网卡名> -l <第二个实体网卡名>
```
例如：
```bash
bash recover_mdbs_bond4_configure_backend.sh -f p1p1 -l p1p2
```

# q
mdbs_bond4_configure_only.sh 与 mdbs_bond4_configure_backend.sh 的主要区别是什么？
# a
- `mdbs_bond4_configure_only.sh` 仅创建 bond 虚拟网卡并绑定两个物理网卡，不配置 IP、子网掩码和 MTU，参数为 `-b`、`-f`、`-l`。  
- `mdbs_bond4_configure_backend.sh` 则是一次性完成 bond 创建、IP/掩码配置和 MTU（默认 9000）设置，额外需要 `-i` 和 `-m` 参数。

# q
使用分离脚本分步配置一个带 IP 和自定义 MTU 的 bond4 网卡，完整的命令序列是怎样的？
# a
可分三步完成：
```bash
# 1. 创建 bond 并绑定物理网卡
bash mdbs_bond4_configure_only.sh -b bond4hci -f p1p1 -l p1p2

# 2. 配置 IP 和子网掩码
bash mdbs_bond4_configure_ip.sh -b bond4hci -i 172.22.250.101 -n 255.255.0.0

# 3. 设置 MTU（例如 9000 或 1500）
bash mdbs_bond4_configure_mtu.sh -b bond4hci -m 9000
```
如果需要回滚，则执行对应的 `recover_mdbs_bond4_configure_only.sh`、`recover_mdbs_bond4_configure_ip.sh`、`recover_mdbs_bond4_configure_mtu.sh`。

