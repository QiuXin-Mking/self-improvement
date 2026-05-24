# q
MDBS集群的主配置文件是什么？其典型结构包含哪些部分？
# a
主配置文件为 `/etc/mdbs/mdbs.yaml`，典型结构包含以下部分：
```yaml
cluster:
  uuid: <cluster-uuid>      # 集群唯一标识
node:
  nid: <node-id>            # 当前节点ID
  backend_ipaddr: <ip>      # 后端IP地址
metasvr:
  ssd:                     # SSD相关配置（示例未展开）
  listen_ipaddr: <ip>      # 元数据服务监听IP
metanodes:                 # 元数据节点列表
  - nid: <id>
    ipaddr: <ip>
datanodes: []              # 数据节点列表（当前为空）
```

# q
在配置MDBS节点间SSH免密登录时，需要对哪些文件做什么操作？
# a
需要在所有相关节点（例如示例中的 104、105、106、64）的 `~/.ssh/authorized_keys` 文件中添加各节点的公钥，以实现集群内部的 SSH 互信。

