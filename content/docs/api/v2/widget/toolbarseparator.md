---
tags: [api]
title: widget.ToolbarSeparator
slug: toolbarseparator

aliases:
- /api/v2/widget/toolbarseparator.html
- /api/v2.0/widget/toolbarseparator
- /api/v2.0/widget/toolbarseparator.html
- /api/v2.1/widget/toolbarseparator
- /api/v2.1/widget/toolbarseparator.html
- /api/v2.2/widget/toolbarseparator
- /api/v2.2/widget/toolbarseparator.html
- /api/v2.3/widget/toolbarseparator
- /api/v2.3/widget/toolbarseparator.html
- /api/v2.4/widget/toolbarseparator
- /api/v2.4/widget/toolbarseparator.html
- /api/v2.5/widget/toolbarseparator
- /api/v2.5/widget/toolbarseparator.html
- /api/v2.6/widget/toolbarseparator
- /api/v2.6/widget/toolbarseparator.html
- /api/v2.7/widget/toolbarseparator
- /api/v2.7/widget/toolbarseparator.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type ToolbarSeparator

```go
type ToolbarSeparator struct{}
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
