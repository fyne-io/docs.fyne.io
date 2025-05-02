---
layout: page
tags: [api]
title: Fyne API "widget.Icon"
package: fyne.io/fyne/v2/widget
---

# widget.Icon
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Icon

```go
type Icon struct {
	BaseWidget

	Resource fyne.Resource // The resource for this icon
}
```

Icon widget is a basic image component that load's its resource to match the theme.

#### func  NewIcon

```go
func NewIcon(res fyne.Resource) *Icon
```
NewIcon returns a new icon widget that displays a themed icon resource

#### func (*Icon) CreateRenderer

```go
func (i *Icon) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Icon) MinSize

```go
func (i *Icon) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Icon) SetResource

```go
func (i *Icon) SetResource(res fyne.Resource)
```
SetResource updates the resource rendered in this icon widget
