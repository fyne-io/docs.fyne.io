---
tags: [api]
title: dialog.ConfirmDialog
slug: confirmdialog

aliases:
- /api/v2.7/dialog/confirmdialog

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

#

###

```go
type ConfirmDialog struct {
}
```

ConfirmDialog is like the standard Dialog but with an additional confirmation button

###

```go
func NewConfirm(title, message string, callback func(bool), parent fyne.Window) *ConfirmDialog
```
NewConfirm creates a dialog over the specified window for user confirmation. The title is used for the dialog window and message is the content. The callback is executed when the user decides. After creation you should call Show().

###

```go
func NewCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window,
) *ConfirmDialog
```
NewCustomConfirm creates and returns a dialog over the specified application using custom content. The cancel button will have the dismiss text set and the "OK" will use the confirm text. The response callback is called on user action. The MinSize() of the CanvasObject passed will be used to set the size of the window.

###

```go
func (d *ConfirmDialog) Confirm()
```
Confirm instructs the dialog to close agreeing with whatever content was displayed.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (d ConfirmDialog) Dismiss()
```

###

```go
func (d ConfirmDialog) Hide()
```

###

```go
func (d ConfirmDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (d ConfirmDialog) Refresh()
```

###

```go
func (d ConfirmDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

###

```go
func (d *ConfirmDialog) SetConfirmImportance(importance widget.Importance)
```
SetConfirmImportance sets the importance level of the confirm button.

Since 2.4

###

```go
func (d *ConfirmDialog) SetConfirmText(label string)
```
SetConfirmText allows custom text to be set in the confirmation button

###

```go
func (d ConfirmDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

###

```go
func (d ConfirmDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

###

```go
func (d ConfirmDialog) Show()
```
