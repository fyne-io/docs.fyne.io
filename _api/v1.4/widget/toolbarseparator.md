---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarSeparator"
---

# widget.ToolbarSeparator
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ToolbarSeparator

```go
type ToolbarSeparator struct {
}
```

ToolbarSeparator is a thin, visible divide that can be added to a Toolbar. This is typically used to assist visual grouping of ToolbarItems.

#### func (*ToolbarSeparator) ToolbarObject

```go
func (t *ToolbarSeparator) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the visible line object for this ToolbarSeparator
