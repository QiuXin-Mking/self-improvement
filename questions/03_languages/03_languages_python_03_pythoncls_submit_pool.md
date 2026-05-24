# q
cls_submit_pool 在代码示例中的核心作用是什么？
# a
将指定函数（如 `sub_cmd`、`verify_rpm`、`unpack_rpm`、`run_playbook`）及其参数提交到线程池中执行，实现并发任务调度。

# q
调用 cls_submit_pool 时，入参的顺序规则是什么？
# a
第一个入参是待执行的函数，后续入参依次为该函数所需的参数。例如 `cls_submit_pool(self.verify_rpm, file_path)` 中 `verify_rpm` 是函数，`file_path` 是其参数。

