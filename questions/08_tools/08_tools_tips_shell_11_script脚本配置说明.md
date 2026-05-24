# q
如何一键配置bond4模式，绑定物理网卡p1p1和p1p2为bond4hci，设置IP 172.22.251.104、掩码255.255.0.0并采用默认MTU 9000？
# a
执行命令：
```bash
bash mdbs_bond4_configure_backend.sh -b bond4hci -i 172.22.251.104 -f p1p1 -l p1p2 -m 255.255.0.0
```
该脚本会配置bond4、IP、网关并默认设置MTU为9000。回滚使用：
```bash
bash mdbs_bond4_configure_backend_recover.sh -f p1p1 -l p1p2
```

# q
如何分步配置bond4（先创建bond，再配置IP和掩码，最后设置MTU为1500），并说明各步骤的回滚方式？
# a
分步执行：
1. 创建bond（不配IP和MTU）：
   ```bash
   bash mdbs_bond4_configure_only.sh -b bond4hci -f p1p1 -l p1p2
   ```
2. 配置IP和掩码（使用 `-m` 指定掩码）：
   ```bash
   bash mdbs_bond4_configure_ip.sh -b bond4hci -i 172.22.250.101 -m 255.255.0.0
   ```
3. 设置MTU：
   ```bash
   bash mdbs_bond4_configure_mtu.sh -b bond4hci -m 1500
   ```
对应的回滚命令：
- bond创建出错：`bash mdbs_bond4_configure_only_recover.sh -f p1p1 -l p1p2`
- IP配置出错：`bash mdbs_bond4_configure_ip_recover.sh -b bond4hci`
- MTU配置出错：`bash mdbs_bond4_configure_mtu_recover.sh -b bond4hci`

# q
如何将已创建的bond4hci切换为bond模式1（或2），并如何回滚切换？
# a
切换bond模式（例如切换到模式1）：
```bash
bash mdbs_switch_bond_configure.sh -b bond4hci -n 1
```
回滚切换：
```bash
bash mdbs_switch_bond_configure_recover.sh -b bond4hci
```
参数 `-n` 可指定目标模式编号，如1、2等。

# q
在配置bond前，如何检测实体网卡名称是否存在，以避免脚本因不存在的网卡退出？
# a
使用命令：
```bash
ip addr show dev <网卡名称>
```
如果网卡不存在，相应配置脚本会直接退出，因此执行前应先用该命令验证。

