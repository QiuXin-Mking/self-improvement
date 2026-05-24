# q
如何从Ceph RGW日志中批量提取导致HTTP 404错误的上传ID？
# a
使用以下两条命令：
```bash
ansible ceph_osd -m shell -a 'cat /var/log/ceph/radosgw/ceph-client.rgw.$(hostname).log | grep 404' | grep Audit > 404_case
awk '{ for(i=1;i<=NF;i++) { if($i=="\"404\"") { if(i>1) print $(i-1); } } }' 404_case > upload_id
```
第一条命令通过Ansible在所有`ceph_osd`节点上执行，从远程RGW日志中过滤包含`404`和`Audit`的行，保存到本地文件`404_case`；第二条命令用awk扫描每一行，找到字段内容为`"404"`的位置，并输出前一列（即对应的上传ID），结果重定向到`upload_id`文件。

