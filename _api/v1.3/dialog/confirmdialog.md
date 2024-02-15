---
layout: page
tags: [api]
title: Fyne API dialog
---

# dialog
---
```go
import "fyne.io/fyne/dialog"
```

## Usage

#### type ConfirmDialog

```go
type ConfirmDialog struct {
}
```

ConfirmDialog is like the standard Dialog but with an additional confirmation button

#### func  NewConfirm

```go
func NewConfirm(title, message string, callback func(bool), parent fyne.Window) *ConfirmDialog
```
NewConfirm creates a dialog over the specified window for user confirmation. The title is used for the dialog window and message is the content. The callback is executed when the user decides. After creation you should call Show().

#### func (ConfirmDialog) Hide

```go
func (d ConfirmDialog) Hide()
```

#### func (ConfirmDialog) Layout

```go
func (d ConfirmDialog) Layout(obj []fyne.CanvasObject, size fyne.Size)
```

#### func (ConfirmDialog) MinSize

```go
func (d ConfirmDialog) MinSize(obj []fyne.CanvasObject) fyne.Size
```

#### func (*ConfirmDialog) SetConfirmText

```go
func (d *ConfirmDialog) SetConfirmText(label string)
```
SetConfirmText allows custom text to be set in the confirmation button

#### func (ConfirmDialog) SetDismissText

```go
func (d ConfirmDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (ConfirmDialog) SetOnClosed

```go
func (d ConfirmDialog) SetOnClosed(closed func())
```
SetOnClosed allows to set a callback function that is called when the dialog is closed

#### func (ConfirmDialog) Show

```go
func (d ConfirmDialog) Show()
```
