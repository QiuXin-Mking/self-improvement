# q
如何查看Ceph集群当前的CRUSH规则列表？
# a
使用命令 `ceph osd crush rule list`

# q
如何导出和修改Ceph的CRUSH map？
# a
1. 导出二进制map：`ceph osd getcrushmap -o crushmap.bin`
2. 解码为文本：`crushtool -d crushmap.bin -o crushmap.txt`
3. 编辑 `crushmap.txt` 后，重新编译：`crushtool -c crushmap.txt -o crushmap.bin`
4. 应用新map：`ceph osd setcrushmap -i crushmap.bin`

# q
创建Ceph pool时遇到 `ERANGE: total pgs exceeds max` 错误的原因和解决思路是什么？
# a
总PG数计算公式为 `pg_num × size`，需小于 `mon_max_pg_per_osd × num_in_osds`。示例中 `pg_num 2048 size 3` 产生6147个PG，超过配置上限5000，因此需要减小 `pg_num`（如改为1024）或调整 `mon_max_pg_per_osd` 参数。

# q
如何使用RBD为Lustre创建并映射块设备？
# a
1. 创建Ceph pool（例如 `ceph osd pool create lustre_pool 1024 1024 replicated replicated_rule 3`）
2. 创建RBD镜像：`rbd create --size 204800 lustre_pool/lustre_mgt_mdt --image-feature layering`（大小为MB）
3. 映射到本地设备：`sudo rbd map lustre_pool/lustre_mgt_mdt --id admin`
   映射后会在 `/dev` 下产生类似 `/dev/rbd0` 的块设备。

# q
如何删除Ceph pool并取消所有RBD映射？
# a
- 删除pool：`ceph osd pool delete lustre_pool lustre_pool --yes-i-really-really-mean-it`
- 批量取消映射：
```bash
for dev in $(rbd showmapped | awk 'NR>1 {print $5}'); do
    umount $dev
    rbd unmap $dev
done
```

