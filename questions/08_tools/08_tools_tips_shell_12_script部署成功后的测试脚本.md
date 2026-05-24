# q
该部署后测试脚本的主要功能包括哪些步骤？
# a
脚本执行以下步骤：
1. 通过 SSH 远程复制 CA 证书（从 172.22.251.104 到 172.22.6.65）
2. 在三台主机（104、105、106）上触发 SCSI 磁盘扫描
3. 使用 `ntpdate 172.17.251.101` 同步时间
4. 设置 Python 环境变量并激活虚拟环境 `/opt/macrosan/mdbs/py_env/bin/activate`
5. 使用 `nosetests` 运行测试脚本，并将输出记录到带时间戳的日志，同时在后台跟踪日志尾部

# q
如何在 Shell 脚本中通过 SSH 在远程主机上执行多行命令块？
# a
使用 Here Document 语法，将命令块重定向到 `ssh` 命令，示例：
```bash
ssh root@172.22.251.104 > /dev/null 2>&1 << eeooff
echo "- - -" > /sys/class/scsi_host/host0/scan
eeooff
```
`eeooff` 为自定义结束标记，`> /dev/null 2>&1` 将 ssh 的输出和错误重定向到空设备。

# q
脚本中扫描磁盘上线所用的命令是什么？
# a
```bash
echo "- - -" > /sys/class/scsi_host/host0/scan
```
该命令向 SCSI 主机适配器发送扫描信号，使系统重新发现新添加的 SCSI 设备（如磁盘）。脚本中对 104、105、106 三台主机各执行了一次。

# q
脚本如何将 `nosetests` 的输出保存到带时间戳的日志文件并实时查看？
# a
首先用 `time=date +%Y_%m_%d_%H_%M` 生成时间戳，然后以后台进程方式运行：
```bash
nosetests /home/qiuxin/hci1/hci_test/script/tools/test_tool.py >> /home/qiuxin/hci1/log/mdbs_test_${time}.log 2>&1 &
sleep 2
tail -f /home/qiuxin/hci1/log/mdbs_test_${time}.log
```
`>>` 将标准输出和标准错误追加入日志文件，`tail -f` 实时跟踪日志输出。

