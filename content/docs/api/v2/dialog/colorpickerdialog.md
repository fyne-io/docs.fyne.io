---
tags: [api]
title: dialog.ColorPickerDialog
slug: colorpickerdialog

aliases:
- /api/v2/dialog/colorpickerdialog.html
- /api/v2.0/dialog/colorpickerdialog
- /api/v2.0/dialog/colorpickerdialog.html
- /api/v2.1/dialog/colorpickerdialog
- /api/v2.1/dialog/colorpickerdialog.html
- /api/v2.2/dialog/colorpickerdialog
- /api/v2.2/dialog/colorpickerdialog.html
- /api/v2.3/dialog/colorpickerdialog
- /api/v2.3/dialog/colorpickerdialog.html
- /api/v2.4/dialog/colorpickerdialog
- /api/v2.4/dialog/colorpickerdialog.html
- /api/v2.5/dialog/colorpickerdialog
- /api/v2.5/dialog/colorpickerdialog.html
- /api/v2.6/dialog/colorpickerdialog
- /api/v2.6/dialog/colorpickerdialog.html
- /api/v2.7/dialog/colorpickerdialog
- /api/v2.7/dialog/colorpickerdialog.html

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type ColorPickerDialog

```go
type ColorPickerDialog struct {
	Advanced bool
}
```

ColorPickerDialog is a simple dialog window that displays a color picker.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewColorPicker

```go
func NewColorPicker(title, message string, callback func(c color.Color), parent fyne.Window) *ColorPickerDialog
```
NewColorPicker creates a color dialog and returns the handle. Using the returned type you should call Show() and then set its color through SetColor(). The callback is triggered when the user selects a color.


<div class="since">Since: <code>
1.4</code></div>

#### func (ColorPickerDialog) Dismiss

```go
func (d ColorPickerDialog) Dismiss()
```

#### func (ColorPickerDialog) Hide

```go
func (d ColorPickerDialog) Hide()
```

#### func (ColorPickerDialog) MinSize

```go
func (d ColorPickerDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

#### func (*ColorPickerDialog) Refresh

```go
func (p *ColorPickerDialog) Refresh()
```
Refresh causes this dialog to be updated

#### func (ColorPickerDialog) Resize

```go
func (d ColorPickerDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (*ColorPickerDialog) SetColor

```go
func (p *ColorPickerDialog) SetColor(c color.Color)
```
SetColor updates the color of the color picker.

#### func (ColorPickerDialog) SetDismissText

```go
func (d ColorPickerDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (ColorPickerDialog) SetOnClosed

```go
func (d ColorPickerDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (*ColorPickerDialog) Show

```go
func (p *ColorPickerDialog) Show()
```
Show causes this dialog to be displayed
