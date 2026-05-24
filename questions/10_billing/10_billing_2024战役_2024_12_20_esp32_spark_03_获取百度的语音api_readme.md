# q
在ESP-IDF项目中，如何通过menuconfig配置百度语音服务的Access Token？
# a
运行 `idf.py menuconfig`，在配置界面中找到 `BAIDU AUDIO ACCESS TOKEN` 选项，将从百度语音技术平台获取的 `access_token` 填入该项。

# q
如何获取百度语音技术API的Access Token？
# a
1. 在百度智能云控制台进入“语音技术”产品，创建应用获取 API Key 和 Secret Key。  
2. 通过控制台的 API Explorer 选择“鉴权认证机制”，调用 `oauth/2.0/token` 接口（POST 方法），参数为 `grant_type=client_credentials`、`client_id`（API Key）、`client_secret`（Secret Key），从响应中提取 `access_token`。

# q
千帆ModelBuilder的Access Token应该如何获取并配置到ESP32项目中？
# a
在千帆ModelBuilder平台创建应用获得 AK 和 SK，然后通过鉴权接口 `oauth/2.0/token`（POST）获取 `access_token`。在 `idf.py menuconfig` 中找到 `BAIDUllM ACCESS TOKEN`（或类似千帆LLM Token配置项）并填入该值。

