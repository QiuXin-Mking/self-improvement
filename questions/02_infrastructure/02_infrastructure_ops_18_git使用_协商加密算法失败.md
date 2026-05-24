# q
SSH连接报错“no matching key exchange method found. Their offer: diffie-hellman-group14-sha1,diffie-hellman-group1-sha1”的根本原因是什么？
# a
这不是密钥认证或权限问题，而是SSH的密钥交换算法协商阶段失败。服务器只支持旧的`diffie-hellman-group14-sha1`和`diffie-hellman-group1-sha1`算法，现代SSH客户端默认禁用了这些不安全算法，导致协商失败，根本无法进入身份认证阶段。

# q
如何临时启用已禁用的旧算法，使SSH客户端能连接只支持`diffie-hellman-group14-sha1`等旧算法的服务器？
# a
使用`-o`选项追加允许的算法。常见参数组合如下：
```bash
ssh -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 \
    -oHostKeyAlgorithms=+ssh-rsa \
    -oPubkeyAcceptedKeyTypes=+ssh-rsa \
    -p 29418 qiux1@10.3.196.2
```
`+`表示在默认基础上追加，多个算法用逗号分隔。

# q
如何在使用`ssh-copy-id`向老旧服务器复制公钥时，传递相同的算法协商选项？
# a
同样通过`-o`选项传递，命令示例：
```bash
ssh-copy-id -i ~/.ssh/id_rsa.pub \
  -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 \
  -oHostKeyAlgorithms=+ssh-rsa \
  -oPubkeyAcceptedKeyTypes=+ssh-rsa \
  -p 29418 qiux1@10.3.196.2
```

