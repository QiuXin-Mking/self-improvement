# q
osc层中 `cio_io_submit()` 方法的核心流程是什么？
# a
`cio_io_submit()` 是 `cl_io_operations` 的一个方法实现，用于osc层。它迭代in-queue中的页面，对每个页面先调用 `cl_page_prep()` 进行IO准备，然后通过 `osc_io_submit_page()` 提交；若页面已提交，则通过 `osc_set_async_flags()` 修改osc标志位。

# q
`osc_set_async_flags` 的作用是什么？
# a
`osc_set_async_flags` 用于修改osc（Object Storage Client）相关的异步标志，当页面已经提交时，通过该函数更新标志而不再重复提交页面。

# q
`cl_page_io_start` 的功能是什么？
# a
`cl_page_io_start` 表示客户端页面IO开始，将页面排入IO队列，并更改其状态为排队等待IO。

