# q
`net.core.rmem_default` 和 `net.core.wmem_default` 的作用是什么？
# a
分别设置套接字接收和发送缓冲区的默认大小（字节）。当新套接字创建且未显式指定缓冲区大小时，使用这些值作为默认值。

# q
`net.ipv4.tcp_rmem` 和 `net.ipv4.tcp_wmem` 的三个值分别代表什么？
# a
格式为 `min default max`，分别定义 TCP 缓冲区自动调整的最小值、默认值和最大值。例如 `4096 8738000 104857600` 表示最小 4096 字节，默认约 8.7 MB，最大约 100 MB。

# q
如何查询当前系统的 `net.ipv4.tcp_rmem` 参数？
# a
使用命令 `sysctl net.ipv4.tcp_rmem` 即可查看。

# q
如何让内核网络参数修改立即生效但不持久化？
# a
使用 `sudo sysctl -w <参数名>=<值>`，例如：
```bash
sudo sysctl -w net.core.rmem_default=26214400
```
此修改在重启后失效。

# q
如何使内核网络参数在系统重启后依然生效？
# a
将参数写入 `/etc/sysctl.conf`，例如：
```
net.core.rmem_default = 26214400
net.ipv4.tcp_rmem = 4096 8738000 104857600
```
然后执行 `sudo sysctl -p` 应用配置文件。

