---
tags: [api]
title: dialog.ColorPickerDialog
slug: colorpickerdialog

aliases:
- /api/v2.7/dialog/colorpickerdialog

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

#

###

```go
type ColorPickerDialog struct {
	Advanced bool
}
```

ColorPickerDialog is a simple dialog window that displays a color picker.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewColorPicker(title, message string, callback func(c color.Color), parent fyne.Window) *ColorPickerDialog
```
NewColorPicker creates a color dialog and returns the handle. Using the returned type you should call Show() and then set its color through SetColor(). The callback is triggered when the user selects a color.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (d ColorPickerDialog) Dismiss()
```

###

```go
func (d ColorPickerDialog) Hide()
```

###

```go
func (d ColorPickerDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (p *ColorPickerDialog) Refresh()
```
Refresh causes this dialog to be updated

###

```go
func (d ColorPickerDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

###

```go
func (p *ColorPickerDialog) SetColor(c color.Color)
```
SetColor updates the color of the color picker.

###

```go
func (d ColorPickerDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

###

```go
func (d ColorPickerDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

###

```go
func (p *ColorPickerDialog) Show()
```
Show causes this dialog to be displayed
