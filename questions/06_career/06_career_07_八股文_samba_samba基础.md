# q
查询samba的状态
# a
smbstatus
systemctl status smb.service
netstat -ntlp | grep -E '(:445|:139)'
[root@lustre_back1 smbtest]# ss -tnlp | grep -E '(:445|:139)'
LISTEN     0      50           *:445                      *:*                   users:(("smbd",pid=56925,fd=36))
LISTEN     0      50           *:139                      *:*                   users:(("smbd",pid=56925,fd=37))
LISTEN     0      50        [::]:445                   [::]:*                   users:(("smbd",pid=56925,fd=34))
LISTEN     0      50        [::]:139                   [::]:*                   users:(("smbd",pid=56925,fd=35))

# q
samba 配置文件在什么位置
# a
/etc/samba/smb.conf

