# q
SQL中row_number()、rank()、dense_rank()这三个窗口函数的区别是什么？
# a
**row_number()**：给每一行分配唯一的连续序号，即使排序字段值相同也会分配不同的序号

**rank()**：相同值会获得相同的排名，但下一个排名会跳过相应的位置（跳跃排名）

**dense_rank()**：相同值会获得相同的排名，但下一个排名连续不跳跃（密集排名）

**对比示例：**
假设数据按score排序：90, 85, 85, 80

| score | row_number() | rank() | dense_rank() |
|-------|--------------|--------|--------------|
| 90    | 1            | 1      | 1            |
| 85    | 2            | 2      | 2            |
| 85    | 3            | 2      | 2            |
| 80    | 4            | 4      | 3            |

**关键区别：**
- `row_number()`：总是连续（1,2,3,4）
- `rank()`：相同值相同排名，下一个跳跃（1,2,2,4）
- `dense_rank()`：相同值相同排名，下一个连续（1,2,2,3）

# q
什么是SQL窗口函数？窗口函数的基本语法是什么？
# a
**窗口函数（Window Function）**：也称为分析函数，是对一组行（称为窗口）进行计算并返回结果给每一行的函数。

**基本语法：**
```sql
窗口函数名() OVER (
    [PARTITION BY 列名]  -- 可选：定义窗口分区
    [ORDER BY 列名]      -- 可选：定义窗口内的排序
    [ROWS/RANGE BETWEEN ...]  -- 可选：定义窗口范围
)
```

**窗口函数分类：**
1. **排名函数**：`row_number()`, `rank()`, `dense_rank()`, `ntile()`
2. **聚合函数**：`sum()`, `avg()`, `count()`, `max()`, `min()`（作为窗口函数使用）
3. **偏移函数**：`lag()`, `lead()`, `first_value()`, `last_value()`

**示例：**
```sql
-- 按部门分区，计算每个部门内工资排名
SELECT 
    name,
    department,
    salary,
    RANK() OVER (PARTITION BY department ORDER BY salary DESC) as dept_rank
FROM employees;
```

# q
SQL窗口函数中的PARTITION BY和ORDER BY分别有什么作用？
# a
**PARTITION BY**：将数据分成多个分区（窗口），窗口函数在每个分区内独立计算

**ORDER BY**：在窗口内定义排序规则，决定窗口函数计算的顺序

**示例对比：**

```sql
-- 不使用PARTITION BY：在整个结果集上计算排名
SELECT 
    name,
    salary,
    RANK() OVER (ORDER BY salary DESC) as overall_rank
FROM employees;
-- 结果：所有员工按工资全局排名

-- 使用PARTITION BY：在每个部门内计算排名
SELECT 
    name,
    department,
    salary,
    RANK() OVER (PARTITION BY department ORDER BY salary DESC) as dept_rank
FROM employees;
-- 结果：每个部门内员工按工资排名（每个部门的排名从1开始）
```

**关键点：**
- `PARTITION BY` 类似于 `GROUP BY`，但不会减少行数
- `ORDER BY` 定义窗口内的计算顺序
- 两者可以单独使用，也可以组合使用
- `PARTITION BY` 后可以跟多个列，用逗号分隔

# q
SQL中rank()和dense_rank()在什么场景下使用？有什么区别？
# a
**使用场景：**

**rank()适用场景：**
- 体育比赛排名：如马拉松比赛，相同时间的人并列第3，下一个人排第5
- 学生成绩排名：需要知道有多少人比自己排名靠前
- 股票排名：显示实际的排名位置

**dense_rank()适用场景：**
- 等级评定：如学生等级A/B/C，即使有多个人是B级，下一个等级仍然是C
- 分层统计：需要连续的排名编号
- 百分比计算：计算前10%的数据（使用dense_rank更容易）

**区别示例：**
假设10个学生按分数排名，其中3个学生都是85分：

| 分数 | rank() | dense_rank() | 说明 |
|------|--------|--------------|------|
| 100  | 1      | 1            |      |
| 90   | 2      | 2            |      |
| 85   | 3      | 3            | 3个学生都是85分 |
| 85   | 3      | 3            |      |
| 85   | 3      | 3            |      |
| 80   | 6      | 4            | rank()跳到6，dense_rank()连续为4 |
| 70   | 7      | 5            |      |

**关键区别：**
- `rank()`：反映实际排名位置，相同值后下一个排名会跳跃
- `dense_rank()`：反映等级层次，相同值后下一个排名连续

