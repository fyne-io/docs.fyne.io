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

Package dialog defines standard dialog windows for application GUIs

## Usage

#### func  ShowConfirm

```go
func ShowConfirm(title, message string, callback func(bool), parent fyne.Window)
```
ShowConfirm shows a dialog over the specified window for a user confirmation. The title is used for the dialog window and message is the content. The callback is executed when the user decides.

#### func  ShowCustom

```go
func ShowCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window)
```
ShowCustom shows a dialog over the specified application using custom content. The button will have the dismiss text set. The MinSize() of the CanvasObject passed will be used to set the size of the window.

#### func  ShowCustomConfirm

```go
func ShowCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window)
```
ShowCustomConfirm shows a dialog over the specified application using custom content. The cancel button will have the dismiss text set and the "OK" will use the confirm text. The response callback is called on user action. The MinSize() of the CanvasObject passed will be used to set the size of the window.

#### func  ShowError

```go
func ShowError(err error, parent fyne.Window)
```
ShowError shows a dialog over the specified window for an application error. The title and message are extracted from the provided error.

#### func  ShowInformation

```go
func ShowInformation(title, message string, parent fyne.Window)
```
ShowInformation shows a dialog over the specified window for user information. The title is used for the dialog window and message is the content.

#### types

 * [ConfirmDialog](confirmdialog.html)
 * [Dialog](dialog.html)
 * [ProgressDialog](progressdialog.html)
 * [ProgressInfiniteDialog](progressinfinitedialog.html)
