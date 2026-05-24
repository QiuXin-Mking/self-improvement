# q
如何通过pprof查看Go程序中所有的goroutine堆栈信息？
# a
访问 `http://<ip>:<port>/debug/pprof/goroutine?debug=2` 端点，返回完整的goroutine堆栈（包括所有goroutine的详细信息）。例如：
```
172.22.6.71:9988/debug/pprof/goroutine?debug=2
```

# q
pprof端点中 `debug=2` 参数的作用是什么？
# a
`debug=2` 返回所有goroutine的完整堆栈信息（文本格式），便于直接查看每个goroutine的状态和调用链，适合通过浏览器或curl快速诊断并发问题。

