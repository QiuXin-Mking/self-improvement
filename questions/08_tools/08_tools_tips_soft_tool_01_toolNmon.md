# q
如何从压缩包安装 nmon 并使其可执行？
# a
```bash
cd ~
mkdir nmon
cd nmon/
tar zxvf nmon16m_helpsystems.tar.gz
mv nmon_x86_64_centos7 nmon
chmod +x nmon
./nmon
```

# q
在 nmon 的实时监控界面中，按什么键可以分别查看 CPU 和内存统计数据？
# a
按 `c` 键查看 CPU 统计数据，按 `m` 键查看内存统计数据。

# q
如何使用 nmon 进行长时间性能监控并自动保存采集数据？写出示例命令及参数含义。
# a
示例命令：
```
./nmon -s10 -c380 -f -m ~/nmon
```
参数含义：
- `-s10`：每隔10秒采集一次数据
- `-c380`：总共采集380次数据
- `-f`：生成文件，文件名格式为“主机名+当前时间.nmon”
- `-m ~/nmon`：指定文件保存目录为 ~/nmon

# q
如何停止已批量运行的 nmon 监控进程？
# a
```bash
ps -ef | grep nmon
kill -9 <nmon进程PID>
```
先通过 `ps -ef | grep nmon` 查找 nmon 进程的 PID，再使用 `kill -9` 强制终止。

