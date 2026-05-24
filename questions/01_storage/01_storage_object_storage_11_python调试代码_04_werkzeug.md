# q
`werkzeug.serving` 模块是什么，它的主要用途是什么？
# a
`werkzeug.serving` 是 Werkzeug Web 框架中用于提供开发用简单 HTTP 服务器的模块。通常用于在开发环境中快速启动一个轻量级 WSGI 服务器来测试 Web 应用程序，不适合生产环境。

# q
`run_simple` 函数的作用是什么，它接受哪些参数？
# a
`run_simple` 是 `werkzeug.serving` 模块中的函数，用于快速启动一个 WSGI 服务器。它接受三个基本参数：主机名（如 `'localhost'`）、端口号（如 `5000`）以及 WSGI 应用程序对象。

# q
如何使用 `werkzeug.serving` 启动一个返回 “Hello, World!” 的简单 WSGI 应用？
# a
```python
from werkzeug.serving import run_simple
from werkzeug.wrappers import Request, Response

@Request.application
def application(request):
    return Response('Hello, World!')

if __name__ == '__main__':
    run_simple('localhost', 5000, application)
```
通过 `@Request.application` 装饰器将函数转换为 WSGI 应用，再调用 `run_simple` 传入地址、端口和应用即可启动。

# q
为什么 `werkzeug.serving` 不适合生产环境，应该用什么替代？
# a
`werkzeug.serving` 仅为开发测试设计，缺乏生产环境所需的稳定性、安全性和并发处理能力。生产环境应使用 Gunicorn、uWSGI 等专业 WSGI 服务器。

