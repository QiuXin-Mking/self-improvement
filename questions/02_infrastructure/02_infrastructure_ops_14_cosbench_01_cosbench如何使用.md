# q
如何解压COSBench安装包？
# a
使用命令 `tar -zxvf cosbench-0.4.2.9.c4-linux.tar.gz` 解压，然后进入解压后的目录 `cd cosbench-0.4.2.9.c4-linux`。

# q
如何启动COSBench的控制器和驱动器？
# a
在解压后的目录下，分别执行 `./start-controller.sh` 和 `./start-driver.sh` 脚本。

# q
如何停止COSBench的控制器和驱动器？
# a
执行 `./stop-driver.sh` 停止驱动器，执行 `./stop-controller.sh` 停止控制器。

# q
如何检查COSBench控制器或驱动器进程是否在运行？
# a
使用 `ps aux | grep controller` 检查控制器进程，使用 `ps aux | grep driver` 检查驱动器进程。

# q
如何检查COSBench的默认端口监听状态？
# a
使用命令 `netstat -tuln | grep 19088` 检查19088端口的监听状态。

