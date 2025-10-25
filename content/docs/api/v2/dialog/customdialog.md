---
tags: [api]
title: dialog.CustomDialog
slug: customdialog

aliases:
- /api/v2/dialog/customdialog.html
- /api/v2.0/dialog/customdialog
- /api/v2.0/dialog/customdialog.html
- /api/v2.1/dialog/customdialog
- /api/v2.1/dialog/customdialog.html
- /api/v2.2/dialog/customdialog
- /api/v2.2/dialog/customdialog.html
- /api/v2.3/dialog/customdialog
- /api/v2.3/dialog/customdialog.html
- /api/v2.4/dialog/customdialog
- /api/v2.4/dialog/customdialog.html
- /api/v2.5/dialog/customdialog
- /api/v2.5/dialog/customdialog.html
- /api/v2.6/dialog/customdialog
- /api/v2.6/dialog/customdialog.html
- /api/v2.7/dialog/customdialog
- /api/v2.7/dialog/customdialog.html

package: fyne.io/fyne/v2/dialog
---


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

#### func (CustomDialog) Dismiss

```go
func (d CustomDialog) Dismiss()
```

#### func (CustomDialog) Hide

```go
func (d CustomDialog) Hide()
```

#### func (CustomDialog) MinSize

```go
func (d CustomDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


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
SetButtons sets the row of buttons at the bottom of the dialog. Passing an empty slice will result in a dialog with no buttons.


<div class="since">Since: <code>
2.4</code></div>

#### func (CustomDialog) SetDismissText

```go
func (d CustomDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (*CustomDialog) SetIcon

```go
func (d *CustomDialog) SetIcon(icon fyne.Resource)
```
SetIcon sets an icon to be shown in the top right of the dialog. Passing a nil resource will remove the icon from the dialog.


<div class="since">Since: <code>
2.6</code></div>

#### func (CustomDialog) SetOnClosed

```go
func (d CustomDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (CustomDialog) Show

```go
func (d CustomDialog) Show()
```
