---
tags: [api]
title: fyne.Menu
slug: menu

aliases:
- /api//menu
- /api//menu.html
- /api/v2.0//menu
- /api/v2.0//menu.html
- /api/v2.1//menu
- /api/v2.1//menu.html
- /api/v2.2//menu
- /api/v2.2//menu.html
- /api/v2.3//menu
- /api/v2.3//menu.html
- /api/v2.4//menu
- /api/v2.4//menu.html
- /api/v2.5//menu
- /api/v2.5//menu.html
- /api/v2.6//menu
- /api/v2.6//menu.html
- /api/v2.7//menu
- /api/v2.7//menu.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Menu

```go
type Menu struct {
	Label string
	Items []*MenuItem
}
```

Menu stores the information required for a standard menu. A menu can pop down from a [MainMenu] or could be a pop out menu.

#### func  NewMenu

```go
func NewMenu(label string, items ...*MenuItem) *Menu
```
NewMenu creates a new menu given the specified label (to show in a [MainMenu]) and list of items to display.

#### func (*Menu) Refresh

```go
func (m *Menu) Refresh()
```
Refresh will instruct this menu to update its display.


<div class="since">Since: <code>
2.2</code></div>
