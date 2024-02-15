---
layout: page
tags: [api]
title: Fyne API "dialog.ProgressDialog"
---

# dialog.ProgressDialog
---
```go
import "fyne.io/fyne/dialog"
```

## Usage

#### type ProgressDialog

```go
type ProgressDialog struct {
}
```

ProgressDialog is a simple dialog window that displays text and a progress bar.

#### func  NewProgress

```go
func NewProgress(title, message string, parent fyne.Window) *ProgressDialog
```
NewProgress creates a progress dialog and returns the handle. Using the returned type you should call Show() and then set its value through SetValue().

#### func (ProgressDialog) Hide

```go
func (d ProgressDialog) Hide()
```

#### func (ProgressDialog) Layout

```go
func (d ProgressDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ProgressDialog) MinSize

```go
func (d ProgressDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (ProgressDialog) Refresh

```go
func (d ProgressDialog) Refresh()
```

#### func (ProgressDialog) Resize

```go
func (d ProgressDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (ProgressDialog) SetDismissText

```go
func (d ProgressDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ProgressDialog) SetOnClosed

```go
func (d ProgressDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (*ProgressDialog) SetValue

```go
func (p *ProgressDialog) SetValue(v float64)
```
SetValue updates the value of the progress bar - this should be between 0.0 and 1.0.

#### func (ProgressDialog) Show

```go
func (d ProgressDialog) Show()
```
