# q
如何从 GitHub 克隆 2048 游戏源码？
# a
```
git clone https://github.com/gabrielecirulli/2048.git
```

# q
在 Ubuntu 上安装 Nginx 的命令是什么？
# a
```
sudo apt update
sudo apt install nginx
```

# q
如何将 2048 项目部署到 Nginx 的默认站点根目录？
# a
```
sudo cp -r 2048 /var/www/html/
```

# q
如何测试 Nginx 配置语法并让新配置生效？
# a
使用 `nginx -t` 进行语法验证，然后执行 `sudo systemctl restart nginx` 重启服务。

# q
如何递归地将 2048 目录的所有权变更为 nginx 用户和组？
# a
```
sudo chown -R nginx:nginx /var/www/html/2048
```

