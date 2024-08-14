---
layout: page
tags: [api]
title: Fyne API "dialog.CustomDialog"
package: fyne.io/fyne/v2/dialog
---

# dialog.CustomDialog
---
```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type CustomDialog

```go
type CustomDialog struct {
}
```

CustomDialog implements a custom dialog.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewCustom

```go
func NewCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window) *CustomDialog
```
NewCustom creates and returns a dialog over the specified application using custom content. The button will have the dismiss text set. The MinSize() of the CanvasObject passed will be used to set the size of the window.

#### func  NewCustomWithoutButtons

```go
func NewCustomWithoutButtons(title string, content fyne.CanvasObject, parent fyne.Window) *CustomDialog
```
NewCustomWithoutButtons creates a new custom dialog without any buttons. The MinSize() of the CanvasObject passed will be used to set the size of the window.


<div class="since">Since: <code>
2.4</code></div>

#### func (CustomDialog) Hide

```go
func (d CustomDialog) Hide()
```

#### func (CustomDialog) MinSize

```go
func (d CustomDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below


<div class="since">Since: <code>
2.1</code></div>

#### func (CustomDialog) Refresh

```go
func (d CustomDialog) Refresh()
```

#### func (CustomDialog) Resize

```go
func (d CustomDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (*CustomDialog) SetButtons

```go
func (d *CustomDialog) SetButtons(buttons []fyne.CanvasObject)
```
SetButtons sets the row of buttons at the bottom of the dialog. Passing an empy slice will result in a dialog with no buttons.


<div class="since">Since: <code>
2.4</code></div>

#### func (CustomDialog) SetDismissText

```go
func (d CustomDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (CustomDialog) SetOnClosed

```go
func (d CustomDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (CustomDialog) Show

```go
func (d CustomDialog) Show()
```
