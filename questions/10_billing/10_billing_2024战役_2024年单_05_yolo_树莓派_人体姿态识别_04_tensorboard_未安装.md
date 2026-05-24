# q
在 Windows PowerShell 虚拟环境中运行 `tensorboard --logdir=logs --port=6007` 时，出现 “无法将'tensorboard'项识别为 cmdlet” 错误，最可能的原因是什么？
# a
tensorboard 未安装在当前虚拟环境中，或虚拟环境未正确激活，导致命令不可用。

# q
上述错误信息中出现的完整命令是什么？
# a
```powershell
tensorboard --logdir=logs --port=6007
```

