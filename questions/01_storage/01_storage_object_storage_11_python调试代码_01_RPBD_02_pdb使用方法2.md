# q
在pdb中，如何查看可用的调试命令列表？
# a
在pdb提示符下输入 `help` 命令，会列出所有文档化的命令，例如断点命令 (`b`/`break`)、执行控制命令 (`c`/`continue`, `n`/`next`, `s`/`step`)、显示命令 (`list`, `ll`/`longlist`, `p`/`pp`)、堆栈查看命令 (`w`/`where`, `u`/`up`, `d`/`down`)、退出命令 (`q`/`quit`) 等。

# q
pdb的 `help` 命令输出的杂项帮助主题（Miscellaneous help topics）包括哪些？
# a
包括 `exec` 和 `pdb`。

# q
在pdb帮助命令列表中，断点相关的命令有哪些简写？
# a
断点相关命令简写包括 `b`、`break`、`tbreak`（临时断点）、`condition`、`ignore`、`enable`、`disable`、`clear`、`commands`。

# q
pdb中用于程序执行控制的命令有哪些简写？
# a
包括 `c`/`cont`/`continue`（继续执行）、`n`/`next`（执行下一行）、`s`/`step`（步入函数）、`r`/`return`（执行到当前函数返回）、`j`/`jump`（跳转到指定行）、`until`（执行到指定行或循环结束）、`unt`（`until`的别名）。

