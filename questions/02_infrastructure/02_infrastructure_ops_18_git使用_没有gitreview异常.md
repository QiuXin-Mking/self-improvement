# q
git-review 提示 “No '.gitreview' file found” 是什么原因，如何定位？
# a
表示当前 Git 仓库根目录下缺少 `.gitreview` 配置文件，`git-review` 无法获知 Gerrit 服务器地址和仓库路径。此时需要手工创建该文件，或手动添加名为 `gerrit` 的远程仓库。典型错误信息如下：
```
No '.gitreview' file found in this repository. We don't know where
your gerrit is.
Please manually create a remote named "gerrit" or rename the default
one and try again.
```

# q
git-review 连接 Gerrit 时出现 “no matching key exchange method found. Their offer: diffie-hellman-group14-sha1,diffie-hellman-group1-sha1” 错误意味着什么？
# a
这是 SSH 密钥交换算法协商失败的错误。客户端与 Gerrit 服务器无法就密钥交换方法达成一致，服务器只提供了两种较旧的算法：`diffie-hellman-group14-sha1` 和 `diffie-hellman-group1-sha1`，而客户端默认可能已禁用这些弱算法。错误示例：
```
Unable to negotiate with 10.3.196.2 port 29418: no matching key exchange method found. 
Their offer: diffie-hellman-group14-sha1,diffie-hellman-group1-sha1
```
要解决此问题，通常需要在客户端的 `~/.ssh/config` 中为对应主机添加允许的旧算法，例如 `KexAlgorithms +diffie-hellman-group14-sha1,diffie-hellman-group1-sha1`。

