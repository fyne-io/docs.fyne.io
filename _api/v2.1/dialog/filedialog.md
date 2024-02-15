---
layout: page
tags: [api]
title: Fyne API "dialog.FileDialog"
---

# dialog.FileDialog
---
```go
import "fyne.io/fyne/v2/dialog"
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
NewFileOpen creates a file dialog allowing the user to choose a file to open. The callback function will run when the dialog closes. The URI will be nil when the user cancels or when nothing is selected.

The dialog will appear over the window specified when Show() is called.

#### func  NewFileSave

```go
func NewFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window) *FileDialog
```
NewFileSave creates a file dialog allowing the user to choose a file to save to (new or overwrite). If the user chooses an existing file they will be asked if they are sure. The callback function will run when the dialog closes. The URI will be nil when the user cancels or when nothing is selected.

The dialog will appear over the window specified when Show() is called.

#### func  NewFolderOpen

```go
func NewFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window) *FileDialog
```
NewFolderOpen creates a file dialog allowing the user to choose a folder to open. The callback function will run when the dialog closes. The URI will be nil when the user cancels or when nothing is selected.

The dialog will appear over the window specified when Show() is called.


<div class="since">Since: <code>
1.4</code></div>

#### func (*FileDialog) Hide

```go
func (f *FileDialog) Hide()
```
Hide hides the file dialog.

#### func (*FileDialog) MinSize

```go
func (f *FileDialog) MinSize() fyne.Size
```
MinSize returns the size that this dialog should not shrink below


<div class="since">Since: <code>
2.1</code></div>

#### func (*FileDialog) Refresh

```go
func (f *FileDialog) Refresh()
```
Refresh causes this dialog to be updated

#### func (*FileDialog) Resize

```go
func (f *FileDialog) Resize(size fyne.Size)
```
Resize dialog to the requested size, if there is sufficient space. If the parent window is not large enough then the size will be reduced to fit.

#### func (*FileDialog) SetDismissText

```go
func (f *FileDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the confirmation button

#### func (*FileDialog) SetFileName

```go
func (f *FileDialog) SetFileName(fileName string)
```
SetFileName sets the filename in a FileDialog in save mode. This is normally called before the dialog is shown.

#### func (*FileDialog) SetFilter

```go
func (f *FileDialog) SetFilter(filter storage.FileFilter)
```
SetFilter sets a filter for limiting files that can be chosen in the file dialog.

#### func (*FileDialog) SetLocation

```go
func (f *FileDialog) SetLocation(u fyne.ListableURI)
```
SetLocation tells this FileDirectory which location to display. This is normally called before the dialog is shown.


<div class="since">Since: <code>
1.4</code></div>

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
