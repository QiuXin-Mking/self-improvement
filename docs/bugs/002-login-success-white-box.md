# 002 — 登录成功后出现白框

**状态**：已修复  
**日期**：2026-05-30  
**严重度**：低  

## 现象

登录（或点击一键体验）成功后，页面跳转前出现一个白框闪烁。

## 复现

1. 打开登录页
2. 点击登录或「🎯 一键体验」
3. 观察页面跳转到 dashboard 之前
4. 看到白色矩形一闪而过

## 根因

`LoginView.vue` 在登录成功后调用了 `showSuccessToast({ position: 'top' })` 紧接着 `router.push('/dashboard')`。Vant Toast 组件生成白色背景的通知条，路由立即切换导致 toast 无法正常展示和消失，残留为白框闪烁。

## 修复

移除登录/体验成功时的 `showSuccessToast` 调用。跳转到 dashboard 本身就是最明确的成功反馈。

```diff
  if (result.success) {
-   showSuccessToast({ message: '登录成功', position: 'top' })
    router.push('/dashboard')
  }
```

## 验证

```
# 登录后应直接跳转 dashboard，无白框
curl -s -X POST http://115.190.235.149:4430/api/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"demo","password":"demo123"}' | grep '"success":true'
```
