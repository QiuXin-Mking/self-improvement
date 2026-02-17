# Grouping and Ordering

Groups, z-ordering, and element organization.

## Groups

Elements share group membership via `groupIds`:

```json
{
  "id": "element-1",
  "groupIds": ["group-abc"]
}
{
  "id": "element-2",
  "groupIds": ["group-abc"]
}
```

### Nested Groups

Elements can belong to multiple nested groups (deepest to shallowest order):

```json
{
  "groupIds": ["inner-group", "outer-group"]
}
```

First entry is the most deeply nested group.

### Group Selection

When selecting a grouped element:

1. First click selects the entire group
2. Double-click to select individual elements within

### Polygon Labels

Polygons can't use `containerId` for text binding. Put both polygon and label
in the same group so they move together when dragged.

## Z-Ordering

Element order is controlled by position in the `elements` array.

- Elements later in array appear on top
- First element is at the back

```json
{
  "elements": [
    { "id": "background" },
    { "id": "middle-layer" },
    { "id": "foreground" }
  ]
}
```

### Index Property

For multiplayer scenarios, use fractional index:

```json
{
  "index": "a1"
}
```

Follows rocicorp/fractional-indexing format.

## Visual Containers

For visual boundaries around related elements (not actual groups):

```json
{
  "type": "rectangle",
  "strokeStyle": "dashed",
  "backgroundColor": "transparent",
  "strokeColor": "#868e96"
}
```

## Ordering Best Practices

1. Place background elements (boundaries, containers) first in array
2. Add main content elements in the middle
3. Place foreground elements (labels, annotations) last
4. Keep related elements adjacent in the array

## Lock State

Prevent accidental modification:

```json
{
  "locked": true
}
```

Locked elements:

- Cannot be moved or resized
- Can still be selected
- Must be unlocked to edit
