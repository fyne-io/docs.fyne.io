---
tags: [api]
title: dialog.ProgressDialog
slug: progressdialog

aliases:
- /api/v2.7/dialog/progressdialog

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

#

###

```go
type ProgressDialog struct {
}
```

ProgressDialog is a simple dialog window that displays text and a progress bar.


<div class="deprecated">
Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBar() inside.</div>

###

```go
func NewProgress(title, message string, parent fyne.Window) *ProgressDialog
```
NewProgress creates a progress dialog and returns the handle. Using the returned type you should call Show() and then set its value through SetValue().


<div class="deprecated">
Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBar() inside.</div>

###

```go
func (d ProgressDialog) Dismiss()
```

###

```go
func (d ProgressDialog) Hide()
```

###

```go
func (d ProgressDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func (d ProgressDialog) Refresh()
```

###

```go
func (d ProgressDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

###

```go
func (d ProgressDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

###

```go
func (d ProgressDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

###

```go
func (p *ProgressDialog) SetValue(v float64)
```
SetValue updates the value of the progress bar - this should be between 0.0 and 1.0.

###

```go
func (d ProgressDialog) Show()
```
