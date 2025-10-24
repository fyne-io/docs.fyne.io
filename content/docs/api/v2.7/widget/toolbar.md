---
tags: [api]
title: widget.Toolbar
slug: toolbar

aliases:
- /api/v2.7/widget/toolbar

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Toolbar struct {
	BaseWidget
	Items []ToolbarItem
}
```

Toolbar widget creates a horizontal list of tool buttons

###

```go
func NewToolbar(items ...ToolbarItem) *Toolbar
```
NewToolbar creates a new toolbar widget.

###

```go
func (t *Toolbar) Append(item ToolbarItem)
```
Append a new ToolbarItem to the end of this Toolbar

###

```go
func (t *Toolbar) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (t *Toolbar) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (t *Toolbar) Prepend(item ToolbarItem)
```
Prepend a new ToolbarItem to the start of this Toolbar
