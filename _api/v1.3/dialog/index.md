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

#### func  NewFileIcon

```go
func NewFileIcon(uri fyne.URI) fyne.CanvasObject
```
NewFileIcon takes a filepath and creates an icon with an overlayed label using the detected mimetype and extension

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

#### func  ShowFileOpen

```go
func ShowFileOpen(callback func(fyne.URIReadCloser, error), parent fyne.Window)
```
ShowFileOpen creates and shows a file dialog allowing the user to choose a file to open. The dialog will appear over the window specified.

#### func  ShowFileSave

```go
func ShowFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window)
```
ShowFileSave creates and shows a file dialog allowing the user to choose a file to save to (new or overwrite). If the user chooses an existing file they will be asked if they are sure. The dialog will appear over the window specified.

#### func  ShowInformation

```go
func ShowInformation(title, message string, parent fyne.Window)
```
ShowInformation shows a dialog over the specified window for user information. The title is used for the dialog window and message is the content.

#### types

 * [ConfirmDialog](confirmdialog.html)
 * [Dialog](dialog.html)
 * [FileDialog](filedialog.html)
 * [ProgressDialog](progressdialog.html)
 * [ProgressInfiniteDialog](progressinfinitedialog.html)
