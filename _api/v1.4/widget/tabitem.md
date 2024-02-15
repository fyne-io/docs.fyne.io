---
layout: page
tags: [api]
title: Fyne API "widget.TabItem"
---

# widget.TabItem
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TabItem

```go
type TabItem struct {
	Text    string
	Icon    fyne.Resource
	Content fyne.CanvasObject
}
```

TabItem represents a single view in a TabContainer. The Text and Icon are used for the tab button and the Content is shown when the corresponding tab is active.


<div class="deprecated">
Deprecated: use container.TabItem instead.</div>

#### func  NewTabItem

```go
func NewTabItem(text string, content fyne.CanvasObject) *TabItem
```
NewTabItem creates a new item for a tabbed widget - each item specifies the content and a label for its tab.

#### func  NewTabItemWithIcon

```go
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem
```
NewTabItemWithIcon creates a new item for a tabbed widget - each item specifies the content and a label with an icon for its tab.
