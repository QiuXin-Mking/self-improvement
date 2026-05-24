# q
如何基于 Docker 创建一个用于编译 Ceph 的 Ubuntu 24.10 容器？
# a
使用以下命令拉取镜像并启动容器，挂载本地目录：
```bash
docker pull ubuntu:24.10
docker run -it --name ceph-compile -v D:\ceph_docker:/mnt/ceph_docker ubuntu:24.10 /bin/sh
```
若需后台运行，可改用：
```bash
docker run -d --name ceph-compile -v D:\ceph_docker:/mnt/ceph_docker ubuntu:24.10 tail -f /dev/null
```
进入容器使用：
```bash
docker exec -it ceph-compile bash
```

# q
如何克隆 Ceph 的指定版本（比如 octopus）源代码？
# a
使用 `git clone` 并指定 `--branch` 参数，例如：
```bash
git clone --branch octopus https://github.com/ceph/ceph.git ceph-octopus
```

# q
编译 Ceph 前如何解决软件依赖？
# a
先执行源码目录下的依赖安装脚本：
```bash
./install-deps.sh
```
若某些依赖未自动安装，可手动用 `apt` 补充，常见缺失包安装命令如：
```bash
apt update && apt install -y cmake
apt update && apt install -y libibverbs-dev librdmacm-dev libudev-dev libblkid-dev libkeyutils-dev libldap2-dev libaio-dev libfuse-dev libleveldb-dev xfslibs-dev pkg-config libsnappy-dev liblz4-dev libcurl4-openssl-dev liboath-dev liblttng-ust-dev libbabeltrace-ctf-dev python3-cython python3-pip
```

# q
Ceph 源码编译的核心步骤是什么？
# a
1. 配置 CMake：
   ```bash
   ./do_cmake.sh
   ```
   或带构建类型参数：
   ```bash
   ARGS="-DCMAKE_BUILD_TYPE=RelWithDebInfo" ./do_cmake.sh
   ```
2. 进入 build 目录编译（如使用 96 核并行）：
   ```bash
   cd build/
   make -j96
   ```

