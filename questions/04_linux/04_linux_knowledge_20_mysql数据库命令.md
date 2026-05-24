# q
如何通过docker exec进入MariaDB Galera集群的MySQL客户端？
# a
使用命令：
```bash
docker exec -it mdbs-mariadb-galera mysql -uroot -ppasswd
```
之后可以执行SQL语句，例如切换数据库和查询数据。

# q
在MySQL客户端中，如何查看sts数据库中rseek表的数据？
# a
先切换到sts数据库，然后执行查询：
```bash
use sts;
select * from rseek;
```
也可以使用 `show tables;` 查看所有表，使用 `select count(*) from rseek;` 查看当前记录数。

