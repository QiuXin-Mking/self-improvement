# q
yumdownloader --resolve 命令的作用是什么？
# a
`yumdownloader --resolve` 用于下载指定的RPM软件包及其所有依赖项。它会解析依赖关系，自动将主包和所需的依赖包一并下载到当前目录。例如：
```
yumdownloader --resolve libmount-devel libyaml-devel e2fsprogs-devel e2fsprogs
```

