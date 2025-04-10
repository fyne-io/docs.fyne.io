---
layout: page
tags: [api]
title: Fyne API "layout.CustomPaddedLayout"
package: fyne.io/fyne/v2/layout
---

# layout.CustomPaddedLayout
---
```go
import "fyne.io/fyne/v2/layout"
```

## Usage

#### type CustomPaddedLayout

```go
type CustomPaddedLayout struct {
	TopPadding    float32
	BottomPadding float32
	LeftPadding   float32
	RightPadding  float32
}
```

CustomPaddedLayout is a layout similar to PaddedLayout, but uses custom values for padding on each side, rather than the theme padding value.


<div class="since">Since: <code>
2.5</code></div>

#### func (CustomPaddedLayout) Layout

```go
func (c CustomPaddedLayout) Layout(objects []fyne.CanvasObject, size fyne.Size)
```
Layout is called to pack all child objects into a specified size. For CustomPaddedLayout this sets all children to the full size passed minus the given paddings all around.

#### func (CustomPaddedLayout) MinSize

```go
func (c CustomPaddedLayout) MinSize(objects []fyne.CanvasObject) (min fyne.Size)
```
MinSize finds the smallest size that satisfies all the child objects. For CustomPaddedLayout this is determined simply as the MinSize of the largest child plus the given paddings all around.
