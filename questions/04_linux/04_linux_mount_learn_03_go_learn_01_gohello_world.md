# q
一个完整的 Go hello world 程序由哪些关键部分组成？
# a
由以下部分组成：  
- `package main`：声明该文件属于 main 包，可执行程序的入口包  
- `import "fmt"`：导入 fmt 包以使用格式化输出函数  
- `func main() {}`：程序入口函数，无参数无返回值  
- `fmt.Println("hello world")`：调用 fmt 包的 Println 函数向标准输出打印字符串并换行

# q
如何使用 `go build` 编译并运行一个 Go 程序？
# a
1. 编译：执行 `go build hello.go`，会在当前目录生成与源文件同名（去掉 `.go`）的二进制可执行文件（如 `hello`）  
2. 执行：运行生成的可执行文件，如 `./hello`，即可输出结果

# q
如何使用 `go run` 直接运行 Go 源代码？
# a
执行 `go run hello.go`，该命令会在内存中临时编译并直接执行源代码，不会在磁盘上保留生成的可执行文件，运行结束后会输出结果。

