# q
如何配置CentOS 7使用阿里云yum源？
# a
执行以下命令将仓库配置替换为阿里云镜像：
```bash
sudo curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```

# q
如何清理yum缓存并重新生成缓存？
# a
执行以下两条命令：
```bash
sudo yum clean all
sudo yum makecache
```

# q
如何安装ESP-IDF在CentOS上所需的依赖包？
# a
先更新系统，然后安装必要的软件包：
```bash
sudo yum update
sudo yum install git wget flex bison gperf python3 python3-pip python3-setuptools cmake ninja-build ccache libffi-dev libssl-dev dfu-util
```

# q
如何下载ESP-IDF开发框架？
# a
创建目录并克隆仓库（包含子模块）：
```bash
mkdir -p ~/esp
cd ~/esp
git clone --recursive https://github.com/espressif/esp-idf.git
```

