# q
httpd是什么，它与Apache HTTP服务器有什么关系？
# a
httpd 是 Apache HTTP 服务器的可执行文件名，通常用于指代 Apache HTTP 服务器本身。Apache HTTP 服务器（简称 Apache）是全球使用最广泛的开源 HTTP 服务器软件之一，由 Apache Software Foundation 开发和维护，用于托管网站和 web 应用程序。

# q
在Linux系统上如何使用yum安装并启动httpd服务？
# a
```bash
yum install httpd -y
systemctl start httpd
systemctl status httpd
```

# q
httpd的主配置文件路径是什么？
# a
```bash
/etc/httpd/conf/http.conf
```

