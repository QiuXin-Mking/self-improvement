# q
在 VS Code 中使用 clangd 格式化 C++ 代码时，如何将 tab 缩进宽度改为 4 个空格？
# a
在 VS Code 工作区根目录创建 `.clangd` 配置文件，添加以下内容：
```yaml
IndentWidth: 4
TabWidth: 4
UseTab: Never
```
并可根据团队习惯指定基础风格，如 `BasedOnStyle: Microsoft`。

