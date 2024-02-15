---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarSpacer"
---

# widget.ToolbarSpacer
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ToolbarSpacer

```go
type ToolbarSpacer struct {
}
```

ToolbarSpacer is a blank, stretchable space for a toolbar. This is typically used to assist layout if you wish some left and some right aligned items. Space will be split evebly amongst all the spacers on a toolbar.

#### func (*ToolbarSpacer) ToolbarObject

```go
func (t *ToolbarSpacer) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the actual spacer object for this ToolbarSpacer
