# q
static_cast 的主要用途和特点是什么？
# a
static_cast 主要用于良性、安全的转换，例如基本数据类型之间的转换、父子类指针间的向上转型。它不会进行运行时类型检查。

# q
dynamic_cast 的特点和使用场景是什么？
# a
dynamic_cast 用于多态类型之间的转换，支持运行时类型检查。常用于将父类指针或引用安全转换为子类指针或引用；转换失败时指针版本返回 nullptr，引用版本抛出异常。

# q
const_cast 的作用是什么？有什么限制？
# a
const_cast 用于增加或去除变量的 const/volatile 限定符。它不能用于不同类型之间的转换，只能修改 cv 限定。

# q
reinterpret_cast 的特点和风险是什么？
# a
reinterpret_cast 用于不同类型指针或引用间的底层强制转换，直接重新解释内存中的二进制数据。它对类型安全没有任何保证，使用时需非常谨慎。

