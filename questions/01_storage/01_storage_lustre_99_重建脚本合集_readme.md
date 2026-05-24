# q
在 Lustre 中如何持久化修改的参数？
# a
使用 `lctl conf_param` 或 `lctl set_param -P` 命令可以将参数持久化。

# q
Rebuild_ceph_lustre_case1.sh 与 rebuild_ceph_lustre_case2.sh 脚本的区别是什么？
# a
case1.sh 是从零开始（从0到有）建立 Ceph + Lustre，case2.sh 是在已有基础上（从有到有）建立 Ceph + Lustre。

