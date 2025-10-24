---
tags: [api]
title: widget.RadioGroup
slug: radiogroup

aliases:
- /api/v2.7/widget/radiogroup

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type RadioGroup struct {
	DisableableWidget
	Horizontal bool
	Required   bool
	OnChanged  func(string) `json:"-"`
	Options    []string
	Selected   string
}
```

RadioGroup widget has a list of text labels and checks check icons next to each. Changing the selection (only one can be selected) will trigger the changed func.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewRadioGroup(options []string, changed func(string)) *RadioGroup
```
NewRadioGroup creates a new radio group widget with the set options and change handler


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (r *RadioGroup) Append(option string)
```
Append adds a new option to the end of a RadioGroup widget.

###

```go
func (r *RadioGroup) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (r *RadioGroup) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (r *RadioGroup) Refresh()
```

###

```go
func (r *RadioGroup) SetSelected(option string)
```
SetSelected sets the radio option, it can be used to set a default option.
