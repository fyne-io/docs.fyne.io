---
layout: page
tags: [api]
title: Fyne API "container.AppTabs"
---

# container.AppTabs
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type AppTabs

```go
type AppTabs struct {
	widget.BaseWidget

	Items     []*TabItem
	OnChanged func(tab *TabItem)
}
```

AppTabs container is used to split your application into various different areas identified by tabs. The tabs contain text and/or an icon and allow the user to switch between the content specified in each TabItem. Each item is represented by a button at the edge of the container.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewAppTabs

```go
func NewAppTabs(items ...*TabItem) *AppTabs
```
NewAppTabs creates a new tab container that allows the user to choose between different areas of an app.


<div class="since">Since: <code>
1.4</code></div>

#### func (*AppTabs) Append

```go
func (c *AppTabs) Append(item *TabItem)
```
Append adds a new TabItem to the rightmost side of the tab panel

#### func (*AppTabs) CreateRenderer

```go
func (c *AppTabs) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*AppTabs) CurrentTab

```go
func (c *AppTabs) CurrentTab() *TabItem
```
CurrentTab returns the currently selected TabItem.

#### func (*AppTabs) CurrentTabIndex

```go
func (c *AppTabs) CurrentTabIndex() int
```
CurrentTabIndex returns the index of the currently selected TabItem.

#### func (*AppTabs) ExtendBaseWidget

```go
func (c *AppTabs) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.


<div class="deprecated">
Deprecated: Support for extending containers is being removed</div>

#### func (*AppTabs) MinSize

```go
func (c *AppTabs) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*AppTabs) Remove

```go
func (c *AppTabs) Remove(item *TabItem)
```
Remove tab by value

#### func (*AppTabs) RemoveIndex

```go
func (c *AppTabs) RemoveIndex(index int)
```
RemoveIndex removes tab by index

#### func (*AppTabs) SelectTab

```go
func (c *AppTabs) SelectTab(item *TabItem)
```
SelectTab sets the specified TabItem to be selected and its content visible.

#### func (*AppTabs) SelectTabIndex

```go
func (c *AppTabs) SelectTabIndex(index int)
```
SelectTabIndex sets the TabItem at the specific index to be selected and its content visible.

#### func (*AppTabs) SetItems

```go
func (c *AppTabs) SetItems(items []*TabItem)
```
SetItems sets the container's items and refreshes.

#### func (*AppTabs) SetTabLocation

```go
func (c *AppTabs) SetTabLocation(l TabLocation)
```
SetTabLocation sets the location of the tab bar

#### func (*AppTabs) Show

```go
func (c *AppTabs) Show()
```
Show this widget, if it was previously hidden
