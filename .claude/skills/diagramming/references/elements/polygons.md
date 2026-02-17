# Polygon Elements

Custom shapes from line elements with closed paths.

## Basic Polygon

Set `polygon: true` on a line element:

```json
{
  "type": "line",
  "points": [
    [0, 0],
    [100, 0],
    [100, 100],
    [0, 100],
    [0, 0]
  ],
  "polygon": true
}
```

Include the first point again at the end to close the path.

## Custom Shapes

### Triangle

```json
{
  "type": "line",
  "points": [
    [50, 0],
    [100, 86],
    [0, 86],
    [50, 0]
  ],
  "polygon": true,
  "backgroundColor": "#a5d8ff",
  "fillStyle": "solid"
}
```

### Pentagon

```json
{
  "type": "line",
  "points": [
    [50, 0],
    [97, 34],
    [79, 90],
    [21, 90],
    [3, 34],
    [50, 0]
  ],
  "polygon": true
}
```

### Hexagon

For pointy-side hexagons with consistent 60° angles regardless of width:

```
endCap = height × 0.577      // keeps angle constant
minWidth = height × 1.155    // below this, shape collapses

Points for width W, height H:
  [endCap, 0], [W - endCap, 0],  // top flat edge
  [W, H/2],                       // right point
  [W - endCap, H], [endCap, H],  // bottom flat edge
  [0, H/2],                       // left point
  [endCap, 0]                     // close
```

Example (width=100, height=50):

```json
{
  "type": "line",
  "points": [
    [29, 0],
    [71, 0],
    [100, 25],
    [71, 50],
    [29, 50],
    [0, 25],
    [29, 0]
  ],
  "polygon": true
}
```

### Star (5-point)

Alternating outer and inner vertices:

```json
{
  "type": "line",
  "points": [
    [50, 0],
    [61, 35],
    [98, 35],
    [68, 57],
    [79, 91],
    [50, 70],
    [21, 91],
    [32, 57],
    [2, 35],
    [39, 35],
    [50, 0]
  ],
  "polygon": true
}
```

### Arrow Shape

```json
{
  "type": "line",
  "points": [
    [0, 30],
    [60, 30],
    [60, 0],
    [100, 50],
    [60, 100],
    [60, 70],
    [0, 70],
    [0, 30]
  ],
  "polygon": true
}
```

## Calculating Vertices

### Regular Polygons

For n-sided regular polygon with radius r:

```
angle = 2π / n
x[i] = r * cos(i * angle - π/2)
y[i] = r * sin(i * angle - π/2)
```

Subtract π/2 to start from top center.

### Stars

For n-pointed star with outer radius R and inner radius r:

```
angle = π / n
x[2i] = R * cos(2i * angle - π/2)     // outer points
y[2i] = R * sin(2i * angle - π/2)
x[2i+1] = r * cos((2i+1) * angle - π/2)  // inner points
y[2i+1] = r * sin((2i+1) * angle - π/2)
```

Inner radius typically 0.38 \* R for classic 5-point star.

## Styling

All fill and stroke styles work with polygons:

```json
{
  "fillStyle": "solid",
  "backgroundColor": "#a5d8ff",
  "strokeColor": "#1971c2",
  "strokeWidth": 2,
  "strokeStyle": "solid",
  "roughness": 1
}
```

## Labeling Polygons

Unlike rectangles and ellipses, polygons don't support text binding with
`containerId`. To add a label that moves with a polygon, put both elements
in the same group using `groupIds`.

```json
[
  {
    "id": "hexagon-1",
    "type": "line",
    "polygon": true,
    "groupIds": ["hexagon-1-group"],
    ...
  },
  {
    "id": "hexagon-1-text",
    "type": "text",
    "text": "Label",
    "containerId": null,
    "groupIds": ["hexagon-1-group"],
    ...
  }
]
```

Position the text manually at the polygon's center. Estimate the text size
(see text.md) and offset accordingly.

## Arrow Bindings

**Polygons cannot have arrow bindings directly.** Excalidraw only allows arrows
to bind to: rectangle, diamond, ellipse, image, iframe, embeddable, frame,
magicframe, and uncontained text.

When a polygon is grouped with a text label, bind arrows to the **text element**
instead. Since the text shares the same `groupIds` as the polygon, moving the
group will move both the polygon and the arrow attachment point together.

```json
[
  {
    "id": "hexagon-shape",
    "type": "line",
    "x": 100,
    "y": 200,
    "width": 100,
    "height": 50,
    "polygon": true,
    "groupIds": ["hexagon-group"],
    "boundElements": null,
    "points": [
      [29, 0],
      [71, 0],
      [100, 25],
      [71, 50],
      [29, 50],
      [0, 25],
      [29, 0]
    ]
  },
  {
    "id": "hexagon-label",
    "type": "text",
    "x": 125,
    "y": 217,
    "text": "Label",
    "containerId": null,
    "groupIds": ["hexagon-group"],
    "boundElements": [{ "type": "arrow", "id": "arrow-to-hexagon" }]
  },
  {
    "id": "arrow-to-hexagon",
    "type": "arrow",
    "startBinding": {
      "elementId": "hexagon-label",
      "fixedPoint": [0.5, 0],
      "mode": "orbit"
    },
    "endBinding": null
  }
]
```

Key points:

- The polygon has `boundElements: null` (it cannot accept bindings)
- The grouped text label receives the arrow binding instead
- Arrows use the text element's ID in `startBinding`/`endBinding`
- Moving the group moves polygon, label, and attached arrows together

## Composition Patterns

### Layered Shapes

Stack shapes with decreasing size for depth effect.

### Complex Outlines

For shapes with holes, layer a smaller shape on top with background color.

## Validation

### Minimum Points

Polygons require at least 4 points (3 vertices + closing point).

### Point Format

Each point is `[x, y]` relative to element position.

### Self-Intersection

Avoid self-intersecting paths for predictable fill behavior.

## Common Shapes Reference

| Shape     | Points | Formula                       |
| --------- | ------ | ----------------------------- |
| Triangle  | 4      | 60° angles                    |
| Square    | 5      | 90° angles                    |
| Pentagon  | 6      | 72° angles                    |
| Hexagon   | 7      | 60° angles                    |
| Octagon   | 9      | 45° angles                    |
| 5-pt Star | 11     | Alternating R/r vertices      |
| 6-pt Star | 13     | Alternating R/r vertices      |
| Arrow     | 8      | Chevron with rectangular tail |
