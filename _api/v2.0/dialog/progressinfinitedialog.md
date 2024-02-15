---
layout: page
tags: [api]
title: Fyne API "dialog.ProgressInfiniteDialog"
---

# dialog.ProgressInfiniteDialog
---
```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type ProgressInfiniteDialog

```go
type ProgressInfiniteDialog struct {
}
```

ProgressInfiniteDialog is a simple dialog window that displays text and a infinite progress bar.


<div class="deprecated">
Deprecated: Create a new custom dialog with a widget.ProgressBarInfinite() inside.</div>

#### func  NewProgressInfinite

```go
func NewProgressInfinite(title, message string, parent fyne.Window) *ProgressInfiniteDialog
```
NewProgressInfinite creates a infinite progress dialog and returns the handle. Using the returned type you should call Show().


<div class="deprecated">
Deprecated: Create a new custom dialog with a widget.ProgressBarInfinite() inside.</div>

#### func (*ProgressInfiniteDialog) Hide

```go
func (d *ProgressInfiniteDialog) Hide()
```
Hide this dialog and stop the infinite progress goroutine

#### func (ProgressInfiniteDialog) Layout

```go
func (d ProgressInfiniteDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ProgressInfiniteDialog) MinSize

```go
func (d ProgressInfiniteDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (ProgressInfiniteDialog) Refresh

```go
func (d ProgressInfiniteDialog) Refresh()
```

#### func (ProgressInfiniteDialog) Resize

```go
func (d ProgressInfiniteDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (ProgressInfiniteDialog) SetDismissText

```go
func (d ProgressInfiniteDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ProgressInfiniteDialog) SetOnClosed

```go
func (d ProgressInfiniteDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (ProgressInfiniteDialog) Show

```go
func (d ProgressInfiniteDialog) Show()
```
