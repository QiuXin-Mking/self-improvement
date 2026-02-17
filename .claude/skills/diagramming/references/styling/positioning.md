# Positioning and Layout

Coordinates, dimensions, and rotation.

## Properties

| Property | Type   | Description           |
| -------- | ------ | --------------------- |
| `x`      | number | X position (top-left) |
| `y`      | number | Y position (top-left) |
| `width`  | number | Element width         |
| `height` | number | Element height        |
| `angle`  | number | Rotation in radians   |

## Rotation

Angle in radians (not degrees):

```json
{
  "angle": 0.785398
}
```

| Degrees | Radians |
| ------- | ------- |
| 0°      | 0       |
| 45°     | 0.785   |
| 90°     | 1.571   |
| 180°    | 3.142   |
| 270°    | 4.712   |

## Alignment

### Center Calculations

```
centerX = element.x + element.width / 2
centerY = element.y + element.height / 2
```

### Distribution

Space elements evenly by calculating intervals between centers.

## Edge Points

For arrows connecting to shape edges:

| Edge   | X             | Y              |
| ------ | ------------- | -------------- |
| Top    | `x + width/2` | `y`            |
| Bottom | `x + width/2` | `y + height`   |
| Left   | `x`           | `y + height/2` |
| Right  | `x + width`   | `y + height/2` |

## Grid Layout

Example 2x4 grid starting at (100, 100) with 200px horizontal and 130px vertical spacing:

```
     100    300    500    700
100  [A]    [B]    [C]    [D]
230  [E]    [F]    [G]    [H]
```

## Layout Tips

- Pick a consistent flow direction (left-to-right or top-to-bottom)
- Reorder elements to avoid crossing arrows when possible

### Spacing Minimums

| Context                     | Minimum |
| --------------------------- | ------- |
| Between sibling elements    | 40px    |
| Between rows                | 60px    |
| Frame inner padding         | 40px    |
| Arrow clearance from shapes | 20px    |
