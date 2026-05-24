# q
RPM spec文件中的`%prep`阶段的作用是什么，通常包含什么命令？
# a
`%prep`阶段用于解压源码包并进入源码目录，通常使用 `%setup -q` 宏解压源码tarball。

# q
在spec文件中，如何定义软件包安装后和卸载前执行的脚本？
# a
使用 `%post` 和 `%preun` 脚本部分。`%post` 定义安装完成后执行的脚本，`%preun` 定义卸载前执行的脚本。

# q
`BuildRequires`字段的作用是什么？
# a
`BuildRequires`用于列出构建该RPM包过程中必须依赖的软件包。

# q
Spec文件的包基础信息部分包含哪些核心字段？
# a
包含 `Name`（包名称）、`Version`（版本号）、`Release`（发布版本号）、`Summary`（简短描述）等字段。

