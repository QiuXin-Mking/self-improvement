# q
在 CentOS/RHEL 系统上安装 Samba 服务端需要执行哪些命令？
# a
使用 yum 安装 samba、samba-common 和 samba-client 三个包：
```sh
sudo yum install samba samba-common samba-client
```

# q
smb.conf 配置文件中 [global] 段通常包含哪些核心参数？
# a
[global] 段的核心参数包括：
- `workgroup`: 工作组名称（例如 SAMBA）
- `security`: 安全模式（例如 user）
- `passdb backend`: 密码数据库后端（例如 tdbsam）
此外还可能包含打印相关配置如 `printing`、`printcap name` 等。

# q
在 smb.conf 中如何配置一个允许匿名访问的共享目录？
# a
定义一个共享节（如 `[share]`），并设置以下参数：
```ini
[share]
    path = /srv/samba/share
    browseable = yes
    writable = yes
    guest ok = yes
    force user = nobody
```
其中 `guest ok = yes` 允许匿名访问，`force user = nobody` 强制以 nobody 用户身份操作文件。

# q
testparm 命令在 Samba 配置中的作用是什么？
# a
`testparm` 用于验证 smb.conf 配置文件的语法是否正确，是修改配置后推荐的检测工具。执行后会输出解析后的配置并提示潜在错误。

