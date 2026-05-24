# q
rbd map 或 rbd unmap 命令阻塞无返回的典型现象是什么？
# a
执行 ```rbd map <rbd_name>``` 后命令一直阻塞，无法正常返回；同样 ```rbd unmap <rbd_name>``` 也会出现阻塞，无法完成映射或取消映射。

# q
如何定位 rbd map 阻塞问题？
# a
使用 strace 跟踪命令，对比正常环境。执行 ```strace rbd map <rbd_name>```，观察系统调用差异，找出阻塞点。

