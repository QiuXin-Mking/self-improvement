# q
sort命令的-u选项有什么作用？
# a
`sort -u` 选项的作用是：**排序并去除重复行（只保留唯一的行）**

用法：
```bash
sort -u file.txt                 # 对文件排序并去重，输出到终端
sort -u file.txt > sorted.txt   # 排序并去重后保存到新文件
```

示例：
原始文件内容：
```
apple
banana
apple
cherry
banana
```
执行 `sort -u file.txt` 输出：
```
apple
banana
cherry
```

特点：
- `-u` 表示 `--unique`，先排序后去重，确保所有重复行都被去除
- 比 `sort | uniq` 更简洁，一次完成两个操作

# q
uniq命令的特点是什么？为什么使用uniq前通常要先排序？
# a
`uniq` 命令的特点：**只能去除相邻的重复行**，不处理不相邻的重复行。

常用用法：
```bash
uniq file.txt                # 直接去除相邻重复行
sort file.txt | uniq         # 先排序，使所有相同行相邻，再去重
```

常用选项：
```bash
uniq -c        # 在行前显示重复次数
uniq -d        # 只显示重复的行
uniq -u        # 只显示不重复的行
```

必须排序的原因：
- 如果重复行不相邻，`uniq` 无法识别并去除
- 例如文件内容为 `apple\nbanana\napple`，`uniq` 会保留两行 apple，只有先 `sort` 让相同行相邻，`uniq` 才能正确去重

# q
Linux中如何对文件内容进行排序并去重？列出常用命令方法。
# a
常用两种方法：
```bash
# 方法一：使用 sort -u（推荐）
sort -u file.txt > sorted.txt

# 方法二：排序后通过 uniq 去重
sort file.txt | uniq > sorted.txt
```

两种方法等效，`sort -u` 更简洁。

`sort` 的常用参数：
- `-u`：去重
- `-r`：降序排序
- `-n`：按数值排序（默认按字典序）
- `-k N`：按第 N 列排序
- `-t ','`：指定分隔符

例如按第2列数字升序排列：
```bash
sort -k 2 -n file.txt
```

