# q
Controller 在 Web 开发中主要负责什么？它不做什么？
# a
Controller 的职责：
- 解析请求参数（路径参数、Query、Body 等）
- 调用 Service/业务逻辑层
- 组装并写回响应（成功/失败、JSON 数据、状态码）

它不做的事情：
- 不直接连数据库（属于 Model/Repository）
- 不做复杂业务计算（属于 Service）

# q
MVC 分层架构各层的职责是什么？
# a
- **Model（模型）**：处理数据、业务逻辑和存储校验，对应实体类、Service、Mapper/DAO。类比餐厅后厨。
- **View（视图）**：展示数据并与用户交互，对应前端页面或模板（如 Vue/React、Thymeleaf）。类比餐厅餐桌/菜单。
- **Controller（控制器）**：接收用户请求，协调 Model 和 View 的交互，对应 Controller 类。类比餐厅服务员。

# q
“best effort” 在系统设计中的含义是什么？
# a
“best effort”（尽力而为）指系统或组件会尽最大能力完成目标，但**不保证**成功、时效、顺序或完整性。简单说就是“我会拼尽全力，但成不成、何时成不承诺”。例如普通快递尽力送但不承诺具体送达时间，遇到意外可延迟不赔付。

