# q
libcimbar 的典型传输速率和单帧容量是多少？使用 zstd 压缩后文件大小上限在不同方案下分别是多少？
# a
速率约 850kbps，单帧约 7.5KB。方案 A（HTML/WASM）压缩后文件上限约 10MB，方案 B（源码编译）可达 33MB。

# q
方案 A（HTML/WASM）和方案 B（源码编译）在文件上限、使用方式和准备要求上的核心区别是什么？
# a
方案 A 只需一个 `cimbar_js.html` 文件，浏览器打开即用，零依赖，但文件大小限制约 10MB；方案 B 需离线下载源码 ZIP 和所有依赖 deb 包，在内网编译出 CLI 工具 `cimbar`，可编码为静态 PNG，压缩后文件上限可达 33MB。

# q
在内网 Linux 完全离线且没有安装图形浏览器的情况下，如何通过方案 A 使用 cimbar？
# a
在外网下载 Firefox 免安装包（如 `firefox-137.0.2.tar.xz`）放入 U 盘；内网解压后用 `~/firefox/firefox --no-sandbox ~/cimbar_js.html` 命令打开 HTML 文件，无需安装或 root 权限即可使用。

# q
要离线编译 libcimbar，外网机器上应如何一次性下载所有依赖的 deb 包？
# a
在外网与内网相同 Ubuntu 版本的机器上执行：
```bash
mkdir -p ~/cimbar_debs && cd ~/cimbar_debs
sudo apt install --download-only -y build-essential cmake pkg-config libopencv-dev libglfw3-dev libgles2-mesa-dev libzstd-dev
cp /var/cache/apt/archives/*.deb ~/cimbar_debs/
```
若缓存中缺少已安装包的 deb，可用 `for pkg in ...; do apt download $(apt-cache depends --recurse --no-recommends "$pkg" | grep "^\w" | sort -u); done` 递归下载。

# q
使用 CFC App 扫描屏幕条码时，出现识别不到或解码失败的常见原因及对应处理方法是什么？
# a
- 闪退：APK 仅支持 arm64，确认手机是 arm64 架构。
- 识别不到：距离/角度不当或强光反光，应正对屏幕 30-50cm，避开强光。
- 颜色偏差：屏幕开启了护眼模式/夜间模式，需关闭这两种模式。
- 解码失败：PNG 被微信/相册等二次压缩，应避免经过这些应用。
- 文件不完整：漏扫某些帧，需重新扫描遗漏的帧。

