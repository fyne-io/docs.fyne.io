---
tags: [api]
title: container.TabItem
slug: tabitem

aliases:
- /api/v2.7/container/tabitem

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

#

###

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

###

```go
func NewTabItem(text string, content fyne.CanvasObject) *TabItem
```
NewTabItem creates a new item for a tabbed widget - each item specifies the content and a label for its tab.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewTabItemWithIcon(text string, icon fyne.Resource, content fyne.CanvasObject) *TabItem
```
NewTabItemWithIcon creates a new item for a tabbed widget - each item specifies the content and a label with an icon for its tab.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (ti *TabItem) Disabled() bool
```
Disabled returns whether or not the TabItem is disabled.


<div class="since">Since: <code>
2.3</code></div>
