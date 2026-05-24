# q
如何通过Ansible批量卸载所有Lustre客户端的/mnt/lustre挂载点？
# a
```bash
ansible lustre -m shell -a 'umount /mnt/lustre'
```

# q
如何手动挂载Lustre文件系统nas_test到本地/mnt/lustre？
# a
```bash
mount -t lustre 192.168.6.172@tcp:192.168.6.174@tcp:192.168.6.175@tcp:/nas_test /mnt/lustre
```

# q
如何通过Ansible批量挂载Lustre文件系统nas_test到所有客户端？
# a
```bash
ansible lustre -m shell -a 'mount -t lustre 192.168.6.172@tcp:192.168.6.174@tcp:192.168.6.175@tcp:/nas_test /mnt/lustre'
```

# q
如何检查Lustre客户端是否成功挂载nas_test文件系统？
# a
```bash
ansible lustre -m shell -a 'df -h | grep nas_test'
```
或手动在客户端执行
```bash
df -h | grep lustre
```

