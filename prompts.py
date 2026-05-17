"""Prompt templates for converting knowledge base content to Q&A pairs."""

SYSTEM_PROMPT = """你是一个技术知识提炼助手。你的任务是将技术文档内容转换为简洁、准确的问答对，用于间隔重复学习系统。

输出格式要求（严格遵守）：
- 每个问答对必须按以下格式输出：
# q
问题内容（一句话，清晰具体）
# a
答案内容（简洁但完整，保留关键技术细节、命令、配置参数）
- `# q` 和 `# a` 必须独占一行，问题和答案内容在下一行
- 代码块使用 ``` 包裹
- 每个文件提取 2-5 个最核心的问答对

不要编造内容。所有信息必须来自源文件。如果文件内容太短或质量太低（纯日志、空文件），返回 SKIP。
"""

TECH_NOTE_PROMPT = """## 文件类型：技术笔记

源文件路径: {file_path}
领域: {domain}

从以下内容提取核心概念和知识点，生成问答对。问题形式：
- "XX是什么 / 如何理解XX"
- "XX的核心原理是什么"
- "XX与YY的区别是什么"
- "XX包含哪些关键组件"

源内容:
{content}"""

COMMAND_REF_PROMPT = """## 文件类型：命令参考

源文件路径: {file_path}
领域: {domain}

将命令和操作步骤整理为操作类问答对。问题形式：
- "如何执行XX操作"
- "XX命令的作用是什么"
- "如何排查XX问题"

保留完整的命令和参数。代码块用 ```bash 包裹。

源内容:
{content}"""

PROBLEM_CASE_PROMPT = """## 文件类型：问题案例

源文件路径: {file_path}
领域: {domain}

将问题排查经验整理为诊断类问答对。问题形式：
- "XX问题的典型根因是什么"
- "如何从日志定位XX问题"
- "解决XX问题的标准流程是什么"

保留关键的日志片段、错误信息和排查步骤。

源内容:
{content}"""

SUMMARY_PROMPT = """## 文件类型：总结/表格

源文件路径: {file_path}
领域: {domain}

从总结性或表格型内容中提取独立的知识点问答对。每行/每节有意义的内容生成一个问答。问题形式：
- "XX的核心要点是什么"
- "XX有哪些关键特征"

源内容:
{content}"""

GENERAL_PROMPT = """## 文件类型：通用文档

源文件路径: {file_path}
领域: {domain}

从以下内容提取最有价值的知识点，生成问答对。根据内容类型灵活选择问题形式。

源内容:
{content}"""

def get_prompt(file_path, domain, content):
    """Return (system_prompt, user_prompt) based on file classification."""
    path_lower = file_path.lower()

    if any(kw in path_lower for kw in ['command', 'commands', '命令']):
        prompt = COMMAND_REF_PROMPT
    elif any(kw in path_lower for kw in ['problem', 'case', '问题', 'cases']):
        prompt = PROBLEM_CASE_PROMPT
    elif any(kw in path_lower for kw in ['summary', 'table', '_summary_', '_table']):
        prompt = SUMMARY_PROMPT
    elif any(domain.startswith(d) for d in ['01', '02', '03', '04']):
        prompt = TECH_NOTE_PROMPT
    else:
        prompt = GENERAL_PROMPT

    user_msg = prompt.format(file_path=file_path, domain=domain, content=content)
    return SYSTEM_PROMPT, user_msg


def extract_domain(file_path):
    """Extract domain from qiuxin_aliyun_back file path.

    e.g. '01_storage/ceph/03_BlueStore.md' -> '01_storage'
         'COMMANDS_TABLE.md' -> '00_summaries'
    """
    parts = file_path.split('/')
    for part in parts:
        if part[0:2].isdigit() and '_' in part:
            return part
    return '00_summaries'
