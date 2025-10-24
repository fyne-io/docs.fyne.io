---
tags: [api]
title: dialog.FormDialog
slug: formdialog

aliases:
- /api/v2.7/dialog/formdialog

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

#

###

```go
type FormDialog struct {
}
```

FormDialog is a simple dialog window for displaying FormItems inside a form.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func NewForm(title, confirm, dismiss string, items []*widget.FormItem, callback func(bool), parent fyne.Window) *FormDialog
```
NewForm creates and returns a dialog over the specified application using the provided FormItems. The cancel button will have the dismiss text set and the confirm button will use the confirm text. The response callback is called on user action after validation passes. If any Validatable widget reports that validation has failed, then the confirm button will be disabled. The initial state of the confirm button will reflect the initial validation state of the items added to the form dialog.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func (d FormDialog) Dismiss()
```

###

```go
func (d FormDialog) Hide()
```

###

```go
func (d FormDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (d FormDialog) Refresh()
```

###

```go
func (d FormDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

###

```go
func (d FormDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

###

```go
func (d FormDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

###

```go
func (d FormDialog) Show()
```

###

```go
func (d *FormDialog) Submit()
```
Submit will submit the form and then hide the dialog if validation passes.


<div class="since">Since: <code>
2.4</code></div>
