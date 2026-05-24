# q
如何从源节点导出 keystone 数据库中的 user、project、ec2_credential 和 assignment 表？
# a
使用 `mysqldump` 导出指定表：
```bash
mysqldump -uroot -p keystone user project ec2_credential assignment > partial_keystone.sql
```

# q
如何将导出的 `partial_keystone.sql` 文件传输到目标节点 keystone2？
# a
使用 `scp` 拷贝文件到目标节点的 `/root/` 目录：
```bash
scp partial_keystone.sql root@keystone2:/root/
```

# q
如何在目标节点将 `partial_keystone.sql` 导入到 keystone 数据库？
# a
使用 `mysql` 命令导入 SQL 文件（若存在主键冲突，可提前处理或使用 `INSERT IGNORE`/`REPLACE` 等方式）：
```bash
mysql -uroot -p keystone < /root/partial_keystone.sql
```

# q
导入数据后，如何重启 keystone 服务使更改生效？
# a
执行以下命令重启 OpenStack Keystone 服务：
```bash
systemctl restart openstack-keystone
```

