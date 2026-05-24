# q
zdev有哪两种类型？
# a
HDD的zdev和cdisk的zdev。

# q
ocache和zdev之间的关系是什么？
# a
ocache管理着HDD的zdev和cdisk的zdev，并且zdev的响应（rsp）会返回给ocache。

# q
哪个函数负责处理来自zdev的响应？
# a
```ocache_io_process_zdev_rsp``` 函数，用于处理来自HDD zdev或cdisk zdev的响应。

