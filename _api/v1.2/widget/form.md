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

#### type Form

```go
type Form struct {
	BaseWidget

	Items    []*FormItem
	OnSubmit func()
	OnCancel func()
}
```

Form widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown. Setting OnSubmit will set the submit button to be visible and call back the function when tapped. Setting OnCancel will do the same for a cancel button.

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
