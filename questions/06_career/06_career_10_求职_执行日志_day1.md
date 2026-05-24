# q
在设计呼叫中心（Call Center）系统时，可以使用什么方式表示固定的角色和状态？
# a
使用枚举类（Enum）来定义人员角色（如主管、操作员、主任）、呼叫状态（call_status）和呼叫阶段（call_state），以及等级（Rank）等固定常量。

# q
在 system-design-primer 的 call_center 实践中，需要重点抽象出哪些核心概念？
# a
重点抽象出人员类别（如主管、操作员、主任）和呼叫（Call）动作，将不同的角色与呼叫处理流程封装为可扩展的类。

