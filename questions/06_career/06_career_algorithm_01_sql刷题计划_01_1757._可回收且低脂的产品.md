# q
在 Products 表中，如何查询既是低脂又是可回收的产品编号？
# a
使用 `WHERE` 子句同时过滤 `low_fats = 'Y'` 和 `recyclable = 'Y'`，查询语句为：
```sql
SELECT product_id FROM Products WHERE low_fats = 'Y' AND recyclable = 'Y';
```

# q
题目要求返回的结果是否有顺序要求？
# a
无顺序要求，返回结果中产品编号的顺序可以任意排列。

# q
Products 表中的 low_fats 和 recyclable 列分别有哪些取值？
# a
两者都是枚举类型，可取值为 `'Y'`（是）和 `'N'`（否）。

