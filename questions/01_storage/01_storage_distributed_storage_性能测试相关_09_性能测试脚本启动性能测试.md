# q
如何在前端任意节点激活 Python 虚拟环境？
# a
执行命令：
```shell
source /opt/macrosan/mdbs/py_env/bin/activate
```

# q
在性能测试的 hosts 文件中，vdbench 与 mcloud 的配置有何区别？
# a
vdbench 配置使用真实的节点管理 IP 映射 hostname 到 localhost，例如：
```
172.22.101.189 node-189 localhost #used by vdbench
```
并注释掉 127.0.0.1 映射该 hostname 的条目。  
mcloud 配置则强制将 hostname 映射回 127.0.0.1，例如：
```
127.0.0.1 node-189 localhost #addbyMCloud
```
并注释掉真实 IP 的映射行。两者不能同时生效，需按实际用途二选一。

# q
如何启动 vdbench 性能测试并实时查看输出日志？
# a
使用 nohup 后台启动 vdbench，并指定配置文件和输出目录：
```shell
nohup ./vdbench -f 23_ym_mo_dx_all -o 23_ym_mo_dx_all.tod &
```
随后通过 tail 命令监控日志：
```shell
tail -f nohup.out
```

