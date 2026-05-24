# q
如何使用 sed 将文件 file.txt 中所有的 apple 替换为 orange？
# a
```bash
sed 's/apple/orange/g' file.txt
```
该命令会全局替换所有匹配的 apple 为 orange。

# q
如何使用 sed 仅打印文件的第三行？
# a
```bash
sed -n '3p' file.txt
```
`-n` 禁止默认输出，`3p` 显式打印第三行。

# q
如何使用 sed 过滤并显示包含 “error” 的行？
# a
```bash
sed -n '/error/p' log.txt
```
`-n` 抑制自动打印，`/error/p` 仅输出匹配 “error” 的行。

# q
如何使用 sed 删除文件中的空行？
# a
```bash
sed '/^$/d' file
```
该命令直接删除空行，无需 `-n` 选项。

# q
sed 命令中 -n 选项的作用是什么？请举例说明它与默认行为的区别。
# a
- 未使用 `-n`：sed 会**自动打印每一行**（即使执行了其他操作），因此像 `sed '3p' file` 会先打印所有行，并将第三行额外打印一次。
- 使用 `-n`：sed **仅打印显式指定的行**（通过 `p` 命令或匹配条件的行），例如 `sed -n '3p' file` 只会输出第三行。

