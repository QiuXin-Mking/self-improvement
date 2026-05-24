# q
为什么在这个目录编码工具中选择 Pillow 而不是 OpenCV？
# a
OpenCV 的二进制依赖较重，可能导致打包后体积大或运行时缺少 .dll/.so 文件；Pillow 对基础颜色操作和图像生成已足够，且打包更可靠，更易实现“即拷即用”的单文件分发。

# q
这个工具最终使用什么工具打包成单文件跨平台二进制？
# a
使用 PyInstaller 打包为单文件、跨平台的二进制可执行程序，并为 Windows、macOS、Linux 分别生成独立二进制。

# q
自动化构建脚本 build.sh/build.bat 的核心流程是什么？
# a
流程包括：创建虚拟环境 → 安装依赖 → 运行单元测试 → 使用 PyInstaller 打包（可启用 UPX 压缩，排除调试符号和未使用模块）。

# q
该工具推荐的技术栈中，负责纠错码的库是什么？
# a
reedsolo，一个纯 Python 实现的 Reed-Solomon 库，易于集成且无需外部依赖。

