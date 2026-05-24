# q
off_t类型在Linux中是什么，主要用途是什么？
# a
off_t类型用于指示文件的偏移量，通常就是`long`类型。在gcc编译时默认为32位整数（`long int`），在64位Linux系统中会被编译为`long long int`（64位整数）。它的定义位于`<unistd.h>`头文件中。

# q
在64位Linux系统中，off_t的具体类型如何受编译宏影响？
# a
off_t的类型定义取决于`__USE_FILE_OFFSET64`宏：
- 若未定义该宏，`off_t`定义为`__off_t`（通常是`long int`）；
- 若定义了该宏，`off_t`定义为`__off64_t`（通常是`long long int`）。
此外，`__USE_LARGEFILE64`宏还会控制`off64_t`类型是否定义。示例代码片段如下：
```c
242 # ifndef __off_t_defined
243 #  ifndef __USE_FILE_OFFSET64
244 typedef __off_t off_t;
245 #  else
246 typedef __off64_t off_t;
247 #  endif
248 #  define __off_t_defined
249 # endif
250 # if defined __USE_LARGEFILE64 && !defined __off64_t_defined
251 typedef __off64_t off64_t;
252 #  define __off64_t_defined
253 # endif
```

