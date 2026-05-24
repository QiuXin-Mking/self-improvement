# q
如何下载并安装树莓派1.44寸LCD屏幕所需的bcm2835库？
# a
```bash
wget http://www.airspayce.com/mikem/bcm2835/bcm2835-1.60.tar.gz
tar zxvf bcm2835-1.60.tar.gz
cd bcm2835-1.60/
sudo ./configure && sudo make && sudo make check && sudo make install
```

# q
使用该屏幕时需要安装哪些Python依赖库？
# a
```bash
sudo apt-get update
sudo apt-get install python-pip
sudo pip install RPi.GPIO
sudo pip install spidev
```

# q
如何在树莓派上编译并运行C语言编写的屏幕示例程序？
# a
进入示例代码目录后：
```bash
cd c
make clean
make
sudo ./main
```

# q
为了让1.44寸屏幕的按键正常工作，需要在/boot/config.txt中添加什么配置？
# a
```
gpio=6,19,5,26,13,21,20,16=pu
```
这会将指定GPIO设置为上拉模式。

# q
如何下载并解压屏幕的官方示例代码？
# a
```bash
sudo apt-get install p7zip-full
sudo wget http://www.waveshare.net/w/upload/f/fa/1.44inch-LCD-HAT-Code.7z
sudo 7z x 1.44inch-LCD-HAT-Code.7z
sudo chmod 777 -R 1.44inch-LCD-HAT-Code
```

