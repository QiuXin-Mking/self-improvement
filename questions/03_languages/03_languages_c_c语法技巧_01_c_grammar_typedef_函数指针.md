# q
如何使用typedef声明一个函数指针类型？
# a
使用语法 `typedef 返回类型 (*新类型名)(参数列表);`。例如：
```c
typedef void (*func)(void);
```
这样就定义了一个名为 `func` 的类型，它指向无参数、无返回值的函数。

# q
typedef定义的函数指针如何用于实现任务调度/多态？
# a
通过函数指针数组可以构建通用调度器，根据索引调用不同实现。例：
```c
typedef void (*func)(void);

void Task1(void) { printf("I'm Task1.\n"); }
void Task2(void) { printf("I'm Task2.\n"); }
void Task3(void) { printf("I'm Task3.\n"); }

void gTask(char i){
    func vTask[3] = {&Task1, &Task2, &Task3};
    func fun = vTask[i];
    (*fun)();
}
```
调用 `gTask(0)` 输出 `I'm Task1.`，实现了类似接口的多态调用。

# q
在实际项目中（如socket库），typedef函数指针的典型用途是什么？
# a
用于声明回调函数类型，解耦模块。例如：
```c
typedef void (*zsocket_recv_fn_t)(zsocket_t _sock, msg_t msg);
typedef void (*zsocket_err_fn_t)(zsocket_t sock, sock_type_e sock_type);
typedef void (*zsocket_accept_fn_t)(zsocket_t sock, zsocket_t new_sock);
typedef void (*zsocket_connect_fn_t)(zsocket_t sock, int32_t err);
```
这些类型可以注册回调函数，当socket收到数据、出错、接受新连接或连接完成时，通过函数指针调用用户提供的处理函数。

