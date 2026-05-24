# q
这个C语言示例中异步回调是如何实现的？
# a
通过POSIX线程模拟异步操作：主线程使用 `pthread_create` 创建一个新线程，将回调函数作为参数传入；新线程执行模拟的异步操作（如网络请求），操作完成后调用传入的回调函数处理结果；主线程继续执行，最后用 `pthread_join` 等待线程结束。

# q
示例中如何定义和使用回调函数类型？
# a
使用 `typedef` 定义函数指针类型：
```c
typedef void (*CallbackFunc)(int);
```
该类型指向返回 `void`、接受一个 `int` 参数的函数。`fetchDataFromServer` 的参数声明为 `CallbackFunc callback`，在函数内部通过 `callback(data)` 调用。

# q
`pthread_create` 在示例中如何传递回调函数？
# a
```c
pthread_create(&thread, NULL, (void *(*)(void *))fetchDataFromServer, (void *)handleData);
```
将 `handleData` 函数指针强制转换为 `void *` 作为第四个参数传给线程函数，线程函数 `fetchDataFromServer` 的入口被强制转换以适配 `void *(*)(void *)` 原型，内部通过参数获取回调并调用。

# q
主程序如何与异步线程同步？
# a
使用 `pthread_join(thread, NULL)` 阻塞等待异步线程完成，确保主程序在退出前取得异步操作结果。

