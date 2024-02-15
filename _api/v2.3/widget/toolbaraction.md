---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarAction"
---

# widget.ToolbarAction
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ToolbarAction

```go
type ToolbarAction struct {
	Icon        fyne.Resource
	OnActivated func() `json:"-"`
}
```

ToolbarAction is push button style of ToolbarItem

#### func  NewToolbarAction

```go
func NewToolbarAction(icon fyne.Resource, onActivated func()) *ToolbarAction
```
NewToolbarAction returns a new push button style ToolbarItem

#### func (*ToolbarAction) SetIcon

```go
func (t *ToolbarAction) SetIcon(icon fyne.Resource)
```
SetIcon updates the icon on a ToolbarItem


<div class="since">Since: <code>
2.2</code></div>

#### func (*ToolbarAction) ToolbarObject

```go
func (t *ToolbarAction) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets a button to render this ToolbarAction
