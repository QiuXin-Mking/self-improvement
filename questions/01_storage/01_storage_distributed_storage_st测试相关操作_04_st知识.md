# q
如何激活宏杉分布式存储测试的Python虚拟环境？
# a
执行以下命令：
```bash
export PYTHONPATH=/home/qiuxin/hci_python
source /etc/profile
source /opt/macrosan/mdbs/py_env/bin/activate
```

# q
如何快速运行指定的测试用例（如osd_test中的test_osd_case9）？
# a
使用nosetests命令，格式为：
```bash
nosetests /root/code/hci_test/script/test/base_test/osd_test.py:test_osd_case9
```

# q
节点隔离和重建的etcdctl命令分别是什么？
# a
- 节点隔离300秒：`etcdctl --endpoints=127.0.0.1:23790 put /ttl/isolate 300`
- 节点重建（关闭心跳后重建）300秒：`etcdctl --endpoints=127.0.0.1:23790 put /ttl/heartbeat_off/rebuild 300`
- 磁盘关闭后重建60秒：`etcdctl --endpoints=127.0.0.1:23790 put /ttl/disk_off/rebuild 60`

# q
MDBS（宏杉分布式块存储）相关的crt证书文件存放在什么路径？
# a
`/opt/macrosan/mdbs/config/openssl/` 目录下，具体证书文件为 `ca_root.crt`。

# q
如何查看MDBS的Python应用日志？
# a
使用命令：`tail -f /var/log/mdbs/python_app.log`

