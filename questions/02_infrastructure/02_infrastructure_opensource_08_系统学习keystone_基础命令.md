# q
如何为指定用户添加 project 角色？请写出完整的 keystone 命令（假设用户名 wsops，项目名 admin，角色 admin）
# a
```bash
keystone user-role-add --user wsops --tenant admin --role admin
```
注意：需要提前确认项目、用户和角色名称，角色可以是 `admin` 或 `Member` 等。

# q
如何查看当前 keystone 中所有可用的角色名称？
# a
```bash
keystone role-list
```
根据结果选择一个常用角色（如 `admin` 或 `Member`），再为用户赋予对应项目权限。

