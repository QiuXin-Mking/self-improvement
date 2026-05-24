# q
在树莓派上如何创建并激活一个 Python 虚拟环境？
# a
首先安装 python3-venv 包：`sudo apt install python3-venv`。  
然后创建虚拟环境：`python3 -m venv myenv`（myenv 为环境目录名）。  
激活虚拟环境：`source myenv/bin/activate`。  
在激活后的环境中可使用 pip 安装包，例如 `pip install requests`。

# q
如何使用国内镜像源安装 opencv-contrib-python？
# a
使用豆瓣源示例：
```
pip install opencv-contrib-python -i http://pypi.douban.com/simple --trusted-host pypi.douban.com
```
使用阿里云源示例：
```
pip install opencv-contrib-python -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
```

# q
列出几个常用的 pip 国内镜像源地址。
# a
清华大学：`https://pypi.tuna.tsinghua.edu.cn/simple`  
中国科学技术大学：`https://pypi.mirrors.ustc.edu.cn/simple`  
豆瓣：`http://pypi.douban.com/simple/`  
阿里云：`http://mirrors.aliyun.com/pypi/simple/`

