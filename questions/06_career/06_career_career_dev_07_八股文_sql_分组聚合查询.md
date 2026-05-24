# q
SQL中GROUP BY的作用是什么？什么情况下需要使用GROUP BY？
# a
**GROUP BY的作用：** 将查询结果按照指定的列进行分组，通常与聚合函数（如COUNT、SUM、AVG、MAX、MIN）一起使用。

**基本语法：**
```sql
SELECT 列名, 聚合函数(列名)
FROM 表名
GROUP BY 列名;
```

**使用场景：**
1. **统计汇总**：需要按某个维度统计数量、总和、平均值等
2. **分组计数**：统计每个分组中的记录数
3. **分组计算**：计算每个分组的聚合值

**示例：**

**1. 统计每个学生的选课数量：**
```sql
SELECT name, COUNT(*) as course_count
FROM course
GROUP BY name;
```

**重要规则：**
- SELECT子句中的非聚合列必须在GROUP BY中出现
- GROUP BY可以包含多个列
- GROUP BY在WHERE之后、HAVING之前执行

# q
SQL中HAVING和WHERE的区别是什么？什么时候用HAVING，什么时候用WHERE？
# a
**HAVING和WHERE的核心区别：**

**1. 执行时机不同：**
- **WHERE**：在分组（GROUP BY）之前执行，过滤原始行
- **HAVING**：在分组（GROUP BY）之后执行，过滤分组结果

**2. 适用范围不同：**
- **WHERE**：只能过滤原始行的列，不能使用聚合函数
- **HAVING**：可以过滤分组后的结果，可以使用聚合函数

**3. 使用位置不同：**
- **WHERE**：位于 `FROM` 和 `GROUP BY` 之间
- **HAVING**：位于 `GROUP BY` 和 `ORDER BY` 之间

**SQL执行顺序：**
```
FROM → WHERE → GROUP BY → HAVING → SELECT → ORDER BY
```

**示例：**
```sql
-- WHERE过滤原始行
SELECT name, AVG(score) as avg_score
FROM scores
WHERE subject = '数学' AND score > 80
GROUP BY name;

-- HAVING过滤分组结果
SELECT name, COUNT(*) as course_count
FROM course
GROUP BY name
HAVING COUNT(*) > 2;
```

**总结规则：**
- 过滤原始行 → 使用 `WHERE`
- 过滤分组结果 → 使用 `HAVING`
- 过滤聚合函数结果 → 必须使用 `HAVING`
- 过滤普通列条件 → 优先使用 `WHERE`（效率更高）

# q
SQL查询语句的执行顺序是什么？
# a
SQL查询语句的**书写顺序**和**执行顺序**不同：

**执行顺序（实际处理顺序）：**
```
1. FROM         -- 选择数据源表
2. WHERE        -- 过滤原始行
3. GROUP BY     -- 分组
4. HAVING       -- 过滤分组结果
5. SELECT       -- 选择字段
6. ORDER BY     -- 排序
7. LIMIT        -- 限制结果数量
```

**完整示例：**
```sql
SELECT name, COUNT(*) as course_count
FROM course
WHERE language != 'C++'
GROUP BY name
HAVING COUNT(*) > 2
ORDER BY course_count DESC
LIMIT 5;
```

**执行流程：**
1. FROM course：选择course表
2. WHERE language != 'C++'：过滤掉C++课程
3. GROUP BY name：按学生姓名分组
4. HAVING COUNT(*) > 2：筛选选课数大于2的学生
5. SELECT name, COUNT(*) as course_count：选择字段
6. ORDER BY course_count DESC：按课程数降序排序
7. LIMIT 5：只返回前5条结果

# q
SQL中COUNT(*)和COUNT(列名)有什么区别？
# a
**区别：**
- `COUNT(*)`：统计所有行，包含NULL值
- `COUNT(列名)`：统计指定列中非NULL值的数量，忽略NULL

**示例：**
假设表中有4行数据，其中name列有1行NULL：
```sql
SELECT COUNT(*) FROM students;       -- 结果：4
SELECT COUNT(name) FROM students;    -- 结果：3（排除NULL）
```

**使用场景：**
- 需要统计总行数 → 使用 `COUNT(*)`
- 需要统计某列非NULL值的数量 → 使用 `COUNT(列名)`
- 统计不重复值 → `COUNT(DISTINCT 列名)`

**性能：**
`COUNT(*)` 通常性能更好，因为不需要检查列值是否为NULL。

