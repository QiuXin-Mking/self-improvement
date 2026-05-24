# q
showmount -e 命令的功能是什么？
# a
用于列出指定 NFS 服务器上可用的共享目录，格式为 `showmount -e <NFS服务器IP或主机名>`，方便客户端挂载和管理员监控共享情况。

# q
管理 NFS 共享时，常用哪些配置命令？
# a
编辑 `/etc/exports` 文件定义导出目录，然后执行 `exportfs -ra` 重新导出所有共享。

