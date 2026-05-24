# q
Go程序的可执行入口必须满足什么条件？
# a
必须包含一个名为 `main` 的包（`package main`），并且在该包中定义一个 `main()` 函数，程序启动时会首先执行该函数（若存在 `init()` 函数则先于 `main()` 执行）。

# q
Go中如何导入并使用外部包提供的功能？以格式化输出为例说明。
# a
使用 `import "包名"` 语句导入，例如 `import "fmt"` 导入格式化IO包。然后通过 `包名.函数名` 调用，如 `fmt.Println("Hello, World!")` 输出内容并自动追加换行符 `\n`。

# q
`fmt.Println` 和 `fmt.Print` 的主要区别是什么？
# a
`println` 输出后会自动追加换行符 `\n`，而 `print` 不会。使用 `fmt.Print("hello, world\n")` 可以得到与 `fmt.Println("hello, world")` 相同的结果。两者都支持直接输出变量，使用默认格式。

# q
Go语言中标识符的大小写如何决定其可见性？
# a
以大字母开头的标识符（如 `Group1`）是可导出的，能被外部包访问（类似 `public`）；以小写字母开头的标识符对包外不可见，但包内可用（类似 `protected`）。

