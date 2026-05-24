# q
chronyd是什么
# a
chronyd是chrony套件的守护进程，用于时间同步。

# q
如何临时添加时钟源并立即同步时间
# a
先停止chronyd服务：
```bash
systemctl stop chronyd
```
然后执行一次性同步命令：
```bash
chronyd -q "server 172.29.0.14 iburst"
```

# q
如何通过配置文件永久添加时钟源
# a
编辑 `/etc/chrony.conf`，在文件中添加一行：
```
server 172.29.0.14 iburst
```
保存后重启服务生效：
```bash
systemctl restart chronyd
```

# q
如何查看当前配置的NTP服务器及其状态
# a
确保chronyd服务正在运行，然后执行：
```bash
chronyc sources
```

