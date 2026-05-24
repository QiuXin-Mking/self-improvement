# q
在GDB中如何查看当前函数的所有局部变量？
# a
使用命令 `info locals` 或简写 `i locals`。

# q
在GDB中如何查看当前的调用堆栈？
# a
使用命令 `bt`（backtrace）。

# q
在GDB中如何设置断点？
# a
使用 `break` 命令，常见形式：`break 文件名:函数名`（如 `break a.cpp:foo`）或 `break 文件名:行号`（如 `break a.cpp:42`）。

# q
在GDB中如何删除或禁用断点？
# a
先用 `info breakpoints` 查看断点编号，然后用 `delete 编号` 删除（如 `delete 2`），用 `disable 编号` 禁用（如 `disable 1`）。

# q
在GDB中如何查看当前函数的参数？
# a
使用命令 `info args`。

