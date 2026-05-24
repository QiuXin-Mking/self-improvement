# q
vector 的 end() 迭代器指向什么位置？
# a
end() 指向“尾后”位置的迭代器，不指向有效元素，是最后一个元素的下一个位置。

# q
在 vector 中，如何通过 begin() 和 end() 判断容器为空？
# a
空容器时，begin() == end()。

