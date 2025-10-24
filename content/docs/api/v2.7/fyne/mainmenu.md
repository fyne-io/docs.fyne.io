---
tags: [api]
title: fyne.MainMenu
slug: mainmenu

aliases:
- /api/v2.7//mainmenu

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type MainMenu struct {
	Items []*Menu
}
```

MainMenu defines the data required to show a menu bar (desktop) or other appropriate top level menu.

###

```go
func NewMainMenu(items ...*Menu) *MainMenu
```
NewMainMenu creates a top level menu structure used by fyne.Window for displaying a menubar (or appropriate equivalent).

###

```go
func (m *MainMenu) Refresh()
```
Refresh will instruct any rendered menus using this struct to update their display.


<div class="since">Since: <code>
2.2</code></div>
