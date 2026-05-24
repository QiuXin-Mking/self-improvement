# q
如何查看Redis服务的运行状态和进程信息？
# a
使用以下命令：
- `rpm -qa | grep redis`：查询已安装的Redis包
- `systemctl -a | grep redis`：查看Redis系统服务状态
- `ps aux | grep redis`：查看Redis进程号与端口号

# q
如何通过命令行连接Redis并操作键值对？
# a
使用 `redis-cli -h 10.1.0.3 -p 6379 -a <密码>` 登录Redis，密码在配置文件 `/etc/redis/redis_6075.conf` 中查看。设置键值对：`set myKey abc`，获取键值对：`get myKey`。

