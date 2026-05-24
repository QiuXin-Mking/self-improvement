# q
yumdownloader 命令的作用是什么？
# a
yumdownloader 用于从已配置的 yum 源直接下载指定软件包的 RPM 文件到当前目录，不执行安装操作。示例：
```bash
yumdownloader libyaml
```

# q
如何利用 Python 快速启动一个 HTTP 文件服务器并在后台运行？
# a
使用以下命令可在后台启动监听 8080 端口的 HTTP 服务器，提供当前目录的文件列表，且所有输出被丢弃，终端关闭后进程继续运行：
```bash
nohup python -m SimpleHTTPServer 8080 &>/dev/null &
```

