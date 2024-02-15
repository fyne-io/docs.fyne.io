---
layout: page
tags: [api]
title: Fyne API "widget.RadioGroup"
---

# widget.RadioGroup
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type RadioGroup

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

RadioGroup widget has a list of text labels and radio check icons next to each. Changing the selection (only one can be selected) will trigger the changed func.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewRadioGroup

```go
func NewRadioGroup(options []string, changed func(string)) *RadioGroup
```
NewRadioGroup creates a new radio widget with the set options and change handler


<div class="since">Since: <code>
1.4</code></div>

#### func (*RadioGroup) Append

```go
func (r *RadioGroup) Append(option string)
```
Append adds a new option to the end of a RadioGroup widget.

#### func (*RadioGroup) CreateRenderer

```go
func (r *RadioGroup) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*RadioGroup) MinSize

```go
func (r *RadioGroup) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*RadioGroup) Refresh

```go
func (r *RadioGroup) Refresh()
```
Refresh causes this widget to be redrawn in it's current state.


<div class="implements">Implements: <code>
fyne.CanvasObject</code></div>

#### func (*RadioGroup) SetSelected

```go
func (r *RadioGroup) SetSelected(option string)
```
SetSelected sets the radio option, it can be used to set a default option.
