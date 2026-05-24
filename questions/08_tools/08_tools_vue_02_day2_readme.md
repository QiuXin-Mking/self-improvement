# q
如何启动一个已有的 Vue 3 开发环境并访问应用？
# a
1. 进入容器：`docker exec -it vue3-dev bash`
2. 进入项目目录：`cd demo1/test1/`
3. 安装依赖（如需要）：`npm install`
4. 启动开发服务器：`npm run dev`
5. 浏览器访问：`http://localhost:5173/`

# q
当别人的 Vue 项目运行不起来时，推荐的解决方案是什么？
# a
删除 `node_modules` 目录，然后重新执行 `npm install` 安装依赖。

# q
在 Vue 3 项目中，`public` 目录和 `src` 目录分别存放什么内容？
# a
- `public`：存放静态资源，如 `favicon.ico`（网站角标）等文件。
- `src`：存放源码，包括 TypeScript（`.ts`）等业务代码。

# q
Vue 3 官方推荐的编辑器插件是什么？为什么不再使用 Volar？
# a
官方推荐使用 VS Code + Vue Official 插件。不再使用 Volar 是因为它已停止维护。`Vue Official` 是官方后续开发的替代品。

# q
`index.html` 在 Vue 3 项目中的作用是什么？访问 `http://localhost:5173/index.html` 和 `http://localhost:5173/` 有何关系？
# a
`index.html` 是应用入口文件。开发服务器将根路径 `/` 映射到 `index.html`，因此访问这两个地址效果等价。

