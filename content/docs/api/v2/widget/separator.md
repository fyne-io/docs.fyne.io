---
tags: [api]
title: widget.Separator
slug: separator

aliases:
- /api/v2.0/widget/separator
- /api/v2.1/widget/separator
- /api/v2.2/widget/separator
- /api/v2.3/widget/separator
- /api/v2.4/widget/separator
- /api/v2.5/widget/separator
- /api/v2.6/widget/separator
- /api/v2.7/widget/separator

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Separator

```go
type Separator struct {
	BaseWidget
}
```

Separator is a widget for displaying a separator with themeable color.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewSeparator

```go
func NewSeparator() *Separator
```
NewSeparator creates a new separator.


<div class="since">Since: <code>
1.4</code></div>

#### func (*Separator) CreateRenderer

```go
func (s *Separator) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the separator.

#### func (*Separator) MinSize

```go
func (s *Separator) MinSize() fyne.Size
```
MinSize returns the minimal size of the separator.
