# q
cosbench的控制器配置文件(controller.conf)中如何配置driver节点？
# a
在`[controller]`段中设置`drivers = <数量>`，然后为每个driver添加`[driverN]`段，指定`name`和`url`，例如：
```ini
[controller]
drivers = 3
log_level = INFO
...
[driver1]
name = driver1
url = http://10.1.1.2:18088/driver
[driver2]
name = driver2
url = http://10.1.1.3:18088/driver
```

# q
cosbench的Web控制台如何访问？
# a
控制节点需要具备外部可访问的IP，在浏览器里访问`http://<控制节点IP>:19088/controller/index.html`。

# q
cosbench的driver节点需要哪些统一配置？
# a
所有成为cosbench客户端的节点（即driver节点）都需配置相同的`controller.conf`文件，其中应包含所有driver的URL信息；driver节点上使用`./stop-driver.sh`脚本进行相关操作（文档中描述为“启动”）。

