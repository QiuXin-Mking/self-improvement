# q
什么是事件循环？
# a
事件循环是 JavaScript 运行时处理异步操作的机制，通过调用栈、任务队列和微任务队列协调执行。

# q
宏任务和微任务有什么区别？
# a
宏任务包括 setTimeout、setInterval、I/O；微任务包括 Promise.then、MutationObserver。微任务优先级更高，会在当前宏任务结束后立即执行。
