# q
如何安装Samba？
# a
使用yum命令安装：`yum install samba`

# q
Samba的主配置文件路径是什么？
# a
`/etc/samba/smb.conf`

# q
在smb.conf中定义一个共享目录需要配置哪些关键参数？
# a
示例配置：
```bash
[shared]
   path = /srv/samba/shared
   browsable = yes
   writable = yes
   guest ok = no
   read only = no
   create mask = 0755
```

# q
如何为Samba添加一个新的访问用户？
# a
使用命令：`sudo smbpasswd -a yourusername`

# q
启动Samba服务并设置开机自启的命令是什么？
# a
启动服务：
```bash
sudo systemctl start smbd
sudo systemctl start nmbd
```
启用开机自启：
```bash
sudo systemctl enable smbd
sudo systemctl enable nmbd
```

