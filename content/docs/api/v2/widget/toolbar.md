---
tags: [api]
title: widget.Toolbar
slug: toolbar

aliases:
- /api/v2.0/widget/toolbar
- /api/v2.1/widget/toolbar
- /api/v2.2/widget/toolbar
- /api/v2.3/widget/toolbar
- /api/v2.4/widget/toolbar
- /api/v2.5/widget/toolbar
- /api/v2.6/widget/toolbar
- /api/v2.7/widget/toolbar

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Toolbar

```go
type Toolbar struct {
	BaseWidget
	Items []ToolbarItem
}
```

Toolbar widget creates a horizontal list of tool buttons

#### func  NewToolbar

```go
func NewToolbar(items ...ToolbarItem) *Toolbar
```
NewToolbar creates a new toolbar widget.

#### func (*Toolbar) Append

```go
func (t *Toolbar) Append(item ToolbarItem)
```
Append a new ToolbarItem to the end of this Toolbar

#### func (*Toolbar) CreateRenderer

```go
func (t *Toolbar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Toolbar) MinSize

```go
func (t *Toolbar) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Toolbar) Prepend

```go
func (t *Toolbar) Prepend(item ToolbarItem)
```
Prepend a new ToolbarItem to the start of this Toolbar
