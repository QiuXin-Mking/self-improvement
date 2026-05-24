# q
执行keystone命令时出现“DeprecationWarning: The keystone CLI is deprecated”警告的典型根因是什么？
# a
该警告是因为系统安装的`keystone`命令行工具已经过时，官方推荐改用`python-openstackclient`（OpenStack统一客户端）。  
警告日志示例：
```
/usr/lib/python2.7/site-packages/keystoneclient/shell.py:65: DeprecationWarning: The keystone CLI is deprecated in favor of python-openstackclient. For a Python library, continue using python-keystoneclient.
```
解决方法：使用`openstack user show`等相应命令替代`keystone`命令。

# q
如何通过命令查看某个keystone用户的Access Key和Secret Key？
# a
使用`keystone user-get <用户名>`可以查看用户基本信息，但不会直接列出AK/SK。  
要获取用户的AK/SK，需通过其他命令或API，样例输出中用户信息下方直接附带了租户的access/secret列表，例如：
```
+----------------+----------------------------------+----------------------------------+
|     tenant     |              access              |              secret              |
+----------------+----------------------------------+----------------------------------+
| PMIaaS_project | 098242d25b0c491e813f9ac54d124991 | dbdb9e13bd1e457b8d4e03887b07beb3 |
|    linliang    | 67ba0a56fe7940ddb65b271996d5b291 | 1984c2526251404e9928e4ee834affd8 |
...
```
从中可直接提取对应租户的`AK`和`SECRET`。若只关心某个用户，可结合租户名称过滤提取。

