---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TabContainer

```go
type TabContainer struct {
	BaseWidget

	Items []*TabItem
}
```

TabContainer widget allows switching visible content from a list of TabItems. Each item is represented by a button at the top of the widget.

#### func  NewTabContainer

```go
func NewTabContainer(items ...*TabItem) *TabContainer
```
NewTabContainer creates a new tab bar widget that allows the user to choose between different visible containers

#### func (*TabContainer) Append

```go
func (t *TabContainer) Append(item *TabItem)
```
Append adds a new TabItem to the rightmost side of the tab panel

#### func (*TabContainer) CreateRenderer

```go
func (t *TabContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*TabContainer) CurrentTab

```go
func (t *TabContainer) CurrentTab() *TabItem
```
CurrentTab returns the currently selected TabItem.

#### func (*TabContainer) CurrentTabIndex

```go
func (t *TabContainer) CurrentTabIndex() int
```
CurrentTabIndex returns the index of the currently selected TabItem.

#### func (*TabContainer) MinSize

```go
func (t *TabContainer) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*TabContainer) Remove

```go
func (t *TabContainer) Remove(item *TabItem)
```
Remove tab by value

#### func (*TabContainer) RemoveIndex

```go
func (t *TabContainer) RemoveIndex(index int)
```
RemoveIndex removes tab by index

#### func (*TabContainer) SelectTab

```go
func (t *TabContainer) SelectTab(item *TabItem)
```
SelectTab sets the specified TabItem to be selected and its content visible.

#### func (*TabContainer) SelectTabIndex

```go
func (t *TabContainer) SelectTabIndex(index int)
```
SelectTabIndex sets the TabItem at the specific index to be selected and its content visible.

#### func (*TabContainer) SetTabLocation

```go
func (t *TabContainer) SetTabLocation(l TabLocation)
```
SetTabLocation sets the location of the tab bar

#### func (*TabContainer) Show

```go
func (t *TabContainer) Show()
```
Show this widget, if it was previously hidden
