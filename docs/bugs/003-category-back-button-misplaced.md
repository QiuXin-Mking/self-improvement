# 分类选择页底部返回按钮不协调

## 状态

✅ 已修复 (2026-06-13)

## 现象

分类选择页（CategorySelectView）底部有一个孤零零的 `[返回首页]` 小按钮，与整体 UI 风格不搭。其他页面（LearningView、DashboardView）的导航都在顶部，只有分类页把返回按钮放在最下面，像临时拼凑的。

## 原因

CategorySelectView 用 `<van-button type="default" size="small">` 在底部居中，没有顶部导航栏，和其他页面的模式不一致。

## 修复

将底部按钮替换为顶部导航栏 `← 首页`，匹配 LearningView `✕ 退出` 的风格：

```diff
+ <div class="category-top-bar">
+   <button class="btn-back" @click="goBack">← 首页</button>
+ </div>

- <div class="back-link">
-   <van-button type="default" size="small" @click="goBack">返回首页</van-button>
- </div>
```

## 涉及文件

- `frontend/src/views/CategorySelectView.vue`
