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

#### type Dialog

```go
type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
}
```

Dialog is the common API for any dialog window with a single dismiss button

#### func  NewError

```go
func NewError(err error, parent fyne.Window) Dialog
```
NewError creates a dialog over the specified window for an application error. The title and message are extracted from the provided error. After creation you should call Show().

#### func  NewInformation

```go
func NewInformation(title, message string, parent fyne.Window) Dialog
```
NewInformation creates a dialog over the specified window for user information. The title is used for the dialog window and message is the content. After creation you should call Show().
