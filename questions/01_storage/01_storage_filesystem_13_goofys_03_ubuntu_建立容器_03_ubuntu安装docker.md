# q
在Ubuntu上安装Docker之前，为什么需要安装apt-transport-https等软件包？
# a
默认APT使用HTTP协议获取软件包，但Docker官方仓库通过HTTPS提供，为了保证安全的传输和验证，需要安装`apt-transport-https`、`ca-certificates`等包以启用HTTPS支持。

# q
如何将Docker的官方GPG密钥添加到APT密钥环中？
# a
使用以下命令下载并添加密钥：
```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

# q
在Ubuntu Jammy上添加Docker官方APT仓库的命令是什么？
# a
```bash
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

# q
安装Docker CE（社区版）的命令是什么？
# a
```bash
sudo apt-get install docker-ce
```

# q
安装完成后如何快速验证Docker是否成功安装？
# a
运行`docker --version`命令，若返回Docker版本信息则表示安装成功。

