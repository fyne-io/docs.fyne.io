---
tags: [api]
title: dialog (package)
slug: (package)

aliases:
- /api/v2.7/dialog

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

Package dialog defines standard dialog windows for application GUIs.

#

###

```go
func ShowColorPicker(title, message string, callback func(c color.Color), parent fyne.Window)
```
ShowColorPicker creates and shows a color dialog. The callback is triggered when the user selects a color.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func ShowConfirm(title, message string, callback func(bool), parent fyne.Window)
```
ShowConfirm shows a dialog over the specified window for a user confirmation. The title is used for the dialog window and message is the content. The callback is executed when the user decides.

###

```go
func ShowCustom(title, dismiss string, content fyne.CanvasObject, parent fyne.Window)
```
ShowCustom shows a dialog over the specified application using custom content. The button will have the dismiss text set. The MinSize() of the CanvasObject passed will be used to set the size of the window.

###

```go
func ShowCustomConfirm(title, confirm, dismiss string, content fyne.CanvasObject,
	callback func(bool), parent fyne.Window,
)
```
ShowCustomConfirm shows a dialog over the specified application using custom content. The cancel button will have the dismiss text set and the "OK" will use the confirm text. The response callback is called on user action. The MinSize() of the CanvasObject passed will be used to set the size of the window.

###

```go
func ShowCustomWithoutButtons(title string, content fyne.CanvasObject, parent fyne.Window)
```
ShowCustomWithoutButtons shows a dialog, without buttons, over the specified application using custom content. The MinSize() of the CanvasObject passed will be used to set the size of the window.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func ShowEntryDialog(title, message string, onConfirm func(string), parent fyne.Window)
```
ShowEntryDialog creates a new entry dialog and shows it immediately.


<div class="deprecated">
Deprecated: Use dialog.ShowForm() with a widget.Entry inside instead.</div>

###

```go
func ShowError(err error, parent fyne.Window)
```
ShowError shows a dialog over the specified window for an application error. The message is extracted from the provided error (should not be nil).

###

```go
func ShowFileOpen(callback func(reader fyne.URIReadCloser, err error), parent fyne.Window)
```
ShowFileOpen creates and shows a file dialog allowing the user to choose a file to open.

The callback function will run when the dialog closes and provide a reader for the chosen file. The reader will be nil when the user cancels or when nothing is selected. When the reader isn't nil it must be closed by the callback.

The dialog will appear over the window specified.

###

```go
func ShowFileSave(callback func(writer fyne.URIWriteCloser, err error), parent fyne.Window)
```
ShowFileSave creates and shows a file dialog allowing the user to choose a file to save to (new or overwrite). If the user chooses an existing file they will be asked if they are sure.

The callback function will run when the dialog closes and provide a writer for the chosen file. The writer will be nil when the user cancels or when nothing is selected. When the writer isn't nil it must be closed by the callback.

The dialog will appear over the window specified.

###

```go
func ShowFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window)
```
ShowFolderOpen creates and shows a file dialog allowing the user to choose a folder to open. The callback function will run when the dialog closes. The URI will be nil when the user cancels or when nothing is selected.

The dialog will appear over the window specified.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func ShowForm(title, confirm, dismiss string, content []*widget.FormItem, callback func(bool), parent fyne.Window)
```
ShowForm shows a dialog over the specified application using the provided FormItems. The cancel button will have the dismiss text set and the confirm button will use the confirm text. The response callback is called on user action after validation passes. If any Validatable widget reports that validation has failed, then the confirm button will be disabled. The initial state of the confirm button will reflect the initial validation state of the items added to the form dialog. The MinSize() of the CanvasObject passed will be used to set the size of the window.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func ShowInformation(title, message string, parent fyne.Window)
```
ShowInformation shows a dialog over the specified window for user information. The title is used for the dialog window and message is the content.

###

 * [ColorPickerDialog](colorpickerdialog.html)
 * [ConfirmDialog](confirmdialog.html)
 * [CustomDialog](customdialog.html)
 * [Dialog](dialog.html)
 * [EntryDialog](entrydialog.html)
 * [FileDialog](filedialog.html)
 * [FormDialog](formdialog.html)
 * [ProgressDialog](progressdialog.html)
 * [ProgressInfiniteDialog](progressinfinitedialog.html)
 * [ViewLayout](viewlayout.html)
