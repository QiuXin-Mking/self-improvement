# q
如何查看根目录下各子目录的大小并按大小排序？
# a
使用以下命令：
```bash
sudo du -h --max-depth=1 / | sort -hr | head -n 10
```
- `du -h --max-depth=1 /`：以人类可读格式递归显示根目录下一级子目录的磁盘使用量
- `sort -hr`：按人类可读的数字大小逆序排序
- `head -n 10`：取前10行结果

