# q
如何通过 wget 下载 RocksDB 指定版本的源码压缩包？
# a
使用以下命令下载指定版本（例如 v9.7.3）的 tar.gz 包：
```bash
wget https://github.com/facebook/rocksdb/archive/refs/tags/v9.7.3.tar.gz -O rocksdb-9.7.3.tar.gz
```
下载后解压并进入目录：
```bash
tar -xzf rocksdb-9.7.3.tar.gz
cd rocksdb-9.7.3
```

# q
如何通过 Git 克隆 RocksDB 仓库并切换到特定稳定分支或标签？
# a
克隆代码仓库：
```bash
git clone https://github.com/facebook/rocksdb.git
cd rocksdb
```
切换到主分支或特定版本标签（如 v9.7.3）：
```bash
git checkout main
# 或
git checkout v9.7.3
```

