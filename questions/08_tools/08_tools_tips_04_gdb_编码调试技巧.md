# q
如何在Python日志中打印当前函数名？
# a
使用 `str(__name__)` 获取函数名，并通过格式化字符串传入日志函数，例如：
```python
LOG.info("function:%s " % str(__name__))
```

# q
在C语言中，如何通过日志宏自动记录当前函数名？
# a
使用预定义宏 `__FUNCTION__`，它会在编译时替换为当前函数的名字。示例：
```c
LOG_DBG("%s process##qx blk_in_osd:%u cmd_begin:%lu",
        (char *)(__FUNCTION__), blk_in_osd, ocache_cmd->cmd_begin);
LOG_ERROR("%s process##qx1 blk_in_osd:%u cmd_begin:%lu",
          (char *)(__FUNCTION__));
```
日志宏 `LOG_DBG`、`LOG_DBG_MSG`、`LOG_ERROR` 均支持传入 `__FUNCTION__`。

# q
如何手动编写一个能产生coredump的C程序？
# a
编写一个简单的C程序，解引用空指针引发段错误（SIGSEGV），系统将生成核心转储文件。示例代码：
```c
#include <stdio.h>
#include <stdlib.h>

int main(){
    int *a = NULL;
    int b = 0;
    b = *a;   // 触发空指针解引用
    return 0;
}
```
编译并运行该程序，如果系统的coredump配置已开启，就会生成 core 文件。

