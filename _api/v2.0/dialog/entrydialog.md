---
layout: page
tags: [api]
title: Fyne API "dialog.EntryDialog"
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
}
```

EntryDialog is a variation of a dialog which prompts the user to enter some text.


<div class="deprecated">
Deprecated: Use dialog.NewFormDialog() or dialog.ShowFormDialog() with a widget.Entry inside instead.</div>

#### func  NewEntryDialog

```go
func NewEntryDialog(title, message string, onConfirm func(string), parent fyne.Window) *EntryDialog
```
NewEntryDialog creates a dialog over the specified window for the user to enter a value.

onConfirm is a callback that runs when the user enters a string of text and clicks the "confirm" button. May be nil.


<div class="deprecated">
Deprecated: Use dialog.NewFormDialog() with a widget.Entry inside instead.</div>

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
