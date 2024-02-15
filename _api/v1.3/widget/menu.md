---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type Menu

```go
type Menu struct {
	widget.Base
	Items     []fyne.CanvasObject
	OnDismiss func()
}
```

Menu is a widget for displaying a fyne.Menu.

#### func  NewMenu

```go
func NewMenu(menu *fyne.Menu) *Menu
```
NewMenu creates a new Menu.

#### func (*Menu) CreateRenderer

```go
func (m *Menu) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the menu.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) DeactivateChild

```go
func (m *Menu) DeactivateChild()
```
DeactivateChild deactivates the active child menu.

#### func (*Menu) Dismiss

```go
func (m *Menu) Dismiss()
```
Dismiss dismisses the menu by dismissing and hiding the active child and performing OnDismiss.

#### func (*Menu) Hide

```go
func (m *Menu) Hide()
```
Hide hides the menu.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) MinSize

```go
func (m *Menu) MinSize() fyne.Size
```
MinSize returns the minimal size of the menu.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) Move

```go
func (m *Menu) Move(pos fyne.Position)
```
Move sets the position of the widget relative to its parent.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) Refresh

```go
func (m *Menu) Refresh()
```
Refresh triggers a redraw of the menu.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) Resize

```go
func (m *Menu) Resize(size fyne.Size)
```
Resize has no effect because menus are always displayed with their minimal size.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) Show

```go
func (m *Menu) Show()
```
Show makes the menu visible.

<div class="implements">Implements: <code> fyne.Widget</code></div>

#### func (*Menu) Tapped

```go
func (m *Menu) Tapped(*fyne.PointEvent)
```
Tapped catches taps on separators and the menu background. It doesn't perform any action.

<div class="implements">Implements: <code> fyne.Tappable</code></div>
