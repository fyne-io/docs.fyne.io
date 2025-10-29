---
tags: [api]
title: widget.RadioGroup
slug: radiogroup

aliases:
- /api/widget/radiogroup
- /api/widget/radiogroup.html
- /api/v2.0/widget/radiogroup
- /api/v2.0/widget/radiogroup.html
- /api/v2.1/widget/radiogroup
- /api/v2.1/widget/radiogroup.html
- /api/v2.2/widget/radiogroup
- /api/v2.2/widget/radiogroup.html
- /api/v2.3/widget/radiogroup
- /api/v2.3/widget/radiogroup.html
- /api/v2.4/widget/radiogroup
- /api/v2.4/widget/radiogroup.html
- /api/v2.5/widget/radiogroup
- /api/v2.5/widget/radiogroup.html
- /api/v2.6/widget/radiogroup
- /api/v2.6/widget/radiogroup.html
- /api/v2.7/widget/radiogroup
- /api/v2.7/widget/radiogroup.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
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

RadioGroup widget has a list of text labels and checks check icons next to each. Changing the selection (only one can be selected) will trigger the changed func.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewRadioGroup

```go
func NewRadioGroup(options []string, changed func(string)) *RadioGroup
```
NewRadioGroup creates a new radio group widget with the set options and change handler


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

#### func (*RadioGroup) SetSelected

```go
func (r *RadioGroup) SetSelected(option string)
```
SetSelected sets the radio option, it can be used to set a default option.
