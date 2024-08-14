---
layout: page
tags: [api]
title: Fyne API "dialog.EntryDialog"
package: fyne.io/fyne/v2/dialog
---

# dialog.EntryDialog
---
```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type EntryDialog

```go
type EntryDialog struct {
	*FormDialog
}
```

EntryDialog is a variation of a dialog which prompts the user to enter some text.


<div class="deprecated">
Deprecated: Use dialog.NewForm() or dialog.ShowForm() with a widget.Entry inside instead.</div>

#### func  NewEntryDialog

```go
func NewEntryDialog(title, message string, onConfirm func(string), parent fyne.Window) *EntryDialog
```
NewEntryDialog creates a dialog over the specified window for the user to enter a value.

onConfirm is a callback that runs when the user enters a string of text and clicks the "confirm" button. May be nil.


<div class="deprecated">
Deprecated: Use dialog.NewForm() with a widget.Entry inside instead.</div>

#### func (EntryDialog) Hide

```go
func (d EntryDialog) Hide()
```

#### func (EntryDialog) MinSize

```go
func (d EntryDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below


<div class="since">Since: <code>
2.1</code></div>

#### func (EntryDialog) Refresh

```go
func (d EntryDialog) Refresh()
```

#### func (EntryDialog) Resize

```go
func (d EntryDialog) Resize(size fyne.Size)
```
Resize dialog, call this function after dialog show

#### func (EntryDialog) SetDismissText

```go
func (d EntryDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button This is a no-op for dialogs without dismiss buttons.

#### func (*EntryDialog) SetOnClosed

```go
func (i *EntryDialog) SetOnClosed(callback func())
```
SetOnClosed changes the callback which is run when the dialog is closed, which is nil by default.

The callback is called unconditionally whether the user confirms or cancels.

Note that the callback will be called after onConfirm, if both are non-nil. This way onConfirm can potential modify state that this callback needs to get the user input when the user confirms, while also being able to handle the case where the user cancelled.

#### func (*EntryDialog) SetPlaceholder

```go
func (i *EntryDialog) SetPlaceholder(s string)
```
SetPlaceholder defines the placeholder text for the entry

#### func (*EntryDialog) SetText

```go
func (i *EntryDialog) SetText(s string)
```
SetText changes the current text value of the entry dialog, this can be useful for setting a default value.

#### func (EntryDialog) Show

```go
func (d EntryDialog) Show()
```
