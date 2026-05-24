# q
什么是Go语言中的goroutine？
# a
goroutine 是 Go 语言中的轻量级线程实现，由 Go 运行时（runtime）管理。Go 程序会智能地将 goroutine 中的任务合理地分配给每个 CPU，实现并发执行，开发者无需手动维护线程池。

# q
goroutine 解决了传统线程池方案的什么问题？
# a
传统线程池需要开发者在 Socket 网络编程等场景中手动维护线程与 CPU 数量的对应关系，以避免频繁线程切换损失效率。goroutine 让使用者只需分配足够多的任务，系统自动将任务分配到 CPU 上并发运行，更直观方便。

# q
Go 程序启动时默认的 goroutine 是如何创建的？
# a
Go 程序从 `main` 包的 `main()` 函数开始执行，在程序启动时，Go 运行时会自动为 `main()` 函数创建一个默认的 goroutine。

