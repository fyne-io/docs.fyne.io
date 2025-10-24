---
tags: [api]
title: widget.ToolbarSeparator
slug: toolbarseparator

aliases:
- /api/v2.7/widget/toolbarseparator

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type ToolbarSeparator struct{}
```

ToolbarSeparator is a thin, visible divide that can be added to a Toolbar. This is typically used to assist visual grouping of ToolbarItems.

###

```go
func NewToolbarSeparator() *ToolbarSeparator
```
NewToolbarSeparator returns a new separator item for a Toolbar to assist with ToolbarItem grouping

###

```go
func (t *ToolbarSeparator) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the visible line object for this ToolbarSeparator
