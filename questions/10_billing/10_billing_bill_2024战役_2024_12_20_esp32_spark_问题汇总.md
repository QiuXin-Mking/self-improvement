# q
如何在ESP32项目中获取天气服务的API密钥？
# a
天气服务密钥的获取步骤详见文档 `02_获取天气的密钥/readme.md`，通常涉及在天气服务提供商（如和风天气、OpenWeatherMap）的开发者平台注册应用并生成密钥。

# q
如何为ESP32语音功能获取百度语音服务的API？
# a
百度语音服务的API获取流程参考文档 `03_获取百度的语音api/readme.md`，主要步骤包括在百度智能云控制台创建应用，并开通语音识别/合成服务以获取 `AppID`、`API Key` 和 `Secret Key`。

# q
如何在WSL环境中使用idf.py进行ESP32开发？
# a
在WSL中使用idf.py的方法详见 `01_wsl如何使用/idf_如何打开wsl.md`。通常需要先安装ESP-IDF工具链，然后在WSL终端中进入项目目录，运行 `idf.py build`、`idf.py flash` 等命令。确保串口设备已正确映射到WSL。

