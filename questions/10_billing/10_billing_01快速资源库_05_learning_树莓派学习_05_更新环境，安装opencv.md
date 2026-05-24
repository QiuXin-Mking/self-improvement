# q
在树莓派上如何安装 python3-pip 和 python3-venv？
# a
使用以下命令：
```bash
sudo apt install python3-pip
sudo apt install python3-venv
```

# q
如何创建并激活 Python 虚拟环境？
# a
```bash
# 创建虚拟环境
python3 -m venv myenv
# 激活虚拟环境
source myenv/bin/activate
```

# q
如何在虚拟环境中安装 opencv-contrib-python？
# a
激活虚拟环境后使用 pip 安装，例如：
```bash
pip install opencv-contrib-python
```

# q
如何使用国内镜像源加速 pip 安装 opencv-contrib-python？
# a
指定镜像源安装：
- 豆瓣源：
  ```bash
  pip install opencv-contrib-python -i http://pypi.douban.com/simple --trusted-host pypi.douban.com
  ```
- 阿里云源：
  ```bash
  pip install opencv-contrib-python -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
  ```

# q
常见的 pip 国内基础镜像源有哪些？
# a
- 清华大学：https://pypi.tuna.tsinghua.edu.cn/simple
- 中国科学技术大学：https://pypi.mirrors.ustc.edu.cn/simple
- 豆瓣：http://pypi.douban.com/simple/
- 阿里云：http://mirrors.aliyun.com/pypi/simple/

