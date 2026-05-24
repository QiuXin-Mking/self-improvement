# q
如何通过静态资源路径判断一个网站使用了 Next.js 框架？
# a
如果页面中 JS 文件的路径包含 `/_next/static/` 等标志性前缀，通常表明该网站使用了 Next.js。此外，URL 结构和页面服务端渲染（SSR）特征也是辅助判断依据。

# q
类似 LanguageLeap 的文档/知识库网站，有哪些开箱即用的开源项目推荐？
# a
推荐的轻量级或功能齐全的开源项目包括：
- **Nextra**：基于 Next.js + MDX，超轻量级文档引擎，支持多语言和全文搜索。
- **Docusaurus**：React 驱动的文档站点生成器，适合知识库/教程，支持插件和 PWA。
- **Outline**：知识库协作工具，React + Node.js (Koa) + PostgreSQL，支持账号、目录、标签、搜索。
- **Wiki.js**：专业 wiki 系统，支持多前端渲染，用户管理功能全面。
- **Docsium**：基于 React 和 Next.js 的文档库工具。
- **PaperMod**：Hugo 主题，适合纯静态 Markdown 站点。
- **思源笔记**：支持私有部署的 Notion 替代型笔记工具。

# q
如果需要用户注册、收藏和内容管理功能，应优先选择哪些开源知识库项目？
# a
应优先考虑 **Outline** 或 **Wiki.js**。Outline 提供完整的账号体系、收藏、多人协作、全文搜索和管理后台；Wiki.js 则提供更专业的用户管理和权限控制，二者均适合需要动态交互的知识库系统。

