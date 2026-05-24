# q
数据库、表、主键、外键这些术语分别指什么？
# a
数据库是一组具有相关数据的表；表是带有数据的矩阵；主键是用来唯一标识表中每一行记录的字段，一个表通常只有一个主键；外键是连接两个表之间的链接。

# q
在 DESCRIBE 命令输出的 Key 列中，PRI 和 MUL 分别表示什么含义？
# a
PRI 表示该字段是主键（Primary Key），用来唯一标识表中的每一行记录；MUL 表示该字段是在一个索引中的非唯一键列（Multiple），即可以有重复值的索引列。

# q
MySQL 中 auto_increment 属性有什么作用？
# a
auto_increment 表示在插入新记录时，自动为该字段生成递增的唯一值，通常用于主键列。

# q
MySQL 中常见的数据类型可分为哪几类？请列举各类并给出例子。
# a
常见数据类型可分为六类：整数类型（如 INT、BIGINT、SMALLINT）、浮点数类型（如 FLOAT、DOUBLE）、字符串类型（如 CHAR、VARCHAR、TEXT）、日期和时间类型（如 DATE、TIME、DATETIME、TIMESTAMP）、布尔类型（通常表示为 BOOLEAN 或 BIT）、二进制类型（如 BLOB、BYTEA）。

