# q
如何查询 Views 表中所有浏览过自己文章的作者 ID，并按升序排列？
# a
使用 `WHERE author_id = viewer_id` 过滤出作者浏览自己文章的行，再用 `DISTINCT` 去重，最后 `ORDER BY` 升序排列，并将列重命名为 `id`：
```sql
SELECT DISTINCT author_id AS id
FROM Views
WHERE author_id = viewer_id
ORDER BY id;
```

# q
为什么在此查询中需要使用 `DISTINCT`？
# a
因为 `Views` 表可能包含重复行（没有主键），同一作者可能多次浏览自己的文章，使用 `DISTINCT` 可以确保结果中每个作者只出现一次。

# q
在 SQL 中如何给查询结果列取别名？
# a
使用 `AS` 关键字，格式为 `column_name AS alias`，例如 `author_id AS id`。

# q
`ORDER BY id` 在这道题里按什么顺序排列？
# a
默认按升序（ASC）排列，将作者 ID 从小到大输出。

