---
layout: page
tags: [api]
title: Fyne API "widget.Form"
package: fyne.io/fyne/v2/widget
---

# widget.Form
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Form

```go
type Form struct {
	BaseWidget

	Items      []*FormItem
	OnSubmit   func() `json:"-"`
	OnCancel   func() `json:"-"`
	SubmitText string
	CancelText string
}
```

Form widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown. Setting OnSubmit will set the submit button to be visible and call back the function when tapped. Setting OnCancel will do the same for a cancel button. If you change OnSubmit/OnCancel after the form is created and rendered, you need to call Refresh() to update the form with the correct buttons. Setting OnSubmit/OnCancel to nil will remove the buttons.

#### func  NewForm

```go
func NewForm(items ...*FormItem) *Form
```
NewForm creates a new form widget with the specified rows of form items and (if any of them should be shown) a form controls row at the bottom

#### func (*Form) Append

```go
func (f *Form) Append(text string, widget fyne.CanvasObject)
```
Append adds a new row to the form, using the text as a label next to the specified Widget

#### func (*Form) AppendItem

```go
func (f *Form) AppendItem(item *FormItem)
```
AppendItem adds the specified row to the end of the Form

#### func (*Form) CreateRenderer

```go
func (f *Form) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Form) Disable

```go
func (f *Form) Disable()
```
Disable disables submitting this form.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Form) Disabled

```go
func (f *Form) Disabled() bool
```
Disabled returns whether submitting the form is disabled. Note that, if the form fails validation, the submit button may be disabled even if this method returns true.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Form) Enable

```go
func (f *Form) Enable()
```
Enable enables submitting this form.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Form) MinSize

```go
func (f *Form) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Form) Refresh

```go
func (f *Form) Refresh()
```
Refresh updates the widget state when requested.

#### func (*Form) SetOnValidationChanged

```go
func (f *Form) SetOnValidationChanged(callback func(error))
```
SetOnValidationChanged is intended for parent widgets or containers to hook into the validation. The function might be overwritten by a parent that cares about child validation (e.g. widget.Form)

#### func (*Form) Validate

```go
func (f *Form) Validate() error
```
Validate validates the entire form and returns the first error that is encountered.
