# q
如何使用ansible命令在远程主机上执行shell命令？
# a
使用 `ansible <主机组或主机> -m shell -a '<命令>'`，例如：
```bash
ansible stg -m shell -a 'systemctl status telegraf'
```
其中 `-m shell` 指定使用shell模块，`-a` 传递要执行的命令。

# q
Ansible默认的主机清单文件路径是什么？
# a
默认路径为 `/etc/ansible/hosts`，可以使用 `cat /etc/ansible/hosts` 查看内容。

# q
如何对特定的主机组（如stg或ceph_osd）执行ansible命令？
# a
直接在命令中指定主机组名称，例如：
```bash
ansible stg -m shell -a 'systemctl status ems'
ansible ceph_osd -m shell 'sar -d -f /var/log/sa/sa29 -s 17:30:00 -e 18:30:00 | grep dev259-0'
```
ansible会根据主机清单文件中的组定义自动定位到对应主机。

