# q
如何进入 Ceph 编译所需的容器环境？
# a
先通过云桌面跳板登录 172.16.140.70（root/openstack），再 SSH 到 ees24，最后执行 `docker exec -it 2f083d60fb74 bash` 进入编译容器。

# q
编译 Ceph 15.2.10 所需的源代码如何获取？
# a
在 ees24 的 `/home/ceph_rpm` 下创建目录（如 `rpm`），进入后执行：
```bash
git clone ssh://liangwl@10.3.196.2:29418/edgeray_storage/ceph_15.2.10
```

# q
编译 Ceph RPM 包的核心脚本是什么？如何执行？
# a
编译脚本为 `auto_build_rpms_octopus.sh`。先将其拷贝到 `ceph_15.2.10` 的同级目录，然后在容器的 `/mnt/rpm` 路径下设置版本号（例如 `version="15.2.10"`）并执行该脚本。

# q
编译生成的 RPM 包最终存放在哪里？
# a
成功编译后，包位于路径 `ceph-15.2.10.7/rpmbuild/RPMS/` 下。

