# q
Markdown 中如何创建不同级别的标题？
# a
使用 `#` 号，一个 `#` 表示一级标题，`##` 表示二级标题，依此类推，最多支持六级标题（`######`）。

# q
Markdown 中链接的两种写法是什么？
# a
- 行内式：`[显示文字](URL)`，例如 `[Makedown 的 csdn](https://blog.csdn.net/afei__/article/details/80717153)`
- 参考式：`[显示文字][标识]`，并在文档任意处用 `[标识]: URL` 定义，例如 `[第二种写法][1]` 和 `[1]: https://www.baidu.com/`

# q
如何在 Markdown 表格中控制列的文本对齐方式？
# a
在表格分隔行使用冒号。左对齐 `:---`，右对齐 `---:`，居中对齐 `:---:`。示例：
```markdown
| 商品 | 数量 | 单价 |
| ---- | ----:|:----:|
| 苹果 |   10 | \$1  |
```

# q
如何在 Markdown 中插入带语法高亮的代码块？
# a
使用三个反引号后紧跟语言标识（如 `c`、`java`、`sh`），然后写入代码，最后以三个反引号结束。例如：
```c
#include<cstdio>
using namespace std;  
int main()
{
	return 0 ; 
}
```

# q
Obsidian 中如何实现文本高亮？
# a
使用两个等号包围要突出的文字，例如 `==highlight==`。

