# q
在 SQL 中为什么使用 `referee_id != 2` 作为条件时，`referee_id` 为 `NULL` 的记录会被排除？
# a
因为在 SQL 中，任何值与 `NULL` 用比较运算符（包括 `!=`、`=`、`<>` 等）比较，结果都不是 `TRUE` 或 `FALSE`，而是 `UNKNOWN`。`WHERE` 子句只返回条件结果为 `TRUE` 的行，所以 `NULL` 记录会被过滤掉。例如：`NULL != 2` 的结果是 `UNKNOWN`，该行不会被返回。

# q
如何查询 `Customer` 表中所有**没有被** `id = 2` 的客户推荐的客户姓名？
# a
需要同时考虑 `referee_id` 不等于 2 以及 `referee_id` 为 `NULL` 的行。可以使用以下查询（任一种）：
```sql
SELECT name FROM Customer
WHERE referee_id != 2 OR referee_id IS NULL;
```
或使用 `IFNULL()` / `COALESCE()` 将 `NULL` 转换为一个不会出现的值：
```sql
SELECT name FROM Customer
WHERE IFNULL(referee_id, 0) != 2;
```

# q
在 SQL 中，当过滤条件涉及可能为 `NULL` 的列时，应该如何处理？
# a
必须显式使用 `IS NULL` 或 `IS NOT NULL` 来处理 `NULL` 值。常见方法：
1. 用 `OR` 将 `IS NULL` 与其他条件结合：`WHERE column != value OR column IS NULL`
2. 使用 `COALESCE(column, 默认值)` 或 `IFNULL(column, 默认值)` 将 `NULL` 替换后比较：`WHERE COALESCE(column, -1) != value`
需要注意的是，不能漏掉 `NULL`，否则符合条件但值为 `NULL` 的行会被错误排除。

