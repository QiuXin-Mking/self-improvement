# q
如何发现 iSCSI 存储目标？
# a
使用 `sendtargets` 模式执行目标发现，指定门户 IP 和端口（默认 3260）：
```bash
iscsiadm -m discovery -t sendtargets -p <IP>[:<Port>]

# 示例
iscsiadm -m discovery -t st -p 192.168.3.109
iscsiadm -m discovery -t st -p 172.22.251.60:3260
```
发现并同时登录可加 `-l` 参数：
```bash
iscsiadm -m discovery -t st -p <IP> -l
```

# q
如何登录指定的 iSCSI Target？
# a
使用节点模式（`-m node`），指定目标 IQN（`-T`）和门户（`-p`），并执行登录（`-l`）：
```bash
iscsiadm -m node -T <Target_IQN> -p <IP>:<Port> -l

# 示例
iscsiadm -m node -T iqn.2010-05.com.macrosan.target:fw1:83985bcd2b574f20a73a59c8a2e80754:7320 -p 172.22.251.60:3260 -l
```
登录所有可用路径可使用：
```bash
iscsiadm -m node -L all
```

# q
如何断开所有已登录的 iSCSI 会话？
# a
使用节点模式下的 `--logoutall` 参数：
```bash
iscsiadm -m node -U all
```
如需断开单个目标，则使用 `-u` 并指定 IQN 和门户：
```bash
iscsiadm -m node -T <Target_IQN> -p <IP> -u
```

# q
如何永久删除本地的 iSCSI 节点记录？
# a
使用节点模式（`-m node`），配合 `-o delete` 操作。删除指定目标记录：
```bash
iscsiadm -m node -T <Target_IQN> -o delete
```
删除所有已记录的 Target 节点信息：
```bash
iscsiadm -m node -o delete all
```
添加 `-p <IP>` 可指定门户进一步限定删除范围。

