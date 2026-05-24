# q
异地mount Lustre OST时，日志报"cannot find UUID by nid"错误的典型根因是什么？
# a
根因是OST在新节点mount后，其本地配置中记录的failover nid（如192.168.6.175@tcp）在MGS上不存在对应的UUID，导致MGC在应用恢复日志（mgc_apply_recover_logs）时找不到该nid对应的节点。通常因为OST原servicenode或failover node与MGS中的注册信息不一致。
```
LustreError: 99898:0:(mgc_request.c:1567:mgc_apply_recover_logs()) mgc: cannot find UUID by nid '192.168.6.175@tcp': rc = -2
Lustre: 99898:0:(mgc_request.c:1788:mgc_process_recover_nodemap_log()) MGC192.168.6.174@tcp: error processing recovery log nas_test-cliir: rc = -2
```

# q
如何从系统日志定位Lustre OST异地mount后MGC连接失败的具体错误？
# a
检查目标节点（mount OST的节点）的kernel日志，搜索"cannot find UUID by nid"或"error processing recovery log"。关键位置在mgc_request.c，日志会显示具体的nid和返回码-2，表明MGS中缺少该nid对应的UUID，导致客户端无法通过MGC获取正确的OST连接信息。

# q
解决Lustre OST异地mount后客户端持续报"Resource temporarily unavailable"的标准流程是什么？
# a
1. 在OST所在节点使用`tunefs.lustre --erase-params`清除旧参数；
2. 重新指定正确的servicenode和所有mgsnode，例如：
```bash
tunefs.lustre --erase-params --servicenode=192.168.6.175@tcp \
  --mgsnode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp --mgsnode=192.168.6.175@tcp   /dev/rbd5
```
3. 重新mount OST；
4. 在客户端验证`lfs check all`恢复正常。

# q
客户端执行`lfs check all`时对某个OST报"Resource temporarily unavailable (11)"及"Input/output error (5)"，这代表什么问题？
# a
这表示客户端侧的OSC（对应OST0007）无法与实际的OST服务建立连接或通信失败。常见原因包括：OST未启动、网络不可达，或者OST在异地mount后其failover NID未在MGS上正确注册，导致客户端仍然请求旧的NID而无法连接。

