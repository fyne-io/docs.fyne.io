---
layout: page
tags: [api]
title: Fyne API "fyne.MenuItem"
package: fyne.io/fyne/v2
---

# fyne.MenuItem
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type MenuItem

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

#### func  NewMenuItem

```go
func NewMenuItem(label string, action func()) *MenuItem
```
NewMenuItem creates a new menu item from the passed label and action parameters.

#### func  NewMenuItemSeparator

```go
func NewMenuItemSeparator() *MenuItem
```
NewMenuItemSeparator creates a menu item that is to be used as a separator.

#### func  NewMenuItemWithIcon

```go
func NewMenuItemWithIcon(label string, icon Resource, action func()) *MenuItem
```
NewMenuItemWithIcon creates a new menu item from the passed label, icon, and action parameters.


<div class="since">Since: <code>
2.7</code></div>
