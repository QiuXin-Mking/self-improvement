# q
如何排查在 Linux 中执行 `shutdown`、`fdisk` 等管理命令时出现 “bash: command not found” 错误？
# a
首先检查当前用户的 `$PATH` 环境变量是否包含 `/sbin`、`/usr/sbin` 等命令所在目录。可以用 `echo $PATH` 查看。如果缺失，可以临时添加：
```bash
PATH=$PATH:$HOME/bin:/sbin:/usr/bin:/usr/sbin
```
若需永久生效，应将此配置写入 `/etc/profile` 或用户的 `~/.bashrc` 中。

# q
`/bin` 和 `/sbin` 目录分别主要存放什么类型的命令？
# a
- `/bin`：存放系统必备的基础用户命令，如 `cat`、`cp`、`ls`、`mkdir`、`rm` 等。
- `/sbin`：存放超级用户（root）使用的系统管理必备工具，如 `fdisk`、`shutdown`、`ifconfig`、`reboot` 等。

# q
`/usr/bin` 和 `/usr/sbin` 目录分别主要存放什么类型的命令？
# a
- `/usr/bin`：存放后期安装的应用软件必备执行文件，例如 `gcc`、`make`、`wget`、`passwd` 等。
- `/usr/sbin`：存放用户安装的系统管理守护进程或服务相关程序，例如 `httpd`、`named`、`samba`、`tcpdump` 等。

# q
如何将自己编写的 Shell 脚本转化为可以直接使用的系统命令？
# a
将写好的脚本去掉 `.sh` 后缀，然后复制到 `/usr/bin` 目录下，并赋予可执行权限：
```bash
sudo cp myscript /usr/bin/myscript
sudo chmod +x /usr/bin/myscript
```
之后即可在终端中直接输入脚本名称执行。

