# 001 — 页面刷新后用户名消失

**状态**：已修复  
**日期**：2026-05-30  
**严重度**：中  

## 现象

Dashboard 页面刷新后，顶部的用户名（`user?.username`）消失。

## 复现

1. 登录 demo 账户
2. 进入 `/dashboard`，看到「欢迎，demo」
3. `Cmd+R` 刷新页面
4. 用户名消失

## 根因

`router/index.ts` 守卫只检查 `localStorage` 中 token 存在就放行，但 Pinia 的 `user` 状态在页面刷新后丢失。Dashboard 显示 `user?.username` 为 `null`。

## 修复

在路由守卫受保护路由分支中，若 `!authStore.user`，先调用 `loadUserProfile()` 从服务端恢复用户信息再放行。

```diff
  if (token) {
-   next()
+   if (!authStore.user) {
+     const ok = await authStore.loadUserProfile()
+     if (!ok) { next('/login'); return }
+   }
+   next()
  }
```

## 验证

```
curl -s http://115.190.235.149:4430/api/profile \
  -H "Authorization: Bearer $TOKEN" | grep '"username":"demo"'
```
