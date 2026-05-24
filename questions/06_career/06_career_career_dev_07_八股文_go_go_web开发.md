# q
Controller 在 Web 开发中的主要职责和不做什么？
# a
负责：解析请求参数（路径参数、Query、Body 等）；调用 Service/业务逻辑层；组装并写回响应（成功/失败、JSON 数据、状态码）。不负责：直接连接数据库（属于 Model/Repository），不做复杂业务计算（属于 Service）。

# q
MVC 分层架构中各层（Model、View、Controller）的职责分别是什么？
# a
MVC 是 Model-View-Controller 架构，针对交互型应用设计，职责严格划分：
- Model：处理数据（存储、校验、业务逻辑），如实体类、Service、Mapper/DAO
- View：展示数据，与用户交互，如前端页面（Vue/React）、Thymeleaf 模板
- Controller：接收用户请求，协调 Model 和 View 交互，如 Controller 类

# q
“best effort”（尽力而为）在系统设计中意味着什么？如何理解？
# a
指系统或组件在处理请求/传输数据时，尽最大能力去完成目标，但不做成功、时效、顺序或完整性的承诺。类比普通快递：尽力送，但不保证具体送达时间，遇到天气或交通问题延迟也不负责赔偿。

