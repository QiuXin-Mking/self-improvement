# q
如何使用 AWS CLI 查询所有安全组的 ID 和名称？
# a
```bash
aws ec2 describe-security-groups --query "SecurityGroups[*].[GroupId,GroupName]" --output table
```

# q
使用 AWS CLI 创建一个 db.t3.micro 规格的 MySQL RDS 实例需要哪些关键参数？请给出完整命令示例。
# a
创建命令示例（关键参数：标识符、实例类、引擎、主用户凭证、存储、安全组、公开访问、备份保留期）：
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
创建 RDS 实例后，如何查看其运行状态与连接端点？
# a
使用以下命令查看实例详情，重点关注 `DBInstanceStatus`、`Endpoint.Address` 和 `Port`：
```bash
aws rds describe-db-instances --db-instance-identifier imageappdb
```
返回示例中状态为 `available`，端点为 `imageappdb.czdhlnsaqrsi.us-east-1.rds.amazonaws.com:3306`。

