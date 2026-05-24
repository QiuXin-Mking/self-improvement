# q
iscsiadm 执行 discovery 时出现“Connection refused”错误的可能原因有哪些？
# a
原因可能包括：目标 IP 地址没有建立网络接口（网口），或者目标节点没有启动对应的 iSCSI 服务。

# q
iscsiadm 执行 discovery 时出现“No portals found”提示表示什么？
# a
表示 iscsiadm 已成功连接到目标节点，但该节点上没有发现任何 iSCSI 门户（portals），即目标端未配置任何可供发现的 target。

