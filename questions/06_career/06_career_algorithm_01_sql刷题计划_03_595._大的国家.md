# q
编写 SQL 查询从 `World` 表中找出“大国”的名称、人口和面积，大国的条件是什么？
# a
大国条件：面积至少为 300 万平方公里（`area >= 3000000`）或者人口至少为 2500 万（`population >= 25000000`）。  
查询语句如下：
```sql
SELECT name, population, area
FROM World
WHERE area >= 3000000 OR population >= 25000000;
```

