# q
在树莓派上使用哪个命令可以检查摄像头的检测状态？
# a
使用 `vcgencmd get_camera` 命令，输出中 `supported=1` 表示支持摄像头，`detected=1` 表示已检测到摄像头。

# q
当 `vcgencmd get_camera` 显示 `supported=1 detected=0` 时，最简单的解决办法是什么？
# a
执行 `sudo reboot` 重启系统，重新检查摄像头连接和驱动加载。

# q
如何在命令行用 fswebcam 拍摄一张图片？
# a
先安装 fswebcam：
```bash
sudo apt-get install fswebcam
```
然后直接执行：
```bash
fswebcam image.jpg
```
即可拍摄一张图片并保存为 `image.jpg`。

