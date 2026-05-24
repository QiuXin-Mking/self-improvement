# q
什么是回调函数？它的核心原理是什么？
# a
回调函数是通过函数指针调用的函数。将函数指针作为参数传递给另一个函数，当这个指针被用来调用其所指向的函数时，就实现了回调。其核心原理是函数指针的传递和调用。

# q
如何用typedef定义函数指针类型？请给出示例。
# a
使用typedef关键字定义函数指针类型，语法为 `typedef 返回值类型 (*类型名)(参数列表);`。例如：
```c
typedef int (*Fun1)(int);           // 参数一个int，返回int
typedef void (*msg_callback_t)(service_t svc, msg_t rsp, void* context); // 复杂参数
```
之后就可以用 `Fun1` 等作为类型名声明变量。

# q
在C语言中，如何通过函数指针实现回调功能？
# a
首先定义一个函数指针类型，然后编写一个接受该函数指针作为参数的函数（如 callFun），在该函数内部通过指针调用传入的函数。调用时将具体函数名（函数地址）作为参数传递。例如：
```c
void callFun(FunType fp, int x) {
    fp(x);  // 执行回调
}
callFun(myFun, 100);
```

# q
在嵌入式开发中，回调函数常与状态机结合使用。请根据M26示例说明其应用方式。
# a
通过结构体将状态值和对应的处理函数绑定，形成状态集合表。例如：
```c
typedef struct {
    uint8_t mStatus;
    uint8_t (* Funtion)(void);
} M26_WorkStatus_TypeDef;
M26_WorkStatus_TypeDef M26_WorkStatus_Tab[] = {
    {GPRS_NETWORK_CLOSE, M26_PWRKEY_Off},
    ...
};
```
状态机函数遍历表，根据当前状态调用对应的回调函数：`return M26_WorkStatus_Tab[i].Funtion();`，实现了基于回调的状态处理。

