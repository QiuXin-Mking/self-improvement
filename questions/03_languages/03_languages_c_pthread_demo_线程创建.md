# q
如何使用 pthread_create 创建新线程？需要传递哪些参数？
# a
使用 `pthread_create` 函数创建新线程：
```c
int pthread_create(pthread_t *thread, const pthread_attr_t *attr,
                   void *(*start_routine)(void *), void *arg);
```
- `thread`：输出参数，保存新线程的ID
- `attr`：线程属性，通常传 `NULL`
- `start_routine`：线程执行的入口函数，必须符合 `void *func(void *)` 签名
- `arg`：传递给入口函数的参数，可以为任意指针
示例：
```c
pthread_t thread_id;
char *str = "hello world";
pthread_create(&thread_id, NULL, workThreadEntry, str);
```

# q
pthread_join 的作用是什么？该示例中如何使用？
# a
`pthread_join` 用于等待指定线程结束并回收其资源，避免僵尸线程。原型：
```c
int pthread_join(pthread_t thread, void **retval);
```
- `thread`：要等待的线程ID
- `retval`：可获取线程返回值，传 `NULL` 表示忽略
示例中主线程调用 `pthread_join(thread_id, NULL);` 阻塞等待子线程 `workThreadEntry` 执行完后再继续。

# q
在 pthread 线程函数中如何获取传入的参数？
# a
线程入口函数固定为 `void *func(void *args)` 接口，创建线程时传入的参数会以 `void*` 形式传递。函数内部需要将 `args` 强制转换回原始类型再使用。例如：
```c
void *workThreadEntry(void *args) {
    char *str = (char*)args;
    printf("threadId:%lu, argv:%s\n", pthread_self(), str);
    return NULL;
}
```

# q
pthread_self() 的作用是什么？如何输出其值？
# a
`pthread_self()` 返回当前调用线程的线程ID，类型为 `pthread_t`。在 Linux 中通常可以用 `%lu` 格式将其转为无符号长整型输出：
```c
printf("threadId=%lu\n", pthread_self());
```
用于在线程内部标识自身或进行线程比较。

