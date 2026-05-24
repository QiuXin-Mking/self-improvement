# q
如何将Linux会话的编码临时改为en_US.UTF-8？
# a
依次执行以下命令：
```bash
locale-gen en_US.UTF-8
update-locale LANG=en_US.UTF-8
export LANG=en_US.UTF-8
export LANGUAGE=en_US:en
export LC_ALL=en_US.UTF-8
```

# q
如何将Linux会话的编码临时改为zh_CN.GBK？
# a
依次执行以下命令：
```bash
locale-gen zh_CN.GBK
update-locale LANG=zh_CN.GBK
export LANG=zh_CN.GBK
export LANGUAGE=zh_CN:zh
export LC_ALL=zh_CN.GBK
```

# q
在Dockerfile中如何配置系统默认编码为en_US.UTF-8？
# a
需要安装语言包、生成locale并设置环境变量，关键配置片段如下：
```bash
RUN apt-get update && apt-get install -y locales && \
    locale-gen en_US.UTF-8 && \
    update-locale LANG=en_US.UTF-8

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8
```

