---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarAction"
---

# widget.ToolbarAction
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ToolbarAction

```go
type ToolbarAction struct {
	Icon        fyne.Resource
	OnActivated func()
}
```

ToolbarAction is push button style of ToolbarItem

#### func (*ToolbarAction) ToolbarObject

```go
func (t *ToolbarAction) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets a button to render this ToolbarAction
