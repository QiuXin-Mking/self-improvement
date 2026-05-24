# q
MDBS 添加数据节点的基本流程是什么？
# a
1. 确保集群节点能 ping 通新节点的后端网络。
2. 使用 `deploy mgt add` 命令添加管理节点，指定 IP 地址和用户名（如 `--ipaddr 192.168.1.102 --username root`）并输入密码。
3. 使用 `cluster node join --ipaddr 192.168.1.102` 将节点加入集群，成功后会返回主机名、IP 和节点 ID（nid）。

# q
移除 MDBS 集群节点前，必须先执行什么操作？
# a
必须先让节点隔离（isolate）。如果直接使用 `cluster node remove --nid <nid>` 会报错 `node is not isolated`。可以使用 `node offline set_isolate_time --time <时间>` 设置隔离时间，并通过 `cluster node query_isolate` 确认节点已处于隔离状态后，才能成功移除节点。

# q
`deploy mgt add` 命令的必要参数有哪些？
# a
必须指定 `--ipaddr`（目标节点 IP）和 `--username`（登录用户名），例如：
```
deploy mgt add --ipaddr 192.168.1.102 --username root
```
执行后会提示输入对应用户的密码。

# q
如何查询 MDBS 节点的隔离状态？
# a
使用命令 `cluster node query_isolate`。如果节点已隔离，输出会包含该节点的 `nid`、`hostname`、`ipaddr` 和 `version`；未隔离时输出可能为空或只显示命令成功信息。

