# q
如何配置一个自定义的远程yum源？
# a
1. 将 `.repo` 文件放入 `/etc/yum.repos.d/` 目录。
2. 文件中需包含以下配置块：
```ini
[nf]
name=nf
baseurl=http://172.22.122.14/centos/
enabled=1
gpgcheck=0
```
3. 建议先将系统默认的 `.repo` 文件备份到子目录（如 `bak`）中，以避免冲突。

# q
`yum list updates` 命令的作用是什么？
# a
列出资源库中所有可以更新的 rpm 包。

# q
如何清除 yum 的缓存？
# a
- `yum clean packages` ：清除暂存中的 rpm 包文件
- `yum clean headers` ：清除暂存中的 rpm 头文件
- `yum clean all` 或 `yum clearn` ：清除所有暂存的包文件和头文件

# q
`yum remove` 和 `yum erase` 在删除包时有什么特点？
# a
`yum remove` 会删除指定的 rpm 包，并且会自动移除依赖该包的其他包（联动删除依赖关系）。例如 `yum remove perl*` 会删除所有以 `perl-` 开头的包及其依赖方。

