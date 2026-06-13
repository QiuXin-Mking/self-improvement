当用户说 `/record-bug` 时，按以下流程创建 bug 记录：

1. 查看 `docs/bugs/` 下已有最大编号，确定新编号 NNN（3 位补零）
2. 创建 `docs/bugs/NNN-short-slug.md`，内容模板：

```markdown
# Title

## 状态
✅ 已修复 (YYYY-MM-DD)

## 现象

## 原因

## 修复

## 涉及文件
- `path/to/file`
```

3. 更新 `docs/bugs/README.md` 的索引表，添加新行
4. `git add` 并 commit
