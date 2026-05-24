# q
在爱捷云环境中进行 Ceph 编包时，如何进入用于构建的 Docker 容器？
# a
先通过跳板机 `172.16.140.70` 登录，再 SSH 到 ees 节点，然后执行以下操作：
```bash
docker ps -aq                   # 查询所有容器 ID
docker exec -it 2f083d60fb74 bash   # 进入指定容器
```

# q
Ceph 构建脚本的名称及执行路径是什么？
# a
脚本名称：`auto_build_rpms_octopus.sh`。  
在容器内将脚本从 `./ceph_15.2.10/` 拷贝到上一级目录，然后进入 `/mnt/rpm` 目录，设置版本号（例如 `version="15.2.10"`），再拷贝脚本到上一级目录后执行：
```bash
cp ./ceph_15.2.10/auto_build_rpms_octopus.sh ./
cd /mnt/rpm
version="15.2.10"
cp ../auto_build_rpms_octopus.sh ../
# 执行脚本（原文未给出确切执行命令，通常为 ./auto_build_rpms_octopus.sh）
```

# q
构建完成后，生成的 Ceph RPM 包存放在哪个目录？
# a
```bash
ceph-15.2.10.7/rpmbuild/RPMS/
```

