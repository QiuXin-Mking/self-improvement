# q
Shell 脚本中使用 test 命令进行数值比较时，有哪些操作符及其含义？
# a
| 参数 | 含义 |
|:---:|:---|
| -eq | 等于则为真（equal） |
| -ne | 不等于则为真（not equal） |
| -gt | 大于则为真（greater than） |
| -ge | 大于等于则为真（greater equal） |
| -lt | 小于则为真（less than） |
| -le | 小于等于则为真（less equal） |

示例：
```shell
num1=100
num2=100
if test $[num1] -eq $[num2]
then
    echo '两个数相等！'
else
    echo '两个数不相等！'
fi
```

# q
如何在 Shell 中将 case 语句用作多分支选择结构？请说明其基本语法和执行逻辑。
# a
case 语句模拟其他语言中的 switch … case 结构，使用 `case … esac` 包裹。取值后必须跟 `in`，每个分支以右圆括号结束，执行体末尾使用 `;;` 表示 break，跳出整个 case 结构。未匹配时可用 `*` 捕获并执行默认操作。

基本语法：
```shell
case 取值 in
    模式1) 命令1 ;;
    模式2) 命令2 ;;
    *) 默认命令 ;;
esac
```
示例：
```shell
echo '输入 1 到 4 之间的数字:'
read aNum
case $aNum in
    1) echo '你选择了 1' ;;
    2) echo '你选择了 2' ;;
    3) echo '你选择了 3' ;;
    4) echo '你选择了 4' ;;
    *) echo '你没有输入 1 到 4 之间的数字' ;;
esac
```

# q
在 while 无限循环中，如何使用 break 和 continue 控制流程？
# a
- **break**：直接终止整个循环，常用于错误输入后退出循环。  
  示例：当输入不属于 1~5 时，输出提示并执行 `break` 跳出循环。
  
- **continue**：终止本次循环，进入下一次循环，其后的命令不再执行。  
  示例：当输入不合法时，输出提示后执行 `continue`，后面的 `echo "游戏结束"` 不会被执行，循环继续。

break 示例片段：
```shell
while :
do
    read aNum
    case $aNum in
        1|2|3|4|5) echo "你输入的数字为 $aNum!" ;;
        *) echo "你输入的数字不是 1 到 5 之间的! 游戏结束"
            break ;;
    esac
done
```

continue 示例片段：
```shell
while :
do
    read aNum
    case $aNum in
        1|2|3|4|5) echo "你输入的数字为 $aNum!" ;;
        *) echo "你输入的数字不是 1 到 5 之间的!"
            continue
            echo "游戏结束" ;;
    esac
done
```

