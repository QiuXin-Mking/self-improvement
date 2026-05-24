# q
如何检查树莓派的摄像头是否被系统识别？
# a
使用命令 `vcgencmd get_camera`，输出中 `supported=1` 表示支持，`detected=1` 表示已检测到摄像头。如果 `detected=0`，可能需要检查连接或重启。

# q
在树莓派上查看视频设备节点的命令是什么？
# a
```bash
ls /dev/video*
```
该命令列出所有可用的视频设备文件，例如 `/dev/video0`。

# q
如何使用 fswebcam 在树莓派上捕获图像？
# a
先安装 fswebcam：
```bash
sudo apt-get install fswebcam
```
然后执行：
```bash
fswebcam image.jpg
```
即可拍摄一张照片并保存为 `image.jpg`。

