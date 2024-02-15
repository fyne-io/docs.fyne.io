---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type MainMenu

```go
type MainMenu struct {
	Items []*Menu
}
```

MainMenu defines the data required to show a menu bar (desktop) or other appropriate top level menu.

#### func  NewMainMenu

```go
func NewMainMenu(items ...*Menu) *MainMenu
```
NewMainMenu creates a top level menu structure used by fyne.Window for displaying a menubar (or appropriate equivalent).
