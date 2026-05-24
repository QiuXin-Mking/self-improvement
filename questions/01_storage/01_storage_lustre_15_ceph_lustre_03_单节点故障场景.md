# q
在Lustre中，如何检查所有OSC、MDC和MGC的状态？
# a
使用命令 `lfs check all`，它会列出所有服务（如OST对应的OSC、MDT对应的MDC、MGC等）及其状态，正常时显示 "active"。  
示例输出：
```
nas_test-OST0000-osc-MDT0000 active.
nas_test-OST0000-osc-ffff996372a0f800 active.
nas_test-MDT0000-mdc-ffff996372a0f800 active.
```

# q
当Lustre的单个OST被卸载后，执行 `lfs check all` 对该OST关联的OSC会报告什么错误信息？
# a
会报告两类错误：
- `lfs check: error: check 'nas_test-OST0003-osc-MDT0000': Resource temporarily unavailable (11)`
- `lfs check: error: check 'nas_test-OST0003-osc-ffff996372a0f800': Input/output error (5)`

表示对应的OST已经不可用，客户端无法访问。

# q
Lustre客户端通过哪个组件与OST进行交互，该组件在 `lfs check all` 输出中如何标识？
# a
客户端通过OSC (Object Storage Client) 与OST交互。在 `lfs check all` 输出中，对应条目格式为 `<文件系统名>-OST<编号>-osc-<标识>`，例如 `nas_test-OST0003-osc-MDT0000` 表示与MDT0000关联的OSC。

