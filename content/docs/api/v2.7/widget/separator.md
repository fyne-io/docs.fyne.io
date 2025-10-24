---
tags: [api]
title: widget.Separator
slug: separator

aliases:
- /api/v2.7/widget/separator

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Separator struct {
	BaseWidget
}
```

Separator is a widget for displaying a separator with themeable color.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewSeparator() *Separator
```
NewSeparator creates a new separator.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (s *Separator) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the separator.

###

```go
func (s *Separator) MinSize() fyne.Size
```
MinSize returns the minimal size of the separator.
