# q
如何使用 `unison` 命令同步两个目录？
# a
使用 `unison <目录1> <目录2>`，例如：
```
unison /home/ubuntu/cpp_project/05_bill/01快速资源库/01_esp32_code /home/ubuntu/cpp_project/05_bill/2025_04_13_基于单片机曲仓远程系统设计/01_esp32_code
```
该命令会让两个目录的内容保持一致（双向同步）。

# q
如何配置 crontab 每隔 5 分钟执行一次脚本并记录日志？
# a
1. 为脚本添加执行权限：
   ```
   chmod +x /path/to/script.sh
   ```
2. 编辑 crontab：
   ```
   crontab -e
   ```
3. 添加定时任务行：
   ```
   */5 * * * * /path/to/script.sh > /tmp/script.log 2>&1
   ```
   含义：每 5 分钟执行脚本，并将标准输出和标准错误都重定向到 `/tmp/script.log`。

