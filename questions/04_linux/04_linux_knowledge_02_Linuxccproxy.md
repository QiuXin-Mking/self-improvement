# q
如何在Linux终端中通过环境变量设置HTTP代理？
# a
使用 `export` 命令设置 `http_proxy` 和 `https_proxy` 环境变量，格式为 `http://代理IP:端口`，例如：
```sh
export http_proxy=172.17.8.69:808
export https_proxy=172.17.8.69:808
```

# q
设置代理后如何验证网络连通性？
# a
使用 `curl` 命令访问一个外部网站，例如 `curl www.baidu.com`。如果返回了完整的HTML内容（如百度首页源码），说明代理配置成功，网络连通正常。

