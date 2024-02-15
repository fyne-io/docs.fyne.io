---
layout: page
tags: [api]
title: Fyne API "widget.PopUpMenu"
---

# widget.PopUpMenu
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type PopUpMenu

```go
type PopUpMenu struct {
	*Menu
}
```

PopUpMenu is a Menu which displays itself in an OverlayContainer.

#### func  NewPopUpMenu

```go
func NewPopUpMenu(menu *fyne.Menu, c fyne.Canvas) *PopUpMenu
```
NewPopUpMenu creates a new, reusable popup menu. You can show it using ShowAtPosition.


<div class="since">Since: <code>
2.0</code></div>

#### func (*PopUpMenu) CreateRenderer

```go
func (p *PopUpMenu) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the pop-up menu.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*PopUpMenu) Hide

```go
func (p *PopUpMenu) Hide()
```
Hide hides the pop-up menu.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*PopUpMenu) Move

```go
func (p *PopUpMenu) Move(pos fyne.Position)
```
Move moves the pop-up menu. The position is absolute because pop-up menus are shown in an overlay which covers the whole canvas.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*PopUpMenu) Resize

```go
func (p *PopUpMenu) Resize(size fyne.Size)
```
Resize changes the size of the pop-up menu.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*PopUpMenu) Show

```go
func (p *PopUpMenu) Show()
```
Show makes the pop-up menu visible.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*PopUpMenu) ShowAtPosition

```go
func (p *PopUpMenu) ShowAtPosition(pos fyne.Position)
```
ShowAtPosition shows the pop-up menu at the specified position.
