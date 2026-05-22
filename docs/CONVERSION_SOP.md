# 知识库 → 间隔重复学习 转换 SOP

## 概述

将 `qiuxin_aliyun_back` 个人存储技术知识库（~2228 个 .md 文件）通过 Claude API 批量转换为 `# q` / `# a` 格式的问答对，导入 `Self-improvement` 间隔重复学习系统。

## 文件清单

| 文件 | 作用 |
|------|------|
| `convert_qa.py` | 主转换脚本 |
| `prompts.py` | 按文档类型的 Prompt 模板 |
| `.progress.json` | 断点续传状态（自动生成） |
| `requirements.txt` | Python 依赖 |
| `.env.example` | API Key 配置模板 |

## 依赖安装

```bash
cd /Users/qiuxin/code/Self-improvement
pip3 install -r requirements.txt
```

## 环境配置

需要以下环境变量（通常已由 Claude Code 自动设置）：

```bash
ANTHROPIC_AUTH_TOKEN=sk-...       # API 密钥
ANTHROPIC_BASE_URL=https://api.deepseek.com/anthropic  # API 代理地址
ANTHROPIC_MODEL=deepseek-v4-pro   # 模型名称
```

或创建 `.env` 文件（从 `.env.example` 复制）。

## 运行全量转换

```bash
cd /Users/qiuxin/code/Self-improvement
python3 convert_qa.py
```

### 运行特性

- **断点续传**：中断后重跑自动跳过已处理文件，从上次断点继续
- **每 10 个文件** 自动保存 checkpoint 到 `.progress.json`
- **API 重试**：遇到 rate limit 自动等待重试（最多 3 次）
- **进度查看**：运行时实时输出当前文件和处理状态

### 预计耗时

- 2228 个文件，每个 ~3-5 秒 API 调用
- 总计约 **3-5 小时**（取决于 API 响应速度）
- 后台任务可能因 timeout 被 kill，直接重跑即可恢复

## 查看进度

```bash
python3 -c "
import json
p = json.load(open('/Users/qiuxin/code/Self-improvement/.progress.json'))
total = len(p['completed'])
done = sum(1 for v in p['completed'].values() if v['status'] == 'done')
print(f'Processed: {total}/{2228} | With Q&A: {done} | Total Q&A: {p.get(\"total_qa\", 0)}')
"
```

## 转换策略

脚本根据文件路径自动选择 Prompt 类型：

| 文档类型 | 识别规则 | Q&A 策略 |
|---------|---------|---------|
| 技术笔记 | 路径含 `01`-`04` 域 | 提取核心概念 |
| 命令参考 | 文件名含 `command`/`命令` | 操作类问答 |
| 问题案例 | 路径含 `problem`/`case`/`问题` | 诊断排查问答 |
| 总结表格 | 路径含 `summary`/`table` | 知识点提取 |
| 通用文档 | 其他 | 灵活提取 |

### 排除规则

- 目录：`99_archive`、`.git`、`.claude`、`.ralph`、`.vscode`
- 文件：`投资.md`、`服务器列表.md`
- 内容过短：< 50 字符的文件跳过
- 内容过长：> 8000 字符自动截断

## 输出结构

```
Self-improvement/questions/
├── 00_summaries/     # _summary_*.md、*_TABLE.md、根目录文件
├── 01_storage/       # Ceph/Lustre/RocksDB...
├── 03_languages/     # C/Python/Go...
├── 05_problems/      # 问题案例
├── 06_career/        # 八股文/面试/算法
├── 08_tools/         # 工具技巧
├── 09_ai/            # AI 相关
├── 10_billing/       # 计费/接单项目
└── 01_调试命令集合.md # 其他根级文件
```

## 导入到 Self-improvement

转换完成后，更新 `question_input` 并导入：

```bash
cd /Users/qiuxin/code/Self-improvement

# 1. 配置扫描路径
echo "questions/" > question_input
echo "questions/" > question_input_linux

# 2. 用 Web 解析器验证（不能用 CLI --init，CLI 解析器有 bug）
cat > verify_parse.go << 'GOEOF'
package main

import (
    "fmt"
    "self-improvement/internal/parser"
)

func main() {
    qp, _ := parser.NewQuestionParser("questions")
    questions, err := qp.ParseAllFiles()
    if err != nil {
        fmt.Printf("ERROR: %v\n", err)
        return
    }
    fmt.Printf("Total Q&A extracted: %d\n", len(questions))
}
GOEOF
go run verify_parse.go
rm verify_parse.go

# 3. 启动 Web 服务后，通过 API 导入：
# POST /api/protected/init（需 JWT Token）
```

## 已知问题

1. **CLI `--init` 有 bug**：CLI 的简化解析器 Q&A 配对逻辑错误，必须用 Web 解析器 
   (`internal/parser/parser.go`) 导入

2. **后台运行被 kill**：长时间运行可能被 timeout 中断，重跑 `python3 convert_qa.py` 即可
   断点续传会自动跳过已处理文件

3. **API 响应格式**：DeepSeek 模型返回 ThinkingBlock + TextBlock，脚本已适配
