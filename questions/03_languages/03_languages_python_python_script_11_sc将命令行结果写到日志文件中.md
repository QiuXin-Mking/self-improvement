# q
如何用 Python 执行 shell 命令并捕获输出？
# a
使用 `subprocess.check_output`，设置 `shell=True` 以支持 shell 语法，`universal_newlines=True` 使输出以字符串形式返回。  
示例函数：
```python
import subprocess

def run_command(command):
    try:
        result = subprocess.check_output(command, shell=True, universal_newlines=True)
        return result.strip()  # 去除末尾换行符
    except subprocess.CalledProcessError as e:
        print(f"Command execution failed with error: {e}")
        return None
```

# q
如何将命令输出以追加模式写入日志文件？
# a
使用 `open(output_file, "a")` 打开文件，并调用 `file.write(output)` 写入内容。  
示例：
```python
output = run_command("ls -l")
if output is not None:
    with open("output.txt", "a") as file:
        file.write(output)
```

# q
如何捕获并处理 `subprocess.check_output` 执行失败时的异常？
# a
捕获 `subprocess.CalledProcessError` 异常，打印错误信息并返回 `None`，避免程序崩溃。  
示例：
```python
try:
    result = subprocess.check_output(command, shell=True, universal_newlines=True)
except subprocess.CalledProcessError as e:
    print(f"Command execution failed with error: {e}")
    return None
```

