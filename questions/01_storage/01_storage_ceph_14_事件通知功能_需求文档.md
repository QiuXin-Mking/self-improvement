# q
Ceph RGW 桶事件通知功能是什么？
# a
桶事件通知功能允许在 RGW（RADOS Gateway）的存储桶上配置事件通知，当桶内发生 S3 操作（如对象创建、删除等）时，自动向指定的外部端点发送事件消息。目前支持将通知发送到 HTTP、AMQP 0.9.1 和 Kafka 端点。

# q
Ceph 桶事件通知支持哪些目标端点？
# a
当前支持三种目标端点：HTTP(S) 端点、AMQP 0.9.1 协议的消息队列（如 RabbitMQ），以及 Apache Kafka 消息系统。

