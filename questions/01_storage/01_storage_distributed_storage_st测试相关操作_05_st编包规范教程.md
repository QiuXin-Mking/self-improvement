# q
如何挂载远程CIFS共享到本地Linux目录？
# a
使用 `mount -t cifs` 命令，指定共享路径和本地挂载点，并通过 `-o` 传递用户名和密码：
```bash
mount -t cifs //172.17.8.69/001_gitlab /home/qiuxin -o username=John,pass=5i6106100258
```

# q
在执行编译脚本前如何设置HTTP代理？
# a
通过设置环境变量 `http_proxy` 和 `https_proxy` 指向代理服务器（例如 `172.17.8.69:808`）：
```bash
export http_proxy=172.17.8.69:808
export https_proxy=172.17.8.69:808
```
之后执行 `curl www.baidu.com` 测试连通性，再运行编译脚本。

# q
如何为RPM包添加数字签名？
# a
使用 `rpm --addsign` 命令，后跟RPM包路径，然后输入密码完成签名：
```bash
rpm --addsign /home/qiuxin/hci1/build/mdbs-1.1.16-Linux.rpm
```
提示 `Enter pass phrase:` 时输入密码（示例中为 `macrosan`）。

