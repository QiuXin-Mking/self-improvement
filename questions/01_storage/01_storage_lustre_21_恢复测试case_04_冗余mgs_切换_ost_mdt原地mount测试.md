# q
如何检查Lustre节点与MGS的连接状态？
# a
在所有节点上执行 `lfs check all | grep MGC`，正常输出为 `MGC<NID>@tcp active.`。若出现错误如 `lfs check: error: check 'MGC192.168.6.172@tcp': Cannot send after transport endpoint shutdown (108)`，则表明该节点无法与对应MGS正常通信。

# q
Lustre MGS从节点A迁移到节点B后，原MGS节点执行lfs check报“Cannot send after transport endpoint shutdown”的典型根因是什么？
# a
MGS迁移后，原MGS节点上的MGS服务已停止，但该节点客户端仍使用旧的MGS NID尝试通信，导致连接已关闭，报出 `Cannot send after transport endpoint shutdown (108)`。需在客户端更新MGS配置或重新挂载以指向新MGS节点。

# q
卸载Lustre OST后，客户端lfs check会显示哪些典型错误？
# a
客户端执行 `lfs check all` 会出现针对被卸载OST的错误，例如：
```
lfs check: error: check 'nas_test-OST0007-osc-MDT0001': Resource temporarily unavailable (11)
lfs check: error: check 'nas_test-OST0007-osc-ffff99f77f63e800': Transport endpoint is not connected (107)
```
表示该OST已不可用，出现资源暂时不可用和传输端点未连接错误。

