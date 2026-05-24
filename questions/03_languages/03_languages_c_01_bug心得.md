# q
多个独立if语句与if-else if结构在处理条件分支时有什么区别？
# a
多个独立if（如 ```if(...)``` 后跟 ```if(...)```）会依次判断所有条件，若都满足则会执行所有对应代码块，可能导致后面的执行覆盖前面的结果。而 if-else if 结构只会执行第一个满足条件的分支，语义是互斥的，能避免意外覆盖，明确优先级。  
例如：
```c
// 错误：serial和devname都不为NULL时，devname分支会覆盖serial分支的结果
if (NULL != serial) { disk = udev_find_disk_by_serial(udev_svc, serial); }
if (NULL != devname) { disk = udev_find_disk_by_devname(udev_svc, devname); }

// 正确：优先使用serial，只有serial为NULL时才使用devname
if (NULL != serial) { disk = udev_find_disk_by_serial(udev_svc, serial); }
else if (NULL != devname) { disk = udev_find_disk_by_devname(udev_svc, devname); }
```

# q
将空指针传递给C标准库函数（如strncpy）会引发什么问题？
# a
会导致段错误（Segmentation fault），程序崩溃退出。  
例如：
```c
char devname[48];
char *dev_name = NULL; 
strncpy(devname, dev_name, sizeof(devname));  // dev_name为NULL，崩溃
```
因此调用库函数前必须检查指针是否为空。

# q
在Python中使用`%d`格式化时，如果变量为None会引发什么错误？如何安全地记录含变量的日志？
# a
会引发`TypeError: %d format: a number is required, not NoneType`错误。  
推荐的安全日志方式是分段输出变量值和类型，避免直接内嵌格式化：
```python
LOG.info("ocache id is start")
LOG.info(ocache_id)
LOG.info(type(ocache_id))
LOG.info("ocache id is end")
```
这样可以清晰看到实际内容与类型，防止因None或类型不匹配导致的格式化异常。

