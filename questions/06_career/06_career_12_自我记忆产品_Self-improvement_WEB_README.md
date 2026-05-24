# q
如何启动Web版本的学习系统？
# a
运行命令
```bash
python web_app.py
```
默认会在 `http://0.0.0.0:5000` 启动服务器。

# q
如何在手机上访问Web学习系统？
# a
确保手机和电脑在同一局域网（WiFi）下，在手机浏览器中访问 `http://你的电脑IP地址:5000`。

# q
如何查看当前电脑的IP地址？
# a
- Windows：在命令提示符中输入 `ipconfig`，查看 IPv4 地址
- Mac/Linux：在终端中输入 `ifconfig` 或 `ip addr`，查看 IP 地址

# q
在Web学习流程中，反馈记忆程度有哪些选项？
# a
四个选项：
- ✅ 熟练 - 记得很清楚
- 👍 一般 - 记得但不熟练
- 😐 忘记 - 忘记了部分内容
- ❌ 完全忘记 - 完全不记得

# q
Web版本和命令行版本是否共享学习数据？
# a
是的，两者共享 `data/learning_data.json` 文件，可以混用。

