# q
如何编译和替换 MDBS 的 C 共享库？
# a
1. 进入编译服务器（如 251.101），mount 整个 hci 代码。
2. 进入修改模块的 build 目录，依次执行：
   ```shell
   cmake ..
   make
   ```
3. 编译完成后，将生成的 .so 文件（如 libbase.so）通过 scp 分发到所有集群节点的 `/opt/macrosan/mdbs/lib/` 目录，例如：
   ```shell
   scp -r /home/qiuxin/lib/libbase.so 172.22.251.104:/opt/macrosan/mdbs/lib/
   scp -r /home/qiuxin/lib/libbase.so 172.22.251.105:/opt/macrosan/mdbs/lib/
   scp -r /home/qiuxin/lib/libbase.so 172.22.251.106:/opt/macrosan/mdbs/lib/
   ```

# q
如何停止和重启 MDBS 集群的所有服务？
# a
1. SSH 到集群服务器（如 251.104）。
2. 激活 Python 虚拟环境并执行停止/重启脚本：
   ```shell
   source /opt/macrosan/mdbs/py_env/bin/activate
   python /opt/macrosan/mdbs/bin/tools/py/restart_services.py cluster stop all
   # 重启
   python /opt/macrosan/mdbs/bin/tools/py/restart_services.py cluster restart all
   ```
   也可通过 Python 封装函数远程重启：
   ```python
   def restart_services(ipaddr, i, info):
       proxy_info = 'ssh root@%s "source /opt/macrosan/mdbs/py_env/bin/activate;python /opt/macrosan/mdbs/web/restart_services.py %s %s all"' % (ipaddr, i, info)
       e2, o = subprocess.getstatusoutput(proxy_info)
       return o
   ```

# q
如何编译并部署 MDBS 的 Go 程序？
# a
1. 进入 Go 模块目录（如 mdbs_meta、mdbs_engine_agent），执行 `go build` 生成可执行文件。
   ```shell
   cd /home/qiuxin/hci_feat_8k_gran_4110/hci_golang/mdbs_meta
   go build
   cd /home/qiuxin/hci_feat_8k_gran_4110/hci_golang/mdbs_engine_agent
   go build
   ```
2. 将生成的可执行文件 scp 到所有节点的 `/opt/macrosan/mdbs/bin/golang/` 目录，例如：
   ```shell
   scp -r mdbs_engine_agent 172.22.251.62:/opt/macrosan/mdbs/bin/golang/
   scp -r mdbs_meta 172.22.251.62:/opt/macrosan/mdbs/bin/golang/
   ```

# q
节点被隔离后如何恢复？
# a
1. 查询被隔离的节点：
   ```shell
   etcdctl get /node/offline --prefix
   ```
2. 删除对应节点的隔离记录（例如节点 3）：
   ```shell
   etcdctl del /node/offline/3
   ```
3. 重启服务使节点重新加入集群。

# q
替换 Python 文件后应重启哪些服务？如何查看新增日志？
# a
1. 替换 Python 文件通常涉及 `/opt/macrosan/mdbs/web/` 目录（如 `cluster/api/`）。
2. 重启 Web 相关服务：
   ```shell
   systemctl stop nd_app
   systemctl stop cluster_app
   systemctl restart nd_app
   systemctl restart cluster_app
   ```
3. 查看新增日志：
   ```shell
   tail -f /var/log/mdbs/python_app.log
   tail -f /var/log/mdbs/engine.log
   tail -f /var/log/messages
   ```

