# Fill and Stroke Styles

Visual styling properties for elements.

## Fill Styles

| Value         | Description           |
| ------------- | --------------------- |
| `solid`       | Solid fill            |
| `hachure`     | Diagonal line pattern |
| `cross-hatch` | Cross pattern         |
| `zigzag`      | Zigzag pattern        |

```json
{
  "fillStyle": "solid",
  "backgroundColor": "#a5d8ff"
}
```

## Stroke Styles

| Value    | Description     |
| -------- | --------------- |
| `solid`  | Continuous line |
| `dashed` | Dash pattern    |
| `dotted` | Dot pattern     |

```json
{
  "strokeStyle": "dashed",
  "strokeColor": "#1971c2"
}
```

## Stroke Width

| Name      | Value | Use Case           |
| --------- | ----- | ------------------ |
| thin      | 1     | Subtle lines       |
| bold      | 2     | Standard (default) |
| extraBold | 4     | Emphasis, headings |

## Roughness

Hand-drawn effect level:

| Value | Name       | Description               |
| ----- | ---------- | ------------------------- |
| 0     | Architect  | Clean, precise lines      |
| 1     | Artist     | Slight hand-drawn effect  |
| 2     | Cartoonist | Exaggerated sketch effect |

```json
{
  "roughness": 1
}
```

Use `roughness: 0` for technical diagrams. Default is 1.

## Roundness

Controls corner rounding:

```json
{
  "roundness": { "type": 3 }
}
```

Set `roundness: null` for sharp corners.

### Roundness Types

| Type | Name                | Default Radius | Used For                   |
| ---- | ------------------- | -------------- | -------------------------- |
| 1    | LEGACY              | -              | Backward compatibility     |
| 2    | PROPORTIONAL_RADIUS | 25% of size    | Lines, arrows, diamonds    |
| 3    | ADAPTIVE_RADIUS     | 32px           | Rectangles, images, frames |

### By Element Type

| Element   | Roundness Type            |
| --------- | ------------------------- |
| Rectangle | 3 (ADAPTIVE)              |
| Diamond   | 2 (PROPORTIONAL)          |
| Ellipse   | 2 (PROPORTIONAL)          |
| Line      | null (sharp) or 2 (round) |
| Arrow     | null (sharp) or 2 (round) |
| Frame     | 3 (ADAPTIVE)              |

## Complete Style Example

```json
{
  "strokeColor": "#1971c2",
  "backgroundColor": "#a5d8ff",
  "fillStyle": "solid",
  "strokeWidth": 2,
  "strokeStyle": "solid",
  "roughness": 1,
  "opacity": 100,
  "roundness": { "type": 3 }
}
```
