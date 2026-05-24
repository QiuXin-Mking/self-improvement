# q
grubby是什么？它主要解决什么问题？
# a
grubby 是一个用于管理 Linux 内核启动参数和内核版本的命令行实用工具，可以方便地在 GRUB 等引导加载程序的配置文件中添加、删除或修改内核参数，简化内核的指定、信息查看和默认设置等操作，尤其适合需要频繁更新内核或调整启动参数的系统管理员。

# q
如何使用 grubby 列出所有可用的内核？
# a
使用命令 `sudo grubby --list`

# q
如何用 grubby 将特定内核设置为默认启动内核并查看某个内核的详细信息？
# a
- 设置默认内核：`sudo grubby --set-default /boot/vmlinuz-<版本号>`（例如 `sudo grubby --set-default /boot/vmlinuz-5.15.0-76-generic`）
- 查看内核详细信息：`sudo grubby --info=/boot/vmlinuz-<版本号>`（例如 `sudo grubby --info=/boot/vmlinuz-5.15.0-76-generic`）

