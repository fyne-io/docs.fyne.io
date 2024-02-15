---
layout: page
tags: [api]
title: Fyne API "widget.Separator"
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
	widget.Base
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

#### func (*Separator) Hide

```go
func (s *Separator) Hide()
```
Hide hides the separator.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) MinSize

```go
func (s *Separator) MinSize() fyne.Size
```
MinSize returns the minimal size of the separator.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) Move

```go
func (s *Separator) Move(pos fyne.Position)
```
Move sets the position of the separator relative to its parent.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) Refresh

```go
func (s *Separator) Refresh()
```
Refresh triggers a redraw of the separator.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) Resize

```go
func (s *Separator) Resize(size fyne.Size)
```
Resize changes the size of the separator.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*Separator) Show

```go
func (s *Separator) Show()
```
Show makes the separator visible.


<div class="implements">Implements: <code>
fyne.Widget</code></div>
