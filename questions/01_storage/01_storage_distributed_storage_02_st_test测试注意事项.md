# q
磁盘全部离线时如何执行扫描上线操作？
# a
先用 `lsblk` 检查磁盘状态，再通过 `realpath /sys/block/sda` 查询对应 host，最后执行扫描命令：
```bash
echo "- - -" > /sys/class/scsi_host/host0/scan
```

# q
如何使用 etcdctl 查询和修改节点重建时间参数？
# a
查询所有 ttl 相关配置：
```bash
etcdctl get /ttl --prefix
```
修改关键重建时间（单位：秒）：
```bash
etcdctl --endpoints=127.0.0.1:23790 put /ttl/heartbeat_off/rebuild 300
etcdctl --endpoints=127.0.0.1:23790 put /ttl/isolate 300
etcdctl --endpoints=127.0.0.1:23790 put /ttl/disk_off/rebuild 60
```

# q
如何激活 MDBS 测试所需的 Python 环境并执行单个用例？
# a
激活环境：
```bash
export PYTHONPATH=/home/qiuxin/hci_python
source /etc/profile
source /opt/macrosan/mdbs/py_env/bin/activate
```
执行单个测试用例：
```bash
nosetests /home/qiuxin/hci_test/script/tools/snap_test.py:test_ts_snap_case1
```

