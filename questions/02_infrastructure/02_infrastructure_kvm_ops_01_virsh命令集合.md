# q
如何列出宿主机上所有虚拟机（包括未运行的）？
# a
```bash
virsh list --all
```

# q
如何仅列出虚拟机名称？
# a
```bash
virsh list --name
```

# q
如何查看指定虚拟机的网络接口及MAC地址？
# a
```bash
virsh domiflist <虚拟机名称>
```

# q
如何在宿主机上通过MAC地址查找对应虚拟机的IP？
# a
```bash
arp -an | grep <MAC地址>
```
或者直接查看hosts文件：
```bash
cat /etc/hosts
```

# q
如何使用命令行直接编辑并生效虚拟机配置？
# a
```bash
virsh edit <vmname>
```

# q
如何通过XML文件定义、取消定义和启动虚拟机？
# a
定义虚拟机：
```bash
virsh define vm1.xml
```
取消定义：
```bash
virsh undefine vm1
```
启动虚拟机：
```bash
virsh start vm1
```

