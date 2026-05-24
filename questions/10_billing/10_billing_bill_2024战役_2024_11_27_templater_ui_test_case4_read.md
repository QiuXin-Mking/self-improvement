# q
摇杆组件的基本 HTML 和 CSS 结构是什么？
# a
HTML 需要一个圆形容器和可拖动的摇杆元素：
```html
<div id="joystick-container">
    <div id="joystick"></div>
</div>
```
CSS 设置容器为圆形（`border-radius: 50%`）、固定尺寸，摇杆使用绝对定位并居中于容器：
```css
#joystick-container {
    position: relative;
    width: 200px;
    height: 200px;
    background-color: #ccc;
    border-radius: 50%;
}
#joystick {
    position: absolute;
    width: 50px;
    height: 50px;
    background-color: #555;
    border-radius: 50%;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    cursor: pointer;
}
```

# q
如何通过 JavaScript 限制摇杆拖拽时不能超出圆形区域？
# a
计算容器半径减去摇杆半径得到最大允许移动距离 `maxMove`。在 `mousemove` 中先求出鼠标相对于容器中心的偏移 `x, y`，计算距离 `distance`；若 `distance > maxMove`，则通过角度将坐标约束在边界上：
```javascript
const maxMove = containerRect.width / 2 - joystick.offsetWidth / 2;
// ... 在 mousemove 中
let x = event.clientX - containerRect.left - containerRect.width / 2;
let y = event.clientY - containerRect.top - containerRect.height / 2;
const distance = Math.sqrt(x * x + y * y);
if (distance > maxMove) {
    const angle = Math.atan2(y, x);
    x = maxMove * Math.cos(angle);
    y = maxMove * Math.sin(angle);
}
joystick.style.transform = `translate(${x}px, ${y}px)`;
```

