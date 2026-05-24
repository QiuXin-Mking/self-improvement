# q
如何使用AWS CLI查询当前区域的所有安全组ID和名称？
# a
```bash
aws ec2 describe-security-groups --query "SecurityGroups[*].[GroupId,GroupName]" --output table
```

# q
通过AWS CLI创建名为`imageappdb`的MySQL RDS实例，使用`db.t3.micro`实例类型，管理员密码为`Admin1234!`，安全组为`sg-0783450f5d3042229`，并启用公网访问，命令是什么？
# a
```bash
aws rds create-db-instance \
  --db-instance-identifier imageappdb \
  --db-instance-class db.t3.micro \
  --engine mysql \
  --master-username admin \
  --master-user-password 'Admin1234!' \
  --allocated-storage 20 \
  --vpc-security-group-ids sg-0783450f5d3042229 \
  --publicly-accessible \
  --backup-retention-period 0
```

# q
如何查看名为`imageappdb`的RDS实例的详细信息（包括终端节点地址和端口）？
# a
```bash
aws rds describe-db-instances --db-instance-identifier imageappdb
```
返回的`Endpoint.Address`即为数据库连接地址，`Endpoint.Port`通常为3306。

# q
新建RDS实例时，为了让同VPC的EC2能够访问数据库，需要满足哪两个关键配置？
# a
1. 子网选择与EC2所在的同一个VPC；
2. 安全组允许来自EC2（或Lambda）的3306端口入站流量。

