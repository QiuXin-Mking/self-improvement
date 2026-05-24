# q
如何加速 Ceph 的 `make` 编译？
# a
```bash
make -j 64
```

# q
在 Ceph 编译容器内如何启动默认测试集群？
# a
```bash
cd build
make vstart
../src/vstart.sh --debug --new -x --localhost --bluestore
```

# q
如何启动指定节点数量的 Ceph 测试集群？
# a
```bash
MON=3 OSD=3 MDS=1 MGR=1 RGW=1 ../src/vstart.sh -n -d
```

# q
启动集群后，如何查看集群状态？
# a
```bash
./bin/ceph -s
```

# q
如何停止 Ceph 测试集群？
# a
```bash
../src/stop.sh
```

