# q
如何检查计费服务节点上的监控计量模块(ins_monitor_meter)的状态？
# a
使用命令 `ll /var/lib/ins/ins_monitor_meter` 查看该目录是否存在及其内容。

# q
计费问题中，如何从日志确认“开始获取实例对象”的任务是否被执行？
# a
分别对 es_agw 和 nd_agent 日志进行关键字搜索：
```bash
fgrep "Begin get ins obj" /var/log/ees_manager/es_agw.log
fgrep "Begin get ins obj" /var/log/ees_manager/nd_agent.log
```

# q
如何验证当前节点是否为计费任务的主节点？
# a
在 es_agw 和 nd_agent 日志中搜索 “Not master node”：
```bash
fgrep "Not master node" /var/log/ees_manager/es_agw.log
fgrep "Not master node" /var/log/ees_manager/nd_agent.log
```
如果找到该字符串，则说明当前节点不是主节点。

# q
计费相关服务被中断时，日志中会出现哪些关键警告信息？
# a
日志片段如下，表明 API 网关因外部中断而关闭：
```
[WARNING] Interrupted by "Ctl + C" or "system kill", Api Gateway will go down!
[WARNING] timer service stop...
[INFO ] Scheduler has been shut down
```

