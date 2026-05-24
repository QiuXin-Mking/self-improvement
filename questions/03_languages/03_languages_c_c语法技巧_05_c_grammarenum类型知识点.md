# q
C语言中枚举类型的默认值规则是什么？
# a
默认情况下，枚举的第一个元素值为0，后续每个元素依次在前一个元素的基础上递增1。例如：
```c
enum { A, B, C };
// A = 0, B = 1, C = 2
```

# q
当给枚举的第一个元素指定初始值后，后续元素的值如何确定？
# a
后续元素仍然遵循“在前一个元素基础上递增1”的规则。例如：
```c
enum week { Mon = 1, Tues, Wed, Thurs, Fri, Sat, Sun };
// Tues = 2, Wed = 3, Thurs = 4, Fri = 5, Sat = 6, Sun = 7
```

# q
如何在switch-case结构中配合枚举类型使用？
# a
可以将枚举常量直接作为case的标签，枚举变量通过输入或赋值后与枚举值比较。示例：
```c
enum week { Mon = 1, Tues, Wed, Thurs, Fri, Sat, Sun } day;
scanf("%d", &day);
switch(day) {
    case Mon: puts("Monday"); break;
    case Tues: puts("Tuesday"); break;
    // ... 其他case
    default: puts("Error!");
}
```

# q
如何结合typedef定义枚举类型的别名？
# a
使用`typedef enum { ... } 别名;`语法，可以创建新的类型名，方便后续声明变量。例如：
```c
typedef enum {
    OSD_CREATE,
    OSD_SCAN,
} osd_mgt_msg_type_e;

osd_mgt_msg_type_e msg;
// OSD_CREATE 默认值为0，OSD_SCAN 为1
```

