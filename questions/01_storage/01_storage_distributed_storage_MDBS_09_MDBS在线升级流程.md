# q
MDBS在线升级流程中，如何对生成的RPM包添加数字签名，以及删除签名的命令是什么？
# a
添加签名需在编译服务器上执行，密码为 `macrosan`：
```sh
rpm --addsign /home/qiuxin/master/build/mdbs-1.1.18-Linux.rpm
```
按提示输入密码 `macrosan`。

删除签名使用：
```sh
rpm --delsign /home/qiuxin/master/build/mdbs-1.1.18-Linux.rpm
```

# q
MDBS在线升级的压缩包 `upgrade.tar.gz` 应包含哪些文件？如何生成？
# a
压缩包应包含MDBS安装包和env包，例如：
```sh
tar -czf upgrade.tar.gz mdbs-1.1.18-Linux.rpm mdbs-env-1.0.23-Linux.rpm
```
将两个RPM包打包为一个升级文件。

# q
MDBS升级管理命令 `upgrade mgt` 提供了哪些关键操作？
# a
- 上传升级包：`upgrade mgt upload --file <路径>`
- 检查升级包：`upgrade mgt check`
- 继续升级：`upgrade mgt continue_upgrade`

# q
如何构建MDBS的软件安装包？
# a
在源码目录执行打包脚本：
```sh
./hci_script/hci_make.sh
```

