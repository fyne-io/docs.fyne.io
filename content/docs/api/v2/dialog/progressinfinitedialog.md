---
tags: [api]
title: dialog.ProgressInfiniteDialog
slug: progressinfinitedialog

aliases:
- /api/v2.0/dialog/progressinfinitedialog
- /api/v2.1/dialog/progressinfinitedialog
- /api/v2.2/dialog/progressinfinitedialog
- /api/v2.3/dialog/progressinfinitedialog
- /api/v2.4/dialog/progressinfinitedialog
- /api/v2.5/dialog/progressinfinitedialog
- /api/v2.6/dialog/progressinfinitedialog
- /api/v2.7/dialog/progressinfinitedialog

package: fyne.io/fyne/v2/dialog
---


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
Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBarInfinite() inside.</div>

#### func  NewProgressInfinite

```go
func NewProgressInfinite(title, message string, parent fyne.Window) *ProgressInfiniteDialog
```
NewProgressInfinite creates a infinite progress dialog and returns the handle. Using the returned type you should call Show().


<div class="deprecated">
Deprecated: Use NewCustomWithoutButtons() and add a widget.ProgressBarInfinite() inside.</div>

#### func (ProgressInfiniteDialog) Dismiss

```go
func (d ProgressInfiniteDialog) Dismiss()
```

#### func (*ProgressInfiniteDialog) Hide

```go
func (d *ProgressInfiniteDialog) Hide()
```
Hide this dialog and stop the infinite progress goroutine

#### func (ProgressInfiniteDialog) MinSize

```go
func (d ProgressInfiniteDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below.


<div class="since">Since: <code>
2.1</code></div>

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
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (ProgressInfiniteDialog) SetOnClosed

```go
func (d ProgressInfiniteDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (ProgressInfiniteDialog) Show

```go
func (d ProgressInfiniteDialog) Show()
```
