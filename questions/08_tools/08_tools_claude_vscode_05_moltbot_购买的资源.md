# q
阿里云DashScope兼容模式API的base_url是什么？支持哪些模型代码？
# a
base_url 为 `https://dashscope.aliyuncs.com/compatible-mode/v1`，可使用的模型代码如 `qwen-plus`、`qwen3-max`。

# q
如何为clawdbot配置阿里云通义千问API密钥？
# a
在 `~/.bash_profile` 中设置环境变量 `DASHSCOPE_API_KEY`，如：
```bash
echo "export DASHSCOPE_API_KEY='你的API密钥'" >> ~/.bash_profile
source ~/.bash_profile
```
重启终端或执行 `source ~/.bash_profile` 使其生效。

# q
如何启动clawdbot控制面板？如何在远程服务器上通过本地浏览器访问它？
# a
执行 `clawdbot dashboard`（或 `moltbot dashboard`）启动Web控制面板。在远程服务器上，使用SSH本地端口转发：
```bash
ssh -N -L 18789:127.0.0.1:18789 root@服务器IP
```
然后在本地浏览器打开 `http://localhost:18789/?token=...` 访问。

# q
如何使用clawdbot命令行测试对模型的提问？
# a
使用以下命令向主智能体发送消息并获取回复：
```bash
clawdbot agent --agent main --message "你的问题"
```

