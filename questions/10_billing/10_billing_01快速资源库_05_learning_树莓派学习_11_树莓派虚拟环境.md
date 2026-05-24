# q
如何在树莓派上使用 venv 创建一个名为 myenv 的 Python 虚拟环境？
# a
使用命令 `python3 -m venv myenv`，这会在当前目录下生成 myenv 文件夹，包含独立的 Python 解释器和库。

# q
如何激活刚才创建的 myenv 虚拟环境？
# a
运行 `source myenv/bin/activate`，激活后命令行提示符前会出现 `(myenv)`。

# q
如何查找当前用户主目录下所有名称以 “env” 结尾的虚拟环境目录？
# a
运行 `find ~ -type d -name "*env"`，该命令会递归搜索家目录，列出所有以 “env” 结尾的目录路径。

# q
在树莓派上如何将 pip 的默认包索引源永久更换为清华大学镜像？
# a
执行 `pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple`，该命令会在 `~/.config/pip/pip.conf` 中写入全局配置，后续 pip 安装将从清华源获取包。

