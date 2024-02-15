---
layout: page
tags: [api]
title: Fyne API "widget.ToolbarItem"
---

# widget.ToolbarItem
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ToolbarItem

```go
type ToolbarItem interface {
	ToolbarObject() fyne.CanvasObject
}
```

ToolbarItem represents any interface element that can be added to a toolbar

#### func  NewToolbarAction

```go
func NewToolbarAction(icon fyne.Resource, onActivated func()) ToolbarItem
```
NewToolbarAction returns a new push button style ToolbarItem

#### func  NewToolbarSeparator

```go
func NewToolbarSeparator() ToolbarItem
```
NewToolbarSeparator returns a new separator item for a Toolbar to assist with ToolbarItem grouping

#### func  NewToolbarSpacer

```go
func NewToolbarSpacer() ToolbarItem
```
NewToolbarSpacer returns a new spacer item for a Toolbar to assist with ToolbarItem alignment
