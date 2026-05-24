# q
如何在树莓派上安装 BCM2835 库以支持 1.44 寸小屏幕的底层 GPIO 操作？
# a
```bash
wget http://www.airspayce.com/mikem/bcm2835/bcm2835-1.60.tar.gz
tar zxvf bcm2835-1.60.tar.gz
cd bcm2835-1.60/
sudo ./configure && sudo make && sudo make check && sudo make install
```
官网：http://www.airspayce.com/mikem/bcm2835/

# q
树莓派 1.44 寸屏幕示例代码依赖哪些 Python 库，如何安装？
# a
安装 pip 后使用以下命令安装：
```bash
sudo apt-get update
sudo apt-get install python-pip
sudo pip install RPi.GPIO
sudo pip install spidev
```
依赖 `RPi.GPIO` 和 `spidev` 两个库。

# q
如何下载并解压微雪（Waveshare）1.44 寸 LCD HAT 的示例代码？
# a
```bash
sudo apt-get install p7zip-full
sudo wget http://www.waveshare.net/w/upload/f/fa/1.44inch-LCD-HAT-Code.7z
sudo 7z x 1.44inch-LCD-HAT-Code.7z
sudo chmod 777 -R 1.44inch-LCD-HAT-Code
cd 1.44inch-LCD-HAT-Code/RaspberryPi/
```

# q
在树莓派上安装 wiringPi 并使其与 1.44 寸屏幕配合所需的步骤是什么？
# a
安装 wiringPi 并编译 C 示例：
```bash
sudo apt-get install wiringpi
wget https://project-downloads.drogon.net/wiringpi-latest.deb
sudo dpkg -i wiringpi-latest.deb
gpio -v   # 验证安装
```
然后进入示例代码的 C 目录编译运行：
```bash
cd c
make clean
make
sudo ./main
```

# q
为了启用 1.44 寸 LCD 屏幕，需要在 `/boot/config.txt` 中添加什么配置？
# a
使用命令 `sudo nano /boot/config.txt` 编辑文件，添加以下内容以配置 GPIO 上拉：
```
gpio=6,19,5,26,13,21,20,16=pu
```
此行为屏幕正常工作提供必要的引脚上拉设置。

