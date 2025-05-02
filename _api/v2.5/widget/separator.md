---
layout: page
tags: [api]
title: Fyne API "widget.Separator"
package: fyne.io/fyne/v2/widget
---

# widget.Separator
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


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) MinSize

```go
func (s *Separator) MinSize() fyne.Size
```
MinSize returns the minimal size of the separator.


<div class="implements">Implements: <code>
fyne.Widget</code></div>
