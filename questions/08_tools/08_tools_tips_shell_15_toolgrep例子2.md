# q
在 Shell 脚本中使用 `getopts` 解析命令行选项时，如何获取带参数的选项值？
# a
在 `getopts` 解析到需要参数的选项时，通过内置变量 `$OPTARG` 获取该选项的参数值。例如：
```sh
while getopts "p:h" arg; do
  case $arg in
    p) parameter="$OPTARG" ;;
    h) call_help ;;
    *) call_help; exit 1 ;;
  esac
done
```
选项字符串 `p:` 表示 `-p` 需要一个参数，该参数会被存入 `$OPTARG`。

# q
示例脚本中如何检查是否没有传递任何命令行参数，并触发帮助？
# a
通过检查位置参数总数 `$#` 是否等于 0 来判断：
```sh
[ $# -eq 0 ] && call_help && exit 1
```
如果没有参数，调用 `call_help` 并退出。

# q
脚本中条件 `[ x"$parameter" = x'y' ]` 中的 `x` 前缀有什么作用？这种写法常见于什么场景？
# a
`x` 前缀用于防止变量为空时测试命令语法错误（如 `[ = y ]`）。如果 `$parameter` 为空，比较变成 `[ x = xy ]`，仍是一个合法表达式。这种写法常见于传统的 Bourne Shell 脚本中，以提高字符串比较的健壮性。

# q
根据示例脚本，当选项 `-p` 的值为 `n` 时会发生什么？
# a
当 `-p n` 时，`parameter` 被赋值为 `n`，后续判断：
```sh
[ x"$parameter" = x'n' ] && call_help
```
因此会直接调用帮助函数，不执行其他操作，最后执行 `sync`。

