# q
COSBench工作负载XML配置中的五个标准阶段（workstage）分别是什么，各自执行什么任务？
# a
五个阶段及其作用：
- **init**：初始化容器，为测试做准备。
- **prepare**：准备测试数据，按配置创建指定大小和数量的对象。
- **main**：执行主要负载操作（如读/写），通常定义运行时长和操作比例。
- **cleanup**：清理测试过程中产生的对象。
- **dispose**：删除容器，完成环境回收。

# q
在COSBench配置中，`containers=r(1,16)` 里的 `r` 表示什么含义？
# a
`r` 表示随机选择指定区间内的值。例如 `containers=r(1,16)` 表示从1到16号容器中**随机**选取进行操作。

# q
如何用COSBench的`<storage>`元素配置S3存储连接？
# a
通过 `<storage>` 元素指定存储类型和连接参数，示例：
```xml
<storage type="s3" config="accesskey=your_access_key;secretkey=your_secret_key;endpoint=https://your.endpoint:port" />
```
必须包含 `type="s3"`，并在 `config` 中以分号分隔设置 `accesskey`、`secretkey` 和 `endpoint`。

