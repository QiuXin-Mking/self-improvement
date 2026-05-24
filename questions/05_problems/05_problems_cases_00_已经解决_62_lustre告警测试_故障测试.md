# q
在Lustre环境中，对客户端执行umount后，使用lfs check all检查会出现什么现象？
# a
lfs check all检查会失败，提示客户端未挂载。

# q
进行Lustre故障测试时，通常需要执行哪些基础操作来验证系统的容错能力？
# a
典型操作步骤包括：
1. umount client
2. umount mgs
3. umount mdt
4. umount ost
5. reboot lustre mgt
6. reboot lustre mdt
并可加入rc.local reboot测试。

