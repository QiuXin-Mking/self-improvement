# q
创建私有源，如何使用createrepo命令？
# a
createrepo -pdo /var/www/html/repo/ /path/to/my/rpm/packages/

- `-o /var/www/html/repo/`：将生成的 repodata目录输出到​ /var/www/html/repo/下。执行后，你会在这个目录下看到 repodata/文件夹。
- `-d`：在生成标准 XML 文件的同时，也创建 SQLite 数据库文件（在 repodata/目录内），以优化客户端性能。
- `-p`：生成的 XML 文件是经过格式化、带缩进的，方便人阅读。

```
/var/www/html/repo/
└── repodata/
    ├── filelists.sqlite.bz2
    ├── filelists.xml.gz
    ├── other.sqlite.bz2
    ├── other.xml.gz
    ├── primary.sqlite.bz2
    ├── primary.xml.gz
    └── repomd.xml
```

```
/path/to/my/rpm/packages/
├── package-A.rpm
├── package-B.rpm
└── ... (其他rpm包)
```

# q
如果在createrepo过程中，源包有更新，应该如何继续更新repodata呢？
# a
```
createrepo -update /home/yumrepo
```

# q
如何启动一个服务器，将repo源给其他服务器使用，用python？
# a
在`/home/yumrepo`目录下，执行如下命令：
```
nohup python -m SimpleHTTPServer 8080 &>/dev/null &
```

