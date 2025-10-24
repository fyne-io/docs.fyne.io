---
tags: [api]
title: widget.ToolbarAction
slug: toolbaraction

aliases:
- /api/v2.0/widget/toolbaraction
- /api/v2.1/widget/toolbaraction
- /api/v2.2/widget/toolbaraction
- /api/v2.3/widget/toolbaraction
- /api/v2.4/widget/toolbaraction
- /api/v2.5/widget/toolbaraction
- /api/v2.6/widget/toolbaraction
- /api/v2.7/widget/toolbaraction

package: fyne.io/fyne/v2/widget
---


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

#### func (*ToolbarAction) Disable

```go
func (t *ToolbarAction) Disable()
```
Disable this ToolbarAction so that it cannot be interacted with, updating any style appropriately.


<div class="since">Since: <code>
2.5</code></div>

#### func (*ToolbarAction) Disabled

```go
func (t *ToolbarAction) Disabled() bool
```
Disabled returns true if this ToolbarAction is currently disabled or false if it can currently be interacted with.


<div class="since">Since: <code>
2.5</code></div>

#### func (*ToolbarAction) Enable

```go
func (t *ToolbarAction) Enable()
```
Enable this ToolbarAction, updating any style or features appropriately.


<div class="since">Since: <code>
2.5</code></div>

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
