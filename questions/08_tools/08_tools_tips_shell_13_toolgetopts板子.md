# q
在Shell脚本中，如何用getopts实现命令行选项的基本解析框架？
# a
使用 `while getopts "选项字符串" 变量` 循环，并在 `case` 语句中匹配每个选项字符。示例如下：
```shell
while getopts "p:h" arg
do
    case $arg in
    p)
        parameter="$OPTARG"
        ;;
    h)
        call_help
        ;;
    ?)
        call_help
        exit 1
        ;;
    esac
done
```

# q
getopts选项字符串中的冒号有什么作用？如何定义一个没有参数的长选项？
# a
选项字符后跟一个冒号表示该选项需要附加一个参数，参数值会保存在 `$OPTARG` 中。例如 `"p:"` 表示 `-p` 必须后接参数。  
如果选项不需要参数，则不要在字符后加冒号，例如 `"h"`，此时 `-h` 仅作为开关存在，`$OPTARG` 不会被赋值。

# q
在getopts的case分支中，`?` 通常用来处理什么情况？
# a
`?` 分支用于处理用户输入了未定义的选项（即不在选项字符串中的字符）时的错误。通常会在该分支中调用帮助函数并退出脚本，如：
```shell
?)
    call_help
    exit 1
    ;;
```

