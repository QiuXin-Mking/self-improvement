# q
如何查询指定用户的配额信息？
# a
使用 `radosgw-admin user info` 命令查看用户详情，输出中包含配额相关字段。示例：
```bash
radosgw-admin user info --uid qx1
```
该命令会显示该用户的 bucket 配额和用户配额（包括启用状态、最大对象数、最大容量等）。

# q
如何为已有用户设置并启用用户级别的配额（限制对象数量）？
# a
使用 `radosgw-admin quota set` 命令，指定用户范围并启用配额。示例（设置用户配额，最多 1000 个对象）：
```bash
radosgw-admin quota set --uid=qx1 --quota-scope=user --enable=true --max-objects=1000
```
若需同时设置对象数和容量限制，可添加 `--max-size` 参数，命令中的 `--max-objects=500` 为笔误，通常根据需求指定一个值即可。

