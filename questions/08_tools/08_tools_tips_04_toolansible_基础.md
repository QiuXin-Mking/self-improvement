# q
Ansible是什么？它适用于哪些场景？
# a
Ansible是一个配置管理工具和自动化运维工具，可用于批量任务和重复性工作，例如同时在100台服务器上安装并启动nginx、将文件一次性拷贝到100台服务器、持续为新服务器部署服务等。

# q
Ansible的“幂等性”是指什么？
# a
幂等性指Ansible以结果为导向，自动判断当前状态是否与目标状态一致：如果一致则不操作，不一致则调整为期望状态，从而保证重复执行相同操作得到相同结果。

# q
如何检查Ansible playbook的语法错误并模拟运行？
# a
使用命令 ```ansible-playbook --syntax-check <playbook.yml>``` 检查语法；使用 ```ansible-playbook --check <playbook.yml>``` 模拟运行（不会真正执行任务）。

# q
在Ansible playbook的tasks中，为模块参数赋值有哪两种常见格式？
# a
一种是用冒号加空格的YAML格式（例如 ```file: path: /testdir/testfile state: touch mode: 0700```），另一种是等号连接的单行格式（例如 ```file: path=/testdir/testfile state=touch mode=0700```），两者效果相同。

