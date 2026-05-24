# q
如何对文件内容进行排序并去重并输出到新文件？
# a
可以使用以下命令：
```bash
sort -u file.txt > sorted.txt
sort file.txt | uniq > sorted.txt
```
- `sort -u` 直接排序并去重，最简洁高效
- `sort | uniq` 先排序，再去除相邻重复行
- 注意：`uniq` 只能去除相邻重复，因此必须先排序

# q
sort命令的-u选项有什么作用？
# a
`-u`（`--unique`）选项用于排序并去除重复行，只保留唯一的行。示例：
```bash
sort -u file.txt
```
比 `sort | uniq` 更简洁，一个命令完成排序和去重。

# q
uniq命令的特点是什么？为什么使用uniq前通常要先排序？
# a
`uniq` 只能去除**相邻的重复行**。如果重复行不相邻，则无法完全去重。因此，通常先使用 `sort` 排序将相同行排在一起，再使用 `uniq`。典型用法：
```bash
sort file.txt | uniq
```

# q
下列命令中哪个不能实现排序并去重？
A. `cat file.txt | sort -u > sorted.txt`
B. `sort file.txt | uniq > sorted.txt`
C. `sort -u file.txt > sorted.txt`
D. `awk file.txt | sort > sorted.txt`
# a
**错误选项：D**  
- `awk file.txt` 语法错误，`awk` 必须指定操作；且 `sort` 无 `-u` 时不执行去重。  
- 选项 A、C 使用 `sort -u` 排序去重，选项 B 通过 `sort | uniq` 实现去重，均正确。

