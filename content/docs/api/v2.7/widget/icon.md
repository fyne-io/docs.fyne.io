---
tags: [api]
title: widget.Icon
slug: icon

aliases:
- /api/v2.7/widget/icon

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Icon struct {
	BaseWidget

	Resource fyne.Resource // The resource for this icon
}
```

Icon widget is a basic image component that load's its resource to match the theme.

###

```go
func NewIcon(res fyne.Resource) *Icon
```
NewIcon returns a new icon widget that displays a themed icon resource

###

```go
func (i *Icon) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (i *Icon) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (i *Icon) SetResource(res fyne.Resource)
```
SetResource updates the resource rendered in this icon widget
