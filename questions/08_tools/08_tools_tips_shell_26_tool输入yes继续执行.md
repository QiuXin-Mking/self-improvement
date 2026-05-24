# q
如何在Bash脚本中提示用户输入并读取输入？
# a
使用 `read` 命令读取用户输入到变量，例如：
```bash
echo -n "请输入 'yes' 来继续执行接下来的命令: "
read input
```
`-n` 选项使 `echo` 输出后不换行，`read input` 将输入保存到变量 `input` 中。

# q
在Bash中如何判断用户输入是否为特定字符串（如 "yes"）？
# a
使用 `if` 语句和字符串比较条件 `[ "$input" = "yes" ]`：
```bash
if [ "$input" = "yes" ]; then
    # 输入为 "yes" 时执行的命令
else
    # 输入不为 "yes" 时执行的命令
fi
```
注意：条件测试中变量需用双引号包裹，等号两边各留空格。

# q
如何在Bash中根据用户输入决定是否执行后续命令？
# a
在条件分支中放置需要执行的命令，如果输入匹配则执行 `then` 块，否则执行 `else` 块或结束脚本。例如：
```bash
if [ "$input" = "yes" ]; then
    echo "用户输入了 'yes'，执行接下来的命令。"
    # 在这里写实际要执行的命令
else
    echo "用户输入不是 'yes'，取消执行。"
fi
```

