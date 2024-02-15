---
layout: page
tags: [api]
title: Fyne API "container.TabItem"
---

# container.TabItem
---
```go
import "fyne.io/fyne/v2/container"
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

TabItem represents a single view in a tab view. The Text and Icon are used for the tab button and the Content is shown when the corresponding tab is active.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTabItem

```go
func NewTabItem(text string, content fyne.CanvasObject) *TabItem
```
NewTabItem creates a new item for a tabbed widget - each item specifies the content and a label for its tab.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTabItemWithIcon

```go
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem
```
NewTabItemWithIcon creates a new item for a tabbed widget - each item specifies the content and a label with an icon for its tab.


<div class="since">Since: <code>
1.4</code></div>
