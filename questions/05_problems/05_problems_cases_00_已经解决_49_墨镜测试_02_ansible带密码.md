# q
如何在 Ansible inventory 文件中为主机配置 SSH 密码认证？
# a
在使用密码连接时，直接在 `/etc/ansible/hosts`（或自定义 inventory）的主机条目后追加连接变量，例如：
```ini
[webserver]
192.168.1.101 ansible_user=root ansible_password=123456 ansible_port=22
```

# q
批量将 IP 列表转换为带 SSH 密码认证的 Ansible 主机清单条目应使用什么命令？
# a
假设 `ip` 文件存放纯 IP 列表（每行一个 IP），可使用 `awk` 命令生成带认证信息的主机组配置：
```bash
awk '{print $0, "ansible_user=root ansible_password=Aa12345678&* ansible_port=22"}' ip > /etc/ansible/hosts
```
注意：输出应写入 inventory 文件（如 `hosts`），而非 `ansible.cfg`。

# q
安装 Ansible 后执行 `ansible` 命令提示 `command not found` 的根本原因及解决方法是什么？
# a
原因：Ansible 通过 pip 安装到用户目录时，可执行文件位于 `~/.local/bin`，而该路径不在 `PATH` 环境变量中。  
解决方法：将路径追加到 `~/.bashrc` 并使其生效：
```bash
echo 'export PATH=$HOME/.local/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
```

