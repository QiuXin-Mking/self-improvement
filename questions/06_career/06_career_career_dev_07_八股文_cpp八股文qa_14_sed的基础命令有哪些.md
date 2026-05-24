# q
如何使用sed将文件中的"apple"全部替换为"orange"?
# a
```bash
sed 's/apple/orange/g' file.txt
```

# q
如何使用sed只打印文件的第三行?
# a
```bash
sed -n '3p' file.txt
```

# q
如何使用sed过滤出包含"error"的行?
# a
```bash
sed -n '/error/p' log.txt
```

# q
如何使用sed删除文件中的空行?
# a
```bash
sed '/^$/d' file
```

# q
sed命令中 -n 选项的作用是什么?
# a
未使用 `-n` 时，sed 会自动打印每一行（无论是否执行操作）。使用 `-n` 时，sed 仅打印显式指定的行（通过 `p` 命令或匹配条件的行）。

