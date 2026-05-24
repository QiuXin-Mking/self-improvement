# q
如何批量将Linux网络接口配置文件的ONBOOT设为yes、BOOTPROTO设为no？
# a
使用以下脚本：
```bash
#!/bin/bash
find /etc/sysconfig/network-scripts/ -type f -name "ifcfg-*" -print0 | while IFS= read -r -d $'\0' file; do
    # 修改ONBOOT配置
    sed -i 's/^ONBOOT=.*/ONBOOT=yes/' "$file"
    # 修改BOOTPROTO配置
    sed -i 's/^BOOTPROTO=.*/BOOTPROTO=no/' "$file"
done
```

# q
如何在当前目录下递归查找所有网卡配置文件并打印路径？
# a
```bash
#!/bin/bash
find ./ -type f -name "ifcfg-*" -print0 | while IFS= read -r -d $'\0' file; do
    echo $file
done
```

# q
在find命令中，`-print0`和`read -d $'\0'`组合的作用是什么？
# a
`-print0`将文件名以空字符（null）分隔输出，`read -d $'\0'`以空字符作为记录分隔符读取，这样可以安全处理文件名中包含空格、换行等特殊字符的情况。

