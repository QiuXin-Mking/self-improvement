# q
位域是什么？如何在C语言中定义？
# a
位域是把一个字节中的二进制位划分为几个不同的区域，每个区域有一个域名，允许按域名直接操作这些位，用于节省存储空间。定义方式与结构体类似，在成员名后加冒号和位数，例如：
```c
struct bs {
    int a:8;
    int b:2;
    int c:6;
};
```

# q
位域不能跨字节存储时，如何强制从下一个存储单元开始？
# a
使用长度为0的无名位域（空域）来强制对齐到下一存储单元。例如a占用5位，b需要4位但当前字节剩余3位不够，可以使用 `unsigned :0;` 使b从下一字节开始：
```c
struct bs {
    unsigned a:5;
    unsigned :0; /* 空域，下一单元开始 */
    unsigned b:4;
};
```

# q
无名的位域有什么作用？
# a
无名的位域用来调整位置或占位，不能被访问。例如 `int :2;` 占用2位但不可使用，常用于对齐或填充：
```c
struct k {
    int a:1;
    int :2;  /* 这2位无法使用 */
    int b:3;
    int c:2;
};
```

# q
位域如何访问和赋值？
# a
位域的用法与结构体成员相同，使用 `.` 或 `->` 访问。可以直接赋值，以及使用按位运算操作：
```c
struct bs { unsigned a:1; unsigned b:3; unsigned c:4; } bit, *pbit;
bit.a = 1;
bit.b = 7;
bit.c = 15;
pbit = &bit;
pbit->a = 0;
pbit->b &= 3;
pbit->c |= 15;
```

# q
如何使用指定初始化器定义结构体常量？
# a
使用C99的指定初始化器语法，在初始化时为特定成员赋值，常用于定义常量：
```c
const static conn_id_t INVALID_CONN_ID = {.type = CONN_NONE, .id = 0, .conn_seq = 0};
```

