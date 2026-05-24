# q
SQL查询哪些学生选了3门以上的课程，正确的SQL语句是什么？
# a
正确答案：
```sql
select name,count(*) as count from course group by name having count > 2 order by count desc;
```
解析：
- `GROUP BY name`：按学生姓名分组。
- `COUNT(*)`：统计每个学生的选课数。
- `HAVING count > 2`：筛选选课数大于2的学生（即至少3门）。
- `ORDER BY count DESC`：按选课数降序排列。
- 选项A缺少 `GROUP BY`，选项B在 `GROUP BY` 中错误使用条件，选项C用 `WHERE` 过滤聚合函数，故D正确。

# q
SQL中 `GROUP BY` 的作用和使用规则是什么？
# a
`GROUP BY` 将结果按指定列分组，常与聚合函数（`COUNT`、`SUM`、`AVG`、`MAX`、`MIN`）配合使用。
- **规则**：`SELECT` 子句中的非聚合列必须出现在 `GROUP BY` 中。
- **执行顺序**：在 `WHERE` 之后、`HAVING` 之前执行。
- **示例**：
```sql
SELECT name, COUNT(*) FROM course GROUP BY name;
```
```sql
SELECT department, gender, COUNT(*) FROM employees GROUP BY department, gender;
```

# q
`HAVING` 和 `WHERE` 有什么区别？
# a
- **执行时机**：`WHERE` 在分组前过滤原始行；`HAVING` 在分组后过滤分组结果。
- **使用限制**：`WHERE` 不能使用聚合函数；`HAVING` 可以使用聚合函数。
- **语法位置**：`WHERE` 在 `FROM` 之后、`GROUP BY` 之前；`HAVING` 在 `GROUP BY` 之后。
- **典型示例**：
```sql
SELECT name, COUNT(*) FROM course
GROUP BY name
HAVING COUNT(*) > 2;
```

# q
`COUNT(*)` 和 `COUNT(列名)` 有什么不同？
# a
- `COUNT(*)` 统计所有行，包括 NULL 值所在的行。
- `COUNT(列名)` 统计该列非 NULL 值的数量，忽略 NULL。
- **性能**：`COUNT(*)` 通常更快，因为无需检查列值。
- **使用场景**：统计总行数用 `COUNT(*)`；统计有效值数量用 `COUNT(列名)`。统计去重数量可用 `COUNT(DISTINCT 列名)`。

# q
SQL查询的书写顺序和执行顺序分别是什么？
# a
**书写顺序**：
```sql
SELECT ... FROM ... WHERE ... GROUP BY ... HAVING ... ORDER BY ... LIMIT ...
```
**执行顺序**：
1. `FROM`   – 确定数据源
2. `WHERE`  – 过滤原始行
3. `GROUP BY` – 分组
4. `HAVING` – 过滤分组结果
5. `SELECT` – 选择字段
6. `ORDER BY` – 排序
7. `LIMIT`  – 限制行数

注意：`WHERE` 运行在分组前，不能使用聚合函数；`HAVING` 运行在分组后，可以使用聚合函数。`ORDER BY` 和 `LIMIT` 最后执行。

