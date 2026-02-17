# Text Elements

Standalone text and labels bound to containers.

## Text Element

```json
{
  "type": "text",
  "x": 0,
  "y": 0,
  "width": 150,
  "height": 25,
  "text": "Label Text",
  "fontSize": 20,
  "fontFamily": 5,
  "textAlign": "center",
  "verticalAlign": "middle",
  "containerId": null,
  "originalText": "Label Text",
  "autoResize": true,
  "lineHeight": 1.25
}
```

Dimensions: `height = 20 × 1.25 = 25`, `width = 10 × 20 × 0.75 = 150`

## Font Families

| ID  | Name            | Style      |
| --- | --------------- | ---------- |
| 1   | Virgil          | Hand-drawn |
| 2   | Helvetica       | Sans-serif |
| 3   | Cascadia        | Monospace  |
| 5   | Excalifont      | Hand-drawn |
| 6   | Nunito          | Sans-serif |
| 7   | Lilita One      | Display    |
| 8   | Comic Shanns    | Monospace  |
| 9   | Liberation Sans | Sans-serif |
| 10  | Assistant       | Sans-serif |

Note: ID 4 is reserved/unused. Default: `DEFAULT_FONT_FAMILY = 5` (Excalifont).

## Font Selection Guide

Use these fonts by default:

| Context                    | Font         | ID  |
| -------------------------- | ------------ | --- |
| General text, labels       | Excalifont   | 5   |
| Code, monospace, technical | Comic Shanns | 8   |
| Clean, formal text         | Nunito       | 6   |

Always prefer Excalifont unless the content calls for monospace (code snippets,
technical identifiers) or a cleaner sans-serif look (formal documentation).

## Font Sizes

**Always use standard sizes.** Do not use custom values.

| Size | Value | Use Case                   |
| ---- | ----- | -------------------------- |
| S    | 16    | Secondary labels, captions |
| M    | 20    | Default for most text      |
| L    | 28    | Headings, emphasis         |
| XL   | 36    | Titles, large headings     |

Default: M (20). Prefer standard sizes for consistency.

## Text Alignment

| Property        | Values              | Default |
| --------------- | ------------------- | ------- |
| `textAlign`     | left, center, right | left    |
| `verticalAlign` | top, middle, bottom | top     |

## Line Height

| Font ID          | Default lineHeight |
| ---------------- | ------------------ |
| 1 (Virgil)       | 1.25               |
| 2 (Helvetica)    | 1.15               |
| 3 (Cascadia)     | 1.2                |
| 5 (Excalifont)   | 1.25               |
| 6 (Nunito)       | 1.15               |
| 8 (Comic Shanns) | 1.2                |

Multiply by fontSize for pixel height: `lineHeightPx = lineHeight * fontSize`

## Estimating Text Size

For standalone (unbound) text, estimate dimensions:

```
height = fontSize × lineHeight × numberOfLines
width  = characterCount × fontSize × 0.6
```

For **bound text** (text with `containerId`), use generous estimates. Excalidraw
will wrap and position the text automatically. Overestimating width is safer
than clipping.

## Multi-line Text

Use `\n` for line breaks:

```json
{
  "text": "Line One\nLine Two",
  "originalText": "Line One\nLine Two"
}
```

## Container Binding

Text can be bound inside shapes (rectangles, diamonds, ellipses, arrows).

**Important**: The `label` property is for Excalidraw's JavaScript API only. In raw JSON files, you must create two separate elements with `containerId` and `boundElements` referencing each other.

### Shape with bound text

```json
{
  "id": "box-1",
  "type": "rectangle",
  "boundElements": [{ "type": "text", "id": "box-1-text" }]
}
```

### Text referencing container

```json
{
  "id": "box-1-text",
  "type": "text",
  "containerId": "box-1",
  "textAlign": "center",
  "verticalAlign": "middle"
}
```

### Container Padding

`BOUND_TEXT_PADDING = 5px` - padding around text inside containers.

## autoResize Behavior

| Value | Behavior                    |
| ----- | --------------------------- |
| true  | Width expands to fit text   |
| false | Text wraps to element width |

## Text Positioning in Containers

For bound text, set `containerId`, `textAlign`, and `verticalAlign`. Excalidraw
computes the actual position based on internal padding and container geometry.

**Do not manually calculate x/y for bound text.** Use the container's center
as an approximation—Excalidraw will adjust:

```
text.x = container.x + container.width / 2
text.y = container.y + container.height / 2
```

Excalidraw applies `BOUND_TEXT_PADDING = 5px` and shape-specific offsets
(ellipses and diamonds have inscribed rectangle calculations).

### Example: Rectangle with Label

```json
{
  "id": "rect-1",
  "type": "rectangle",
  "x": 100,
  "y": 200,
  "width": 180,
  "height": 40,
  "boundElements": [{ "type": "text", "id": "rect-1-text" }]
}
```

```json
{
  "id": "rect-1-text",
  "type": "text",
  "x": 190,
  "y": 220,
  "width": 60,
  "height": 25,
  "text": "status",
  "originalText": "status",
  "fontSize": 20,
  "fontFamily": 5,
  "lineHeight": 1.25,
  "textAlign": "center",
  "verticalAlign": "middle",
  "containerId": "rect-1",
  "autoResize": true
}
```

### Example: Ellipse with Label

For ellipses, use the same center approximation:

## Text Containers

These elements can contain bound text:

- `rectangle`
- `diamond`
- `ellipse`
- `arrow`
