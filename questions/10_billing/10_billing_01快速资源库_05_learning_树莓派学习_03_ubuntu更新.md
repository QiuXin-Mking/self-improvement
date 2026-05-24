# q
在 Ubuntu 上更新系统软件包列表并安装基本开发工具的命令是什么？
# a
```bash
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install build-essential cmake git
```

# q
安装 OpenCV 编译所需图像和视频处理依赖项的命令是什么？
# a
```bash
sudo apt-get install libjpeg-dev libtiff5-dev libjasper-dev libpng-dev libwebp-dev libopenexr-dev
sudo apt-get install libavcodec-dev libavformat-dev libswscale-dev libv4l-dev libxvidcore-dev libx264-dev
sudo apt-get install libgtk-3-dev
sudo apt-get install libatlas-base-dev gfortran
```
若 `libjasper-dev` 报错，删除该包后重新执行第一条命令。

# q
如何通过 cmake 配置 OpenCV 编译选项（生成调试符号、启用示例等）？
# a
```bash
cmake -D CMAKE_BUILD_TYPE=RELEASE \
      -D CMAKE_INSTALL_PREFIX=/usr/local \
      -D INSTALL_C_EXAMPLES=ON \
      -D INSTALL_PYTHON_EXAMPLES=ON \
      -D OPENCV_GENERATE_PKGCONFIG=ON \
      -D BUILD_EXAMPLES=ON ..
```

# q
编译并安装 OpenCV 的命令是什么？
# a
```bash
make -j4      # -j4 表示使用 4 个核心并行编译，可加快速度
sudo make install
```

