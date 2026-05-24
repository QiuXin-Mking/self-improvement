# q
在AWS项目中，无服务器组件（Lambda）通过什么机制触发执行？
# a
当新图片上传至指定的S3桶时，通过S3 ObjectCreated事件自动触发Lambda函数。标注函数和缩略图生成器均依赖此事件源。

# q
自动扩展测试需要提供哪些证据来证明扩展出（Scale Out）和扩展入（Scale In）？
# a
需要提交EC2控制台截图、CloudWatch指标、ALB监控数据，证明：
- 扩展出：ASG因负载增加启动新实例。
- 扩展入：负载降低时自动终止实例。
- 负载均衡：ALB成功将请求分发至多个EC2实例。
可使用Apache Bench或Python并发请求工具模拟高流量访问图片列表页。

# q
标注Lambda函数的核心功能是什么？它涉及哪些AWS服务？
# a
标注函数从S3获取图片，调用Gemini API生成描述，并将描述存储至RDS。涉及服务：S3（读取图片）、Gemini API（外部API调用）、RDS（写入标注）。

# q
AWS Learner Lab环境对Lambda有何关键限制？可能导致什么问题？
# a
Lambda并发执行数上限为10。如果错误配置触发器（例如短时间内大量图片上传），容易超出并发限制，导致函数执行被抑制。

# q
在报告评分标准中，无服务器架构图应展示哪些内容？
# a
需要展示Lambda函数、事件源（S3）、外部API（Gemini）、下游目标（RDS/S3）以及相关服务（如EventBridge）。同时需要标注组件间的交互方式。

