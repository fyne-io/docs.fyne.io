---
tags: [api]
title: widget.CheckGroup
slug: checkgroup

aliases:
- /api/v2.7/widget/checkgroup

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type CheckGroup struct {
	DisableableWidget
	Horizontal bool
	Required   bool
	OnChanged  func([]string) `json:"-"`
	Options    []string
	Selected   []string
}
```

CheckGroup widget has a list of text labels and checkbox icons next to each. Changing the selection (any number can be selected) will trigger the changed func.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewCheckGroup(options []string, changed func([]string)) *CheckGroup
```
NewCheckGroup creates a new check group widget with the set options and change handler


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (r *CheckGroup) Append(option string)
```
Append adds a new option to the end of a CheckGroup widget.

###

```go
func (r *CheckGroup) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (r *CheckGroup) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (r *CheckGroup) Refresh()
```
Refresh causes this widget to be redrawn in it's current state.

###

```go
func (r *CheckGroup) Remove(option string) bool
```
Remove removes the first occurrence of the specified option found from a CheckGroup widget. Return true if an option was removed.


<div class="since">Since: <code>
2.3</code></div>

###

```go
func (r *CheckGroup) SetSelected(options []string)
```
SetSelected sets the checked options, it can be used to set a default option.
