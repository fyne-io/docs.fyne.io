---
layout: page
tags: [api]
title: Fyne API "fyne.MenuItem"
---

# fyne.MenuItem
---
```go
import "fyne.io/fyne"
```

## Usage

#### type MenuItem

```go
type MenuItem struct {
	ChildMenu   *Menu
	IsSeparator bool
	Label       string
	Action      func()
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
