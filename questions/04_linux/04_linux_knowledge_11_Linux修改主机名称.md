# q
如何查看当前 Linux 系统的主机名？
# a
使用以下命令：
```sh
hostnamectl
```
或
```sh
hostname
```

# q
Linux 系统中临时修改主机名和永久修改主机名分别使用什么命令？
# a
- 临时修改（重启后恢复原名）：
  ```sh
  hostname 新主机名
  ```
- 永久修改（重启后依然有效）：
  ```sh
  hostnamectl set-hostname 新主机名
  ```
  或编辑 `/etc/hostname` 文件后重启。

# q
直接修改 /etc/hostname 和使用 hostnamectl set-hostname 在持久化效果上有何异同？
# a
两种方式都能实现永久修改，重启后主机名都保持修改后的值。
区别在于：修改 `/etc/hostname` 后必须手动重启或使用 `reboot` 命令才能生效；而 `hostnamectl set-hostname` 会立即生效，同时更新内核中的主机名，不需要立刻重启（为保证所有服务识别新名称，仍建议重启或重新登录）。

# q
在 openEuler 系统中修改主机名，除了执行 hostnamectl set-hostname，还需要做什么额外操作？
# a
还需要修改 `/etc/hosts` 文件，将其中出现的旧主机名替换为新主机名，然后重启系统使更改全面生效：
```sh
sudo hostnamectl set-hostname new-hostname
sudo sed -i 's/旧主机名/新主机名/g' /etc/hosts
sudo reboot
```

