# q
文档中需要重新制作的Docker镜像有哪些？
# a
需要重新制作的镜像包括：
- `10.5.9.246:10030/base/tensorflow/tensorflow-2.9.3-u20.04:2.9.3-goofysv6`
- `10.5.9.246:10030/base/pytorch/pytorch:2.1.2-cuda11.8-goofys-v3`
- `10.5.9.246:10030/base/ollama/deepseek-r1:14b-goofys-openwebui-v3`

后续传输的镜像版本有 `tensorflow-2.9.3-u20.04:2.9.3-goofysv7`、`pytorch:2.1.2-cuda11.8-goofys-v4` 以及 `ollama/deepseek-r1:14b-goofys-openwebui-v4`。

# q
这些镜像使用的仓库地址和命名结构是什么？
# a
仓库地址为 `10.5.9.246:10030`，镜像命名路径结构为 `/base/<框架>/<基础镜像名>:<标签>`。例如 TensorFlow 镜像的完整路径为 `10.5.9.246:10030/base/tensorflow/tensorflow-2.9.3-u20.04:2.9.3-goofysv6`，其中 `base` 是项目或组织前缀，`tensorflow` 为框架名，`tensorflow-2.9.3-u20.04` 是镜像名称，`2.9.3-goofysv6` 是标签，表示集成了 Goofys 的特定版本。

# q
如何运行这些镜像的容器？
# a
使用 `docker run -it <镜像完整路径> bash` 命令，例如：
```bash
docker run -it 10.5.9.246:10030/base/ollama/deepseek-r1:14b-goofys-openwebui-v4 bash
```
交互式终端会启动并进入容器的 bash 环境，可基于这些镜像进行测试或重新制作。

