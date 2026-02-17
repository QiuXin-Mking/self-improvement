# Image Elements

Embedding images in Excalidraw diagrams.

## Image Element

```json
{
  "type": "image",
  "x": 100,
  "y": 100,
  "width": 200,
  "height": 150,
  "fileId": "abc123def456",
  "status": "saved",
  "scale": [1, 1],
  "crop": null,
  "strokeColor": "transparent"
}
```

## Properties

| Property | Type   | Description                       |
| -------- | ------ | --------------------------------- |
| `fileId` | string | Reference to files object         |
| `status` | string | pending, saved, error             |
| `scale`  | [x, y] | Scale factors (-1 to 1 for flips) |
| `crop`   | object | Crop configuration or null        |

Note: `strokeColor` is always `transparent` for images.

## Files Object

Images are stored in the root `files` object:

```json
{
  "files": {
    "abc123def456": {
      "id": "abc123def456",
      "mimeType": "image/png",
      "dataURL": "data:image/png;base64,iVBORw0KGgo...",
      "created": 1690295874454,
      "lastRetrieved": 1690295874454
    }
  }
}
```

### File Properties

| Property        | Type   | Description             |
| --------------- | ------ | ----------------------- |
| `id`            | string | Unique file identifier  |
| `mimeType`      | string | MIME type of image      |
| `dataURL`       | string | Base64 data URL         |
| `created`       | number | Timestamp (ms)          |
| `lastRetrieved` | number | Last access timestamp   |
| `version`       | number | File version (optional) |

### Supported MIME Types

- `image/svg+xml`
- `image/png`
- `image/jpeg`
- `image/gif`
- `image/webp`
- `image/bmp`
- `image/x-icon`
- `image/avif`
- `image/jfif`

## Cropping

```json
{
  "crop": {
    "x": 10,
    "y": 20,
    "width": 100,
    "height": 80,
    "naturalWidth": 400,
    "naturalHeight": 300
  }
}
```

| Property                        | Description                   |
| ------------------------------- | ----------------------------- |
| `x`, `y`                        | Crop origin in original image |
| `width`, `height`               | Crop dimensions               |
| `naturalWidth`, `naturalHeight` | Original image dimensions     |

## Scaling and Flipping

```json
{
  "scale": [1, 1]     // Normal
  "scale": [-1, 1]    // Flip horizontal
  "scale": [1, -1]    // Flip vertical
  "scale": [-1, -1]   // Flip both
}
```

## Base64 Encoding

Data URLs follow this format:

```
data:[mimeType];base64,[base64Data]
```

Base64 encoding increases file size by ~33%.

## Constants

| Constant                          | Value  |
| --------------------------------- | ------ |
| MAX_ALLOWED_FILE_BYTES            | 4MB    |
| DEFAULT_MAX_IMAGE_WIDTH_OR_HEIGHT | 1440px |

## Arrow Binding

Images can have arrows bound to them:

```json
{
  "type": "image",
  "boundElements": [{ "type": "arrow", "id": "arrow-1" }]
}
```
