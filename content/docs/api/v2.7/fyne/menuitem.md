---
tags: [api]
title: fyne.MenuItem
slug: menuitem

aliases:
- /api/v2.7//menuitem

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type MenuItem struct {
	ChildMenu *Menu
	// Since: 2.1
	IsQuit      bool
	IsSeparator bool
	Label       string
	Action      func() `json:"-"`
	// Since: 2.1
	Disabled bool
	// Since: 2.1
	Checked bool
	// Since: 2.2
	Icon Resource
	// Since: 2.2
	Shortcut Shortcut
}
```

MenuItem is a single item within any menu, it contains a display Label and Action function that is called when tapped.

###

```go
func NewMenuItem(label string, action func()) *MenuItem
```
NewMenuItem creates a new menu item from the passed label and action parameters.

###

```go
func NewMenuItemSeparator() *MenuItem
```
NewMenuItemSeparator creates a menu item that is to be used as a separator.

###

```go
func NewMenuItemWithIcon(label string, icon Resource, action func()) *MenuItem
```
NewMenuItemWithIcon creates a new menu item from the passed label, icon, and action parameters.


<div class="since">Since: <code>
2.7</code></div>
