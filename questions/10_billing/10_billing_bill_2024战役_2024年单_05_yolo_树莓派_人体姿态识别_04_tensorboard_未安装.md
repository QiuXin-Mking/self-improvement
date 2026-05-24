# q
在Windows PowerShell中运行`tensorboard --logdir=logs --port=6007`时提示“CommandNotFoundException”，最可能的原因是什么？
# a
最可能的原因是当前Python虚拟环境（venv）或系统环境中未安装TensorBoard，或tensorboard可执行文件不在PATH路径中。可通过`pip install tensorboard`安装，并确保虚拟环境激活后重新执行命令。

