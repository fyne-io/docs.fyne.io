---
tags: [api]
title: widget.ToolbarSpacer
slug: toolbarspacer

aliases:
- /api/v2.7/widget/toolbarspacer

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ToolbarSpacer struct{}
```

ToolbarSpacer is a blank, stretchable space for a toolbar. This is typically used to assist layout if you wish some left and some right aligned items. Space will be split evebly amongst all the spacers on a toolbar.

###

```go
func NewToolbarSpacer() *ToolbarSpacer
```
NewToolbarSpacer returns a new spacer item for a Toolbar to assist with ToolbarItem alignment

###

```go
func (t *ToolbarSpacer) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the actual spacer object for this ToolbarSpacer
