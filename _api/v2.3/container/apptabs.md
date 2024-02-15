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

	Items []*TabItem

	// Deprecated: Use `OnSelected func(*TabItem)` instead.
	OnChanged    func(*TabItem)
	OnSelected   func(*TabItem)
	OnUnselected func(*TabItem)
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
func (t *AppTabs) Append(item *TabItem)
```
Append adds a new TabItem to the end of the tab bar.

#### func (*AppTabs) CreateRenderer

```go
func (t *AppTabs) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*AppTabs) CurrentTab

```go
func (t *AppTabs) CurrentTab() *TabItem
```
CurrentTab returns the currently selected TabItem.


<div class="deprecated">
Deprecated: Use `AppTabs.Selected() *TabItem` instead.</div>

#### func (*AppTabs) CurrentTabIndex

```go
func (t *AppTabs) CurrentTabIndex() int
```
CurrentTabIndex returns the index of the currently selected TabItem.


<div class="deprecated">
Deprecated: Use `AppTabs.SelectedIndex() int` instead.</div>

#### func (*AppTabs) DisableIndex

```go
func (t *AppTabs) DisableIndex(i int)
```
DisableIndex disables the TabItem at the specified index.


<div class="since">Since: <code>
2.3</code></div>

#### func (*AppTabs) DisableItem

```go
func (t *AppTabs) DisableItem(item *TabItem)
```
DisableItem disables the specified TabItem.


<div class="since">Since: <code>
2.3</code></div>

#### func (*AppTabs) EnableIndex

```go
func (t *AppTabs) EnableIndex(i int)
```
EnableIndex enables the TabItem at the specified index.


<div class="since">Since: <code>
2.3</code></div>

#### func (*AppTabs) EnableItem

```go
func (t *AppTabs) EnableItem(item *TabItem)
```
EnableItem enables the specified TabItem.


<div class="since">Since: <code>
2.3</code></div>

#### func (*AppTabs) ExtendBaseWidget

```go
func (t *AppTabs) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.


<div class="deprecated">
Deprecated: Support for extending containers is being removed</div>

#### func (*AppTabs) Hide

```go
func (t *AppTabs) Hide()
```
Hide hides the widget.


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>

#### func (*AppTabs) MinSize

```go
func (t *AppTabs) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>

#### func (*AppTabs) Remove

```go
func (t *AppTabs) Remove(item *TabItem)
```
Remove tab by value.

#### func (*AppTabs) RemoveIndex

```go
func (t *AppTabs) RemoveIndex(index int)
```
RemoveIndex removes tab by index.

#### func (*AppTabs) Select

```go
func (t *AppTabs) Select(item *TabItem)
```
Select sets the specified TabItem to be selected and its content visible.

#### func (*AppTabs) SelectIndex

```go
func (t *AppTabs) SelectIndex(index int)
```
SelectIndex sets the TabItem at the specific index to be selected and its content visible.

#### func (*AppTabs) SelectTab

```go
func (t *AppTabs) SelectTab(item *TabItem)
```
SelectTab sets the specified TabItem to be selected and its content visible.


<div class="deprecated">
Deprecated: Use `AppTabs.Select(*TabItem)` instead.</div>

#### func (*AppTabs) SelectTabIndex

```go
func (t *AppTabs) SelectTabIndex(index int)
```
SelectTabIndex sets the TabItem at the specific index to be selected and its content visible.


<div class="deprecated">
Deprecated: Use `AppTabs.SelectIndex(int)` instead.</div>

#### func (*AppTabs) Selected

```go
func (t *AppTabs) Selected() *TabItem
```
Selected returns the currently selected TabItem.

#### func (*AppTabs) SelectedIndex

```go
func (t *AppTabs) SelectedIndex() int
```
SelectedIndex returns the index of the currently selected TabItem.

#### func (*AppTabs) SetItems

```go
func (t *AppTabs) SetItems(items []*TabItem)
```
SetItems sets the containers items and refreshes.

#### func (*AppTabs) SetTabLocation

```go
func (t *AppTabs) SetTabLocation(l TabLocation)
```
SetTabLocation sets the location of the tab bar

#### func (*AppTabs) Show

```go
func (t *AppTabs) Show()
```
Show this widget, if it was previously hidden


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>
