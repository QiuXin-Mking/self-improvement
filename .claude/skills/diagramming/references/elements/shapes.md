# Shape Elements

Rectangle, diamond, and ellipse elements.

## Base Properties

Every element requires:

| Property          | Type    | Required | Default       | Description           |
| ----------------- | ------- | -------- | ------------- | --------------------- |
| `id`              | string  | Yes      | -             | Unique identifier     |
| `type`            | string  | Yes      | -             | Element type          |
| `x`               | number  | Yes      | -             | X position (top-left) |
| `y`               | number  | Yes      | -             | Y position (top-left) |
| `width`           | number  | Yes      | -             | Element width         |
| `height`          | number  | Yes      | -             | Element height        |
| `angle`           | number  | Yes      | 0             | Rotation in radians   |
| `strokeColor`     | string  | Yes      | `#1e1e1e`     | Stroke color (hex)    |
| `backgroundColor` | string  | Yes      | `transparent` | Fill color (hex)      |
| `fillStyle`       | string  | Yes      | `solid`       | Fill pattern          |
| `strokeWidth`     | number  | Yes      | 2             | Stroke width (px)     |
| `strokeStyle`     | string  | Yes      | `solid`       | Stroke pattern        |
| `roughness`       | number  | Yes      | 1             | Sketch effect (0-2)   |
| `opacity`         | number  | Yes      | 100           | Opacity (0-100)       |
| `roundness`       | object  | No       | varies        | Corner rounding       |
| `seed`            | number  | Yes      | -             | Random render seed    |
| `version`         | number  | Yes      | 1             | Version counter       |
| `versionNonce`    | number  | Yes      | -             | Random change ID      |
| `isDeleted`       | boolean | Yes      | false         | Soft delete flag      |
| `groupIds`        | array   | Yes      | []            | Group membership      |
| `frameId`         | string  | No       | null          | Parent frame ID       |
| `boundElements`   | array   | No       | null          | Bound text/arrows     |
| `updated`         | number  | Yes      | -             | Timestamp (ms)        |
| `link`            | string  | No       | null          | Clickable URL         |
| `locked`          | boolean | Yes      | false         | Lock state            |
| `index`           | string  | No       | null          | Fractional index      |
| `customData`      | object  | No       | -             | Custom metadata       |

## Rectangle

```json
{
  "type": "rectangle",
  "roundness": { "type": 3 }
}
```

Uses ADAPTIVE_RADIUS (type 3) for smooth corners. Set `roundness: null` for sharp corners.

## Diamond

```json
{
  "type": "diamond",
  "roundness": { "type": 2 }
}
```

Uses PROPORTIONAL_RADIUS (type 2). Supports full arrow binding with 8-sector detection (4 vertices at 15°, 4 diagonals at 75°).

## Ellipse

```json
{
  "type": "ellipse",
  "roundness": { "type": 2 }
}
```

No roundness variations needed—inherently round.

## Roundness Types

| Type | Name                | Default Radius | Used For           |
| ---- | ------------------- | -------------- | ------------------ |
| 1    | LEGACY              | -              | Backward compat    |
| 2    | PROPORTIONAL_RADIUS | 25% of size    | Diamonds, lines    |
| 3    | ADAPTIVE_RADIUS     | 32px           | Rectangles, frames |

## Constraints

| Constant            | Value |
| ------------------- | ----- |
| MIN_WIDTH_OR_HEIGHT | 1     |

## Adding Labels

Shapes can contain text labels via `boundElements`:

```json
{
  "id": "my-shape",
  "type": "rectangle",
  "boundElements": [{ "type": "text", "id": "my-shape-text" }]
}
```

See [text.md](./text.md) for the corresponding text element.

## Arrow Connections

Shapes can have arrows bound to them:

```json
{
  "id": "my-shape",
  "type": "rectangle",
  "boundElements": [
    { "type": "text", "id": "my-shape-text" },
    { "type": "arrow", "id": "arrow-to-shape" }
  ]
}
```

All three shape types (rectangle, diamond, ellipse) are bindable targets for arrows.
