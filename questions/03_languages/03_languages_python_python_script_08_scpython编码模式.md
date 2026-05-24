# q
通过 subprocess 执行 source activate 命令能否激活 Python 虚拟环境？
# a
不能，该方法无法实现虚拟环境激活，且尝试用 Python2 激活 Python3 的环境也是不合理的。

# q
Python 脚本中如何使用 sys.argv 获取命令行参数？
# a
文件名本身：sys.argv[0]；第一个参数：sys.argv[1]；第二个参数：sys.argv[2]。

