# q
在 Ubuntu 上如何下载并解压 frp v0.38.0 的 linux amd64 版本？
# a
使用以下命令：
```
wget https://github.com/fatedier/frp/releases/download/v0.38.0/frp_0.38.0_linux_amd64.tar.gz
tar -zxvf frp_0.38.0_linux_amd64.tar.gz
cd frp_0.38.0_linux_amd64
```

# q
frps 的主要用途是什么？
# a
frps 是 frp 的服务端程序，用于实现内外网穿透（NAT 穿透），帮助外部访问内网服务。

