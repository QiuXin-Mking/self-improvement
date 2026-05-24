# q
C++中的四种类型转换有哪些
# a
C++中的四种类型转换有：static_cast、dynamic_cast、const_cast 和 reinterpret_cast

- **static_cast**：用于良性、安全的转换，如基本数据类型之间的转换、父子类指针间向上转型，不会进行运行时类型检查。
- **dynamic_cast**：用于多态类型之间的转换，支持运行时类型检查，常用于父类指针/引用向子类安全转换，转换失败会返回nullptr或抛异常。
- **const_cast**：用于增加或去除变量的const/volatile限定，不能用于不同类型之间的转换。
- **reinterpret_cast**：用于不同类型指针或引用间的底层强制转换，直接重新解释内存，对类型安全没任何保证，使用时需谨慎。

