# q
如何启动Web版本的Flask服务器？
# a
安装依赖后执行 `python web_app.py`，默认在 `http://0.0.0.0:5000` 启动。

# q
手机如何访问同一局域网内的Web学习系统？
# a
确保手机与电脑连接同一WiFi，在手机浏览器中访问 `http://电脑IP:5000`。电脑IP可通过 `ipconfig`（Windows）或 `ifconfig`/`ip addr`（Mac/Linux）查看。

# q
Web版本学习流程中，用户反馈记忆程度后会发生什么？
# a
用户查看答案后，从“熟练”“一般”“忘记”“完全忘记”中选择反馈，系统会自动进入下一个问题。

# q
Web版本的技术栈是怎样的，数据如何与命令行版共享？
# a
后端使用 Flask，前端使用原生 HTML+CSS+JavaScript（无外部依赖），数据存储于 `data/learning_data.json` 文件，与命令行版本完全共享。

# q
将Web版部署到云服务器时需要注意哪些关键配置？
# a
设置 `host='0.0.0.0'` 以允许外部访问，使用 Gunicorn 等生产级 WSGI 服务器，推荐配置 HTTPS，也可使用 Docker（扩展现有 Dockerfile）进行部署。

