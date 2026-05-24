# q
如何扫描并上线所有磁盘？
# a
使用命令 `echo "- - -" > /sys/class/scsi_host/host0/scan` 扫描并上线所有磁盘。如果执行中断，必须重新执行该命令以确保磁盘全部在线。

# q
如何使用 etcdctl 修改节点重建时间？
# a
通过 etcdctl 设置相关键值：
```shell
etcdctl --endpoints=127.0.0.1:23790 put /ttl/heartbeat_off/rebuild 300
etcdctl --endpoints=127.0.0.1:23790 put /ttl/isolate 300
etcdctl --endpoints=127.0.0.1:23790 put /ttl/disk_off/rebuild 60
```
可使用 `etcdctl get /ttl --prefix` 查询当前值。

# q
如何编译 HCI 的 RPM 包和 env 包？
# a
编译 RPM 包：确保已安装 ccproxy 并设置代理后，进入代码目录执行：
```sh
cd /home/qiuxin/gitlab/hci1/
./hci_script/hci_make.sh
```
编译 env 包：
```sh
export http_proxy=172.17.8.69:808
export https_proxy=172.17.8.69:808
cd /home/qiuxin/gitlab/hci1/
./hci_script/env_makerpm.sh
```

# q
测试前如何检查环境和执行单个用例？
# a
检查环境：
```python
nosetests /root/code/hci_test/script/tools/check_env.py
```
执行单个用例：
```python
nosetests /home/qiuxin/hci_test/script/tools/snap_test.py:test_ts_snap_case1
nosetests /home/qiuxin/hci1/hci_test/script/test/base_test/osd_test.py:test_osd_case1
```

# q
如何激活 MDBS Python 虚拟环境并设置 PYTHONPATH？
# a
执行以下命令激活环境和设置路径：
```shell
export PYTHONPATH=/home/qiuxin/hci1/hci_python
source /etc/profile
source /opt/macrosan/mdbs/py_env/bin/activate
```
根据实际代码路径可能需要调整 `PYTHONPATH`。

