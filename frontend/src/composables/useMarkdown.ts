import { marked } from 'marked'

marked.setOptions({
  breaks: true, // 单个换行 = <br>（中文友好）
  gfm: true,    // 表格、任务列表、删除线
})

export function renderMarkdown(text: string): string {
  return marked.parse(text) as string
}
