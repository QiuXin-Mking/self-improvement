# q
Python中执行相对导入时出现 "No module named '__main__.xxx'; '__main__' is not a package" 报错的核心原因是什么？
# a
因为脚本所在目录或目标模块所在目录缺少 `__init__.py` 文件，导致 Python 无法将其识别为包（package）。相对导入只在包内有效，当脚本作为 `__main__` 运行时，Python 不会将不含 `__init__.py` 的目录视为包。

# q
如何解决 Python 相对导入因缺少 `__init__.py` 导致的 ModuleNotFoundError？
# a
在目标模块所在的目录（例如 `multi/`）以及需要作为包处理的目录中创建空的 `__init__.py` 文件，确保目录被标记为包。然后保持正确的相对导入语句，例如 `from .multi_small_tool import scan_out_time_ms`，并确认目录结构正确。

# q
Python 中的 `__init__.py` 文件主要作用是什么？是否必须包含代码？
# a
`__init__.py` 用于标记一个目录为 Python 包，使该目录下的模块能被导入或支持相对导入。大多数情况下它是一个空文件即可，不需要包含任何代码；可选的包初始化代码（如定义变量或导入模块）不是必需的。

