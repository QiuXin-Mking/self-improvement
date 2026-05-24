# q
如何通过Ollama运行14B的deepseek-r1模型？
# a
使用命令：
```bash
ollama run deepseek-r1:14b
```

# q
在安装open-webui之前需要做什么准备？
# a
需要先将Python升级到3.11版本，才能通过pip安装open-webui。

# q
容器启动时需要如何配置并启动Ollama和open-webui服务？
# a
设置相关环境变量并后台启动两个服务，完整命令如下：
```bash
# 启动Ollama
export OLLAMA_HOST="0.0.0.0:11434"
export OLLAMA_ORIGINS=*
ollama serve > /var/log/ollama.log 2>&1 &

# 启动open-webui
export HF_ENDPOINT=http://0.0.0.0:11434
export OLLAMA_BASE_URL=http://0.0.0.0:11434
open-webui serve --port 3000 > /var/log/open-webui.log 2>&1 &
```

