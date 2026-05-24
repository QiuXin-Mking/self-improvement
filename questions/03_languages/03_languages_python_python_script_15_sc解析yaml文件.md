# q
PyYAML库是什么？如何安装？
# a
PyYAML 是 Python 中用于解析和生成 YAML 数据的第三方库。安装命令：
```bash
pip install pyyaml
```

# q
使用 PyYAML 安全解析 YAML 文件的函数是什么？其核心原理是什么？
# a
使用 `yaml.safe_load()` 函数。它安全地解析 YAML 数据，只加载基本的 Python 对象，防止执行恶意代码，从而避免潜在的安全问题。

# q
yaml.safe_load() 与 yaml.load() 的区别是什么？
# a
- `yaml.safe_load()`：安全加载，仅解析标准 YAML 子集，不执行任意 Python 代码，适合不可信来源的配置文件。
- `yaml.load()`：可加载任意 YAML 标签，可能会执行嵌入的 Python 对象，存在安全风险，仅应在完全信任 YAML 来源时使用。

# q
使用 PyYAML 解析 YAML 文件并访问其数据的基本流程是什么？
# a
```python
import yaml

with open("config.yaml", "r") as yaml_file:
    config_data = yaml.safe_load(yaml_file)

# 解析后的数据为 Python 字典，通过键访问值
server_info = config_data.get("server", {})
host = server_info.get("host")
```

