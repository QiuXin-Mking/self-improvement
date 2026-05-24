# q
在Ubuntu上安装OpenCV前，如何更新系统并安装编译所需的基本工具？
# a
```bash
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install build-essential cmake git
```

# q
从源码编译OpenCV需要安装哪些图像和视频处理相关的依赖库？
# a
```bash
sudo apt-get install libjpeg-dev libtiff5-dev libjasper-dev libpng-dev libwebp-dev libopenexr-dev
sudo apt-get install libavcodec-dev libavformat-dev libswscale-dev libv4l-dev libxvidcore-dev libx264-dev
sudo apt-get install libgtk-3-dev
sudo apt-get install libatlas-base-dev gfortran
```
（如果`libjasper-dev`安装失败，可将其从命令中移除后重试）

# q
编译OpenCV时，如何使用cmake配置生成pkg-config文件并指定安装路径？
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
如何用多核并行编译并安装OpenCV？
# a
```bash
make -j4
sudo make install
```
其中`-j4`使用4个核心并行编译，可加快编译速度。

