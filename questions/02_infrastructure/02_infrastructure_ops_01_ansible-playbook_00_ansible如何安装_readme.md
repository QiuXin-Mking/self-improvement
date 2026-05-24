# q
如何通过pip3安装Ansible并指定使用清华大学镜像源？
# a
```bash
pip3 install --user ansible -i https://pypi.tuna.tsinghua.edu.cn/simple
```
`--user` 会将 Ansible 安装到当前用户的 `~/.local/bin` 目录。

# q
pip安装Ansible后，如何配置环境变量使其可直接使用？
# a
编辑 `~/.bashrc`，添加：
```bash
export PATH=$PATH:~/.local/bin
```
然后执行 `source ~/.bashrc` 立即生效，或重新打开终端。

# q
使用yum安装Ansible会遇到什么问题？
# a
部分仓库中不包含 ansible 包，执行 `yum install ansible` 会报错：
```
No package ansible available.
```

# q
文档中提到的“线上 ansible.bin 可以直接安装”是指什么？
# a
可以通过直接下载预编译的 `ansible.bin` 二进制文件进行安装，无需使用包管理器。

