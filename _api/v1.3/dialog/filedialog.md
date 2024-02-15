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

#### type FileDialog

```go
type FileDialog struct {
}
```

FileDialog is a dialog containing a file picker for use in opening or saving files.

#### func  NewFileOpen

```go
func NewFileOpen(callback func(fyne.URIReadCloser, error), parent fyne.Window) *FileDialog
```
NewFileOpen creates a file dialog allowing the user to choose a file to open. The dialog will appear over the window specified when Show() is called.

#### func  NewFileSave

```go
func NewFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window) *FileDialog
```
NewFileSave creates a file dialog allowing the user to choose a file to save to (new or overwrite). If the user chooses an existing file they will be asked if they are sure. The dialog will appear over the window specified when Show() is called.

#### func (*FileDialog) Hide

```go
func (f *FileDialog) Hide()
```
Hide hides the file dialog.

#### func (*FileDialog) SetDismissText

```go
func (f *FileDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (*FileDialog) SetFilter

```go
func (f *FileDialog) SetFilter(filter storage.FileFilter)
```
SetFilter sets a filter for limiting files that can be chosen in the file dialog.

#### func (*FileDialog) SetOnClosed

```go
func (f *FileDialog) SetOnClosed(closed func())
```
SetOnClosed sets a callback function that is called when the dialog is closed.

#### func (*FileDialog) Show

```go
func (f *FileDialog) Show()
```
Show shows the file dialog.
