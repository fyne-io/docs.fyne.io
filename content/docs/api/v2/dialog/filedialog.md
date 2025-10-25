---
tags: [api]
title: dialog.FileDialog
slug: filedialog

aliases:
- /api/v2/dialog/filedialog.html
- /api/v2.0/dialog/filedialog
- /api/v2.0/dialog/filedialog.html
- /api/v2.1/dialog/filedialog
- /api/v2.1/dialog/filedialog.html
- /api/v2.2/dialog/filedialog
- /api/v2.2/dialog/filedialog.html
- /api/v2.3/dialog/filedialog
- /api/v2.3/dialog/filedialog.html
- /api/v2.4/dialog/filedialog
- /api/v2.4/dialog/filedialog.html
- /api/v2.5/dialog/filedialog
- /api/v2.5/dialog/filedialog.html
- /api/v2.6/dialog/filedialog
- /api/v2.6/dialog/filedialog.html
- /api/v2.7/dialog/filedialog
- /api/v2.7/dialog/filedialog.html

package: fyne.io/fyne/v2/dialog
---


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
func NewFileOpen(callback func(reader fyne.URIReadCloser, err error), parent fyne.Window) *FileDialog
```
NewFileOpen creates a file dialog allowing the user to choose a file to open.

The callback function will run when the dialog closes and provide a reader for the chosen file. The reader will be nil when the user cancels or when nothing is selected. When the reader isn't nil it must be closed by the callback.

The dialog will appear over the window specified when Show() is called.

#### func  NewFileSave

```go
func NewFileSave(callback func(writer fyne.URIWriteCloser, err error), parent fyne.Window) *FileDialog
```
NewFileSave creates a file dialog allowing the user to choose a file to save to (new or overwrite). If the user chooses an existing file they will be asked if they are sure.

The callback function will run when the dialog closes and provide a writer for the chosen file. The writer will be nil when the user cancels or when nothing is selected. When the writer isn't nil it must be closed by the callback.

The dialog will appear over the window specified when Show() is called.

#### func  NewFolderOpen

```go
func NewFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window) *FileDialog
```
NewFolderOpen creates a file dialog allowing the user to choose a folder to open. The callback function will run when the dialog closes. The URI will be nil when the user cancels or when nothing is selected.

The dialog will appear over the window specified when Show() is called.


<div class="since">Since: <code>
1.4</code></div>

#### func (*FileDialog) Dismiss

```go
func (f *FileDialog) Dismiss()
```
Dismiss instructs the dialog to close without any affirmative action.


<div class="since">Since: <code>
2.6</code></div>

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

#### func (*FileDialog) SetConfirmText

```go
func (f *FileDialog) SetConfirmText(label string)
```
SetConfirmText allows custom text to be set in the confirmation button


<div class="since">Since: <code>
2.2</code></div>

#### func (*FileDialog) SetDismissText

```go
func (f *FileDialog) SetDismissText(label string)
```
SetDismissText allows custom text to be set in the dismiss button

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
SetLocation tells this FileDialog which location to display. This is normally called before the dialog is shown.


<div class="since">Since: <code>
1.4</code></div>

#### func (*FileDialog) SetOnClosed

```go
func (f *FileDialog) SetOnClosed(closed func())
```
SetOnClosed sets a callback function that is called when the dialog is closed.

#### func (*FileDialog) SetTitleText

```go
func (f *FileDialog) SetTitleText(label string)
```
SetTitleText allows custom text to be set in the dialog title


<div class="since">Since: <code>
2.6</code></div>

#### func (*FileDialog) SetView

```go
func (f *FileDialog) SetView(v ViewLayout)
```
SetView changes the default display view of the FileDialog This is normally called before the dialog is shown.


<div class="since">Since: <code>
2.5</code></div>

#### func (*FileDialog) Show

```go
func (f *FileDialog) Show()
```
Show shows the file dialog.
