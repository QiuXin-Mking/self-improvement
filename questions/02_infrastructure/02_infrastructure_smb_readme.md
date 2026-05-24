# q
如何下载并解压 Samba 4.21.0 源码？
# a
```bash
wget https://download.samba.org/pub/samba/stable/samba-4.21.0.tar.gz
tar xf samba-4.21.0.tar.gz
cd samba-4.21.0
```

# q
如何通过日志快速了解 Samba 的整体架构？
# a
使用 `log level = 10` 的调试日志，追踪一条读或写路径的执行流程，即可大致理解其内部架构和组件交互。

