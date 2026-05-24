# q
如何查看当前所有环境变量和PATH变量的值？
# a
使用 `export` 命令显示所有环境变量，使用 `echo $PATH` 输出 PATH 变量的值。

# q
使用 `export PATH` 临时添加环境变量有什么特点和注意事项？
# a
特点：立即生效，但仅当前终端有效、窗口关闭后失效，仅对当前用户有效。  
注意事项：添加新路径时必须包含 `$PATH`，例如 `export PATH=$PATH:/new/path`，避免覆盖原有配置。

# q
Linux中用户级和系统级环境变量配置文件分别有哪些？如果想永久且对所有用户生效应该修改哪个文件？
# a
- 用户级配置文件：`~/.bashrc`、`~/.bash_profile`（或 `~/.profile`）  
- 系统级配置文件：`/etc/bashrc`、`/etc/profile`、`/etc/environment`  
永久对所有用户生效需修改系统级文件（如 `/etc/profile` 或 `/etc/bashrc`），需要管理员权限。

# q
Linux系统加载环境变量的顺序是怎样的？
# a
加载顺序：/etc/environment → /etc/profile → /etc/bash.bashrc → /etc/profile.d/ 下的脚本 → 用户 ~/.profile → ~/.bashrc。  
其中 `/etc/profile` 会加载 `/etc/bash.bashrc` 和 `/etc/profile.d/*.sh`，而 `~/.profile` 会加载 `~/.bashrc`。

# q
如何通过自定义文件集中管理项目环境变量，而不直接修改系统配置文件？
# a
创建一个自定义环境变量文件（如 `uusama.profile`），在其中用 `export` 定义所需变量，然后在 `~/.profile` 末尾添加 `source uusama.profile`，这样每次登录都会自动加载这些自定义变量。也可用类似方式通过 `alias` 定义命令别名，比如 `alias rm="rm -i"`。

