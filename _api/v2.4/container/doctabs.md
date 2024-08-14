---
layout: page
tags: [api]
title: Fyne API "container.DocTabs"
package: fyne.io/fyne/v2/container
---

# container.DocTabs
---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type DocTabs

```go
type DocTabs struct {
	widget.BaseWidget

	Items []*TabItem

	CreateTab      func() *TabItem
	CloseIntercept func(*TabItem)
	OnClosed       func(*TabItem)
	OnSelected     func(*TabItem)
	OnUnselected   func(*TabItem)
}
```

DocTabs container is used to display various pieces of content identified by tabs. The tabs contain text and/or an icon and allow the user to switch between the content specified in each TabItem. Each item is represented by a button at the edge of the container.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewDocTabs

```go
func NewDocTabs(items ...*TabItem) *DocTabs
```
NewDocTabs creates a new tab container that allows the user to choose between various pieces of content.


<div class="since">Since: <code>
2.1</code></div>

#### func (*DocTabs) Append

```go
func (t *DocTabs) Append(item *TabItem)
```
Append adds a new TabItem to the end of the tab bar.

#### func (*DocTabs) CreateRenderer

```go
func (t *DocTabs) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*DocTabs) DisableIndex

```go
func (t *DocTabs) DisableIndex(i int)
```
DisableIndex disables the TabItem at the specified index.


<div class="since">Since: <code>
2.3</code></div>

#### func (*DocTabs) DisableItem

```go
func (t *DocTabs) DisableItem(item *TabItem)
```
DisableItem disables the specified TabItem.


<div class="since">Since: <code>
2.3</code></div>

#### func (*DocTabs) EnableIndex

```go
func (t *DocTabs) EnableIndex(i int)
```
EnableIndex enables the TabItem at the specified index.


<div class="since">Since: <code>
2.3</code></div>

#### func (*DocTabs) EnableItem

```go
func (t *DocTabs) EnableItem(item *TabItem)
```
EnableItem enables the specified TabItem.


<div class="since">Since: <code>
2.3</code></div>

#### func (*DocTabs) Hide

```go
func (t *DocTabs) Hide()
```
Hide hides the widget.


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>

#### func (*DocTabs) MinSize

```go
func (t *DocTabs) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>

#### func (*DocTabs) Remove

```go
func (t *DocTabs) Remove(item *TabItem)
```
Remove tab by value.

#### func (*DocTabs) RemoveIndex

```go
func (t *DocTabs) RemoveIndex(index int)
```
RemoveIndex removes tab by index.

#### func (*DocTabs) Select

```go
func (t *DocTabs) Select(item *TabItem)
```
Select sets the specified TabItem to be selected and its content visible.

#### func (*DocTabs) SelectIndex

```go
func (t *DocTabs) SelectIndex(index int)
```
SelectIndex sets the TabItem at the specific index to be selected and its content visible.

#### func (*DocTabs) Selected

```go
func (t *DocTabs) Selected() *TabItem
```
Selected returns the currently selected TabItem.

#### func (*DocTabs) SelectedIndex

```go
func (t *DocTabs) SelectedIndex() int
```
SelectedIndex returns the index of the currently selected TabItem.

#### func (*DocTabs) SetItems

```go
func (t *DocTabs) SetItems(items []*TabItem)
```
SetItems sets the containers items and refreshes.

#### func (*DocTabs) SetTabLocation

```go
func (t *DocTabs) SetTabLocation(l TabLocation)
```
SetTabLocation sets the location of the tab bar

#### func (*DocTabs) Show

```go
func (t *DocTabs) Show()
```
Show this widget, if it was previously hidden


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>
