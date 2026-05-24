# q
如何用 iscsiadm 命令发现后端 iSCSI target？
# a
使用发现模式，指定 SendTargets 方式：

```
iscsiadm -m discovery -t st -p 172.250.80.80:3260
```

参数：
- `-m discovery`：进入发现模式
- `-t st`：使用 SendTargets 发现方式
- `-p`：指定目标 IP 和端口（默认 3260）

# q
iscsiadm 发现命令中的 `-t st` 含义是什么？
# a
`st` 表示 **SendTargets**，一种 iSCSI 发现机制，客户端向指定 IP 发送请求，目标返回可用的 Target IQN 列表。

# q
如何用 iscsiadm 登录一个已发现的 iSCSI target？
# a
使用节点模式并指定 IQN 登录：

```
iscsiadm -m node -T iqn.2010-05.com.macrosan.target:rout2:692bb5d528284af982f6e72c33af5369:1478 -p 172.250.80.81:3260 -l
```

通用命令：

```
iscsiadm -m node -T <iqn> -p <ip>:3260 -l
```

参数：
- `-m node`：节点模式，操作已发现的节点记录
- `-T`：目标 IQN
- `-p`：目标 IP 和端口
- `-l`：执行登录

