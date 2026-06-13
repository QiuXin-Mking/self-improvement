# Markdown 题目渲染

## 概述

问题和答案内容支持完整的 GFM (GitHub Flavored Markdown) 渲染，学习时以富文本格式展示。

## 实现

- **渲染引擎**：[marked](https://github.com/markedjs/marked) (~8KB gzip)
- **安全策略**：marked v5+ 默认不渲染原始 HTML，天然防 XSS
- **标记规范**：题目内容（`# q` / `# a` 之间）使用标准 Markdown 语法
- **无语法高亮**：代码块使用深色背景 + 等宽字体，不引入 highlight.js

## 支持的语法

| 语法 | 效果 |
|------|------|
| `` `inline code` `` | 行内代码，浅灰底色 |
| ` ```lang ... ``` ` | 围栏代码块，深色背景 |
| `- item` / `1. item` | 无序 / 有序列表 |
| ` \| col \| col \| ` | 表格，有边框和斑马纹 |
| `**bold**` / `*italic*` | 粗体 / 斜体 |
| `> quote` | 引用块，左侧竖线 |
| `---` | 水平分割线 |
| `# h1` ~ `###### h6` | 标题 |
| ASCII 框图 | `pre` 块保留空白，等宽字体对齐 |

## 涉及的代码

| 文件 | 用途 |
|------|------|
| `frontend/src/composables/useMarkdown.ts` | marked 配置和渲染函数 |
| `frontend/src/styles/markdown.scss` | 渲染后 HTML 元素的样式 |
| `frontend/src/components/learning/QuestionCard.vue` | 问题卡片，`v-html` 渲染 |
| `frontend/src/components/learning/AnswerCard.vue` | 答案卡片，`v-html` 渲染 |
| `frontend/src/styles/global.scss` | 等宽字体栈（代码块用） |
