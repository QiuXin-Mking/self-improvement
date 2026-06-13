当用户说 `/record-feature` 时，按以下流程创建特性记录：

1. 查看 `docs/features/` 下已有最大编号，确定新编号 NNN（3 位补零）
2. 创建 `docs/features/NNN-short-slug.md`，内容模板：

```markdown
# Feature Name

## 概述

## 功能
| 项 | 说明 |
|----|------|

## 涉及的代码
| 文件 | 用途 |
|------|------|
```

3. `git add` 并 commit
