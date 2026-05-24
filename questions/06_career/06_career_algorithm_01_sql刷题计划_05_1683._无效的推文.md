# q
在 LeetCode 1683 题中，如何查询所有内容字符数严格大于 15 的无效推文的 tweet_id？
# a
```sql
SELECT tweet_id
FROM Tweets
WHERE CHAR_LENGTH(content) > 15;
```
使用 `CHAR_LENGTH(content)` 计算字符串的字符数（而不是字节长度），筛选长度大于 15 的行，返回对应 `tweet_id`。

