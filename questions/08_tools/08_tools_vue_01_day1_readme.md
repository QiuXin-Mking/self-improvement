# q
使用 Docker 运行 Vue3 开发容器时，如何正确挂载本地项目目录并映射 Vite 端口？
# a
使用 `docker run` 命令，关键参数：
- `-v "本地路径:/app"` 挂载项目目录到容器的 `/app`
- `-p 5173:5173` 映射 Vite 默认端口
- `-w /app` 设置容器工作目录
完整示例（Windows 路径）：
```bash
docker run -itd \
  --name vue3-dev \
  -p 5173:5173 \
  -v "C:\Users\Administrator\Desktop\vue3-demo:/app" \
  -w /app \
  node:18 \
  bash
```
注意镜像版本根据实际需求选择（示例中先用了 node:18，后改用 node:20）。

# q
在 Docker 容器中启动 Vite 开发服务器后，宿主机无法通过 localhost:5173 访问，如何解决？
# a
需要让 Vite 监听所有网络接口，而不仅仅是 localhost。启动时添加 `--host 0.0.0.0`：
```bash
npm run dev -- --host 0.0.0.0
```
或者修改 `package.json` 中的 scripts：
```json
"scripts": {
  "dev": "vite --host 0.0.0.0"
}
```
这样容器内外端口映射才能正常工作，宿主机可通过 `http://localhost:5173` 访问。

# q
如何配置 npm 淘宝镜像源，并使其永久生效？
# a
临时配置：
```bash
npm config set registry https://registry.npmmirror.com/
```
永久生效：执行 `npm config edit`，在打开的配置文件中添加：
```
registry=https://registry.npmmirror.com/
```
验证是否生效：`npm config get registry`。

# q
使用 Vite 创建第一个 Vue 3 项目的完整步骤是什么？
# a
1. 确保 Node.js 环境（建议 Node 18 或以上，文档中最终使用 node:20）
2. 执行创建命令：
```bash
npm create vue@latest
```
3. 根据提示填写项目名称（如 `test1`）
4. 进入项目目录并安装依赖：
```bash
cd test1
npm install
```
5. 启动开发服务器：
```bash
npm run dev
```

