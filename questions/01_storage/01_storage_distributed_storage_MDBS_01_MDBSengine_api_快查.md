# q
测试打桩模式中如何使用MAKE_FUNC_RETURN宏拦截函数调用并返回预定义值？
# a
使用`MAKE_FUNC_RETURN`宏将目标函数与一个全局调试指针（如`dbg_msg_rsp_alloc_tracing`）关联，并指定返回次数（如1）。
示例：
```c
MAKE_FUNC_RETURN(__msg_rsp_alloc_tracing, dbg_msg_rsp_alloc_tracing, 1);
MAKE_FUNC_RETURN_OK(msg_reply);
```
需先分配对应的结构体内存并设置全局调试指针，测试结束后释放内存并清理指针。

# q
如何打桩`msg_alloc_io_hdr`函数以返回自定义的`osd_normal_rsp_t`结构体？
# a
1. 分配`osd_normal_rsp_t`内存：`osd_normal_rsp_t *osd_nrm_rsp = (osd_normal_rsp_t *)test_malloc(sizeof(osd_normal_rsp_t));`
2. 将全局调试指针`dbg_osd_rsp`指向该结构体。
3. 使用`MAKE_FUNC_RETURN(msg_alloc_io_hdr, dbg_osd_rsp, 1);`拦截函数返回该指针。
4. 可对结构体字段进行断言验证，如`assert_true(osd_nrm_rsp->osd_mgt_type == OSD_CREATE);`。
5. 测试结束将`dbg_osd_rsp`置为NULL并释放内存。

# q
全局集群元数据`g_cluster_md`在测试前需要如何初始化？
# a
通常无需手动分配释放，但使用前必须确保其内部链表已初始化，例如：
```c
INIT_LIST_HEAD(&g_cluster_md.disk_hd);
```

# q
如何验证消息结构体中`from`和`to`字段的双向地址一致性？
# a
通过断言检查`from`与`to`的`type`和`addr`是否相互匹配：
```c
assert_true(temp_msg->b.from.type == msg->b.to.type);
assert_true(temp_msg->b.from.addr == msg->b.to.addr);
assert_true(temp_msg->b.to.type == msg->b.from.type);
assert_true(temp_msg->b.to.addr == msg->b.from.addr);
```

# q
在测试中，`void *`与`struct msg *`的隐式转换有何现象？
# a
在特定测试上下文下，传递`void *`和`struct msg *`给同一函数效果一致，例如：
```c
void fun(struct msg *b) {...}
fun(b); // b为struct msg *
fun(a); // a为void *
```
生产代码建议显式转换以避免警告或类型安全问题。

