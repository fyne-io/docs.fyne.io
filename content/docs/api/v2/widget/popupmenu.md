---
tags: [api]
title: widget.PopUpMenu
slug: popupmenu

aliases:
- /api/v2/widget/popupmenu.html
- /api/v2.0/widget/popupmenu
- /api/v2.0/widget/popupmenu.html
- /api/v2.1/widget/popupmenu
- /api/v2.1/widget/popupmenu.html
- /api/v2.2/widget/popupmenu
- /api/v2.2/widget/popupmenu.html
- /api/v2.3/widget/popupmenu
- /api/v2.3/widget/popupmenu.html
- /api/v2.4/widget/popupmenu
- /api/v2.4/widget/popupmenu.html
- /api/v2.5/widget/popupmenu
- /api/v2.5/widget/popupmenu.html
- /api/v2.6/widget/popupmenu
- /api/v2.6/widget/popupmenu.html
- /api/v2.7/widget/popupmenu
- /api/v2.7/widget/popupmenu.html

package: fyne.io/fyne/v2/widget
---


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

#### func (*PopUpMenu) FocusGained

```go
func (p *PopUpMenu) FocusGained()
```
FocusGained is triggered when the object gained focus. For the pop-up menu it does nothing.

#### func (*PopUpMenu) FocusLost

```go
func (p *PopUpMenu) FocusLost()
```
FocusLost is triggered when the object lost focus. For the pop-up menu it does nothing.

#### func (*PopUpMenu) Hide

```go
func (p *PopUpMenu) Hide()
```
Hide hides the pop-up menu.

#### func (*PopUpMenu) Move

```go
func (p *PopUpMenu) Move(pos fyne.Position)
```
Move moves the pop-up menu. The position is absolute because pop-up menus are shown in an overlay which covers the whole canvas.

#### func (*PopUpMenu) Resize

```go
func (p *PopUpMenu) Resize(size fyne.Size)
```
Resize changes the size of the pop-up menu.

#### func (*PopUpMenu) Show

```go
func (p *PopUpMenu) Show()
```
Show makes the pop-up menu visible.

#### func (*PopUpMenu) ShowAtPosition

```go
func (p *PopUpMenu) ShowAtPosition(pos fyne.Position)
```
ShowAtPosition shows the pop-up menu at the specified position.

#### func (*PopUpMenu) ShowAtRelativePosition

```go
func (p *PopUpMenu) ShowAtRelativePosition(rel fyne.Position, to fyne.CanvasObject)
```
ShowAtRelativePosition shows the pop-up menu at the position relative to given object.

Since 2.4

#### func (*PopUpMenu) TypedKey

```go
func (p *PopUpMenu) TypedKey(e *fyne.KeyEvent)
```
TypedKey handles key events. It allows keyboard control of the pop-up menu.

#### func (*PopUpMenu) TypedRune

```go
func (p *PopUpMenu) TypedRune(rune)
```
TypedRune handles text events. For pop-up menus this does nothing.
