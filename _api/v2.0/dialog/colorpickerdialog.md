---
layout: page
tags: [api]
title: Fyne API "dialog.ColorPickerDialog"
---

# dialog.ColorPickerDialog
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

#### func (ColorPickerDialog) Hide

```go
func (d ColorPickerDialog) Hide()
```

#### func (ColorPickerDialog) Layout

```go
func (d ColorPickerDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ColorPickerDialog) MinSize

```go
func (d ColorPickerDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

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
SetDismissText allows custom text to be set in the confirmation button

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
