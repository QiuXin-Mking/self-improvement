# q
在MDBS Flask框架中，API类名与Client类名是否必须相同？
# a
不是必须相同，两者的类名可以独立命名。

# q
MDBS Flask中API类和Client类分别继承什么基类？
# a
API类继承 `Resource`，Client类继承 `NodeAPIObject`。

# q
在Flask应用中，如何将URL路径和API处理类关联起来？
# a
通过 `api.add_resource(ClassName, '/urlpath')` 注册端点，例如 `api.add_resource(Osd_total_mgt, "/osds/total")`，其中 `Osd_total_mgt` 是类名，`"/osds/total"` 是URL路径。

