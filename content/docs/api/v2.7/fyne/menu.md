---
tags: [api]
title: fyne.Menu
slug: menu

aliases:
- /api/v2.7//menu

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Menu struct {
	Label string
	Items []*MenuItem
}
```

Menu stores the information required for a standard menu. A menu can pop down from a [MainMenu] or could be a pop out menu.

###

```go
func NewMenu(label string, items ...*MenuItem) *Menu
```
NewMenu creates a new menu given the specified label (to show in a [MainMenu]) and list of items to display.

###

```go
func (m *Menu) Refresh()
```
Refresh will instruct this menu to update its display.


<div class="since">Since: <code>
2.2</code></div>
