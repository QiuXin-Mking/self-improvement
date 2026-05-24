# q
SQLite如何创建一个数据库文件？
# a
无需专门命令，首次连接数据库文件时SQLite会自动创建。例如：
```bash
sqlite3 mydatabase.db
```

# q
在SQLite中如何创建一个包含主键和NOT NULL约束的表？
# a
使用`CREATE TABLE`语句，例如：
```sql
CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    department TEXT
);
```
`INTEGER PRIMARY KEY`表示主键且自动递增，`TEXT NOT NULL`表示文本列不能为空。

# q
如何向SQLite表中插入一行数据？
# a
使用`INSERT INTO`语句，指定列名和对应的值：
```sql
INSERT INTO employees (name, age, department) VALUES ('Alice', 30, 'HR');
```

# q
SQLite中如何查询年龄大于25岁的所有员工？
# a
使用`SELECT`语句配合`WHERE`条件：
```sql
SELECT * FROM employees WHERE age > 25;
```

# q
在SQLite中如何进行事务操作以确保数据原子性？
# a
使用`BEGIN TRANSACTION`开始事务，执行操作后`COMMIT`提交，或`ROLLBACK`回滚。例如：
```sql
BEGIN TRANSACTION;
INSERT INTO employees (name, age, department) VALUES ('Bob', 25, 'IT');
UPDATE employees SET age = 26 WHERE name = 'Bob';
COMMIT;
```

