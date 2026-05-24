# q
C++ `sort` 函数需要包含哪个头文件？
# a
```cpp
#include <algorithm>
```

# q
C++ `sort` 函数的自定义比较函数语法是什么？
# a
```cpp
sort(first, last, comp);
```
其中 `comp` 是返回 `bool` 的比较函数或可调用对象。

# q
如何编写按年龄升序排序 `Student` 对象的比较函数 `compareByAge`？
# a
```cpp
bool compareByAge(const Student &a, const Student &b) {
    return a.age < b.age;
}
```
然后使用 `sort(students.begin(), students.end(), compareByAge);`

# q
如何使用 lambda 表达式按分数降序排序 `Student` 对象？
# a
```cpp
sort(students.begin(), students.end(), 
     [](const Student &a, const Student &b) {
         return a.score > b.score;
     });
```

