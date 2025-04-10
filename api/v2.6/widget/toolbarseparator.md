---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarSeparator"
package: fyne.io/fyne/v2/widget
---

# widget.ToolbarSeparator
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ToolbarSeparator

```go
type ToolbarSeparator struct {
}
```

ToolbarSeparator is a thin, visible divide that can be added to a Toolbar. This is typically used to assist visual grouping of ToolbarItems.

#### func  NewToolbarSeparator

```go
func NewToolbarSeparator() *ToolbarSeparator
```
NewToolbarSeparator returns a new separator item for a Toolbar to assist with ToolbarItem grouping

#### func (*ToolbarSeparator) ToolbarObject

```go
func (t *ToolbarSeparator) ToolbarObject() fyne.CanvasObject
```
ToolbarObject gets the visible line object for this ToolbarSeparator
