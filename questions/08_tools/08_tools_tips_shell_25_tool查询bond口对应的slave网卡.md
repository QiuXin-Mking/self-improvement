# q
如何查询 Linux bond 接口 bond4hci2 的从属网卡列表？
# a
可以使用以下命令查看 `/proc/net/bonding/` 下的对应文件：
```bash
cat /proc/net/bonding/bond4hci2 | grep "Slave Interface"
```
该命令会输出该 bond 接口中所有从属接口（Slave Interface）的名称。

