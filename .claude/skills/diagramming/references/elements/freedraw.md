# Freedraw Elements

Hand-drawn paths with optional pressure sensitivity.

## Freedraw Element

```json
{
  "type": "freedraw",
  "points": [
    [0, 0],
    [10, 5],
    [20, 3]
  ],
  "pressures": [0.5, 0.7, 0.6],
  "simulatePressure": true
}
```

## Properties

| Property           | Type    | Description                         |
| ------------------ | ------- | ----------------------------------- |
| `points`           | array   | Path coordinates as `[x, y]` pairs  |
| `pressures`        | array   | Pressure values (0.0-1.0) per point |
| `simulatePressure` | boolean | Estimate pressure from stroke speed |

## Points

Points are relative to element `x,y`:

```json
{
  "x": 100,
  "y": 200,
  "points": [
    [0, 0],
    [10, 5],
    [20, 10],
    [30, 8]
  ]
}
```

## Pressure

When `simulatePressure: true`, pressure is estimated from drawing speed (faster = thinner). When `false`, uses the `pressures` array directly.

Each pressure value (0.0-1.0) corresponds to a point:

```json
{
  "points": [
    [0, 0],
    [10, 5],
    [20, 3]
  ],
  "pressures": [0.5, 0.7, 0.6]
}
```

## Styling

Freedraw supports standard styling:

```json
{
  "type": "freedraw",
  "strokeColor": "#1971c2",
  "strokeWidth": 2,
  "roughness": 1
}
```

Lower `roughness` (0) produces cleaner lines.

## Use Cases

- Annotations and highlighting
- Hand-drawn emphasis marks
- Organic shapes and curves
- Signature or gesture-like elements
