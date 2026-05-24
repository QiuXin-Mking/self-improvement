# q
如何使用 mount 命令挂载 CIFS（SMB）网络共享？
# a
使用 `-t cifs` 指定文件系统类型，并通过 `-o` 传递挂载选项，基本语法：
```bash
mount -t cifs //服务器IP/共享名 本地挂载点 -o username=用户名,pass=密码
```
示例：
```bash
mount -t cifs //172.17.8.69/001_gitlab /home/qiuxin -o username=John,pass=5i6106100258
```

# q
mount 挂载 CIFS 时，file_mode 和 dir_mode 参数的作用是什么？如何设置？
# a
`file_mode` 和 `dir_mode` 用于限制挂载后文件和目录的访问权限（类似 chmod 的八进制模式），防止文件被意外赋予可执行权限。示例：
```bash
mount -t cifs //172.17.8.69/001_gitlab /home/qiuxin -o file_mode=0644,dir_mode=0755,username=John,pass=5i6106100258
```
其中 `file_mode=0644` 表示文件权限为 rw-r--r--，`dir_mode=0755` 表示目录权限为 rwxr-xr-x。

# q
如何查看已经挂载的文件系统，并卸载一个 CIFS 挂载点？
# a
查看已挂载的文件系统可用：
```bash
df -h
```
卸载挂载点使用 `umount` 命令，后跟挂载目录名（或设备名）：
```bash
umount /home/qiuxin
```
或简写为挂载点名（如笔记中的 `umount qiuxin`，通常为目录路径）。

