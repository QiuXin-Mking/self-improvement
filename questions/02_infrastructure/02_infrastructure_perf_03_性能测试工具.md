# q
iostat 5 3 命令的输出行为是什么？
# a
每隔 5 秒输出一次 I/O 统计信息，总共输出 3 次。

# q
如何使用 rbd bench 测试 RBD 镜像的随机写性能（4K 块大小，16 线程，总数据 1G）？
# a
```sh
rbd bench --pool my_pool --image test_img --io-size 4K --io-threads 16 --io-total 1G --rw randwrite
```

# q
在 fio 配置文件中，direct=1 参数的作用是什么？
# a
绕过系统缓存（直接 I/O），确保测试结果不受缓存干扰，是测试 RBD 性能时的必选参数。

