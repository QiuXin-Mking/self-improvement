# q
如何生成用于Gerrit的SSH密钥对？
# a
使用命令 `ssh-keygen -t rsa -b 4096 -C "your_email@example.com"`，这会生成私钥 `~/.ssh/id_rsa` 和公钥 `~/.ssh/id_rsa.pub`。

# q
SSH密钥对生成后，私钥和公钥分别存放在哪里？
# a
私钥存放在 `~/.ssh/id_rsa`，公钥存放在 `~/.ssh/id_rsa.pub`。

# q
将SSH公钥添加到Gerrit后，可以实现什么功能？
# a
可以实现免密认证，通过SSH协议与Gerrit服务器安全通信。添加方法是将 `~/.ssh/id_rsa.pub` 的内容复制到Gerrit账户的SSH密钥设置中。

