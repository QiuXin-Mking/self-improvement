# q
在SQLite中，如何创建一个包含id、name、age和department字段的employees表，并将id设为主键自动递增？
# a
使用```sql
CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    department TEXT
);
```
其中`INTEGER PRIMARY KEY`表示该列是主键且自动递增，`TEXT NOT NULL`表示该列不能为空。

# q
如何在SQLite的employees表中插入一条员工记录（如Alice，30岁，HR部门）？
# a
使用```sql
INSERT INTO employees (name, age, department) VALUES ('Alice', 30, 'HR');
```

# q
在SQLite中，如何更新employees表中Alice的年龄为31？
# a
使用```sql
UPDATE employees SET age = 31 WHERE name = 'Alice';
```

# q
如何在SQLite中删除employees表中名为Alice的记录？
# a
使用```sql
DELETE FROM employees WHERE name = 'Alice';
```

# q
如何在SQLite中使用事务确保多条操作（如插入Bob并更新其年龄）的原子性？
# a
使用事务控制：
```sql
BEGIN TRANSACTION;
INSERT INTO employees (name, age, department) VALUES ('Bob', 25, 'IT');
UPDATE employees SET age = 26 WHERE name = 'Bob';
COMMIT;
```
`BEGIN TRANSACTION`开始事务，`COMMIT`提交事务使更改生效，如果出错可使用`ROLLBACK`回滚取消所有操作。

