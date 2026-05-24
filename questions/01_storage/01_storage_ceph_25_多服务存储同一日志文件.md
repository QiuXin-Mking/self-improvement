# q
如何查找 Ceph radosgw 进程监听的端口号？
# a
先用 `ps aux | grep radosgw` 获取进程 PID（例如 60215），再执行 `netstat -nltp | grep 60215` 查看该进程对应的监听端口。

# q
如何快速清空 Ceph RGW 的日志文件？
# a
使用 `echo > /var/log/ceph/ceph-rgw.log` 即可将该日志文件内容清空。

# q
如何测试对象存储服务的域名和端口是否可达？
# a
使用 `curl -v www.qiuxin.com:5056` 发起请求，通过详细信息判断服务是否可连通。

